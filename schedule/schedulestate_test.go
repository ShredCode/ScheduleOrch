package schedule_test

import "testing"
import "scheduleorch/schedule"

func TestSchedule_Next(t *testing.T) {
    expected := "Running"
    actual := schedule.ScheduleState.Running.String()
    if expected != actual {
        t.Errorf("Expected %s do not match actual %s", expected, actual)
    }
}

func TestSchedule_Unknown(t *testing.T) {
    expected := "Unknown"
    actual := schedule.ScheduleState.Unknown.String()
    if expected != actual {
        t.Errorf("Expected %s do not match actual %s", expected, actual)
    }
}
