package http

import (
	"net/http"

	"go-common/library/log"
	bm "go-common/library/net/http/blademaster"
)

// ping check server ok.
func ping(c *bm.Context) {
	var err error
	if err = pingSvc.Ping(c); err != nil {
		log.Error("app-show service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
