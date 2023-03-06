package dam

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/imgproxy/imgproxy/v3/config"
)

type transport struct {
	base string
}

func New() (http.RoundTripper, error) {
	return transport{config.DAMEndpoint}, nil
}

func (t transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	damPath := fmt.Sprintf("%s/api/asset%s/data", t.base, req.URL.Path)
	log.Debug("Round tripping DAM " + damPath)
	return http.Get(damPath)
}
