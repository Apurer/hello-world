package hello
	
import (
    "fmt"
    "github.com/Apurer/hello-world/morestrings"
    "github.com/Apurer/hello-world/arithmetics"
)

func Hello() string {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(arithmetics.Sum(4,2))
    return "hello world"
}

// func main() {
//     hello := Hello()
//     fmt.Println(hello)
//     fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
//     fmt.Println(arithmetics.Sum(4,2))
// }