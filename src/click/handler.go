package click

import (
	"git.qasico.com/cuxs/cuxs"
	"github.com/labstack/echo"
	"net/http"
)

// Handler collection handler for user.
type Handler struct{}

// URLMapping declare endpoint with handler function.
func (h *Handler) URLMapping(r *echo.Group) {
	r.GET(":short_url", h.redirect)
}

// get endpoint to handle get http method.
func (h *Handler) redirect(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)

	l, e := GetByShortUrl(ctx.Param("short_url"))

	header := c.Response().Header()
	header.Set("Cache-Control", "no-cache, private, no-store, must-revalidate, max-age=0")
	header.Set("Pragma", "no-cache")
	header.Set("Expires", "0")
	header.Set("X-Accel-Expires", "0")

	if e == nil {
		return ctx.Redirect(http.StatusMovedPermanently, l.LongUrl)
	}

	return ctx.Serve(e)
}
