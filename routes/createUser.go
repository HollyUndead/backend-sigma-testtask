package routes

import (
	"log"
	"net/http"
	. "testtry2/database"
	"testtry2/helper"
	. "testtry2/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var request User
	if err := ctx.BindJSON(&request); err != nil {
		log.Fatal(err)
		return
	}
	validateErr := helper.ValidateList(request)
	if validateErr != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"errors": validateErr})
		return
	}
	result := Create_User(request)
	if result == nil {
		ctx.JSON(http.StatusContinue, gin.H{"message": result})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
