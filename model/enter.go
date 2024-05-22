package model

import (
	"fmt"
	"time"
)

// Reminder 结构体表示提醒信息
type Reminder struct {
	ID        string    `json:"id"`
	CreatorID string    `json:"creator_id"`
	Content   string    `json:"content"`
	DateTime  time.Time `json:"date_time"`
}

var reminders []Reminder

// GetAllReminders 获取所有提醒列表
func GetAllReminders() []Reminder {
	return reminders
}

// GenerateID 生成唯一的提醒ID
func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// AddReminder 添加提醒信息到列表
func AddReminder(reminder *Reminder) {
	reminder.ID = GenerateID()
	reminders = append(reminders, *reminder)
}

// GetRemindersByUserID 根据用户ID获取提醒列表
func GetRemindersByUserID(userID string) []Reminder {
	var userReminders []Reminder
	for _, reminder := range reminders {
		if reminder.CreatorID == userID {
			userReminders = append(userReminders, reminder)
		}
	}
	return userReminders
}

// DeleteReminder 删除提醒
func DeleteReminder(id string) error {
	for i, reminder := range reminders {
		if reminder.ID == id {
			reminders = append(reminders[:i], reminders[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("日历提醒数据ID %s 不存在", id)
}

// UpdateReminder 更新提醒
func UpdateReminder(id string, updatedReminder *Reminder) error {
	for i, reminder := range reminders {
		if reminder.ID == id {
			reminders[i] = *updatedReminder
			return nil
		}
	}
	return fmt.Errorf("日历提醒数据ID %s 不存在", id)
}
