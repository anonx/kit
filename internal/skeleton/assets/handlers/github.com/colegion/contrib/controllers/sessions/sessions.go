// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/sessions"

	"github.com/colegion/goal/strconv"
)

// Sessions is an insance of tSessions that is automatically generated from Sessions controller
// being found at "github.com/colegion/contrib/controllers/sessions/sessions.go",
// and contains methods to be used as handler functions.
//
// Sessions is a controller that makes Session field
// available for your actions when you're using this
// controller as a parent.
var Sessions tSessions

// tSessions is a type with handler methods of Sessions controller.
type tSessions struct {
}

// New allocates (github.com/colegion/contrib/controllers/sessions).Sessions controller,
// then returns it.
func (t tSessions) New() *contr.Sessions {
	c := &contr.Sessions{}
	return c
}

// Before is a dump method that always returns nil.
func (t tSessions) Before(c *contr.Sessions, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// After is a dump method that always returns nil.
func (t tSessions) After(c *contr.Sessions, w http.ResponseWriter, r *http.Request) http.Handler {
	return nil
}

// Initially is a method that is started by every handler function at the very beginning
// of their execution phase.
func (t tSessions) Initially(c *contr.Sessions, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	// Call magic Initially method of (github.com/colegion/contrib/controllers/sessions).Sessions.
	return c.Initially(w, r, a)
}

// Finally is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tSessions) Finally(c *contr.Sessions, w http.ResponseWriter, r *http.Request, a []string) (finish bool) {
	// Call magic Finally method of (github.com/colegion/contrib/controllers/sessions).Sessions.
	defer func() {
		if !finish {
			finish = c.Finally(w, r, a)
		}
	}()
	return
}

// Init is used to initialize controllers of "github.com/colegion/contrib/controllers/sessions"
// and its parents.
func Init() {
	initSessions()
	contr.Init()
}

func initSessions() {
}

func init() {
	_ = strconv.MeaningOfLife
}
