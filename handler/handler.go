// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// HTTP handlers
//
package handler

import (
    // "fmt"
    // "flag"
    // "time"
    // "net/http"
    "github.com/gin-gonic/gin"
    // "github.com/z0rr0/go.t34.me/utils"
)

func Test(c *gin.Context) {
    c.String(200, "ok")
}
func NotFound(c *gin.Context) {
    c.String(404, "custom NotFound")
}
