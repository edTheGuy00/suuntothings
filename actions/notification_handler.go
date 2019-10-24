package actions

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/edTheGuy00/suuntothings/models"
	"github.com/edTheGuy00/suuntothings/util"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
)

// NotificationHandler handle incomint suunto notificaitons
func NotificationHandler(c buffalo.Context) error {
	n := &models.NotificationPayload{}
	if err := c.Bind(n); err != nil {
		return errors.WithStack(err)
	}

	message := fmt.Sprintf("%s: %s %s: %s. %s", "Nofication for", n.UserName, "With Workout ID", n.WorkoutID, "retrieving workout details from Suunto")

	util.LogToSlack(message)

	tx := c.Value("tx").(*pop.Connection)

	u := &models.User{}

	// // find a user with the user_name
	err := tx.Where("user_name = ?", n.UserName).First(u)

	// // helper function to handle bad attempts
	bad := func() error {
		verrs := validate.NewErrors()
		verrs.Add("UserName", "invalid userName")
		return c.Render(422, r.JSON(verrs))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", "https://cloudapi.suunto.com/v2/workouts", n.WorkoutID), nil)
	// ...
	req.Header.Add("content-type", "application/javascript")
	req.Header.Add("authorization", u.RefreshToken)
	req.Header.Add("ocp-apim-subscription-key", envy.Get("SUBSCRIPTION_KEY", ""))
	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	util.LogToSlack(string(body))

	return c.Render(200, r.JSON(""))
}
