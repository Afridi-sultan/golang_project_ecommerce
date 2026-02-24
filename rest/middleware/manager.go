package middleware

import (
	
	"net/http"
)

type Middleware func (http.Handler)http.Handler

type Manager struct{
	globalMiddlewares []Middleware 
}


func NewManager()*Manager{
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}


func (mngr *Manager) Use (middleware ...Middleware){
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
}

 
func (mngr *Manager) With(midllewares ...Middleware) Middleware{
	return func(handler http.Handler) http.Handler {
		h := handler

		for i := len(midllewares)-1; i >= 0; i--{
			middleware := midllewares[i]

			h = middleware(h)
		}
		return h
	}	
}



func (mngr *Manager) WrapMux(handler http.Handler)http.Handler{

	h := handler

	for _,middleware := range mngr.globalMiddlewares{
		h = middleware(h)
	}
	return h
}
	
