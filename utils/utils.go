// Copyright (c) 2015, Alexander Zaytsev. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// Additional utils
//
package utils

import (
    "os"
    "fmt"
    "log"
    "bytes"
    "strings"
    "net/smtp"
    "io/ioutil"
    "html/template"
    "encoding/json"
    "path/filepath"
)

var (
    LoggerError *log.Logger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    LoggerDebug *log.Logger = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
)

type Config struct {
    DbDatabase string     `json:"database"`
    DbUser string         `json:"dbuser"`
    DbPassword string     `json:"dbpassword"`
    DbPort uint           `json:"dbport"`
    EmailUser string      `json:"email_user"`
    EmailPassword string  `json:"email_password"`
    EmailHost string      `json:"email_host"`
    EmailAddr string      `json:"email_addr"`
    EmailAdmin string     `json:"email_admin"`
    Templates string      `json:"templates"`
    Static string         `json:"static"`
}

type SimpleEmail struct {
    Name string
    Msg string
}

// Initialization of Logger handlers
func LoggerInit(debugmode bool) {
    debugHandle := ioutil.Discard
    if debugmode {
        debugHandle = os.Stdout
    }
    LoggerDebug = log.New(debugHandle, "DEBUG: ",
        log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

// It Validates file name, converts it from relative to absolute.
func FilePath(name string) (string, error) {
    var (
        fullpath string
        err error
    )
    fullpath = strings.Trim(name, " ")
    if len(fullpath) < 1 {
        return fullpath, fmt.Errorf("Empty file name")
    }
    if name[0] == '/' {
        return fullpath, nil
    }
    fullpath, err = filepath.Abs(fullpath)
    if err != nil {
        return fullpath, err
    }
    return fullpath, nil
}

func checkFilePaths(paths ...*string) {
    for _, path := range paths {
        if fullpath, err := FilePath(*path); err != nil {
            LoggerError.Panicf("Can't prepare filename: %v / %v", *path, err)
        } else {
            *path = fullpath
        }
        if _, err := os.Stat(*path); err != nil {
            LoggerError.Panicf("File/dir \"%v\" not found: %v", *path, err)
        }
    }
}

// Parse config file from JSON format.
func GetConfig(name *string) Config {
    var (
        cfg Config
        jsondata []byte
    )
    checkFilePaths(name)
    jsondata, err := ioutil.ReadFile(*name)
    if err != nil {
        LoggerError.Panicf("File reading error: %v", err)
    }
    if err := json.Unmarshal(jsondata, &cfg); err != nil {
        LoggerError.Panicf("Can't parse config file: %v", err)
    }
    checkFilePaths(&cfg.Templates, &cfg.Static)
    return cfg
}

// Send HTML email
func SendEmail(cfg *Config, to []string, subject, msg string) error {
    const mime string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
    auth := smtp.PlainAuth(
        "",
        cfg.EmailUser,
        cfg.EmailPassword,
        cfg.EmailHost,
    )
    subject = fmt.Sprintf("Subject: %v\n", subject)
    content := []byte(subject + mime + msg)
    return smtp.SendMail(cfg.EmailAddr, auth, cfg.EmailUser, to, content)
}

func EmailToAdmin(cfg *Config, subject, msg string) error {
    var htmltpl bytes.Buffer
    to := []string{cfg.EmailAdmin}
    data := SimpleEmail{"Administrator", msg}
    template_file := filepath.Join(cfg.Templates, "simple_email.html")
    template, err := template.ParseFiles(template_file)
    if err != nil {
        LoggerError.Printf("HTML parser error: %v\n", err)
        return err
    }
    err = template.Execute(&htmltpl, &data)
    if err != nil {
        LoggerError.Println(err)
        return err
    }
    return SendEmail(cfg, to, subject, htmltpl.String())
}
