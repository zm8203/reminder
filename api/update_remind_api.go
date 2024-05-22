package api

import (
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// UpdateReminderHandler 更新提醒的处理器
func UpdateReminderHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedReminder model.Reminder
	if err := c.BindJSON(&updatedReminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(updatedReminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := model.UpdateReminder(id, &updatedReminder); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "日历修改成功"})
}
