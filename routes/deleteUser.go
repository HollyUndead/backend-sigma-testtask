package routes

import (
	"fmt"
	"log"
	"net/http"
	. "testtry2/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(ctx *gin.Context) {
	request := ctx.Param("Id")
	fmt.Println(request)
	objId, err := primitive.ObjectIDFromHex(request)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}
	result, err := Delete_User(objId)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}
	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User does not exist"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": request})
}
