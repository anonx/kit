package handlers

import (
	"os"
	"os/exec"
	"testing"

	"github.com/colegion/goal/utils/tool"
)

func TestStart(t *testing.T) {
	main(handlers, 0, tool.Data{})

	cmd := exec.Command("go", "install", "github.com/colegion/goal/tools/generate/handlers/testdata/assets/handlers")
	cmd.Stderr = os.Stderr // Show the output of the program we run.
	if err := cmd.Run(); err != nil {
		t.Errorf(`There are problems with generated handlers, error: "%s".`, err)
	}

	// Remove the directory we have created.
	os.RemoveAll(*output)
}

var handlers []tool.Handler

func init() {
	Handler.Flags.Set("input", "./testdata/controllers")
	Handler.Flags.Set("output", "./testdata/assets/handlers")

	handlers = []tool.Handler{Handler}
}
