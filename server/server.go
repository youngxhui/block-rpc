package server

import (
	"sync"

	"github.com/youngxhui/block-rpc/remote"
)

type Service interface {
	Name() string
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
	defer r.mu.RLock()
	service, ok := r.services[serviceName]
	if !ok {
		// TODO 默认需要一个 NotFoundHandler
		return nil, false
	}
	handler, ok := service[methodName]
	return handler, ok
}
