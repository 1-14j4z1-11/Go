package tar

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"myarchive"
	"os"
	"path/filepath"
)

type archiveFile struct {
	path string
}

func init() {
	myarchive.AddArchiveFileFactory(newArchive)
}

func newArchive(path string) myarchive.ArchiveFile {
	af := new(archiveFile)
	af.path = path
	return af
}

func (af *archiveFile) IsValid() bool {
	offset := 257
	magicSize := 5
	magic := "ustar"

	f, err := os.Open(af.path)
	if err != nil {
		return false
	}
	defer f.Close()

	n, err := f.Read(make([]byte, offset))
	if n != offset || err != nil {
		return false
	}

	buf := make([]byte, magicSize)
	n, err = f.Read(buf)
	if n != magicSize || err != nil {
		return false
	}

	return string(buf) == magic
}

func (af *archiveFile) Decompress() error {
	f, err := os.Open(af.path)
	if err != nil {
		return err
	}
	defer f.Close()

	root := "archive"
	os.MkdirAll(root, os.FileMode(777))

	reader := tar.NewReader(f)
	var header *tar.Header
	for {
		header, err = reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, reader)
		if err != nil {
			return err
		}
		path := filepath.Join(root, header.FileInfo().Name())
		os.MkdirAll(path, header.FileInfo().Mode())
		os.Remove(path)
		err = ioutil.WriteFile(path, buf.Bytes(), header.FileInfo().Mode())
		if err != nil {
			return err
		}
	}
	return nil
}
