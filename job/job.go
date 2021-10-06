package job

type Job interface {
    Name string
    Id string
    RunningStatus bool
}
