package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB

func ConnectDB() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", "sql/sample.sqlite3")
}

type User struct {
	ID       string `db:"id"`       // id
	Email    string `db:"email"`    // email
	Salt     string `db:"salt"`     // salt
	Password string `db:"password"` // password
}

type AuthForm struct {
	Identifier string `form:"identifier"`
	Passphrase string `form:"passphrase"`
}

type IndexTemplate struct {
	ID string
}

var store *sessions.CookieStore

func init() {
	secretkey := securecookie.GenerateRandomKey(32)
	store = sessions.NewCookieStore(secretkey)
}
func renderHTML(htmlpath string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(htmlpath)
	if err != nil {
		return err
	}
	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}

func redirectUnauthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := store.Get(c.Request(), "suma_rs")
		if err != nil {
			return c.Redirect(http.StatusFound, "/authenticate")
		}
		id, ok := sess.Values["id"]
		if !ok {
			return c.Redirect(http.StatusFound, "/authenticate")
		}
		_, ok = id.(string)
		if !ok {
			return c.Redirect(http.StatusFound, "/authenticate")
		}
		return next(c)
	}
}
func authenticate(c echo.Context, userId string, password string) (res bool, err error) {
	user := User{}
	query := `SELECT * FROM users WHERE id = ?`
	res = false
	err = db.Get(&user, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Logger().Infof("user id is not found: %v", userId)
		} else {
			c.Logger().Errorf("db error: %v", err)
		}
		return
	}
	hashedPasswordByte := sha256.Sum256([]byte(password + user.Salt))
	hashedPasswordHex := hex.EncodeToString(hashedPasswordByte[:])
	fmt.Printf("hashedPasswordHex: %v\n", hashedPasswordHex)
	fmt.Printf("user.Password: %v\n", user.Password)
	if hashedPasswordHex != user.Password {
		err = fmt.Errorf("incorrect password")
		return
	}
	res = true
	return
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static/css/", "front/css/")
	e.Static("/static/image/", "front/image/")

	var err error
	db, err = ConnectDB()
	if err != nil {
		e.Logger.Fatalf("db connection failed: %v", err)
	}
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		sess, _ := store.Get(c.Request(), "suma_rs")
		id := sess.Values["id"]
		sess.Save(c.Request(), c.Response())
		return renderHTML("front/index.html", c.Response(), IndexTemplate{
			ID: id.(string),
		})
	}, redirectUnauthenticatedUser)
	e.GET("/authenticate", func(c echo.Context) error {
		return c.File("front/authn.html")
	})
	e.POST("/authenticate", func(c echo.Context) error {
		authform := &AuthForm{}
		c.Bind(authform)
		isAuthenticated, err := authenticate(c, authform.Identifier, authform.Passphrase)
		if err != nil {
			c.Logger().Error(err)
			return c.Redirect(http.StatusSeeOther, "/authenticate")
		}
		if isAuthenticated {
			sess, _ := store.Get(c.Request(), "suma_rs")
			sess.Values["id"] = authform.Identifier
			sess.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusFound, "/")
		} else {
			return c.Redirect(http.StatusSeeOther, "/authenticate")
		}
	})
	port := os.Getenv("RS_PORT")
	if port == "" {
		e.Logger.Fatal(e.Start("localhost:10001"))
	} else {
		// for docker environment
		e.Logger.Fatal(e.Start(":" + port))
	}
}
