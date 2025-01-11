package context_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeiss/ghmd/pkg/context"
	"github.com/zeiss/ghmd/pkg/spec"
)

func TestNew(t *testing.T) {
	t.Setenv("FOO", "BAR")
	t.Setenv("BAR", "1")
	ctx := context.New(spec.Spec{})

	require.Equal(t, "BAR", ctx.Env["FOO"])
	require.Equal(t, "1", ctx.Env["BAR"])
	require.Equal(t, runtime.GOOS, ctx.Runtime.Goos)
	require.Equal(t, runtime.GOARCH, ctx.Runtime.Goarch)
}
