package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	lpcServer = "clrhp05.in2p3.fr"
)

var (
	a2ps   = flag.Bool("a2ps", false, "runs the a2ps text->PS conversion")
	server = flag.String("server", lpcServer, "remote print server")

	tmpdir string
)

func main() {
	log.SetPrefix("printer: ")

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		flag.PrintDefaults()
		log.Fatalf("missing input file name")
	}

	var err error
	tmpdir, err = ioutil.TempDir("", "prntr-")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(tmpdir)

	for _, fname := range flag.Args() {
		process(fname)
	}
}

func process(fname string) {
	log.Printf("printing [%s]...\n", fname)

	switch {
	case strings.HasPrefix(fname, "http"):
		req, err := http.Get(fname)
		if err != nil {
			log.Fatalf("error fetching file %q: %v", fname, err)
		}
		defer req.Body.Close()
		tmp, err := os.Create(filepath.Join(tmpdir, "file.tmp"))
		if err != nil {
			log.Fatal(err)
		}
		defer tmp.Close()
		defer os.Remove(tmp.Name())
		_, err = io.Copy(tmp, req.Body)
		if err != nil {
			log.Fatalf("error copying to temporary file: %v", err)
		}
		err = tmp.Close()
		if err != nil {
			log.Fatal(err)
		}

	default:
		f, err := os.Open(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		tmp, err := os.Create(filepath.Join(tmpdir, "file.tmp"))
		if err != nil {
			log.Fatalf("error creating temporary file: %v", err)
		}
		defer tmp.Close()

		_, err = io.Copy(tmp, f)
		if err != nil {
			log.Fatalf("error copying to temporary file: %v", err)
		}
		err = tmp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := exec.Command("scp", "file.tmp", *server+":.")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = tmpdir
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error copying file to remote print server: %v", err)
	}

	run("lpr",
		"-o", "KMDuplex=2Sided",
		"-o", "media=A4", "-o", "PageSize=A4",
		"-PBat6-1-nb1rv",
		"file.tmp",
	)
	run("/bin/rm", "file.tmp")
}

func run(cmd ...string) {
	args := []string{*server, "--"}
	args = append(args, cmd...)

	c := exec.Command("ssh", args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		log.Fatalf("error running remote command %q: %v\n", strings.Join(cmd, " "), err)
	}
}
