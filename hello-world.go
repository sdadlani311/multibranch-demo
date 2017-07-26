package main
import "fmt"
import "time"

func main() {
    fmt.Println("don't panic!")
    fmt.Println(time.Now().Format(time.RFC850))
}
