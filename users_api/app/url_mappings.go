package app

import (
	"github.com/bookstore_app/users_api/controllers/ping"
	"github.com/bookstore_app/users_api/controllers/users"
)

func mapUrls() {
	// ping let's cloud traffic know that we will be running a server on this port
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}
