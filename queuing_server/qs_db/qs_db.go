package qs_db

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

type AuthorizationRequest struct {
	Ticket          string    `db:"ticket" json:"ticket"`
	RequestedScopes string    `db:"requested_scopes" json:"requested_scopes"`
	ClientRequest   string    `db:"client_request" json:"client_request"`
	UserId          int       `db:"user_id"`
	Status          string    `db:"status"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}

type QsId struct {
	QsId   string `db:"qs_id"`
	UserId int    `db:"user_id"`
}

type Ticket struct {
	TicketId string `db:"ticket_id"`
	QsId     string `db:"qs_id"`
	UserId   int    `db:"user_id"`
}

func ConnectDB() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", "qs_db/sample.sqlite3")
}

func init() {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		panic(err)
	}
}

func SelectWaitingRequests(user_id int) ([]AuthorizationRequest, error) {
	waitingRequests := &[]AuthorizationRequest{}
	query := `SELECT requested_scopes,created_at FROM authorization_requests WHERE status = "waiting" and user_id = ?`
	err := DB.Select(waitingRequests, query, user_id)
	if err != nil {
		return nil, err
	}
	return *waitingRequests, nil
}

func InsertRequest(req *AuthorizationRequest) error {
	userId, err := GetUserIdFromTicketId(req.Ticket)
	if err != nil {
		return err
	}
	req.UserId = userId
	req.Status = "pending"
	_, err = DB.NamedExec(`INSERT INTO authorization_requests (ticket, requested_scopes, user_id, status) VALUES (:ticket, :requested_scopes, :user_id, :status)`, req)
	return err
}

func AddClientRequest(ticketId, clientRequest string) error {
	var reqId int
	newStatus := "waiting"
	err := DB.Get(&reqId, `SELECT request_id FROM authorization_requests WHERE ticket = ?`, ticketId)
	if err != nil {
		return err
	}
	_, err = DB.Exec(`UPDATE authorization_requests SET status = ?, client_request = ? WHERE request_id = ?`, newStatus, clientRequest, reqId)
	return err
}

func GetUserIdFromQsId(qs_id string) (int, error) {
	var userId int
	err := DB.Get(&userId, `SELECT user_id FROM qs_ids WHERE qs_id = ?`, qs_id)
	if err != nil {
		return -1, err
	}
	return userId, nil
}
func GetUserIdFromTicketId(ticketId string) (int, error) {
	var userId int
	err := DB.Get(&userId, `SELECT user_id FROM tickets WHERE ticket_id = ?`, ticketId)
	if err != nil {
		return -1, err
	}
	return userId, nil
}
func InsertTicket(ticketId, qsId string) error {
	var userId int
	userId, err := GetUserIdFromQsId(qsId)
	if err != nil {
		return err
	}
	ticket := Ticket{
		TicketId: ticketId,
		QsId:     qsId,
		UserId:   userId,
	}
	DB.NamedExec(`INSERT INTO tickets (ticket_id, qs_id, user_id) VALUES (:ticket_id, :qs_id, :user_id)`, ticket)

	return nil
}
