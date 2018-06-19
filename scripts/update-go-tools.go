package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	verbose = flag.Bool("v", false, "enable verbose mode")
)

func main() {
	log.SetPrefix("update-go-tools: ")
	log.SetFlags(0)

	flag.Parse()

	for _, pkg := range []string{
		"github.com/mdempsky/gocode",
		"golang.org/x/tools/cmd/benchcmp",
		"golang.org/x/tools/cmd/eg",
		"golang.org/x/tools/cmd/fiximports",
		"golang.org/x/tools/cmd/goimports",
		"golang.org/x/tools/cmd/gomvpkg",
		"golang.org/x/tools/cmd/gorename",
		"golang.org/x/tools/cmd/guru",
		"golang.org/x/tools/cmd/heapview",
		"golang.org/x/tools/cmd/present",
		"golang.org/x/tools/cmd/stringer",
		"golang.org/x/vgo",
	} {
		run("go", "get", "-u", "-v", pkg)
	}
	run("gocode", "exit")
}

func run(args ...string) {
	log.Printf("running %s...", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	if *verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
