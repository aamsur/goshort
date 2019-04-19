// 
// 
// 

package engine

import "github.com/aamsur/goshort/src/generate"

func init() {
	handlers["generate"] = &generate.Handler{}
}
