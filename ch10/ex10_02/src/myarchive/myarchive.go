package myarchive

import (
	"fmt"
)

type ArchiveFile interface {
	IsValid() bool
	Decompress() error
}

var factories []func(path string) ArchiveFile

func AddArchiveFileFactory(factory func(path string) ArchiveFile) {
	factories = append(factories, factory)
}

func Archive(path string) error {
	for _, factory := range factories {
		archive := factory(path)
		if !archive.IsValid() {
			continue
		}

		return archive.Decompress()
	}

	return fmt.Errorf("Unsupported archive file : %s", path)
}
