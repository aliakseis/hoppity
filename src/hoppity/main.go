// Hoppity Hop manipulation REST API
//
// The purpose of this application is to provide
// Hoppity Hop manipulation REST API
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"hoppity/rest"
	"log"
)

var (
	Version	= rest.VERSION_V1
	BuildDate	string
	GitCommit   string
)

func main() {
	log.Print("Version: ", Version)
	log.Print("Build Time: ", BuildDate)
	log.Print("GitCommit: ", GitCommit)

	rest.CreateRoutes()
}
