package main
	
import "fmt"

func Hello() string {
    return "hello world"
}

func main() {
    hello := Hello()
    fmt.Println(hello)
}