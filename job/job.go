package job

type Job interface {
    GetName() string
    GetId() string
    RunningStatus() string
}
