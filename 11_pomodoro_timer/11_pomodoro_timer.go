package main

import (
  "time"
  "fmt"
)

func startPomodoro() {
  const periodMinutes int = 25
  ticker := time.NewTicker(time.Second)
  go func() {
      endTime := time.Now().Add(time.Duration(periodMinutes) * time.Minute)
      for _ = range ticker.C {
          timeLeft := endTime.Sub(time.Now())
          if timeLeft > 0 {
            fmt.Printf("\r%02d:%02d", int32(timeLeft.Minutes()), int32(timeLeft.Seconds()) % 60)
          }
      }
  }()

  time.Sleep(time.Duration(periodMinutes) * time.Minute)
  ticker.Stop()
  fmt.Println("\nYour job is done!")
}

func main() {
  startPomodoro()
}
