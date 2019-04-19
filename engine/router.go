// 
// 
// 

package engine

import (
	"git.qasico.com/cuxs/cuxs"
	"github.com/labstack/echo"
)

// RouteHandlers interface of handlers
type RouteHandlers interface {
	URLMapping(r *echo.Group)
}

// handlers register an endpoint with handler here.
// it will automatic registered into routers
var handlers = map[string]RouteHandlers{}

// Router registering all handler into engine router.
func Router() *echo.Echo {
	engine := cuxs.New()
	v := engine.Group("")
	if len(handlers) > 0 {
		for p, h := range handlers {
			h.URLMapping(v.Group(p))
		}
	}
	return engine
}
