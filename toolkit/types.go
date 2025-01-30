package toolkit

type Tools struct {
	MaxFileSize      int
	AllowedFileTypes []string
}

// UploadedFile represents a file that has been uploaded.
type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int64
}
