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
	log.SetPrefix("update-go: ")
	log.SetFlags(0)

	flag.Parse()
	branch := "master"
	if flag.NArg() > 0 {
		branch = flag.Arg(0)
	}

	godir := "/home/binet/sdk/go/src"
	run(godir, "git", "fetch", "--all", "-p")
	run(godir, "git", "checkout", branch)
	run(godir, "git", "pull")
	run(godir, "./make.bash")
	run("/home/binet", "update-go-tools")
}

func run(dir string, args ...string) {
	log.Printf("running %s...", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	if *verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	if dir != "" {
		cmd.Dir = dir
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
