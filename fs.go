package dagger

import (
	"dagger.io/dagger/querybuilder"
	"embed"
)

// These are exported so that they can be used by codegen.

var QueryBuilder embed.FS = querybuilder.Embeds

//go:embed go.mod
var GoMod []byte

//go:embed go.sum
var GoSum []byte
