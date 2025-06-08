package main

import (
	"github.com/ja-howell/stashclone/server"
)

func main() {
	s := server.New()
	s.Start()
}
