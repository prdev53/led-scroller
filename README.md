# LED Scroller

LED Scroller is a command-line tool for simulating an LED scroller display in your terminal.

## Features

* Display a help menu with the `led help` command
* Start the scrolling and listen for changes with the `led serve` command
* Send text to the LED scroller server with the `led send` command
* Start the scrolling without listening for changes with the `led show` command

## Building the project

### Unix

**Note:** This requires having perl installed. Otherwise look at the next section.

From the project's root directory, run:
```console
$ perl make.pl
```

This will generate the `led` executable in the project's root directory.

### All systems

```console
$ cd src
$ go build
```

This will generate the `led-scroller` executable in the `src` folder.
You are free to rename this file to `led` at your convenience.
