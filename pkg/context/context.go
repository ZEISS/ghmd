package context

import (
	stdctx "context"
	"os"
	"runtime"
	"time"

	"github.com/zeiss/ghmd/pkg/spec"
	"github.com/zeiss/pkg/envx"
)

// GitInfo includes tags and diffs used in some point.
type GitInfo struct {
	Branch      string
	CurrentTag  string
	PreviousTag string
	Commit      string
	ShortCommit string
	FullCommit  string
	FirstCommit string
	CommitDate  time.Time
	URL         string
	Summary     string
	TagSubject  string
	TagContents string
	TagBody     string
	Dirty       bool
}

type Runtime struct {
	Goos   string
	Goarch string
}

// Semver represents a semantic version.
type Semver struct {
	Major      uint64
	Minor      uint64
	Patch      uint64
	Prerelease string
}

// Context is a context for the application.
type Context struct {
	stdctx.Context
	Spec    spec.Spec
	GitInfo GitInfo
	Runtime Runtime
	Semver  Semver
	Env     envx.Env
	Date    time.Time
	Skips   map[string]bool
}

// New context.
func New(spec spec.Spec) *Context {
	return Wrap(stdctx.Background(), spec)
}

// Wrap context.
func Wrap(ctx stdctx.Context, spec spec.Spec) *Context {
	return &Context{
		Context: ctx,
		Spec:    spec,
		Env:     envx.ToEnv(os.Environ()),
		Date:    time.Now(),
		Skips:   map[string]bool{},
		Runtime: Runtime{
			Goos:   runtime.GOOS,
			Goarch: runtime.GOARCH,
		},
	}
}
