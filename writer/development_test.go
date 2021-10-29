package writer_test

import (
	"github.com/hyperyuri/logg"
	"github.com/hyperyuri/logg/levels"
	"github.com/hyperyuri/logg/writer"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestWithDebugMode(t *testing.T) {
	logger := logg.NewLogger().
		WithWriter(writer.WithDebugMode()).
		WithDefaultConfig(logg.NewDefaultConfig("team-x", "project-x"))
	logger.SetLevels(levels.NewLevels(logger))

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger.Info("This not will show up!")

	EXPECTED_STRING := time.Now().Format("2006/01/02 15:04:05")  + " development_test.go:24 INFO - This not will show up!\n"

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout


	assert.Equal(t, EXPECTED_STRING, string(out))
}