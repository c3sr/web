// Code generaTed by fileb0x at "2018-11-10 23:58:36.108319 -0600 CST m=+0.956217939" from config file "b0x.yml" DO NOT EDIT.
// modified(2018-11-06 23:55:59.18114811 -0600 CST)
// original path: jaeger-ui-build/build/asset-manifest.json

package static

import (
	"bytes"
	"compress/gzip"
	"io"

	"os"
)

// FileJaegerAssetManifestJSON is "jaeger/asset-manifest.json"
var FileJaegerAssetManifestJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x04\xff\x84\x8e\xd1\x0a\xc2\x30\x0c\x45\xdf\xf7\x15\xa5\xcf\xda\x56\xc7\x90\xf9\x37\x69\x1b\x43\x8b\x75\x62\xc6\x5e\xc4\x7f\x17\xc3\x60\x01\x05\x5f\x73\xcf\xc9\xbd\xcf\xce\x18\xdb\xa0\xdc\x5c\x62\xb6\x67\x63\x79\x86\xb9\x24\x9f\x98\xbd\x9c\x87\x7e\x0c\x21\xc7\x2c\xf9\x4e\xd3\xae\xc1\xfd\x9f\x21\xcc\x66\x55\x5d\x51\xd7\x86\x10\x2f\x39\x0d\x11\x3f\xa9\x46\xc5\xdd\x16\xfd\xc0\x85\x10\x65\x5d\xdd\x30\x17\xf0\x15\x90\xf0\xb1\xbf\x4e\x34\x39\x5e\x48\xfd\xf8\xce\xe1\x14\xc6\x3e\x1e\x8e\x8e\x17\xb2\xdd\xeb\x1d\x00\x00\xff\xff\x43\x1a\x7b\x7a\x0f\x01\x00\x00")

func init() {

	rb := bytes.NewReader(FileJaegerAssetManifestJSON)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err := FS.OpenFile(CTX, "jaeger/asset-manifest.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
