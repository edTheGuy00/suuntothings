package grifts

import (
	"github.com/edTheGuy00/suuntothings/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
