package queuing_server

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

type AuthorizationRequest struct {
	Id               int        `db:"id" json:"id"`
	GrantType        string     `db:"grant_type" json:"grant_type"`
	Ticket           string     `db:"ticket" json:"ticket"`
	ClaimToken       *string    `db:"claim_token" json:"claim_token"`
	ClaimTokenFormat *string    `db:"claim_token_format" json:"claim_token_format"`
	Pct              *string    `db:"pct" json:"pct"`
	Rpt              *string    `db:"rpt" json:"rpt"`
	Scopes           *string     `db:"scopes" json:"scopes"` // scopes separeted by space
	Status           string     `db:"status" json:"status"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
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
