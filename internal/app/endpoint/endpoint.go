package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	d := e.s.DaysLeft()
	s := fmt.Sprintf("Amount of days left until 1 Jan 2025: %d", d)

	return ctx.String(http.StatusOK, s)
}
