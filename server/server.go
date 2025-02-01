package server

import (
	"sync"

	"github.com/youngxhui/block-rpc/remote"
)

type Service interface {
	Name() string
}

// notFoundHandler 404使用
var notFoundHandler Handler = func(req *remote.Request) (*remote.Response, error) {
	return &remote.Response{}, nil
}

type Handler func(req *remote.Request) (*remote.Response, error)

type Router struct {
	mu       sync.RWMutex
	services map[string]map[string]Handler // service.method -> Handler
}

func NewRouter() *Router {
	return &Router{
		services: make(map[string]map[string]Handler),
	}
}

// Register 注册路由
func (r *Router) Register(service Service, handlers map[string]Handler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	serviceName := service.Name()
	if _, ok := r.services[serviceName]; !ok {
		r.services[serviceName] = make(map[string]Handler)
	}

	for method, handler := range handlers {
		r.services[serviceName][method] = handler
	}
}

func (r *Router) GetHandler(serviceName, methodName string) (Handler, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	service, ok := r.services[serviceName]
	if !ok {
		// TODO 默认需要一个 NotFoundHandler
		return notFoundHandler, true
	}
	handler, ok := service[methodName]
	return handler, ok
}
