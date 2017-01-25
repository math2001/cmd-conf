package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "bytes"
    "github.com/fatih/color"
    "unicode/utf8"
)

func real_length(str string) int {
    return utf8.RuneCountInString(str)
}

func get_branch_name() string {
    // from http://stackoverflow.com/a/12142066/6164984
    cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
    var outb, errb bytes.Buffer
    cmd.Stdout = &outb
    cmd.Stderr = &errb
    err := cmd.Run()
    if err != nil {
        return ""
    }
    return strings.Trim(outb.String(), "\n ")
}

func is_clean_git_branch() bool {
    cmd := exec.Command("git", "status", "--short")
    var outb bytes.Buffer
    cmd.Stdout = &outb
    err := cmd.Run()
    if err != nil {
        return false
    }
    return outb.String() == ""
}

func main() {

    defer unset_color()

    // color.Magenta("And many others ..")
    // color.New(color.FgMagenta).Add(color.Underline).Println("hello world")

    wd, _ := os.Getwd()
    wd = strings.Replace(wd, os.ExpandEnv("$USERPROFILE"), "~", 1)
    wd = strings.Replace(wd, "\\", "/", -1)
    color.New(color.FgMagenta).Add(color.Underline).Print(wd)

    var branch string = get_branch_name()
    if branch != "" {
        fmt.Print(" @ ")
        if is_clean_git_branch() {
            color.New(color.FgGreen).Add(color.Underline).Print(branch)
        } else {
            color.New(color.FgRed).Add(color.Underline).Print(branch)
        }
    }

    fmt.Print("\n$ ")
}

func unset_color() {
    color.Unset()
}
