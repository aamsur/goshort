// 
// 
// 

package main

import (
	"github.com/aamsur/goshort/engine"

	"git.qasico.com/cuxs/common/log"
	"git.qasico.com/cuxs/cuxs"
)

// init preparing application instances.
func init() {
	log.DebugMode = cuxs.IsDebug()
	log.Log = log.New()

	if e := cuxs.DbSetup(); e != nil {
		panic(e)
	}
}

// main creating new instances application
// and serving application server.
func main() {
	// starting server
	cuxs.StartServer(engine.Router())
}
