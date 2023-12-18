package querybuilder

import "embed"

//go:embed marshal.go querybuilder.go
var Embeds embed.FS
