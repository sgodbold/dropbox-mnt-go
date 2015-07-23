package main

import (
	"flag"
	"log"

	"github.com/sgodbold/dropbox-mnt/fs"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n dropboxfs MOUNTPOINT")
	}
	fs.MountFs(flag.Arg(0))
}
