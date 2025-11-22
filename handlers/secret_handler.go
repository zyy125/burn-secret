package handlers

import (
	"github.com/gin-gonic/gin"
	"burn-secret/models"
	"burn-secret/store"
	"burn-secret/utils"
	"time"
)


func CreateSecret(c *gin.Context){
	var req models.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"message": "request fomat error",
		})
		return
	}

	newID := utils.GetRandomID(models.IDLenth)
	newTime := time.Now()

	newSecret := models.Secret{
		ID: newID,
		Content: req.Content,
		MaxViews: req.MaxViews,
		ExpiryMinutes: req.ExpiryMinutes,
		CreatedAt: newTime,
		ViewsCount: 0,
	}

	if err := store.StoreSecret(&newSecret); err != nil {
        c.JSON(500, gin.H{"error": "Failed to save secret"})
        return
	}

	accessUrl := "http://localhost:8080/view/" + newID

	c.JSON(200, gin.H{
			"id": newID,              
			"accessUrl": accessUrl,
			"expireAt": newTime,
	})
}

func GetSecret(c *gin.Context) {
	id := c.Param("id")

	secret, err := store.GetSecret(id)
	if secret == nil && err == nil{
		c.JSON(404, gin.H{
			"code": 404,
			"message": "not found",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"message": err,
		})
		return
	}
	
	isBurned := secret.ViewsCount + 1 > secret.MaxViews 
	c.JSON(200, gin.H{
		"content": secret.Content,
		"remainingViews": secret.MaxViews - secret.ViewsCount,
		"isBurned": isBurned,
	})
}
