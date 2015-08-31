// dot-config initializes all symlinks under
//  $HOME to point to
//  $HOME/.config/path/to/file
package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

type Symlink struct {
	Old string
	New string
}

var links = []Symlink{
	{
		New: ".bash_profile",
		Old: "bash/profile",
	},
	{
		New: ".bashrc",
		Old: "bash/bashrc",
	},
	{
		New: ".gitconfig",
		Old: "git/gitconfig",
	},
	{
		New: ".gitcookies",
		Old: "git/gitcookies",
	},
	{
		New: ".gitignore_global",
		Old: "git/gitignore_global",
	},
	{
		New: ".inputrc",
		Old: "inputrc",
	},
	{
		New: ".ssh",
		Old: "ssh",
	},
	{
		New: ".vim",
		Old: "vim/vim",
	},
	{
		New: ".vimrc",
		Old: "vim/vimrc",
	},
	{
		New: ".xinitrc",
		Old: "xresources/initrc",
	},
	{
		New: ".xscreensaver",
		Old: "xresources/xscreensaver",
	},
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("dot-config: could not retrieve current user informations: %v\n", err)
	}
	dir, err := filepath.Abs(usr.HomeDir)
	if err != nil {
		log.Fatalf(
			"dot-config: could not resolve absolute path to $HOME: %v\n",
			err,
		)
	}

	for _, link := range links {
		newname := filepath.Join(dir, link.New)
		oldname := filepath.Join(dir, ".config", link.Old)
		fi, err := os.Lstat(newname)
		switch {
		case os.IsExist(err):
			if fi.Name() == oldname {
				// ok.
				continue
			}
			err = os.Rename(
				newname,
				newname+".bak."+time.Now().Format("20060102030405"),
			)
			if err != nil {
				log.Fatalf(
					"dot-config: could not backup [%s]: %v\n",
					newname,
					err,
				)
			}

		case os.IsNotExist(err):
			// ok. noop.

		default:
			err = os.Remove(newname)
			if err != nil {
				log.Fatalf(
					"dot-config: could not remove symlink [%s]: %v\n",
					newname,
					err,
				)
			}
		}

		_, err = os.Lstat(oldname)
		if os.IsNotExist(err) {
			log.Fatalf(
				"dot-config: file to symlink to [%s] does NOT exist: %v\n",
				oldname,
				err,
			)
		}

		log.Printf(
			"dot-config: symlinking [%s] -> [%s]...\n",
			newname,
			oldname,
		)
		err = os.Symlink(oldname, newname)
		if err != nil {
			log.Fatalf(
				"dot-config: error symlinking [%s] to [%s]: %v\n",
				newname,
				oldname,
			)
		}
	}
}
