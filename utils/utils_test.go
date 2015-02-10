// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
package utils

import (
    "testing"
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
