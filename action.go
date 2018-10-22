package box

import "net/http"

// An A stands for Action
type A struct {
	Attr
	HttpMethod string

	name         string
	bound        bool
	resource     *R
	handler      interface{}
	interceptors []I
}

func Action(handler interface{}) *A {

	// TODO: introspect and change to "POST" if needed

	return &A{
		HttpMethod:   "GET",
		Attr:         Attr{},
		handler:      handler,
		interceptors: []I{},
	}
}

func ActionPost(handler interface{}) *A {

	return &A{
		HttpMethod:   "POST",
		Attr:         Attr{},
		handler:      handler,
		interceptors: []I{},
	}
}

func actionBound(handler interface{}, method string) *A {
	a := Action(handler)
	a.bound = true
	a.HttpMethod = method
	return a
}

// Bind shortcuts:
func Get(handler interface{}) *A {
	return actionBound(handler, http.MethodGet)
}
func Head(handler interface{}) *A {
	return actionBound(handler, http.MethodHead)
}
func Post(handler interface{}) *A {
	return actionBound(handler, http.MethodPost)
}
func Put(handler interface{}) *A {
	return actionBound(handler, http.MethodPut)
}
func Patch(handler interface{}) *A {
	return actionBound(handler, http.MethodPatch)
}
func Delete(handler interface{}) *A {
	return actionBound(handler, http.MethodDelete)
}
func Connect(handler interface{}) *A {
	return actionBound(handler, http.MethodConnect)
}
func Options(handler interface{}) *A {
	return actionBound(handler, http.MethodOptions)
}
func Trace(handler interface{}) *A {
	return actionBound(handler, http.MethodTrace)
}

// WithName overwrite default action name
func (a *A) WithName(name string) *A {
	a.name = name
	return a
}

func (a *A) Bind(method string) *A {
	a.HttpMethod = method
	a.bound = true
	return a
}

func (a *A) WithAttribute(key string, value interface{}) *A {
	a.SetAttribute(key, value)
	return a
}

func (a *A) WithInterceptors(interceptor ...I) *A {

	for _, i := range interceptor {
		a.interceptors = append(a.interceptors, i)
	}

	return a
}
