package routes

import (
	"log"
	"net/http"
	"testtry2/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUserById(ctx *gin.Context) {
	request := ctx.Param("Id")
	objId, err := primitive.ObjectIDFromHex(request)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}
	result := database.Find_User_ById(objId)
	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
