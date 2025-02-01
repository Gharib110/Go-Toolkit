package toolkit

import (
	"fmt"
	"net/http"
	"path"
)

// DownloadStaticFile download a file, and tries to force the browser to avoid
// displaying it in the browser by Content-Disposition
func (t *Tools) DownloadStaticFile(w http.ResponseWriter, r *http.Request,
	p, file, displayName string) {
	fp := path.Join(p, file)
	w.Header().Set("Content-Disposition", fmt.Sprintf("Attachment: filenam=\"%s\"",
		displayName))

	http.ServeFile(w, r, fp)

}
