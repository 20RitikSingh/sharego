package handler

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func zipDir(source string, zipfile *os.File) error {

	archive := zip.NewWriter(zipfile)
	defer func() {
		if err := archive.Close(); err != nil {
			log.Println("error closing archive ", err)
		}
	}()
	return filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Name() == filepath.Base(zipfile.Name()) {
			return nil
		}

		// Get relative path for archive
		relPath, err := filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if relPath == "." {
			return nil // skip root dir
		}

		if d.IsDir() {
			_, err := archive.Create(relPath + "/")
			return err
		}

		// Get os.FileInfo for zip header
		info, err := d.Info()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Method = zip.Deflate

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Println("error closing file ", err)
			}
		}()

		_, err = io.Copy(writer, file)
		return err
	})
}

func ZipHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/downloadzip/")
	// Security check to prevent directory traversal
	if strings.Contains(path, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		log.Println("err with path")
		return
	}
	tmpFile, err := os.CreateTemp(".", "tempZip")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer func() {
		if err := os.Remove(tmpFile.Name()); err != nil {
			log.Println("error removing temp zip file", err)
		}
	}()
	defer func() {
		if err := tmpFile.Close(); err != nil {
			log.Println("error closing temp file ", err)
		}
	}()
	fmt.Printf("compressing ./%s \n", path)
	err = zipDir(filepath.Join("./", path), tmpFile)
	fmt.Println("file ready to download!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+path+"_compressed.zip")
	http.ServeFile(w, r, filepath.Join("./", tmpFile.Name()))
	fmt.Println("file served !!")
}
