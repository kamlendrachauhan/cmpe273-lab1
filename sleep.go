package main

import ("fmt";"time")

func main() {
    for i :=1;i<10;i++{
        sleep(i)
        fmt.Println("Printing after ",i," second.")
        }
}

func sleep(t int){
        timeDuration := time.Duration(t)
        returnedChannel := <- time.After(time.Second * timeDuration)
        fmt.Println("Time of return :", returnedChannel)
}
