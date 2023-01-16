package main

import (
	"anyops/cli/cmd"
	"embed"
	"fmt"
	"github.com/manucorporat/try"
	"os"
	"os/signal"
	"syscall"
)

var (
	//go:embed all:compose
	files embed.FS
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	workingDir, _ := os.MkdirTemp(os.TempDir(), "anyops-*")
	go func() {
		<-c
		cmd.Finalize(workingDir)
		os.Exit(1)
	}()
	try.This(func() {
		fmt.Println("-----------------------------------------------")
		cmd.ExpandTemporarily(workingDir, files)
		cmd.Execute(workingDir)
		fmt.Println("-----------------------------------------------")
	}).Finally(func() {
		cmd.Finalize(workingDir)
	}).Catch(func(e try.E) {
		cmd.Finalize(workingDir)
	})
}
