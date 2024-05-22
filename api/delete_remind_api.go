package api

import (
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteReminderHandler 删除提醒的处理器
func DeleteReminderHandler(c *gin.Context) {
	id := c.Param("id")
	if err := model.DeleteReminder(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "日历删除成功"})
}
