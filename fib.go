package main

import "fmt"

func main(){
    var input int
    fmt.Print("Enter the index of the fibonacci number you would like to retrieve : ")
    fmt.Scanf("%d", &input)
    output := calculateFib(input)
    fmt.Println(output)
}

func calculateFib(num int) int{
    if num < 0 {
        return -1
        }
    if num == 0 {
        return 0
        }
    if num == 1 || num== 2{
        return 1    
    }
    return calculateFib(num -1)+ calculateFib(num -2)
}
