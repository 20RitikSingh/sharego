package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func FileListHandler(w http.ResponseWriter, r *http.Request) {
	// Get current directory entries
	path := strings.TrimPrefix(r.URL.Path, "/list/")
	if strings.HasPrefix(path, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
	}
	path, err := url.PathUnescape(path)
	if err != nil {
		log.Println("error getting unescaped path")
	}
	// log.Println(path)
	entries, err := os.ReadDir("./" + path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading directory: %v", err), http.StatusInternalServerError)
		return
	}

	// Create HTML template for list items
	const listItem = `
    <li class="list-group-item">
		{{if .IsDir}}
        <a href="/{{.Path}}" class="text-decoration-none">
		{{else}}
        <a href="/file/{{.Path}}" download class="text-decoration-none">
		{{end}}
		<i class="bi {{if .IsDir}}bi-folder-fill{{else}}bi-file-earmark{{end}} me-2"></i>
		{{if .IsDir}}📂{{else}}📄{{end}}{{.Name}}
        </a>
    </li>`

	// Generate HTML output
	var output strings.Builder
	output.WriteString(`<ul class="list-group">`)

	for _, entry := range entries {
		name := entry.Name()

		// Skip hidden files
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		tpl := template.Must(template.New("").Parse(listItem))
		err := tpl.Execute(&output, struct {
			Path  string
			Name  string
			IsDir bool
		}{
			Path:  filepath.Join(path, name),
			Name:  name,
			IsDir: entry.IsDir(),
		})

		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating template: %v", err), http.StatusInternalServerError)
			return
		}
	}
	output.WriteString(`</ul>`)

	w.Header().Set("Content-Type", "text/html")
	if _, err = fmt.Fprint(w, output.String()); err != nil {
		log.Println("error writing response ", err)
	}
}

func FileServerHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/file/")
	path, err := url.PathUnescape(path)
	if err != nil {
		log.Println("error getting unescaped path")
	}
	log.Println(path)
	// Security check to prevent directory traversal
	if strings.Contains(path, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		log.Println("err with path")
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(path))
	http.ServeFile(w, r, filepath.Join(".", path))
}
