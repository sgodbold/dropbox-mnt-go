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
	err := fs.LoadConfig()
	if err != nil {
		log.Fatalf("Config fail: %v\n", err)
	}
	fs.CacheInit()
	fs.MountFs(flag.Arg(0))
}
