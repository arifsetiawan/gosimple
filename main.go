// Simple go app with echo framework and couchbase
package main

import (
	"fmt"
	"net/http"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/Sirupsen/logrus"
	"github.com/couchbase/gocb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

var log = logrus.New()
var bucket *gocb.Bucket

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.DebugLevel
}

// Error object
type Error struct {
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// Response JSONAPI object
type Response struct {
	Errors []Error     `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// User is a single user
type User struct {
	Type       string   `json:"_type,omitempty"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	PositionID string   `json:"position_id,omitempty"`
	Position   string   `json:"position,omitempty"`
	Group      []string `json:"groups,omitempty"`
}

func makeErrorResponse(err error, status int) *Response {
	r := new(Response)
	es := make([]Error, 1)
	es[0] = Error{Status: status, Title: err.Error()}
	r.Errors = es
	return r
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, GoBase!")
}

func createUser(c echo.Context) error {
	u := new(User)
	u.Type = "user"

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError,
			makeErrorResponse(err, http.StatusInternalServerError))
	}

	if _, err := bucket.Upsert("user:"+u.Email, u, 0); err != nil {
		return c.JSON(http.StatusInternalServerError,
			makeErrorResponse(err, http.StatusInternalServerError))
	}

	r := new(Response)
	r.Data = u
	return c.JSON(http.StatusOK, r)
}

func getUser(c echo.Context) error {
	u := new(User)
	email := c.Param("email")

	if _, err := bucket.Get(email, &u); err != nil {
		return c.JSON(http.StatusInternalServerError,
			makeErrorResponse(err, http.StatusInternalServerError))
	}

	r := new(Response)
	r.Data = u
	return c.JSON(http.StatusOK, r)
}

func listUsers(c echo.Context) error {
	listUserQuery := gocb.NewN1qlQuery("SELECT u.email, u.name, u.position_id, u.position FROM gobase u WHERE _type='user'")
	rows, err := bucket.ExecuteN1qlQuery(listUserQuery, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			makeErrorResponse(err, http.StatusInternalServerError))
	}

	var user User
	var users []User
	for i := 0; rows.Next(&user); i++ {
		users = append(users, user)
	}
	_ = rows.Close()

	r := new(Response)
	r.Data = users
	return c.JSON(http.StatusOK, r)
}

func deleteUser(c echo.Context) error {
	email := c.Param("email")

	if _, err := bucket.Remove(email, 0); err != nil {
		return c.JSON(http.StatusInternalServerError,
			makeErrorResponse(err, http.StatusInternalServerError))
	}

	r := new(Response)
	return c.JSON(http.StatusOK, r)
}

func main() {

	initializeConfig()

	fmt.Println(viper.GetString("CouchbaseURI"))
	fmt.Println(viper.GetString("CouchbaseBucket"))

	cluster, err := gocb.Connect(viper.GetString("CouchbaseURI"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"app": "gobase",
			"err": err,
		}).Fatal("Failed to connect to database!")
		return
	}

	bucket, err = cluster.OpenBucket(
		viper.GetString("CouchbaseBucket"), viper.GetString("CouchbasePassword"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"app": "gobase",
			"err": err,
		}).Fatal("Failed to open bucket!")
		return
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: &lumberjack.Logger{
			Filename: viper.GetString("LogPath"),
			MaxSize:  2,
		},
	}))
	e.Use(middleware.Recover())

	e.GET("/", home)

	e.POST("/users", createUser)
	e.GET("/users/:email", getUser)
	e.DELETE("/users/:email", deleteUser)
	e.GET("/users", listUsers)

	log.WithFields(logrus.Fields{
		"app": "gobase",
	}).Debug("App listening to " + viper.GetString("Port"))

	e.Run(standard.New(viper.GetString("Port")))
}
