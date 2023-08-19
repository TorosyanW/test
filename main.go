package main

import "fmt"

	
func main () {
  width := 15
  height := 10

  for index := 0; index < height; index++ {
    for j := 0; j < width; j++ {
       if index == 0 || index == height-1 || j == 0 || j == width-1 {
      fmt.Print("*")
      } else {
        fmt.Print (" ")
      }

  }

 fmt.Println()
}

}
