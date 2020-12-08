package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var (
		isPlain = flag.Bool("plain", false, "enable text/plain")
	)

	flag.Parse()

	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, os.Stdin)
	if err != nil {
		panic(fmt.Errorf("could not copy stdin: %+v", err))
	}

	if *isPlain {
		plaintext(buf)
		return
	}

	switch {
	case bytes.HasPrefix(buf.Bytes(), []byte(`<!DOCTYPE`)),
		bytes.HasPrefix(buf.Bytes(), []byte(`<html`)),
		strings.HasPrefix(http.DetectContentType(buf.Bytes()), "text/html"):
		htmltext(buf)
	default:
		plaintext(buf)
	}
}

func plaintext(stdin io.Reader) {
	cmd := exec.Command("awk", "-f",
		filepath.Join(os.Getenv("HOME"), ".config/aerc/filters/plaintext"),
	)
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Errorf("could not run plain-text filter: %+v", err))
	}
}

func htmltext(stdin io.Reader) {
	cmd := exec.Command("/bin/sh",
		filepath.Join(os.Getenv("HOME"), ".config/aerc/filters/html"),
	)
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)

	err := cmd.Run()
	if err != nil {
		panic(fmt.Errorf(
			"could not run html-ifier: %+v",
			err,
		))
	}

}
