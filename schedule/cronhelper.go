import "github.com/robfig/cron/v3"
import "time"


//Next will return the next runtime of a cron string 
func Next(s string) time.Time {
  c := cron.New()
  c.AddFunc(s, func() { fmt.Println("Background task still runninh") })

  return c.ParseStandard(s).Nex()

}
