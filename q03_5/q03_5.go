package main

import "io"

func copyN(dest io.Writer, src io.Reader, n int64) (int64, error) {
	written, err := io.Copy(dest, io.LimitReader(src, n))
	if written == n {
		return n, nil
	}
	if written < n && err == nil {
		err = io.EOF
	}
	return written, err
}
