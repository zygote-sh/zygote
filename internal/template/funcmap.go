package template

import (
	"html/template"

	"github.com/go-sprout/sprout"
	sptConv "github.com/go-sprout/sprout/registry/conversion"
	sptEnc "github.com/go-sprout/sprout/registry/encoding"
	sptEnv "github.com/go-sprout/sprout/registry/env"
	sptFs "github.com/go-sprout/sprout/registry/filesystem"
	sptMaps "github.com/go-sprout/sprout/registry/maps"
	sptNum "github.com/go-sprout/sprout/registry/numeric"
	sptRegex "github.com/go-sprout/sprout/registry/regexp"
	sptSemver "github.com/go-sprout/sprout/registry/semver"
	sptSlices "github.com/go-sprout/sprout/registry/slices"
	sptStd "github.com/go-sprout/sprout/registry/std"
	sptStr "github.com/go-sprout/sprout/registry/strings"
	sptTime "github.com/go-sprout/sprout/registry/time"
	"github.com/zygote-sh/zygote/internal/template/tmplfuncs"
)

func funcMap() template.FuncMap {
	handler := sprout.New()
	handler.AddRegistries(
		// Sprout registries
		sptConv.NewRegistry(),
		sptStd.NewRegistry(),
		sptStr.NewRegistry(),
		sptEnc.NewRegistry(),
		sptEnv.NewRegistry(),
		sptFs.NewRegistry(),
		sptMaps.NewRegistry(),
		sptNum.NewRegistry(),
		sptRegex.NewRegistry(),
		sptSemver.NewRegistry(),
		sptSlices.NewRegistry(),
		sptTime.NewRegistry(),
		// Custom registries
		tmplfuncs.NewRegistry(),
	)
	funcMap := handler.Build()
	return funcMap
}
