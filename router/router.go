package router

import (
	"awesomeProject/api"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	reminderGroup := r.Group("/reminder")
	{
		reminderGroup.POST("", api.AddReminderHandler)
		reminderGroup.GET("", api.GetRemindersHandler)
		reminderGroup.DELETE("/:id", api.DeleteReminderHandler)
		reminderGroup.PUT("/:id", api.UpdateReminderHandler)
	}
	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		api.ServeWebSocket(c.Writer, c.Request)
	})
	return r
}
