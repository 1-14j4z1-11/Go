package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"myarchive"
	"os"
	"path/filepath"
)

type archiveFile struct {
	path   string
	valid  bool
	reader *zip.ReadCloser
}

func init() {
	myarchive.AddArchiveFileFactory(newArchive)
}

func newArchive(path string) myarchive.ArchiveFile {
	af := new(archiveFile)
	af.path = path

	r, err := zip.OpenReader(path)
	if err != nil {
		af.valid = false
	} else {
		af.valid = true
		af.reader = r
	}

	return af
}

func (af *archiveFile) IsValid() bool {
	return af.valid
}

func (af *archiveFile) Decompress() error {
	root := "archive"
	os.MkdirAll(root, os.FileMode(777))

	for _, f := range af.reader.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		if f.FileInfo().IsDir() {
			path := filepath.Join(root, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, rc)
			if err != nil {
				return err
			}
			path := filepath.Join(root, f.Name)
			os.MkdirAll(path, f.Mode())
			os.Remove(path)
			err = ioutil.WriteFile(path, buf.Bytes(), f.Mode())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
