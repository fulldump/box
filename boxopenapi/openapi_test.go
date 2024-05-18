package boxopenapi

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/fulldump/box"
)

func newApiExample() *box.B {

	b := box.NewBox()

	b.Handle("GET", "/users", ListUsers)
	b.Handle("POST", "/users", CreateUser)

	b.Handle("GET", "/users/{userId}", GetUser)

	return b
}

type User struct {
	Id     string   `json:"id" description:"User identifier"`
	Name   string   `json:"name" description:"User name"`
	Tags   []string `json:"tags" description:"User tags"`
	Age    int      `json:"age" description:"User age"`
	Active bool     `json:"active" description:"User active"`
}

func ListUsers() []*User {
	return nil
}

type CreateUserInput struct {
	Id   string `json:"id" description:"If empty a random uuid will be generated"`
	Name string `json:"name"`
}

func CreateUser(input *CreateUserInput) *User {
	return &User{
		Id:   input.Id,
		Name: input.Name,
	}
}

func GetUser() *User {
	return nil
}

// https://editor-next.swagger.io/
func TestOpenApi(t *testing.T) {

	b := newApiExample()

	result := Spec(b)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "    ")
	e.Encode(result)

	expected := JSON{
		"components": JSON{
			"schemas": JSON{
				"CreateUserInput": JSON{
					"properties": JSON{
						"id": JSON{
							"description": "If empty a random uuid will be generated",
							"type":        "string",
						},
						"name": JSON{
							"type": "string",
						},
					},
					"required": []JSON{},
					"type":     "object",
				},
				"User": JSON{
					"properties": JSON{
						"active": JSON{
							"description": "User active",
							"type":        "boolean",
						},
						"age": JSON{
							"description": "User age",
							"type":        "number",
						},
						"id": JSON{
							"description": "User identifier",
							"type":        "string",
						},
						"name": JSON{
							"description": "User name",
							"type":        "string",
						},
						"tags": JSON{
							"description": "User tags",
							"items": JSON{
								"type": "string",
							},
							"type": "array",
						},
					},
					"required": []JSON{},
					"type":     "object",
				},
			},
		},
		"info": JSON{
			"title":   "config",
			"version": "1",
		},
		"openapi": "3.0.0",
		"paths": JSON{
			"/users": JSON{
				"get": JSON{
					"operationId": "listUsers",
					"responses": JSON{
						"default": JSON{
							"content": JSON{
								"application/json": JSON{
									"schema": JSON{
										"items": JSON{
											"$ref": "#/components/schemas/User",
										},
										"type": "array",
									},
								},
							},
							"description": "some human description",
						},
					},
				},
				"post": JSON{
					"operationId": "createUser",
					"requestBody": JSON{
						"content": JSON{
							"application/json": JSON{
								"schema": JSON{
									"$ref": "#/components/schemas/CreateUserInput",
								},
							},
						},
						"description": "TODO",
						"required":    true,
					},
					"responses": JSON{
						"default": JSON{
							"content": JSON{
								"application/json": JSON{
									"schema": JSON{
										"$ref": "#/components/schemas/User",
									},
								},
							},
							"description": "some human description",
						},
					},
				},
			},
			"/users/{userId}": JSON{
				"get": JSON{
					"operationId": "getUser",
					"parameters": []JSON{
						{
							"in":       "path",
							"name":     "userId",
							"required": true,
							"schema": JSON{
								"type": "string",
							},
						},
					},
					"responses": JSON{
						"default": JSON{
							"content": JSON{
								"application/json": JSON{
									"schema": JSON{
										"$ref": "#/components/schemas/User",
									},
								},
							},
							"description": "some human description",
						},
					},
				},
			},
		},
		"servers": []JSON{
			{
				"url": "https://config.hola.cloud",
			},
		},
	}

	if !equalJson(result, expected) {
		t.Error("Output does not match the expected spec, see diff")
		e.Encode(expected)
	}

}

func equalJson(a, b any) bool {

	aserial, _ := json.Marshal(a)
	bserial, _ := json.Marshal(b)

	var aobject, bobject any
	_ = json.Unmarshal(aserial, &aobject)
	_ = json.Unmarshal(bserial, &bobject)

	return reflect.DeepEqual(aobject, bobject)
}
