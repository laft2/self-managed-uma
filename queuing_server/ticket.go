package queuing_server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laft2/self-managed-uma/queuing_server/qs_db"
)

func generateNewTicket(qsId string) (string, error) {
	ticketUuid, _ := uuid.NewUUID()
	ticketId := ticketUuid.String()
	err := qs_db.InsertTicket(ticketId, qsId)
	if err != nil {
		return "", err
	}
	return ticketId, nil
}

func saveAuthorizationRequest(ticketId, requestedScopes string) error {
	authzReq := &qs_db.AuthorizationRequest{
		Ticket:          ticketId,
		RequestedScopes: requestedScopes,
	}
	err := qs_db.InsertRequest(authzReq)
	return err
}

func AddTicketGroup(e *echo.Echo) {
	g := e.Group("ticket")
	g.POST("/", func(c echo.Context) error {
		// TODO: should authenticate qs_id
		// curl example: curl -XPOST -d 'qs_id=sample_qs_id&requested_scopes=thisisencrypted.in.jwe' 'http://localhost:9010/ticket/'

		qsId := c.FormValue("qs_id")
		requestedScopes := c.FormValue("requested_scopes")
		if qsId == "" || requestedScopes == "" {
			return c.JSON(http.StatusBadRequest, ErrorJSON{
				Error: "qs_id or requested_scopes are empty",
			})
		}
		ticket, err := generateNewTicket(qsId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorJSON{
				Error: "invalid qs_id",
			})
		}
		err = saveAuthorizationRequest(ticket, requestedScopes)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"ticket": ticket,
		})
	})
}
