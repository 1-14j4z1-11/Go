package main

import (
	"io"
)

type limitReader struct {
	base  io.Reader
	count int64
	limit int64
}

func (b *limitReader) Read(p []byte) (n int, err error) {
	if (b.base == nil) || (b.base.(io.Reader) == nil) {
		err = io.EOF
		return
	}

	n, err = b.base.Read(p)
	b.count += int64(n)

	if b.count >= b.limit {
		n -= int(b.count - b.limit)
		err = io.EOF
	}

	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	reader := new(limitReader)
	reader.base = r
	reader.limit = n

	return reader
}
