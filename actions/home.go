package actions

import (
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	// tx, ok := c.Value("tx").(*pop.Connection)
	// if !ok {
	// 	return errors.WithStack(errors.New("no transaction found"))
	// }

	// users := &models.Users{}

	// if err := tx.All(users); err != nil {
	// 	return errors.WithStack(err)
	// }

	// log.Println(users)
	return c.Render(200, r.JSON(map[string]string{"Hi": "Welcome to Suunto Things, Created with Buffalo!"}))
}
