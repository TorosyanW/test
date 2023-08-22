package main

import "fmt"

	
func clearScreen() {
  fmt.Print("\033[H\033[2J]]")

}

func main() {
  sizeY := 10
  sizeX := 20

  var arrayStr  [][]string = make([][]string, sizeY)

  clearScreen()

  for i := 0; i < sizeY; i++ {
    arrayStr[i] = make([]string , sizeX)
   for j := 0; j < sizeX; j++ {
    if i > sizeY-3 {
      arrayStr[i][j] = "日"
    } else {
      arrayStr[i][j] = "語"
    }
   }
     fmt.Println()

  }
   
   for _, valueY := range arrayStr{
    for _, valueX := range valueY {
       fmt.Print(valueX)

    }
        fmt.Println()
  }

}