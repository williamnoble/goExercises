package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/williamnoble/goExercises/http/gin/gin-mongo/internal"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Users struct {
	UsersModel internal.UsersModel
}

//func (u Users) CreateUserHandler(c *gin.Context) {
//
//}
//
//func (u Users) UserByIdHandler(c *gin.Context) {
//
//}
//func (u Users) ListUsersHandler(c *gin.Context) {
//
//}
//func (u Users) AuthenicationHandler(c *gin.Context) {
//
//}
//func (u Users) DeleteUserHandler(c *gin.Context) {
//
//}
//func (u Users) UpdateUserHandler(c *gin.Context) {
//
//}

func (u Users) IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Index page!"})
}

func (u Users) CreateUserHandler(c *gin.Context) {
	var user internal.User
	err := c.BindJSON(&user)
	user.ID = bson.NewObjectId()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "something went wrong",
		})
		return
	}
	err = u.UsersModel.Create(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "somethign went wrong when creating a user" + err.Error(),
		},
		)
		return
	}
	c.JSON(200, gin.H{
		"success:": "user created",
	})
}
