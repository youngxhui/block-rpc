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
		h, has := router.GetHandler("testing", "t")
		if !has {
			t.Errorf("Not Found %v ", has)
		}
		_, err := h(new(remote.Request))
		if err != nil {
			t.Errorf("handler execute error:%v", err)
		}
	})
}

type demoService struct{}

func (t demoService) Name() string {
	return "testing"
}

func TestRouter_GetHandler(t *testing.T) {
	t.Run("NotFoundHandler", func(t *testing.T) {
		router := NewRouter()
		_, has := router.GetHandler("demo", "t")
		if !has {
			t.Errorf("Not Found Handler error, has: %v", has)
		}
	})
}
