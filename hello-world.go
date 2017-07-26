package main
import "fmt"
import "time"

func main() {
    fmt.Println("hey, don't panic!")
    fmt.Println(time.Now().Format(time.RFC850))
}
