// Package views is a subcommand of generate that is
// used for scanning directories and composing a list
// of found files. It may be useful to statically check
// names of used templates.
// I.e. instead of using paths of templates in actions
// directely like `RenderTemplate("path/to/template.html")`,
// it is possible to do something like
// `RenderTemplate(views.Path.To.TemplateHTML)`.
package views

import (
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/anonx/sunplate/command"
	"github.com/anonx/sunplate/generation/output"
	"github.com/anonx/sunplate/log"
	p "github.com/anonx/sunplate/path"
)

var fileNamePattern = regexp.MustCompile(
	"^[A-Za-z]{1}\\w*[.]{0,1}\\w*$",
)

// Start is an entry point of listing subcommand.
// It expects two parameters.
// params is a map with the following keys:
// --input defines what directory to analyze ("./views" by-default).
// --output is a path to directory where to create a new package ("./assets/views" by-default).
// --package is what package should be created as a result ("views" by-default).
func Start(params command.Data) {
	inputDir := params.Default("--input", "./views")
	outputDir := params.Default("--output", "./assets/views")
	outPkg := params.Default("--package", "views")

	// Start search of files.
	fs, fn := walkFunc(inputDir)
	filepath.Walk(inputDir, fn)

	// Generate and save a new package.
	t := output.NewType(
		outPkg, filepath.Join(
			p.SunplateDir("generation", "views"), "./views.go.template",
		),
	)
	t.CreateDir(outputDir)
	t.Extension = ".go" // Save generated file as a .go source.
	t.Context = map[string]interface{}{
		"listing": fs,
		"input":   inputDir,
	}
	t.Generate()
}

// walkFunc returns a files listing and a function that may be used for validation
// of found files. Successfully validated ones are stored to the listing variable.
func walkFunc(dir string) (listing, func(string, os.FileInfo, error) error) {
	l := listing{}
	dir = p.Prefixless(dir, "./") // No "./" is allowed at the beginning.

	return l, func(path string, info os.FileInfo, err error) error {
		// Make sure there are no any errors.
		if err != nil {
			log.Warn.Printf(`An error occured while creating a listing: "%s".`, err)
			return err
		}

		// Make sure file name is of supported format.
		ss := strings.Split(path, "/")
		for _, s := range ss {
			if !fileNamePattern.MatchString(s) {
				log.Warn.Printf(`"%s" is ignored as "%s" is of unsupported format.`, path, s)
				return fmt.Errorf(`"%s" is of unsupported type`, s)
			}

			if !ast.IsExported(s) {
				log.Trace.Printf(`"%s" will not be exported as "%s" starts with a lower case letter.`, path, s)
			}
		}

		// Get filepath without the dir path at the beginning.
		// So, when we are scanning "views/app/index.html" our generated
		// result will be "app/index.html" instead.
		rel, _ := filepath.Rel(dir, path)

		// Add the directory to the list (if it is a dir).
		if info.IsDir() {
			l.addDir(rel)
			return nil
		}

		// Otherwise, register the file.
		l.addFile(rel)
		return nil
	}
}