package middleware

import (
	"net/http"

	"github.com/anditakaesar/uwa-back/internal/client"
	"github.com/anditakaesar/uwa-back/internal/log"
)

type Error struct {
	error      error
	httpStatus int
}

type MiddlewareFunc = func(w http.ResponseWriter, r *http.Request) (*http.Request, *Error)

// Middleware is a helper for route to validate a http Request
type Middleware struct {
	Client *client.Client
	// UserDomain domain.UserDomainInterface
	Log log.LoggerInterface
}

func NewAdapter(d Middleware) Middleware {
	middleware := Middleware{
		Client: d.Client,
		// UserDomain: d.UserDomain,
		Log: d.Log,
	}
	return middleware
}

func (m Middleware) Do(h http.Handler, handler MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newRequest, middlewareError := handler(w, r)
		if middlewareError != nil {
			http.Error(w, middlewareError.error.Error(), middlewareError.httpStatus)
		} else {
			h.ServeHTTP(w, newRequest)
		}
	}
}

// Group is Helper for executes bunch of middleware
func (m Middleware) Group(h http.Handler, verifyAll bool, handlers ...MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newRequest := r
		var middlewareError *Error
		isMiddlewarePass := true

		for i := 0; i < len(handlers); i++ {
			newRequest, middlewareError = handlers[i](w, r)

			if middlewareError == nil && !verifyAll {
				isMiddlewarePass = true
				break
			}

			if middlewareError != nil {
				isMiddlewarePass = false

				if verifyAll {
					break
				}
			}
		}

		if !isMiddlewarePass {
			http.Error(w, middlewareError.error.Error(), middlewareError.httpStatus)
			return
		}

		h.ServeHTTP(w, newRequest)
	}
}

// Verify is helper for executes list of middleware from routes
func (m Middleware) Verify(h http.Handler, handlers ...MiddlewareFunc) http.HandlerFunc {
	return m.Group(h, false, handlers...)
}
