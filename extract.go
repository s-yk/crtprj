package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Extract extract zip
func Extract(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fr, err := f.Open()
		if err != nil {
			return err
		}
		defer fr.Close()

		if f.FileInfo().IsDir() {
			// if directory, make directory
			destPath := filepath.Join(dest, f.Name)
			err := os.MkdirAll(destPath, 0777)
			if err != nil {
				return err
			}
		} else {
			// if not directory;means f is file, write file
			buf := make([]byte, f.UncompressedSize)
			_, err := io.ReadFull(fr, buf)
			if err != nil {
				return err
			}

			destPath := filepath.Join(dest, f.Name)
			err = ioutil.WriteFile(destPath, buf, 0777)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
