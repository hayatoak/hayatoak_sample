package main

import "fmt"

func main(){
  for i := 0; i <= 100; i++{ 
    if i % 15 == 0{
      fmt.Print("FizzBuzz\n")
    } else if i % 5 == 0{
      fmt.Print("Buzz\n")
    } else if i % 3 == 0{
      fmt.Print("Fizz\n")
    } else {
      fmt.Printf("%d\n",i)
    }
  }
}

