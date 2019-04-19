// 
// 
// 

package engine

import (
	"github.com/aamsur/goshort/src/click"
)

func init() {
	handlers["/"] = &click.Handler{}
}
