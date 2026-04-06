package cli_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	appcli "github.com/PlatformStackPulse/go-template/internal/cli"
	"github.com/PlatformStackPulse/go-template/internal/feature"
	"github.com/PlatformStackPulse/go-template/internal/logger"
	"github.com/PlatformStackPulse/go-template/pkg/version"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	os.Stdout = w

	fn()

	_ = w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()
	return buf.String()
}

func TestRootCommandMetadata(t *testing.T) {
	log := logger.NewLogger(false)
	fm := feature.NewManager()
	cmd := appcli.NewRootCommand(log, fm)

	assert.Equal(t, "go-template", cmd.Use)
	assert.NotNil(t, cmd.PersistentPreRun)
	assert.NotNil(t, cmd.PersistentPostRun)
}

func TestHelloCommandOutput(t *testing.T) {
	log := logger.NewLogger(false)
	fm := feature.NewManager()
	cmd := appcli.NewHelloCommand(log, fm)

	out := captureStdout(t, func() {
		cmd.SetArgs([]string{"--name", "Dev"})
		err := cmd.Execute()
		assert.NoError(t, err)
	})

	assert.True(t, strings.Contains(out, "Hello, Dev!"))
}

func TestVersionCommandOutput(t *testing.T) {
	origVersion := version.Version
	defer func() { version.Version = origVersion }()
	version.Version = "v9.9.9"

	log := logger.NewLogger(false)
	cmd := appcli.NewVersionCommand(log)

	out := captureStdout(t, func() {
		cmd.SetArgs([]string{})
		err := cmd.Execute()
		assert.NoError(t, err)
	})

	assert.True(t, strings.Contains(out, "Version: v9.9.9"))
}
