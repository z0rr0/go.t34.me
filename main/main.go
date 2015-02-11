// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// Main package
//
package main

import (
    "fmt"
    "flag"
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/z0rr0/go.t34.me/utils"
    "github.com/z0rr0/go.t34.me/handler"
)

const (
    Port uint = 8080
    Name string = "go.t34.me"
    Config string = "config.json"
)
var (
    Version string = "v0.1 git:000000 2015-01-01"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            utils.LoggerError.Println(r)
            fmt.Printf("Program \"%v\" %v is terminated abnormally.\n", Name, Version)
        }
    }()
    port := flag.Uint("port", Port, "port number")
    debug := flag.Bool("debug", false, "debug mode")
    config := flag.String("config", Config, "configuration file")
    version := flag.Bool("version", false, "version info")
    flag.Parse()
    if (*version) {
        fmt.Printf("%v version: %v\n", Name, Version)
        return
    }
    fmt.Printf("Program \"%v\" %v is starting...\n", Name, Version)

    utils.LoggerInit(*debug)
    cfg := utils.GetConfig(config)
    utils.LoggerDebug.Printf("port=%v, database=%v, debug=%v", *port, cfg.DbDatabase, *debug)

    router := gin.Default()
    if *debug {
        gin.SetMode(gin.DebugMode)
    } else {
        gin.SetMode(gin.ReleaseMode)
    }
    addr := fmt.Sprintf("localhost:%v", *port)
    server := &http.Server{
        Addr:           addr,
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    utils.LoggerDebug.Printf("Listen %v", addr)

    router.NoRoute(handler.NotFound)
    router.GET("/test", handler.Test)
    if err := server.ListenAndServe(); err != nil {
        utils.LoggerError.Panicf("Error: %v", err)
    }

    fmt.Printf("Program \"%v\" %v is successfully terminated.\n", Name, Version)
}
