package term

import (
    "log"
    "os"
    "golang.org/x/term"
)


func GetTermWidth() int {
    w, _, err := term.GetSize(int(os.Stdin.Fd()))
    if err != nil {
        log.Fatal(err)
    }
    return w
}


