package main

import (
	"io"
)

type WriterWrapper struct {
	base  io.Writer
	count int64
}

func (b *WriterWrapper) Write(p []byte) (count int, err error) {
	if (b.base != nil) && b.base.(io.Writer) != nil {
		count, err = b.base.Write(p)
	}

	b.count += int64(count)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrapper := new(WriterWrapper)
	wrapper.base = w

	return wrapper, &wrapper.count
}
