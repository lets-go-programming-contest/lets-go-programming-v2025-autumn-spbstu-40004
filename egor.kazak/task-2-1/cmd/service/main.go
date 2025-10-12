go
package main

import (
  "fmt"
  "log"
)

func main() {
  var count, constraints int
  _, err := fmt.Scan(&count, &constraints)
  if err != nil {
    log.Fatal(err)
  }

  for i := 0; i < count; i++ {
    const minTemp = 15
    currentMin := minTemp
    const maxTemp = 30
    currentMax := maxTemp

    for j := 0; j < constraints; j++ {
      var operator string
      var temperature int
      _, err := fmt.Scan(&operator, &temperature)
      if err != nil {
        log.Fatal(err)
      }

      switch operator {
      case ">=":
        if temperature > currentMin {
          currentMin = temperature
        }
      case "<=":
        if temperature < currentMax {
          currentMax = temperature
        }
      }

      if currentMin <= currentMax {
        fmt.Println(currentMin)
      } else {
        fmt.Println(-1)
      }
    }
  }
}