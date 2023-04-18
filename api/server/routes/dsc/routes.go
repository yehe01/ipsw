package dsc

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes adds the download routes to the router
func AddRoutes(rg *gin.RouterGroup) {
	dr := rg.Group("/dsc")

	// dr.GET("/a2f", handler)     // TODO: implement this
	// dr.GET("/a2o", handler)     // TODO: implement this
	// dr.GET("/a2s", handler)     // TODO: implement this
	// dr.GET("/disass", handler)  // TODO: implement this
	// dr.GET("/dump", handler)    // TODO: implement this
	// dr.GET("/extract", handler) // TODO: implement this
	// dr.GET("/ida", handler)     // TODO: implement this
	// dr.GET("/image", handler)   // TODO: implement this

	// swagger:route GET /dsc/imports DSC getDscImports
	//
	// Imports
	//
	// Get list of dylibs that import a given dylib.
	//
	//     Produces:
	//     - application/json
	//
	//     Parameters:
	//       + name: path
	//         in: query
	//         description: path to dyld_shared_cache
	//         required: true
	//         type: string
	//	    + name: dylib
	//         in: query
	//         description: dylib to search for
	//         required: true
	//         type: string
	//     Responses:
	//       200: dscImportsResponse
	//       500: genericError
	dr.GET("/imports", dscImports)
	// swagger:route GET /dsc/info DSC getDscInfo
	//
	// Info
	//
	// Get info about a given DSC
	//
	//     Produces:
	//     - application/json
	//
	//     Parameters:
	//       + name: path
	//         in: query
	//         description: path to dyld_shared_cache
	//         required: true
	//         type: string
	//     Responses:
	//       200: dscInfoResponse
	//       500: genericError
	dr.GET("/info", dscInfo)
	// swagger:route GET /dsc/macho DSC getDscMacho
	//
	// MachO
	//
	// Get MachO info for a given dylib in the DSC.
	//
	//     Produces:
	//     - application/json
	//
	//     Parameters:
	//       + name: path
	//         in: query
	//         description: path to dyld_shared_cache
	//         required: true
	//         type: string
	//	    + name: dylib
	//         in: query
	//         description: dylib to search for
	//         required: true
	//         type: string
	//     Responses:
	//       200: dscMachoResponse
	//       500: genericError
	dr.GET("/macho", dscMacho)
	// dr.GET("/o2a", handler)     // TODO: implement this
	// dr.GET("/objc", handler)    // TODO: implement this
	// dr.GET("/patches", handler) // TODO: implement this
	// dr.GET("/search", handler)  // TODO: implement this
	// dr.GET("/slide", handler)   // TODO: implement this
	// dr.GET("/split", handler)   // TODO: implement this

	// swagger:route GET /dsc/str DSC getDscStrings
	//
	// Strings
	//
	// Get strings in the DSC that match a given pattern.
	//
	//     Produces:
	//     - application/json
	//
	//     Parameters:
	//       + name: path
	//         in: query
	//         description: path to dyld_shared_cache
	//         required: true
	//         type: string
	//	    + name: pattern
	//         in: query
	//         description: regex to search for
	//         required: true
	//         type: string
	//     Responses:
	//       200: dscStringsResponse
	//       500: genericError
	dr.GET("/str", dscStrings)
	// dr.GET("/stubs", handler) // TODO: implement this
	// dr.GET("/swift", handler) // TODO: implement this

	// swagger:operation GET /dsc/symaddr DSC getDscSymbols
	//
	// Symbols
	//
	// Get symbols addresses in the DSC that match a given lookup JSON payload.
	//
	// ---
	// consumes:
	//   - "application/json"
	// produces:
	//   - "application/json"
	// parameters:
	//   - name: lookups
	//     in: body
	//     description: Symbol lookups
	//     required: true
	//     schema:
	//       type: array
	//       items:
	//        $ref: "#/definitions/Symbol"
	//   - name: path
	//     in: query
	//     description: path to dyld_shared_cache
	//     required: true
	//     type: string
	// responses:
	//   '200':
	//     description: symbol lookup response
	//     schema:
	//       $ref: '#/responses/dscSymbolsResponse'
	//   '400':
	//     description: bad request
	//     schema:
	//       $ref: '#/responses/genericError'
	//   '500':
	//     description: error
	//     schema:
	//       $ref: '#/responses/genericError'
	dr.POST("/symaddr", dscSymbols)
	// dr.GET("/tbd", handler)    // TODO: implement this

	// swagger:route GET /dsc/webkit DSC getDscWebkit
	//
	// Webkit
	//
	// Get <code>webkit</code> version from dylib in the DSC.
	//
	//     Produces:
	//     - application/json
	//
	//     Parameters:
	//       + name: path
	//         in: query
	//         description: path to dyld_shared_cache
	//         required: true
	//         type: string
	//     Responses:
	//       200: dscWebkitResponse
	//       500: genericError
	dr.GET("/webkit", dscWebkit) // TODO: implement this
	// dr.GET("/xref", handler)   // TODO: implement this
}
