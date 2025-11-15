package jinrai

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/jinrai-js/go/internal/proxy"
)

func (c *Static) Serve(port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if c.Proxy != nil {
			for prefix, targetURL := range *c.Proxy {
				if strings.HasPrefix(r.URL.Path, prefix) {
					proxy.Handler(w, r, prefix, targetURL, c.Verbose)
					return
				}
			}
		}

		if c.Assets != nil {
			filePath := path.Join(c.Dist, r.URL.Path)

			if r.URL.Path == "/" {
				c.Handler(w, r)
				return
			}

			if _, err := os.Stat(filePath); err == nil {
				fs := http.FileServer(http.Dir(c.Dist))
				fs.ServeHTTP(w, r)
				return
			}
		}

		c.Handler(w, r)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
