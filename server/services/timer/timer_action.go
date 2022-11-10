package timer

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"strconv"
	"strings"
	"time"
)

func checkTimeInterval(currentTime time.Time, storedData models.TriggerData, trigger *models.Trigger) bool {
	interval, err := strconv.ParseInt(storedData.ActionData, 10, 64)
	if err != nil {
		lib.LogError(err)
		return false
	}

	if currentTime.After(storedData.Timestamp.Add(time.Minute * time.Duration(interval))) {
		storedData.Title = "reminder!"
		storedData.Description = storedData.ActionData + " minutes has passed!"
		storedData.Timestamp = time.Now().UTC()
		trigger.Data = lib.EncodeToBytes(storedData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func checkEveryDayTime(currentTime time.Time, storedData models.TriggerData, trigger *models.Trigger) bool {
	reminder, err := time.Parse("15:04", storedData.ActionData)
	if err != nil {
		lib.LogError(err)
		return false
	}

	reminder = reminder.Add(-time.Hour)

	if reminder.Hour() == currentTime.Hour() && (reminder.Minute() == currentTime.Minute()) {
		storedData.Title = "Daily reminder!"
		storedData.Description = "It's time!"
		storedData.Timestamp = time.Now().UTC()
		trigger.Data = lib.EncodeToBytes(storedData)
		database.Trigger.Update(trigger)
		return true
	}

	return false
}

func checkSingleTime(currentTime time.Time, storedData models.TriggerData, trigger *models.Trigger) bool {
	stringTime := strings.Replace(storedData.ActionData, "T", " ", 1)
	reminder, err := time.Parse("2006-01-02 15:04", stringTime)
	if err != nil {
		lib.LogError(err)
		return false
	}

	reminder = reminder.Add(-time.Hour)

	if reminder.Equal(currentTime.Truncate(60 * time.Second)) {
		storedData.Title = "Single reminder!"
		storedData.Description = "It's time!"
		storedData.Timestamp = time.Now().UTC()
		trigger.Data = lib.EncodeToBytes(storedData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}
