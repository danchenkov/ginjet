package main

import (
	"fmt"
	"log"
	"reflect"

	jet "github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./views"),
)

type tTODO struct {
	Text string
	Done bool
}

type doneTODOs struct {
	list map[string]*tTODO
	keys []string
	len  int
	i    int
}

func (dt *doneTODOs) New(todos map[string]*tTODO) *doneTODOs {
	dt.len = len(todos)
	for k := range todos {
		dt.keys = append(dt.keys, k)
	}
	dt.list = todos
	return dt
}

// Range satisfies the jet.Ranger interface and only returns TODOs that are done,
// even when the list contains TODOs that are not done.
func (dt *doneTODOs) Range() (reflect.Value, reflect.Value, bool) {
	for dt.i < dt.len {
		key := dt.keys[dt.i]
		dt.i++
		if dt.list[key].Done {
			return reflect.ValueOf(key), reflect.ValueOf(dt.list[key]), false
		}
	}
	return reflect.Value{}, reflect.Value{}, true
}

func (dt *doneTODOs) ProvidesIndex() bool { return true }

// Render implements jet.Renderer interface
func (t *tTODO) Render(r *jet.Runtime) {
	done := "yes"
	if !t.Done {
		done = "no"
	}
	r.Write([]byte(fmt.Sprintf("TODO: %s (done: %s)", t.Text, done)))
}

func homeIndex(c *gin.Context) {
	view, err := views.GetTemplate("users/index.jet")
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
	}

	var todos = map[string]*tTODO{
		"example-todo-1": {Text: "Add an show todo page to the example project", Done: true},
		"example-todo-2": {Text: "Add an add todo page to the example project"},
		"example-todo-3": {Text: "Add an update todo page to the example project"},
		"example-todo-4": {Text: "Add an delete todo page to the example project", Done: true},
	}

	view.Execute(c.Writer, nil, todos)
}
