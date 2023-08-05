package queuing_server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/laft2/self-managed-uma/queuing_server/qs_db"
)

func AddClientGroup(e *echo.Echo) {
	queueGroup := e.Group("/client")
	queueGroup.POST("/request", func(c echo.Context) error {
		ticket := c.FormValue("ticket")
		clientReq := c.FormValue("request")
		err := qs_db.AddClientRequest(ticket, clientReq)
		if err != nil {
			return c.JSON(InvalidRequest.StatusCode, InvalidRequest)
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"status": "waiting",
		})
	})
}
