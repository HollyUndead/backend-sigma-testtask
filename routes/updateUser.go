package routes

import (
	"log"
	"net/http"
	. "testtry2/database"
	"testtry2/helper"
	. "testtry2/models"

	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx *gin.Context) {
	var request User
	if err := ctx.BindJSON(&request); err != nil {
		log.Fatal(err)
		return
	}
	validateErr := helper.ValidateList(request)
	if validateErr != nil {
		ctx.JSON(http.StatusResetContent, gin.H{"message": validateErr})
		return
	}
	result, status := Update_User(request)
	ctx.JSON(status, gin.H{"message": result})
}
