package main

import (
    "fmt"
    "time"
    "net/http"
    "net/url"
    "os"
    d "led-scroller/dictionnary"
)


var (
    dict d.Dictionnary
    message d.Text
    column int
)


func initialize() {
    dict = d.Build()
    message = dict.StringToText("ready")
    column = 0
}

func runDisplayLoop() {
    for {
        fmt.Print("\033[H\033[2J")
        fmt.Print("\n\n\n")
        message.Print(&column)
        time.Sleep(45 * time.Millisecond)
    }
}

func postMessage(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        input := r.FormValue("text")
        message = dict.StringToText(input)
        column = 0
    }
}

func sendText() {
    values := url.Values{}
    values.Add("text", os.Args[2])
    http.PostForm("http://localhost:3232/", values)
}

func main() {
    if len(os.Args) == 3 && os.Args[1] == "test" {
        initialize()
        input := os.Args[2]
        message = dict.StringToText(input)
        column = 0
        runDisplayLoop()
    } else if len(os.Args) == 2 && os.Args[1] == "help" {
        fmt.Println("LED Scroller v1.0")
        fmt.Println("led help - Brings up this menu")
        fmt.Println("led serve - Starts the scrolling and listens for changes")
        fmt.Println("led send <text> - Sends the text to the LED scroller server")
    } else if len(os.Args) == 2 && os.Args[1] == "serve" {
        initialize()

        go func() {
            http.HandleFunc("/", postMessage)
            err := http.ListenAndServe(":3232", nil)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
        }()

        runDisplayLoop()
    } else if len(os.Args) == 3 && os.Args[1] == "send" {
        sendText()
    }
}
