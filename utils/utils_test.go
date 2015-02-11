// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
package utils

import (
    "os"
    "testing"
    "path/filepath"
)

func TestLoggerInit(t *testing.T) {
    if (LoggerError == nil) || (LoggerDebug == nil) {
        t.Errorf("Incorrect references")
    }
    LoggerInit(false)
    if (LoggerError.Prefix() != "ERROR: ") || (LoggerDebug.Prefix() != "DEBUG: ") {
        t.Errorf("Incorrect loggers settings")
    }
    LoggerInit(true)
    if (LoggerError.Flags() != 19) || (LoggerDebug.Flags() != 21) {
        t.Errorf("Incorrect loggers settings")
    }
}

func TestFilePath(t *testing.T) {
    var name string
    if _, err := FilePath(name); err == nil {
        t.Errorf("File name should not be checked")
    }
    name = "name"
    if val, err := FilePath("name"); err != nil {
        t.Errorf("Invalid response")
    } else {
        absname, abserr := filepath.Abs("name")
        if (val != absname) || (abserr != nil) {
            t.Errorf("Invalid response")
        }
    }
    pwd, _ := os.Getwd()
    name = filepath.Join(pwd, name)
    if val, err := FilePath(name); err != nil {
        t.Errorf("Invalid response")
    } else {
        if val != name {
            t.Errorf("Invalid response")
        }
    }
}

func TestGetConfig(t *testing.T) {
    // correct config file should exist in $GOPATH
    var name string = "config.json"
    pwd := os.Getenv("GOPATH")
    name = filepath.Join(pwd, name)
    cfg := GetConfig(&name)
    if cfg.DbDatabase == "" {
        t.Errorf("Invalid config data")
    }
}