package toolkit

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// UploadOneFile  uploads a single file to the given directory.
func (t *Tools) UploadOneFile(r *http.Request, uploadDir string,
	rename ...bool) (*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	files, err := t.UploadFiles(r, uploadDir, renameFile)
	if err != nil {
		return nil, err
	}

	return files[0], nil
}

// UploadFiles uploads files to the given directory.
func (t *Tools) UploadFiles(r *http.Request, uploadDir string,
	rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	var uploadedFiles []*UploadedFile

	if t.MaxFileSize == 0 {
		t.MaxFileSize = 1024 * 1024 * 1024
	}

	err := t.CreateDirIfNotExist(uploadDir)
	if err != nil {
		return nil, err
	}

	err = r.ParseMultipartForm(int64(t.MaxFileSize))
	if err != nil {
		return nil, errors.New("the Uploaded files is too large")
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, hdr := range fHeaders {
			uploadedFiles, err = func(
				uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
				var uf UploadedFile
				infile, err := hdr.Open()
				if err != nil {
					return nil, err
				}
				defer infile.Close()

				buff := make([]byte, 512)
				_, err = infile.Read(buff)
				if err != nil {
					return nil, err
				}

				allowed := false
				fileType := http.DetectContentType(buff)

				if len(t.AllowedFileTypes) > 0 {
					for _, tt := range t.AllowedFileTypes {
						if strings.EqualFold(fileType, tt) {
							allowed = true
							break
						}
					}
				} else {
					allowed = true
				}

				if !allowed {
					return nil, errors.New("file type not allowed")
				}

				_, err = infile.Seek(0, 0)
				if err != nil {
					return nil, err
				}

				if renameFile {
					uf.NewFileName = fmt.Sprintf("%s%s",
						t.RandomString(10),
						filepath.Ext(hdr.Filename))
				} else {
					uf.NewFileName = hdr.Filename

				}

				uf.OriginalFileName = hdr.Filename

				var outfile *os.File
				defer outfile.Close()
				outfile, err = os.Create(filepath.Join(uploadDir, uf.NewFileName))
				if err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outfile, infile)
					if err != nil {
						return nil, err
					}

					uf.FileSize = fileSize

				}

				uploadedFiles = append(uploadedFiles, &uf)

				return uploadedFiles, nil
			}(uploadedFiles)
			if err != nil {
				return uploadedFiles, err
			}
		}
	}

	return uploadedFiles, nil
}
