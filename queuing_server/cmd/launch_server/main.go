package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laft2/self-managed-uma/queuing_server"
)

type TokenForm struct {
	GrantType        string `form:"grant_type"`
	Ticket           string `form:"ticket"`
	ClaimToken       string `form:"claim_token"`
	ClaimTokenFormat string `form:"claim_token_format"`
	Pct              string `form:"pct"`
	Rpt              string `form:"rpt"`
	Scope            string `form:"scope"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// TODO: manage resource
	e.GET("/user", func(c echo.Context) error {
		return nil
	})
	queuing_server.AddUmaGroup(e)

	// Start server
	e.Logger.Fatal(e.Start("localhost:9010"))
}
