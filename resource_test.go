package box

import (
	"fmt"
	"testing"
)

func TestR_ResourceParts_BaseCase(t *testing.T) {

	root := NewResource()
	r := root.resourceParts([]string{})

	AssertEqual(t, r, root)
}

func TestR_ResourceParts_NextCase(t *testing.T) {

	root := NewResource()
	a := root.resourceParts([]string{"a"})
	b := root.resourceParts([]string{"a", "b"})

	AssertEqual(t, a.Parent, root)
	AssertEqual(t, b.Parent, a)
}

func TestR_Resource_BaseCase(t *testing.T) {

	root := NewResource()
	r := root.Resource("")

	AssertEqual(t, r, root)
}

func TestR_Resource_BaseCaseSlash(t *testing.T) {

	root := NewResource()
	r := root.Resource("/")

	AssertEqual(t, r, root.Children[0])
}

func TestR_Resource_NextCase(t *testing.T) {

	root := NewResource()

	a := root.Resource("/a")
	b := root.Resource("/a/b")

	AssertEqual(t, a.Parent, root)
	AssertEqual(t, b.Parent, a)
}

func TestR_Match_DecodeUriComponent(t *testing.T) {

	r := NewResource()
	r.Resource("/users/{userId}/history")

	parameters := map[string]string{}

	r.Match("/users/a+b%20c/history", parameters)

	AssertEqual(t, parameters["userId"], "a b c")
}

func TestR_Match(t *testing.T) {

	r := NewResource()
	r.Resource("/users/{userId}/history")

	parameters := map[string]string{}

	history := r.Match("/users/Fulanito/history", parameters)

	fmt.Println(history, parameters)
}
