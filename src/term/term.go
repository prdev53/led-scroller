package term

import (
    "fmt"
    "os"
    "golang.org/x/term"
)


func GetTermWidth() int {
    w, _, err := term.GetSize(int(os.Stdin.Fd()))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return w
}


