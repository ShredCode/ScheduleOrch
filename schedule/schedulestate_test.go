packagr schedule_test

import "testing"

func TestSchedule_Next(t *testing.T) {
    expected := "Running"
    actual := ScheduleState.Running.String()
    if expected != actual {
        t.Errorf("Expected %s do not match actual %s", expected, actual)
    }
}
