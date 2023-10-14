package controllers

import (
	"ginrest/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LanguageHandler struct {
	Db *gorm.DB
}

func (h *LanguageHandler) GetAll(c *gin.Context) {
	var languages []models.Language
	h.Db.Find(&languages)
	c.JSON(200, languages)
}
