
//Package schedule is meant to list the states a schedule can be in as well as the times jobs will run
package schedule

//ScheduleState will hold the 
type ScheduleState int64

const (
	Running ScheduleState = iota
	Delayed
	Stopped
	Init
 	Unknown
)

//Function String will return a readable string of the int code of the schedule's state
func (state ScheduleState) String() string {
	switch state {
	case Running:
		return "Running"
	case Delayed:
		return "Delayed"
	case Stopped:
		return "Stopped"
	case Init:
		return "Init"
 	case Unknown:
		return "Unknown"
	}
	return "Unknown"
}
