package queuing_server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// https://docs.kantarainitiative.org/uma/wg/rec-oauth-uma-federated-authz-2.0.html#permission-endpoint

type PermRequest struct {
	ResourceID     string   `json:"resource_id"`
	ResourceScopes []string `json:"resource_scopes"`
}
type PermissionTicketMap map[string]PermissionTicket
type PermissionTicket struct {
	ID          string `json:"ticket"`
	PermRequest PermRequest
}

// Resource Server Rquest to Permission Endpoint
func PostPerm(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	pat := authHeader[len("Bearer "):]
	fmt.Printf("pat: %v\n", pat) // TODO: implement RS authentication with pat

	permRequest := &PermRequest{}
	c.Bind(permRequest)
	// reject if not have required field
	if permRequest.ResourceID == "" || len(permRequest.ResourceScopes) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":             "invalid_resource_id",
			"error_description": "Permission request failed with bad resource ID.",
		})
	}

	permissionTicket := PermissionTicket{}
	tikcetUuid, _ := uuid.NewUUID()
	ticketId := tikcetUuid.String()
	permissionTicket.ID = ticketId
	permissionTicket.PermRequest = *permRequest
	permissionTicketStore[ticketId] = permissionTicket

	return c.JSON(http.StatusCreated, permissionTicket)
}

func GetTicketInfo(ticket string) (*PermissionTicket, error) {
	if ticket == "" {
		return nil, fmt.Errorf("no ticket included")
	}
	ticketInfo, ok := permissionTicketStore[ticket]
	if !ok {
		return nil, fmt.Errorf("invalid ticket")
	}
	return &ticketInfo, nil
}
