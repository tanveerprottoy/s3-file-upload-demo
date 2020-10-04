package app

import "net/http"

// ParseMultipartForm parses MultipartForm
func ParseMultipartForm(r *http.Request) *http.Request {
	// left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil
	}
	return r
}
