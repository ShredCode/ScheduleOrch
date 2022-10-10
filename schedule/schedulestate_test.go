package schedule_test

import "testing"
import ss "scheduleorch/schedule"

func TestSchedule_Next(t *testing.T) {
    expected := "Running"
    actual := ss.ScheduleState.Running.String()
    if expected != actual {
        t.Errorf("Expected %s do not match actual %s", expected, actual)
    }
}

func TestSchedule_Unknown(t *testing.T) {
    expected := "Unknown"
    actual := ss.ScheduleState.Unknown.String()
    if expected != actual {
        t.Errorf("Expected %s do not match actual %s", expected, actual)
    }
}
