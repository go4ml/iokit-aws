package iokit_aws

import (
	"io"
	"sudachen.xyz/pkg/iokit"
)

const Proto = "s3"

func init() {
	iokit.UrlReaderFactory[Proto] = func(url string)interface{Download(io.Writer)error} {
		return Url(url)
	}
	iokit.UrlWriterFactory[Proto] = func(url string)interface{Upload(rd io.Reader, metadata ...map[string]string) error} {
		return Url(url)
	}
}

