package main

import ( "fmt"
          "os" )

func main() {
  if len(os.Args) < 2 {
     fmt.Println("No ARGs given.")
     os.Exit(1)
  }

  fmt.Println("First ARG", os.Args[1])
}
