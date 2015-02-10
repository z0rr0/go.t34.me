// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// Additional utils
//
package utils

import (
    "os"
    "log"
    "io/ioutil"
)

var (
    LoggerError *log.Logger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    LoggerDebug *log.Logger = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
)

// Initialization of Logger handlers
func LoggerInit(debugmode bool) {
    debugHandle := ioutil.Discard
    if debugmode {
        debugHandle = os.Stdout
    }
    LoggerDebug = log.New(debugHandle, "DEBUG: ",
        log.Ldate|log.Lmicroseconds|log.Lshortfile)
}
