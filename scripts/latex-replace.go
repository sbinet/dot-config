package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("latex-replace: ")
	flag.Parse()

	if flag.NArg() <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	fname := flag.Arg(0)
	log.Printf("processing [%s]...\n", fname)
	replace(fname)
}

var (
	replacer = strings.NewReplacer(
		`\'e`, "é",
		`\`+"`e", "è",
		`\`+"`a", "à",
		`\`+"`u", "ù",
		`\^a`, "â",
		`\^e`, "ê",
		`\^i`, "î",
		`\^o`, "ô",
		`\^u`, "û",
		`\"e`, "ë",
		`\"i`, "ï",
		`\"u`, "ü",
		`\"y`, "ÿ",
		`\c c`, "ç",
		`\c C`, "Ç",
		`\c{c}`, "ç",
		`\c{C}`, "Ç",
		`\oe `, "œ",
	)
)

func replace(fname string) {
	raw, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fname + ".tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = replacer.WriteString(f, string(raw))
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename(fname+".tmp", fname)
	if err != nil {
		log.Fatal(err)
	}

	return
}
