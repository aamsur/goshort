// 
// 
// 

package engine

import "github.com/aamsur/goshort/src/auth"

func init() {
	handlers["auth"] = &auth.Handler{}
}
