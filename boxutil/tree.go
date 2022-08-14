package boxutil

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strings"

	"github.com/fulldump/box"
)

// Tree return a string representation of resources hierarchy
func Tree(root *box.R) (s string) {

	traverse(root, func(r *box.R) {

		actions := r.GetActions()

		sortActions(actions)

		if len(r.Interceptors) == 0 && len(actions) == 0 {
			return
		}

		s += getFullPath(r) + "\n"

		for _, i := range r.Interceptors {
			s += fmt.Sprintf("    <%s>\n", getFunctionName(i))
		}

		for _, a := range actions {
			s += fmt.Sprintf("    %s", a.HttpMethod)

			if !a.Bound {
				s += fmt.Sprintf(" :%s", a.Name)
			}

			s += " "

			for _, i := range a.Interceptors {
				s += fmt.Sprintf("<%s>", getFunctionName(i))
			}

			s += "\n"
		}
	})

	return
}

func traverse(r *box.R, f func(r *box.R)) {
	f(r)
	for _, c := range r.Children {
		traverse(c, f)
	}
}

func getFullPath(r *box.R) (s string) {
	s = r.Path
	for {
		if r.Parent == nil {
			return
		}
		s = r.Parent.Path + "/" + s
		r = r.Parent
	}
}

func sortActions(actions []*box.A) {
	// Sort actions:
	// First bound (by method between them)
	// Second unbound (by name between them)
	sort.Slice(actions, func(i, j int) bool {
		a := actions[i]
		b := actions[j]

		if a.Bound {
			if b.Bound {
				return a.HttpMethod < b.HttpMethod
			}
			return true
		} else {
			if b.Bound {
				return false
			}
			return a.Name < b.Name
		}
	})
}

func getFunctionName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	parts := strings.Split(name, ".")

	return parts[len(parts)-1]
}
