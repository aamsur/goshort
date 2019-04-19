// 
// 
// 

package generate

import (
	"git.qasico.com/cuxs/cuxs"
	"github.com/labstack/echo"
)

// Handler collection handler for user.
type Handler struct{}

// URLMapping declare endpoint with handler function.
func (h *Handler) URLMapping(r *echo.Group) {
	r.POST("", h.create)
}

// create endpoint to handle post http method
func (h *Handler) create(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	var r createRequest

	if e = ctx.Bind(&r); e == nil {
		data := r.Transform()
		if e = Save(data); e == nil {
			ctx.Data(data)
		}
	}

	return ctx.Serve(e)
}
