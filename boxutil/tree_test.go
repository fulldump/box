package boxutil

import (
	"context"
	"fmt"
	"testing"

	"github.com/fulldump/box"
)

func TestTree(t *testing.T) {

	root := box.NewResource()
	v1 := root.Resource("/api/v1").
		WithInterceptors(Auth, AccessLog)
	v1.Resource("/users").
		WithActions(
			box.Get(ListUsers).WithInterceptors(AdminRequired, NoAudit),
			box.Post(CreateUser).WithInterceptors(AdminRequired),
		)
	v1.Resource("/users/{userId}").
		WithActions(
			box.Get(GetUser),
			box.Delete(DeleteUser),
			box.ActionPost(BanUser).WithInterceptors(AdminRequired),
			box.ActionPost(EnableUser).WithInterceptors(AdminRequired),
			box.ActionPost(DisableUser).WithInterceptors(AdminRequired),
		)
	v1.Resource("/spots/{spotId}/menus/{menuId}").
		WithActions(
			box.Delete(DeleteMenu).WithInterceptors(AdminRequired, NoAudit),
			box.Patch(UpdateMenu),
		)

	s := Tree(root)

	expected := `/api/v1
    <Auth>
    <AccessLog>
/api/v1/users
    GET <AdminRequired><NoAudit>
    POST <AdminRequired>
/api/v1/users/{userId}
    DELETE 
    GET 
    POST :banUser <AdminRequired>
    POST :disableUser <AdminRequired>
    POST :enableUser <AdminRequired>
/api/v1/spots/{spotId}/menus/{menuId}
    DELETE <AdminRequired><NoAudit>
    PATCH 
`

	if s != expected {
		t.Errorf("Tree output does not match\nExpected: %s\nObtained:%s\n", expected, s)
	}

}

func CreateUser()  {}
func GetUser()     {}
func BanUser()     {}
func DisableUser() {}
func EnableUser()  {}
func DeleteUser()  {}
func ListUsers()   {}
func DeleteMenu()  {}
func UpdateMenu()  {}

func Auth(next box.H) box.H {
	return func(ctx context.Context) {
		next(ctx)
	}
}

func AccessLog(next box.H) box.H {
	return func(ctx context.Context) {
		next(ctx)
	}
}

func AdminRequired(next box.H) box.H {
	return func(ctx context.Context) {
		next(ctx)
	}
}

func NoAudit(next box.H) box.H {
	return func(ctx context.Context) {
		next(ctx)
	}
}

func TestSortActions(t *testing.T) {

	actions := []*box.A{
		{Name: "D", Bound: false},
		{Name: "C", Bound: false},
		{Name: "B", Bound: true, HttpMethod: "B"},
		{Name: "A", Bound: true, HttpMethod: "A"},
	}

	fmt.Println(actions)
	sortActions(actions)
	fmt.Println(actions)

	if actions[0].Name != "A" {
		t.Errorf("A should be action[0]")
	}
	if actions[1].Name != "B" {
		t.Errorf("B should be action[1]")
	}
	if actions[2].Name != "C" {
		t.Errorf("C should be action[2]")
	}
	if actions[3].Name != "D" {
		t.Errorf("D should be action[3]")
	}
}
