package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var tasks sync.WaitGroup
	tick := time.Tick(500 * time.Millisecond)

	for _, root := range roots {
		tasks.Add(1)

		fileSizes := make(chan int64)
		var n sync.WaitGroup
		n.Add(1)
		go walkDir(root, &n, fileSizes)

		go func() {
			n.Wait()
			close(fileSizes)
		}()

		go func(root string) {
			var nfiles, nbytes int64
loop:
			for {
				select {
				case <-done:
					for range fileSizes {
						// Do nothing.
					}
					return
				case size, ok := <-fileSizes:
					if !ok {
						break loop
					}
					nfiles++
					nbytes += size
				case <-tick:
					printDiskUsage(root, nfiles, nbytes)
				}
			}

			printDiskUsage(root, nfiles, nbytes)
			tasks.Done()
		}(root)
	}

	tasks.Wait()
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s : %d files  %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}
