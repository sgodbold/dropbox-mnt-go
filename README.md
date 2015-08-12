dropbox-mnt
===========

A tool to mount a DropBox account as a filesystem

This is currently a work in progress. My end goal is to have an efficient filesystem which uses caching and goroutines to mount my dropbox account.

## Working Features

config generation
cd and ls... woooo!
simple caching when listing files

## Planned Features

Chunked up/downloading in seperate Go routines
Full caching for any read / write
Stagger uploads to reduce network usage

## Setup

      $ go get github.com/hanwen/go-fuse/fuse
      $ go get github.com/scottferg/Dropbox-Go/dropbox
      $ go install

## Usage

Mount to a directory:
      $ ./dropbox-mnt MOUNTPATH

Unmount a directory without breaking everything:
      $ fusermount -u MOUNTPATH

## Testing

copy your config to test/ and run:
      $ go test
