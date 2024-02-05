package routes

import (
	"net/http"
	. "testtry2/database"

	"github.com/gin-gonic/gin"
)

func GetUsersList(ctx *gin.Context) {
	usersList := List_Users()
	ctx.JSON(http.StatusOK, gin.H{"message": usersList})
}
