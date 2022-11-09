package timer

import (
	"Area/database"
	"Area/database/models"
	"Area/lib"
	"strconv"
	"time"
)

func checkTimeInterval(currentTime time.Time, storedData models.TriggerData, trigger *models.Trigger) bool {
	interval, err := strconv.ParseInt(storedData.ActionData, 10, 64)
	lib.CheckError(err)

	if storedData.Timestamp.IsZero() {
		storedData.Timestamp = time.Now()
		trigger.Data = lib.EncodeToBytes(storedData)
		database.Trigger.Update(trigger)
	}

	if currentTime.After(storedData.Timestamp.Add(time.Minute * time.Duration(interval))) {
		storedData.Title = "reminder!"
		storedData.Description = storedData.ActionData + " minutes has passed!"
		storedData.Timestamp = time.Now()
		trigger.Data = lib.EncodeToBytes(storedData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}

func checkEveryDayTime(currentTime time.Time, storedData models.TriggerData, trigger *models.Trigger) bool {
	reminder, err := time.Parse("15:04", storedData.ActionData)
	lib.CheckError(err)

	if reminder.Hour() == currentTime.Hour() && (reminder.Minute() == currentTime.Minute()) {
		storedData.Title = "Daily reminder!"
		storedData.Description = "It's time!"
		storedData.Timestamp = time.Now()
		trigger.Data = lib.EncodeToBytes(storedData)
		database.Trigger.Update(trigger)
		return true
	}

	return false
}

func checkSingleTime(currentTime time.Time, storedData models.TriggerData, trigger *models.Trigger) bool {
	reminder, err := time.Parse("2006-01-02 15:04", storedData.ActionData)
	lib.CheckError(err)

	if reminder.Before(currentTime.Add(time.Hour)) && storedData.Timestamp.IsZero() {
		storedData.Title = "Single reminder!"
		storedData.Description = "It's time!"
		storedData.Timestamp = time.Now()
		trigger.Data = lib.EncodeToBytes(storedData)
		database.Trigger.Update(trigger)
		return true
	}
	return false
}
