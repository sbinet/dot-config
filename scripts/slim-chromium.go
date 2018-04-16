package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	log.SetPrefix("slim-chromium: ")
	log.SetFlags(0)

	doKill := flag.Bool("kill", false, "kill processes")
	flag.Parse()

	dir, err := os.Open("/proc")
	if err != nil {
		log.Fatalf("could not open /proc: %v", err)
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatalf("could not read dirnames from /proc: %v", err)
	}

	var pids []int
	for _, n := range names {
		pid, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			continue
		}
		cmdline, err := ioutil.ReadFile(filepath.Join(dir.Name(), n, "cmdline"))
		if err != nil {
			log.Fatalf("could not read /proc/%s/cmdline: %v", n, err)
		}
		if !bytes.HasPrefix(cmdline, []byte("/usr/lib/chromium/chromium ")) {
			continue
		}
		if bytes.Contains(cmdline, []byte("--type=zygote ")) {
			continue
		}
		if !bytes.Contains(cmdline, []byte("--type=renderer ")) {
			continue
		}
		pids = append(pids, int(pid))
	}

	log.Printf("pids: %d", len(pids))

	if *doKill {
		for _, pid := range pids {
			proc, err := os.FindProcess(pid)
			if err != nil {
				log.Fatal(err)
			}
			if proc == nil {
				log.Fatalf("nil ptr to proc %d", pid)
			}
			err = proc.Kill()
			if err != nil {
				log.Fatalf("could not kill process %d: %v", pid, err)
			}
		}
	}
}
