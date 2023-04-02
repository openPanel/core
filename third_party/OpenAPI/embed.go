package OpenAPI

import (
	_ "embed"

	"github.com/flowchartsman/swaggerui"
)

//go:embed openPanel.swagger.json
var spec []byte

var SwaggerUIHandler = swaggerui.Handler(spec)
