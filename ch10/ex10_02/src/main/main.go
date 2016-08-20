package main

import (
	"fmt"
	"myarchive"
	_ "myarchive/tar"
	_ "myarchive/zip"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <archive_file>", os.Args[0])
		return
	}

	if err := myarchive.Archive(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v", err)
	}
}
