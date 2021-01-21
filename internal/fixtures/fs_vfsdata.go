// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package fixtures

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS statically implements the virtual filesystem provided to vfsgen.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/payout.created.json": &vfsgen۰CompressedFileInfo{
			name:             "payout.created.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 461,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\x90\xe1\x4a\xc3\x30\x10\xc7\xbf\xf7\x29\x8e\x7c\x16\x56\x2b\x6e\xda\x17\x18\x82\x13\x99\x38\x04\x91\x72\x4b\xaf\xac\x68\x93\x72\xb9\xe8\xba\xb1\x77\x97\x64\x74\xa6\x12\xb8\x40\xf2\xfb\xdd\x71\xff\x63\x06\xa0\x9a\x76\x2f\x9e\xc9\xa9\x12\xde\x33\x00\x80\x63\xac\x00\xca\x60\x47\xaa\x04\xd5\xe3\x60\xbd\x54\x9a\x09\xa5\xb5\x46\x5d\x8d\x40\x8f\xb2\x0b\xc0\xec\xfb\x7a\x76\x86\xdc\xdf\x67\x47\xb2\xb3\x75\xf4\xad\x93\x54\x62\xec\xc2\xb4\x71\x4e\x60\x6d\x1d\x27\x3d\xac\x9e\x5f\x2e\x24\x80\xc2\xce\x7a\x23\xaa\x84\x22\xcf\xf3\xe4\x5d\x7b\x66\x32\x7a\x88\xce\xd3\x3a\x55\x7a\xcf\xbd\x75\xb1\x1b\x53\xe3\x4d\x3d\xe9\xa7\x75\x68\x58\x19\xdf\x6d\x89\x03\x53\xdc\x84\x93\xe7\x8b\xfb\xc5\x7c\x31\x2f\xee\x52\x3a\xd8\xd5\xa8\xb4\x71\x95\x06\xab\xe5\xfa\x67\x59\xdc\x1e\xde\x1e\xdb\xcd\xea\x75\x93\xf2\x06\x99\xcf\x09\x95\xa0\x62\x5a\x54\xc3\x76\x00\xc6\x83\xe5\x1e\x87\x3d\xe8\xaf\x76\x22\x58\xa1\x69\x12\x00\xea\x93\xe2\x5a\x42\x4e\xfe\x99\x17\xea\x94\xa5\x77\xa8\x1f\xd9\xe9\x37\x00\x00\xff\xff\x00\x44\xd2\xb7\xcd\x01\x00\x00"),
		},
		"/payout_creation.setup.json": &vfsgen۰CompressedFileInfo{
			name:             "payout_creation.setup.json",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 966,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x93\x41\x6f\xe2\x30\x10\x85\xef\xf9\x15\x23\x6b\x8f\x48\x38\xac\x96\x6c\x39\xf6\xc6\xa1\xa8\x2a\xc7\xaa\x8a\x06\xc7\x94\x08\x62\x5b\x8e\x8d\x70\xa3\xfc\xf7\xca\x26\x49\x1d\xa0\x6a\x85\x90\x88\x92\xf7\xde\xcc\xe4\x9b\xb8\x49\x00\xc8\xb6\x3c\x19\xab\x79\x4d\x16\xf0\x9a\x00\x00\x34\xe1\x1f\x80\x08\xac\x38\x59\x00\x61\x52\x18\x64\x86\x4c\x7a\x41\xa1\xd9\x79\x61\x7a\x4c\xa7\x9d\x58\x7f\xa9\x15\x37\x3b\x59\x78\x5d\xc9\x7a\x94\xd2\x58\xf9\x36\x7d\x83\xa8\xc5\x3b\x5a\x8d\x47\x30\xe8\xf0\x30\x24\x00\x88\x71\x2a\xe8\xbc\x52\x07\xe9\x38\x27\x9d\xd4\x86\x6b\x3b\xb9\x3d\xf0\xd6\x8a\x22\x47\xc6\xa4\x15\xdf\x4c\x1d\x3b\xee\x1c\xbd\x4b\xe7\xfd\x88\x1b\x14\xfb\xab\xa6\x30\xd0\xcb\xcb\x50\xf7\x4f\xd3\xdd\x2f\xca\xa2\x8d\x6d\xa3\x78\xdc\xe8\x67\x4a\x00\xa4\xdc\xd6\xcc\x3b\xd6\x8f\xcb\x15\xa5\x34\x4b\xe9\xbf\xb1\xa1\x9f\x56\xd8\x6a\xc3\xb5\xb7\xa6\x69\x4a\x29\x25\x83\xa9\xfd\x15\x5a\x85\x4e\x5a\x93\x33\xcd\xd1\x94\x52\xdc\xa6\x7b\x36\xdd\xc9\xb5\x92\x45\xe8\xb4\x7c\x7a\x5e\xc7\x80\xb0\xea\xd0\xcc\x28\xa5\x31\x5f\xab\x35\x17\xcc\x85\xcc\xea\x25\x8e\x28\xab\x95\xac\x43\x35\xcd\xfd\xca\x47\xf5\xae\x88\xcc\xfe\xfa\x1f\xa5\xd9\x43\x36\xcf\xe6\xb3\xff\xb1\x3b\xfe\x60\x86\x55\xc6\x0f\x2f\xf7\x29\x50\xeb\x33\x22\x7f\x82\x3c\x2e\x5e\xc0\xc6\x81\xc6\x0f\xa9\x15\xba\x13\xb0\x43\x39\x0a\x48\x13\x0e\x61\x43\xf6\x3c\xbc\x8c\xe1\xb5\xb9\xb0\x5f\xec\x28\x01\x78\x4b\xda\xcf\x00\x00\x00\xff\xff\x98\xda\x2c\xdc\xc6\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/payout.created.json"].(os.FileInfo),
		fs["/payout_creation.setup.json"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
