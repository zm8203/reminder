package api

import (
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// AddReminderHandler 添加提醒信息的处理器
func AddReminderHandler(c *gin.Context) {
	var reminder model.Reminder
	if err := c.BindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reminder.ID = model.GenerateID()
	model.AddReminder(&reminder)
	c.JSON(http.StatusCreated, reminder)
}
