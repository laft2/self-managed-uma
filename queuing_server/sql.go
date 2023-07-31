package queuing_server

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

type AuthorizationRequest struct {
	Id           int        `db:"id" json:"id"`
	ClientDomain string     `db:"client_domain" json:"client_domain"`
	RSDomain     string     `db:"rs_domain" json:"rs_domain"` // resource server domain
	Scopes       string     `db:"scopes" json:"scopes"`       // scopes separeted by space
	Status       string     `db:"status" json:"status"`
	AddedAt      *time.Time `db:"added_at" json:"added_at"`
}

func ConnectDB() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", "sql/sample.sqlite3")
}

func GetWaitingRequests() ([]AuthorizationRequest, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	waitingRequests := &[]AuthorizationRequest{}
	query := `SELECT * FROM authorization_requests WHERE status = "waiting"`
	err = db.Select(waitingRequests, query)
	if err != nil {
		return nil, err
	}
	return *waitingRequests, nil
}

func SaveRequests(req []*AuthorizationRequest) error {
	// sqlquery := `INSERT INTO pooled_requests ()`
	// DB.Query("INSERT INTO ")
	return nil
}
