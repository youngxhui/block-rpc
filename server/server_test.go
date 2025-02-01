package server

import (
	"testing"

	"github.com/youngxhui/block-rpc/remote"
)

func TestRouter_Register(t *testing.T) {
	t.Run("Register_GetHandler", func(t *testing.T) {
		ts := demoService{}
		router := NewRouter()
		handler := make(map[string]Handler, 0)
		handler["t"] = func(req *remote.Request) (*remote.Response, error) {
			return &remote.Response{}, nil
		}
		router.Register(ts, handler)

		// if _, has := router.GetHandler("testing", "t"); has {
		// 	t.Log("Success")
		// }
	})
}

type demoService struct{}

func (t demoService) Name() string {
	return "testing"
}
