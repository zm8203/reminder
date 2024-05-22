package model

import (
	"testing"
	"time"
)

func TestReminderCRUD(t *testing.T) {
	reminder := Reminder{
		CreatorID: "user1",
		Content:   "Test Reminder",
		DateTime:  time.Now(),
	}

	AddReminder(&reminder)

	reminders := GetRemindersByUserID("user1")
	if len(reminders) != 1 {
		t.Errorf("Expected 1 reminder, got %d", len(reminders))
	}

	err := DeleteReminder(reminder.ID)
	if err != nil {
		t.Errorf("Error deleting reminder: %v", err)
	}

	reminders = GetRemindersByUserID("user1")
	if len(reminders) != 0 {
		t.Errorf("Expected 0 reminder after deletion, got %d", len(reminders))
	}
}
