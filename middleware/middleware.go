package middleware

import "crm-sebagian-team/config"

type AppsMiddleware struct {
	JWTKey []byte
}

// InitMiddleware initialize the middleware
func InitMiddleware(cfg *config.Config) *AppsMiddleware {
	return &AppsMiddleware{
		JWTKey: []byte(cfg.JWT.AccessSecret),
	}
}
