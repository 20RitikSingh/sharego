package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// // Limit upload size to 100MB
	// r.ParseMultipartForm(100 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create the file
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file data
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<div class="alert alert-success mt-3">File uploaded successfully</div>`)
}
