// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package manpagejp

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"
)

type JMProject struct {
	uri       *url.URL
	timestamp string
	file      *os.File

	manpageFile map[string][]string
}

const timestampLayout = "20060102"

func init() {
	log.SetFlags(log.Lshortfile)
}

func NewJMProject() *JMProject {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), 15, 0, 0, 0, 0, now.Location())
	timestamp := t.Format(timestampLayout)
	uri, err := url.Parse(fmt.Sprintf("https://linuxjm.osdn.jp/man-pages-ja-%s.tar.gz", timestamp))
	if err != nil {
		log.Fatal(err)
	}

	return &JMProject{
		uri:         uri,
		timestamp:   timestamp,
		manpageFile: make(map[string][]string),
	}
}

func (j *JMProject) URL() string {
	return j.uri.String()
}

func (j *JMProject) Download(dst string) error {
	uri := j.URL()
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Errorf("Error while %s downloded: %v", uri, err)
	}
	defer resp.Body.Close()

	dir, err := os.Stat(dst)
	if err != nil {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}
	if !dir.IsDir() {
		return fmt.Errorf("%s is not directory", dst)
	}

	fi, err := os.Create(filepath.Join(dst, path.Base(uri)))
	if err != nil {
		return err
	}
	defer fi.Close()

	if _, err := io.Copy(fi, resp.Body); err != nil {
		return err
	}
	j.file = fi

	return nil
}

func (j *JMProject) Extract() error {
	fi, err := os.Open(j.file.Name())
	if err != nil {
		return err
	}
	defer fi.Close()

	gz, err := gzip.NewReader(fi)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gz)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		filename := hdr.Name
		switch hdr.Typeflag {
		case tar.TypeDir: // = directory
			os.Mkdir(filename, 0755)
		case tar.TypeReg: // = regular file
			data := make([]byte, hdr.Size)
			_, err := tarReader.Read(data)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			ioutil.WriteFile(filename, data, 0755)
		default:
			fmt.Printf("Unable to figure out type: %c in file %s\n", hdr.Typeflag, filename)
		}
	}

	return nil
}

var manpageFileRe = regexp.MustCompile(`\w\.[1-9]`)

func (j *JMProject) Copy(dst, src string) error {
	srcRoot := filepath.Join(src, "man-pages-ja-"+j.timestamp, "manual")

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && manpageFileRe.MatchString(path) {
			if info.Mode().Perm() != 0644 {
				if err := os.Chmod(path, 0644); err != nil {
					return err
				}
			}

			dir := filepath.Dir(path)
			j.manpageFile[dir] = append(j.manpageFile[dir], filepath.Base(path))
		}

		return nil
	}
	if err := filepath.Walk(srcRoot, walkFunc); err != nil {
		return err
	}

	for srcDir, files := range j.manpageFile {
		for _, f := range files {
			dstDir := filepath.Join(dst, filepath.Base(srcDir))
			if _, err := os.Stat(dstDir); err != nil {
				if err := os.MkdirAll(dstDir, 0755); err != nil {
					return err
				}
			}
			srcFile := filepath.Join(srcDir, filepath.Base(f))
			dstFile := filepath.Join(dstDir, f)
			if err := copyFile(dstFile, srcFile); err != nil {
				return err
			}
		}
	}

	return nil
}

func CompressTarGz(dst string, src string) error {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)

	srcBuf, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	tw.Write(srcBuf)
	defer tw.Close()

	// make sure the gzip writer flushes any pending bytes
	if err = gz.Close(); err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, buf.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

// copyFile copies a file from src to dst.
func copyFile(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		e := s.Close()
		if err == nil {
			err = e
		}
	}()

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		e := d.Close()
		if err == nil {
			err = e
		}
	}()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}

	i, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, i.Mode())
}
