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

func Index(c *gin.Context) {
    obj := gin.H{}
    c.HTML(200, "index.html", obj)
}
func About(c *gin.Context) {
    obj := gin.H{}
    c.HTML(200, "about.html", obj)
}
func GetData(c *gin.Context) {
    c.Request.ParseForm()
    obj := gin.H{"Short": c.Request.PostForm.Get("longurl")}
    c.HTML(200, "result.html", obj)
}
