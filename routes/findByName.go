package routes

import (
	"net/http"
	. "testtry2/database"

	"github.com/gin-gonic/gin"
)

func FindUserByFirstName(ctx *gin.Context) {
	request := ctx.Param("FirstName")
	result := Find_User_ByFirstName(request)
	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
