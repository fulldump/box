package box

import (
	"net/url"
	"reflect"
	"runtime"
	"strings"
)

// R stands for Resource
type R struct {
	Attr

	path     string
	parent   *R
	children []*R

	// actions contains all actions (bound and unbound)
	actionsByName map[string]*A

	// bound contains only bound actions
	actionsByHttp map[string]*A

	// interceptors ...
	interceptors []I
}

func NewResource() *R {
	return &R{
		Attr:          Attr{},
		path:          "",
		children:      []*R{},
		actionsByName: map[string]*A{},
		actionsByHttp: map[string]*A{},
		interceptors:  []I{},
	}
}

// TODO: maybe parameters should be a type `P`
func (r *R) Match(path string, parameters map[string]string) (result *R) {

	parts := strings.SplitN(path, "/", 2)

	current, err := url.PathUnescape(parts[0])
	if nil != err {
		// TODO: maybe log debug "unescape" error ??
		return
	}

	if strings.HasPrefix(r.path, "{") && strings.HasSuffix(r.path, "}") {
		// Match with pattern
		parameter := r.path
		parameter = strings.TrimPrefix(parameter, "{")
		parameter = strings.TrimSuffix(parameter, "}")
		parameters[parameter] = current
	} else if r.path == current {
		// Exact match
	} else if r.path == "*" {
		// Delegation match
		parameters["*"] = path
		return r
	} else {
		// No match
		return nil // TODO: maybe return error no match ?¿?¿?
	}

	if len(parts) == 1 {
		return r
	}

	for _, c := range r.children {
		if result := c.Match(parts[1], parameters); nil != result {
			return result
		}
	}

	return
}

func (r *R) resourceParts(parts []string) *R {

	if len(parts) == 0 {
		return r
	}

	part := parts[0]
	for _, child := range r.children {
		if child.path == part {
			return child.resourceParts(parts[1:])
		}
	}

	child := NewResource()
	child.path = part
	child.parent = r
	r.children = append(r.children, child)

	return child.resourceParts(parts[1:])
}

// Resource defines a new resource below current resource
func (r *R) Resource(locator string) *R {

	locator = strings.TrimPrefix(locator, "/")
	if locator == "" {
		return r
	}

	parts := strings.Split(locator, "/")

	return r.resourceParts(parts)
}

// Add action to this resource
func (r *R) WithActions(action ...*A) *R {

	for _, a := range action {
		if "" == a.name {
			a.name = getFunctionName(a.handler)
			a.name = actionNameNormalizer(a.name)
		}
		a.resource = r
		r.actionsByName[a.name] = a

		h := a.HttpMethod + " "
		if !a.bound {
			h += a.name
		}
		r.actionsByHttp[h] = a
	}

	return r
}

// Add interceptor to this resource
func (r *R) WithInterceptors(interceptor ...I) *R {

	for _, i := range interceptor {
		r.interceptors = append(r.interceptors, i)
	}

	return r
}

func (r *R) WithAttribute(key string, value interface{}) *R {
	r.SetAttribute(key, value)
	return r
}

func actionNameNormalizer(u string) string {
	if len(u) == 0 {
		return u
	}

	return strings.ToLower(u[0:1]) + u[1:]
}

func getFunctionName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	parts := strings.Split(name, ".")

	return parts[1]
}
