package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

type CustomMiddleware struct {
	Log *logrus.Logger
}

func NewCustomMiddleware(log *logrus.Logger) *CustomMiddleware {
	return &CustomMiddleware{Log: log}
}

func (m *CustomMiddleware) DevMode() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			stop := time.Now()

			request := c.Request()
			response := c.Response()
			latency := stop.Sub(start)

			logMessage := fmt.Sprintf(
				"method=%s, path=%s, status=%d, latency=%s, remote_ip=%s, user_agent=%s",
				request.Method,
				request.URL.Path,
				response.Status,
				latency,
				c.RealIP(),
				request.UserAgent(),
			)

			// Log with Logrus
			m.Log.WithFields(logrus.Fields{
				"method":     request.Method,
				"path":       request.URL.Path,
				"status":     response.Status,
				"latency":    latency,
				"remote_ip":  c.RealIP(),
				"user_agent": request.UserAgent(),
			}).Info(logMessage)

			return err
		}
	}
}
