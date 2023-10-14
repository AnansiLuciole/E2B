package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/e2b-dev/infra/packages/envd/internal/user"

	"go.uber.org/zap"
)

const MaxUploadSize = 100 * 1024 * 1024 // 100MB

func Upload(logger *zap.SugaredLogger, w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)

	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		logger.Error("Error parsing multipart form:", err)
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 100MB in size", http.StatusBadRequest)

		return
	}

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		logger.Error("Error retrieving the file from form-data:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			logger.Error("Error closing file:", closeErr)
		}
	}()

	filepath := r.Form.Get("path")

	var newFilePath string

	if filepath == "" {
		// Create a new file in the user's homedir if no path in the form is specified
		_, _, homedir, _, userErr := user.GetUser(user.DefaultUser)
		if userErr != nil {
			logger.Panic("Error getting user home dir:", userErr)
			http.Error(w, userErr.Error(), http.StatusInternalServerError)

			return
		}

		newFilePath = path.Join(homedir, fileHeader.Filename)
	} else {
		newFilePath = filepath
	}

	dst, err := os.Create(newFilePath)
	if err != nil {
		logger.Error("Error creating the file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer func() {
		closeErr := dst.Close()
		if closeErr != nil {
			logger.Error("Error closing file:", closeErr)
		}
	}()

	pr := &Progress{
		TotalSize: fileHeader.Size,
	}

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, io.TeeReader(file, pr))
	if err != nil {
		logger.Error("Error saving file to filesystem:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	fmt.Fprintf(w, "Upload successful")
}
