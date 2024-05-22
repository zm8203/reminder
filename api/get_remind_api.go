package api

import (
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRemindersHandler 获取提醒列表的处理器
func GetRemindersHandler(c *gin.Context) {
	userID := c.Query("user_id")
	var userReminders []model.Reminder
	if userID != "" {
		userReminders = model.GetRemindersByUserID(userID)
	} else {
		userReminders = model.GetAllReminders()
	}
	c.JSON(http.StatusOK, userReminders)
}
