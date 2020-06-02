package middleware

import (
	"database/sql"
	"github.com/gorilla/context"
	"net/http"
)

type ContextInjector struct {
	*sql.DB
}

func (c *ContextInjector) Injector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "postgres", c.DB)
		next.ServeHTTP(w, r)
	})
}
