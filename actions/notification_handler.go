package actions

import (
	"fmt"

	"github.com/edTheGuy00/suuntothings/models"
	"github.com/edTheGuy00/suuntothings/util"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// NotificationHandler handle incomint suunto notificaitons
func NotificationHandler(c buffalo.Context) error {
	n := &models.NotificationPayload{}
	if err := c.Bind(n); err != nil {
		return errors.WithStack(err)
	}

	message := fmt.Sprintf("%s: %s %s: %s", "Nofications for", n.UserName, "With Workout ID", n.WorkoutID)

	util.LogToSlack(message)

	// tx := c.Value("tx").(*pop.Connection)

	// // find a user with the user_name
	// err := tx.Where("user_name = ?", strings.ToLower(strings.TrimSpace(n.UserName))).First(n)

	// // helper function to handle bad attempts
	// bad := func() error {
	// 	verrs := validate.NewErrors()
	// 	verrs.Add("UserName", "invalid userName")
	// 	return c.Render(422, r.JSON(verrs))
	// }

	// if err != nil {
	// 	if errors.Cause(err) == sql.ErrNoRows {
	// 		// couldn't find an user with the supplied email address.
	// 		return bad()
	// 	}
	// 	return errors.WithStack(err)
	// }

	return c.Render(200, r.JSON(""))
}
