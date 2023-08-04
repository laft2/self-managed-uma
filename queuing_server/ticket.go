package queuing_server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/laft2/self-managed-uma/queuing_server/qs_db"
)

func GenerateNewTicket(qsId string)(string, error){
  ticketUuid, _ := uuid.NewUUID()
  ticketId := ticketUuid.String()
  err := qs_db.InsertTicket(ticketId, qsId)
  if err != nil {
    return "", err
  }
  return ticketId, nil
}

func AddTicketGroup(e *echo.Echo){
  g := e.Group("ticket")
  g.GET("/:qs_id", func(c echo.Context) error {
    // TODO: should authenticate qs_id
    qsId := c.Param("qs_id")
    ticket, err := GenerateNewTicket(qsId)
    if err != nil {
      return err
    }
    return c.JSON(http.StatusCreated, map[string]string{
      "ticket": ticket,
    })
  })
}