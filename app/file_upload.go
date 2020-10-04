package app

import "net/http"

// FileUpload fetches the file and uploads to s3
func FileUpload(r *http.Request) (string, error) {
	// Retrieve the file from form data
	f, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	return UploadObject(
		BucketName,
		header.Filename,
		f,
	)
}
