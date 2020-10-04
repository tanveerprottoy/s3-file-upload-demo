package app

import (
	"errors"
	"net/http"
)

// PostFile post new file
func PostFile(w http.ResponseWriter, r *http.Request) {
	// reqMultipartParsed
	rmp := ParseMultipartForm(r)
	if rmp == nil {
		RespondError(w, errors.New("Parse error"), http.StatusInternalServerError)
		return
	}
	url, err := FileUpload(rmp)
	if err != nil {
		RespondError(w, err, http.StatusInternalServerError)
		return
	}
	Respond(w, map[string]string{"url": url})
}
