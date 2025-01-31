package toolkit

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"testing"
)

var uploadTests = []struct {
	name             string
	allowedFileTypes []string
	renameFile       bool
	errExpected      bool
}{
	{name: "Allowed no rename",
		allowedFileTypes: []string{"image/jpeg", "image/png", "image/gif"},
		renameFile:       false, errExpected: false},
	{name: "Allowed rename",
		allowedFileTypes: []string{"image/jpeg", "image/png", "image/gif"},
		renameFile:       true, errExpected: false},
	{name: "Not Allowed",
		allowedFileTypes: []string{"image/jpeg"},
		renameFile:       false, errExpected: true},
}

func TestTools_UploadFiles(t *testing.T) {
	for _, e := range uploadTests {
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer writer.Close()
			defer wg.Done()

			part, err := writer.CreateFormFile("file",
				"./test_data/test.jpg")
			if err != nil {
				t.Error(err)
				return
			}

			f, err := os.Open("./test_data/test.jpg")
			if err != nil {
				t.Error(err)
				return
			}
			defer f.Close()
			img, _, err := image.Decode(f)
			if err != nil {
				t.Error("Error decoding image", err)
				return
			}

			err = png.Encode(part, img)
			if err != nil {
				t.Error("Error encoding image", err)
				return
			}
		}()

		req, err := http.NewRequest("POST", "/", pr)
		req.Header.Add("Content-Type", writer.FormDataContentType())
		if err != nil {
			t.Error(err)
			return
		}
		var testTools Tools
		testTools.AllowedFileTypes = e.allowedFileTypes

		uploadedFiles, err := testTools.UploadFiles(req,
			"./test_data/uploads",
			e.renameFile)
		if err != nil && !e.errExpected {
			t.Error(err)
			return
		}

		if !e.errExpected {
			if _, err := os.Stat(fmt.Sprintf("./test_data/uploads/%s",
				uploadedFiles[0].NewFileName)); os.IsNotExist(err) {
				t.Error("File not found in uploads folder")
				return
			}

			_ = os.Remove(fmt.Sprintf("./test_data/uploads/%s", uploadedFiles[0].NewFileName))
		}

		if !e.errExpected && err != nil {
			t.Error(err, "	", e.name)
			return
		}
		wg.Wait()
	}
}

func TestTools_UploadOneFile(t *testing.T) {
	for _ = range uploadTests {
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)

		go func() {
			defer writer.Close()

			part, err := writer.CreateFormFile("file",
				"./test_data/test.jpg")
			if err != nil {
				t.Error(err)
				return
			}

			f, err := os.Open("./test_data/test.jpg")
			if err != nil {
				t.Error(err)
				return
			}
			defer f.Close()
			img, _, err := image.Decode(f)
			if err != nil {
				t.Error("Error decoding image", err)
				return
			}

			err = png.Encode(part, img)
			if err != nil {
				t.Error("Error encoding image", err)
				return
			}
		}()

		req, err := http.NewRequest("POST", "/", pr)
		req.Header.Add("Content-Type", writer.FormDataContentType())
		if err != nil {
			t.Error(err)
			return
		}
		var testTools Tools

		uploadedFiles, err := testTools.UploadOneFile(req,
			"./test_data/uploads",
			true)
		if err != nil {
			t.Error(err)
			return
		}

		if _, err := os.Stat(fmt.Sprintf("./test_data/uploads/%s",
			uploadedFiles.NewFileName)); os.IsNotExist(err) {
			t.Error("File not found in uploads folder")
			return
		}
		_ = os.Remove(fmt.Sprintf("./test_data/uploads/%s", uploadedFiles.NewFileName))

	}
}
