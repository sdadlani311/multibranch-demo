package main
import "fmt"
import "time"

func main() {
    fmt.Println("hey now, don't panic!")
    fmt.Println(time.Now().Format(time.RFC850))
}
