package main

import (
    "runtime"
    "regexp"
    "fmt"
)

func check(err error, message string) {
    if err != nil {
        panic(message + ": " + err.Error())
    }
}

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func escapeMarkdown(text string, version int) (string, error) {
    var escape_chars string
    if version == 1 {
        escape_chars = "_*`["
    } else if version == 2 {
        escape_chars = "\\_*[]()~`>#+-=|{}.!"
    }

    re, err := regexp.Compile("([" + regexp.QuoteMeta(escape_chars) + "])")
    if err != nil {
        return text, fmt.Errorf("Failed to compile regex: %w", err)
    }

    return re.ReplaceAllString(text, "\\${1}"), nil
}
