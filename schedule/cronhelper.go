package schedule

import "github.com/robfig/cron/v3"
import "time"
import "fmt"


//Next will return the next runtime of a cron string 
func Next(s string) time.Time {
  c := cron.New()
  c.AddFunc(s, func() { fmt.Println("Background task still runninh") })

  return time.Now() //c.ParseStandard(s).Nex()

}
