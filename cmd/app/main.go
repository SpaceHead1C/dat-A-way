package main

import (
	"dataway/pkg/log"
	"os"
)

func main() {
	c := newConfig()
	if err := parse(os.Args[1:], c); err != nil {
		panic(err.Error())
	}
	l, err := log.NewLogger()
	if err != nil {
		panic(err.Error())
	}
	l.Info("dat(A)way service is up")
}
