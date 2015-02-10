// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// Main package
//
package main

import (
    "fmt"
    "flag"
    "github.com/z0rr0/go.t34.me/utils"
)

const (
    Port uint = 8080
    ProgramName string = "go.t34.me"
    ProgramVersion string = "v0.1"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            utils.LoggerError.Println(r)
            fmt.Printf("Program \"%v\" %v is terminated abnormally.\n", ProgramName, ProgramVersion)
        }
    }()
    fmt.Printf("Program \"%v\" %v is starting...\n", ProgramName, ProgramVersion)
    port := flag.Uint("port", Port, "port number")
    debug := flag.Bool("debug", false, "debug mode")
    flag.Parse()
    utils.LoggerInit(*debug)

    utils.LoggerDebug.Printf("port=%v, debug=%v", *port, *debug)
    fmt.Printf("Program \"%v\" %v is successfully terminated.\n", ProgramName, ProgramVersion)
}
