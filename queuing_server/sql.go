package queuing_server

import (
	"time"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type AuthorizationRequest struct {
	ClientDomain string     `db:"client_domain" json:"client_domain"`
	RSDomain     string     `db:"rs_domain" json:"rs_domain"` // resource server domain
	Scopes       string     `db:"scopes" json:"scopes"`       // scopes separeted by space
	Timestamp    *time.Time `json:"timestamp"`
}

func ConnectDB() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", "sql/sample.sqlite3")
}

func GetWaitingRequests() ([]AuthorizationRequest, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	waitingRequests := []AuthorizationRequest{}
	query := `SELECT * FROM waiting_requests WHERE state = "waiting"`
	err = db.Get(waitingRequests, query)
	if err != nil {
		return nil, err
	}
	return waitingRequests, nil
}

func SaveRequests(req []*AuthorizationRequest) error {
	// sqlquery := `INSERT INTO pooled_requests ()`
	// DB.Query("INSERT INTO ")
	return nil
}
