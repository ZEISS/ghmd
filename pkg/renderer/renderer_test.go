package renderer_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeiss/ghmd/pkg/renderer"
	"github.com/zeiss/ghmd/pkg/spec"
)

func TestHeader(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		spec *spec.Spec
		want string
	}{
		{
			name: "empty spec",
			spec: &spec.Spec{},
			want: "",
		},
		{
			name: "spec with header",
			spec: &spec.Spec{
				Header: `# Header`,
			},
			want: `# Header`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := renderer.New(tt.spec)
			var buf bytes.Buffer
			err := r.Render(&buf)
			require.NoError(t, err)
			require.Equal(t, tt.want, buf.String())
		})
	}
}
