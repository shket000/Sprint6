package handlers

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// HomeHandler возвращает HTML-форму из файла index.html.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

// UploadHandler обрабатывает загруженный файл, конвертирует его содержимое и сохраняет результат.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничение размера файла (10 МБ)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file content", http.StatusInternalServerError)
		return
	}

	input := string(data)
	output, err := service.AutoDetectAndConvert(input)
	if err != nil {
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	fileName := time.Now().UTC().Format("20060102150405") + ext
	outFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Unable to create output file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	if _, err := outFile.WriteString(output); err != nil {
		http.Error(w, "Unable to write to output file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(output))
}
