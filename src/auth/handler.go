// 
// 
// 

package auth

import (
	"git.qasico.com/cuxs/cuxs"
	"github.com/labstack/echo"
)

// Handler collection handler for auth.
type Handler struct{}

// URLMapping declare endpoint with handler function.
func (h *Handler) URLMapping(r *echo.Group) {
	r.POST("", h.signin)
	r.GET("/me", h.me, cuxs.Authorized())
	r.GET("/logout", h.logout, cuxs.Authorized())
}

// signin endpoint to handle post http method.
func (h *Handler) signin(c echo.Context) (e error) {
	var r SignInRequest
	var sd *SessionData

	ctx := c.(*cuxs.Context)
	if e = ctx.Bind(&r); e == nil {
		if sd, e = Login(r.User); e == nil {
			ctx.Data(sd)
		}
	}
	return ctx.Serve(e)
}

// me endpoint untuk get sesion data yang lagi login.
func (h *Handler) me(c echo.Context) (e error) {
	var sd *SessionData

	ctx := c.(*cuxs.Context)
	// get current user dan data application menu
	if sd, e = UserSession(ctx); e == nil {
		ctx.Data(sd)
	}
	return ctx.Serve(e)
}

// logout to save the last logout at time on db
func (h *Handler) logout(c echo.Context) (e error) {
	ctx := c.(*cuxs.Context)
	_, e = Logout(ctx)
	return ctx.Serve(e)
}
