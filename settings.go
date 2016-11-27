package main

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
)

//
var Settings struct {
	Debugger bool
	Verbose  bool
	Mode     string

	Addr     string
	Port     int
	Insecure bool

	Server  string
	Powered string
	Version string

	Program struct {
		Name   string
		Folder string
		Full   string
		Data   string
		Config string
	}

	Docs *url.URL
}

//
const MAJOR = "0.1"

// DO NOT UPDATE. This is populated by the build. See the Makefile.
var VERSION = "0"

//
func init() {
	Settings.Program.Full = os.Args[0]
	Settings.Program.Name = filepath.Base(Settings.Program.Full)

	if dir, _ := filepath.Split(Settings.Program.Full); dir == "" {
		if path, err := exec.LookPath(Settings.Program.Full); err != nil {
			log.Println(err)
		} else {
			Settings.Program.Folder = filepath.Base(path)
		}
	} else if ln, err := filepath.EvalSymlinks(Settings.Program.Full); err != nil {
		log.Println(err)
	} else {
		if dir, err := filepath.Abs(filepath.Dir(ln)); err != nil {
			log.Println(err)
		} else {
			Settings.Program.Folder = dir
		}
	}

	data := os.Getenv("DATA")
	if data == "" {
		data = "data"
	}
	Settings.Program.Data = filepath.Join(Settings.Program.Folder, data)

	config := os.Getenv("CONFIG")
	if config == "" {
		Settings.Program.Config = filepath.Join(Settings.Program.Folder, "config/app.yaml")
	} else {
		Settings.Program.Config = config
	}
}
