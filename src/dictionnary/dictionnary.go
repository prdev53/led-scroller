package dictionnary

import (
    "led-scroller/term"
    "strings"
    "fmt"
)


type (
    Text [6]string
    Dictionnary map[string]Text
)


func Build() Dictionnary {
    dict := make(Dictionnary)
    dict["A"] = [6]string{
        " xxx ",
        "x   x",
        "xxxxx",
        "x   x",
        "x   x",
        "     " }
    dict["B"] = [6]string{
        "xxxx ",
        "x   x",
        "xxxxx",
        "x   x",
        "xxxx ",
        "     " }
    dict["C"] = [6]string{
        " xxxx",
        "x    ",
        "x    ",
        "x    ",
        " xxxx",
        "     " }
    dict["D"] = [6]string{
        "xxxx ",
        "x   x",
        "x   x",
        "x   x",
        "xxxx ",
        "     " }
    dict["E"] = [6]string{
        "xxxxx",
        "x    ",
        "xxxx ",
        "x    ",
        "xxxxx",
        "     " }
    dict["F"] = [6]string{
        "xxxxx",
        "x    ",
        "xxx  ",
        "x    ",
        "x    ",
        "     " }
    dict["G"] = [6]string{
        " xxxx ",
        "x     ",
        "x  xx ",
        "x    x",
        "  xxx ",
        "      " }
    dict["H"] = [6]string{
        "x   x",
        "x   x",
        "xxxxx",
        "x   x",
        "x   x",
        "     " }
    dict["I"] = [6]string{
        "xxxxx",
        "  x  ",
        "  x  ",
        "  x  ",
        "xxxxx",
        "     " }
    dict["J"] = [6]string{
        "xxxxx",
        "   x ",
        "   x ",
        "   x ",
        "xxxx ",
        "     " }
    dict["K"] = [6]string{
        "x   x",
        "xx x ",
        "xxx  ",
        "x  x ",
        "x   x",
        "     " }
    dict["L"] = [6]string{
        "x    ",
        "x    ",
        "x    ",
        "x    ",
        "xxxxx",
        "     " }
    dict["M"] = [6]string{
        "xx   xx",
        "x x x x",
        "x  x  x",
        "x     x",
        "x     x",
        "       " }
    dict["N"] = [6]string{
        "x   x",
        "xx  x",
        "x x x",
        "x  xx",
        "x   x",
        "     " }
    dict["O"] = [6]string{
        " xxxx ",
        "x    x",
        "x    x",
        "x    x",
        " xxxx ",
        "      " }
    dict["P"] = [6]string{
        "xxxx ",
        "x   x",
        "xxxx ",
        "x    ",
        "x    ",
        "     " }
    dict["Q"] = [6]string{
        "xxxxx ",
        "x   x ",
        "x   x ",
        "x   x ",
        "xxxxx ",
        "    xx" }
    dict["R"] = [6]string{
        "xxxx ",
        "x   x",
        "xxxx ",
        "x   x",
        "x   x",
        "     " }
    dict["S"] = [6]string{
        " xxxx",
        "x    ",
        " xxx ",
        "    x",
        "xxxx ",
        "       " }
    dict["T"] = [6]string{
        "xxxxx",
        "  x  ",
        "  x  ",
        "  x  ",
        "  x  ",
        "     " }
    dict["U"] = [6]string{
        "x   x",
        "x   x",
        "x   x",
        "x   x",
        "xxxxx",
        "     " }
    dict["V"] = [6]string{
        "xxx   xxx",
        " x     x ",
        "  x   x  ",
        "   x x   ",
        "    x    ",
        "         " }
    dict["W"] = [6]string{
        "x         x",
        "x         x",
        "x    x    x",
        " x  xxx  x ",
        "  xx   xx  ",
        "           " }
    dict["X"] = [6]string{
        "x    x",
        " x  x ",
        "  xx  ",
        " x  x ",
        "x    x",
        "      " }
    dict["Y"] = [6]string{
        "x   x",
        "x   x",
        " xxx ",
        "  x  ",
        "  x  ",
        "     " }
    dict["Z"] = [6]string{
        "xxxxx",
        "   xx",
        "  x  ",
        "xx   ",
        "xxxxx",
        "     " }
    dict["1"] = [6]string{
        "  x  ",
        " xx  ",
        "  x  ",
        "  x  ",
        "xxxxx",
        "     " }
    dict["2"] = [6]string{
        " xxx ",
        "x   x",
        "  xx ",
        " x   ",
        "xxxxx",
        "     " }
    dict["3"] = [6]string{
        "xxxx ",
        "    x",
        " xxx ",
        "    x",
        "xxxx ",
        "     " }
    dict["4"] = [6]string{
        "x   x",
        "x   x",
        "xxxxx",
        "    x",
        "    x",
        "     " }
    dict["5"] = [6]string{
        "xxxxx",
        "x    ",
        "xxxxx",
        "    x",
        "xxxxx",
        "     " }
    dict["6"] = [6]string{
        "xxxxx",
        "x    ",
        "xxxxx",
        "x   x",
        "xxxxx",
        "     " }
    dict["7"] = [6]string{
        "xxxxx",
        "    x",
        "    x",
        "    x",
        "    x",
        "     " }
    dict["8"] = [6]string{
        "xxxxx",
        "x   x",
        "xxxxx",
        "x   x",
        "xxxxx",
        "     " }
    dict["9"] = [6]string{
        "xxxxx",
        "x   x",
        "xxxxx",
        "    x",
        "xxxxx",
        "     " }
    dict["0"] = [6]string{
        "xxxxx",
        "x   x",
        "x   x",
        "x   x",
        "xxxxx",
        "     " }
    dict["#"] = [6]string{
        "   x  x ",
        " xxxxxxx",
        "  x  x  ",
        "xxxxxxx ",
        " x  x   ",
        "        " }
    dict[" "] = [6]string{
        "   ",
        "   ",
        "   ",
        "   ",
        "   ",
        "   " }
    dict["-"] = [6]string{
        "    ",
        "    ",
        "xxxx",
        "    ",
        "    ",
        "    " }
    dict[","] = [6]string{
        "   ",
        "   ",
        "   ",
        "   ",
        "xx ",
        " x " }
    dict["."] = [6]string{
        "  ",
        "  ",
        "  ",
        "  ",
        "x ",
        "  " }
    dict["!"] = [6]string{
        "  x ",
        "  x ",
        "  x ",
        "    ",
        "  x ",
        "    " }
    dict["?"] = [6]string{
        "   xxxx ",
        "  x    x",
        "     x  ",
        "        ",
        "     x  ",
        "        " }
    dict[":"] = [6]string{
        "  ",
        "x ",
        "  ",
        "x ",
        "  ",
        "  " }
    dict[")"] = [6]string{
        "xx  ",
        "  x ",
        "   x",
        "  x ",
        "xx  ",
        "    " }
    dict["("] = [6]string{
        "  xx",
        " x  ",
        "x   ",
        " x  ",
        "  xx",
        "    " }
    dict["'"] = [6]string{
        " x",
        " x",
        "  ",
        "  ",
        "  ",
        "  " }
    return dict
}


func (dict *Dictionnary) StringToText(m string) (text Text) {
    w := term.GetTermWidth()
    for i := range len(text) {
        // Spaces for the width of the terminal window
        text[i] = strings.Repeat(" ", w)
    }

    for _, c := range strings.ToUpper(m) {
        letter := (*dict)[string(c)]
        for i := range len(text) {
            text[i] = text[i] + letter[i] + "  "
        }
    }

    return
}


func (text *Text) Print(i *int) {
    w := term.GetTermWidth()
    for _, r := range *text {
        s := r[*i:]
        if len(s) > w {
            s = s[0:w]
        }

        // NOTE Post-processed to avoid incorrect length: DO NOT MOVE
        fmt.Println(strings.ReplaceAll(s, "x", "\033[46m \033[0m"))
    }

    if *i + 1 == len(text[0]) {
        *i = 0
    } else {
        *i = *i + 1
    }
    return
}

