package boxopenapi

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/fulldump/box"
)

func newApiExample() *box.B {

	b := box.NewBox()

	b.Handle("GET", "/users", ListUsers)
	b.Handle("POST", "/users", CreateUser)
	b.Handle("GET", "/users/{userId}", GetUser)

	b.Handle("GET", "/recursive", ListRecursiveObjects)
	b.Handle("GET", "/scalarValues", GetScalarValues)
	b.Handle("GET", "/time.Time", GetTypeTime)

	b.Handle("GET", "/hidden-action", func() {}).SetAttribute("openapi", false)

	b.Group("/hidden-resource").SetAttribute("openapi", false)
	b.Handle("GET", "/hidden-resource", func() {})

	return b
}

type User struct {
	Id     string   `json:"id" description:"User identifier"`
	Name   string   `json:"name" description:"User name"`
	Tags   []string `json:"tags,omitempty" description:"User tags"`
	Age    int      `json:"age" description:"User age"`
	Active bool     `json:"active" description:"User active"`
	AnonymousFields
}

type AnonymousFields struct {
	CreationDate time.Time `json:"creation_date" description:"Creation date"`
	ModifiedDate time.Time `json:"modifiedd_date" description:"Modification date"`
}

func ListUsers() []*User {
	return nil
}

type RecursiveObject struct {
	RecursiveObject *RecursiveObject   `description:"my recursive object"`
	RecursiveList   []*RecursiveObject `description:"my recursive list"`
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

func ListRecursiveObjects() *RecursiveObject {
	return nil
}

type ScalarsValues struct {
	MyString  string  `description:"my String description"`
	MyBool    bool    `description:"my Bool description"`
	MyInt     int     `description:"my Int description"`
	MyInt64   int64   `description:"my Int64 description"`
	MyInt32   int32   `description:"my Int32 description"`
	MyInt16   int16   `description:"my Int16 description"`
	MyInt8    int8    `description:"my Int8 description"`
	MyUint    uint    `description:"my Uint description"`
	MyUint64  uint64  `description:"my Uint64 description"`
	MyUint32  uint32  `description:"my Uint32 description"`
	MyUint16  uint16  `description:"my Uint16 description"`
	MyUint8   uint8   `description:"my Uint8 description"`
	MyFloat64 float64 `description:"my Float64 description"`
	MyFloat32 float32 `description:"my Float32 description"`
}

func GetScalarValues() *ScalarsValues {
	return nil
}

type TypeTime struct {
	Today time.Time `description:"the time for today :D"`
}

func GetTypeTime() *TypeTime {
	return nil
}

// https://editor-next.swagger.io/
func TestOpenApi(t *testing.T) {

	b := newApiExample()

	result := Spec(b)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "    ")
	e.Encode(result)

	expected := OpenAPI{
		Components: JSON{
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
				"RecursiveObject": JSON{
					"properties": JSON{
						"RecursiveList": JSON{
							"description": "my recursive list",
							"items": JSON{
								"$ref": "#/components/schemas/RecursiveObject",
							},
							"type": "array",
						},
						"RecursiveObject": JSON{
							"$ref":        "#/components/schemas/RecursiveObject",
							"description": "my recursive object",
						},
					},
					"required": []JSON{},
					"type":     "object",
				},
				"ScalarsValues": JSON{
					"properties": JSON{
						"MyBool": JSON{
							"description": "my Bool description",
							"type":        "boolean",
						},
						"MyFloat32": JSON{
							"description": "my Float32 description",
							"type":        "number",
						},
						"MyFloat64": JSON{
							"description": "my Float64 description",
							"type":        "number",
						},
						"MyInt": JSON{
							"description": "my Int description",
							"type":        "number",
						},
						"MyInt16": JSON{
							"description": "my Int16 description",
							"type":        "number",
						},
						"MyInt32": JSON{
							"description": "my Int32 description",
							"type":        "number",
						},
						"MyInt64": JSON{
							"description": "my Int64 description",
							"type":        "number",
						},
						"MyInt8": JSON{
							"description": "my Int8 description",
							"type":        "number",
						},
						"MyString": JSON{
							"description": "my String description",
							"type":        "string",
						},
						"MyUint": JSON{
							"description": "my Uint description",
							"type":        "number",
						},
						"MyUint16": JSON{
							"description": "my Uint16 description",
							"type":        "number",
						},
						"MyUint32": JSON{
							"description": "my Uint32 description",
							"type":        "number",
						},
						"MyUint64": JSON{
							"description": "my Uint64 description",
							"type":        "number",
						},
						"MyUint8": JSON{
							"description": "my Uint8 description",
							"type":        "number",
						},
					},
					"required": []JSON{},
					"type":     "object",
				},
				"TypeTime": JSON{
					"properties": JSON{
						"Today": JSON{
							"description": "the time for today :D",
							"examples": []any{
								"2006-01-02T15:04:05Z07:00",
							},
							"format": "date-time",
							"type":   "string",
						},
					},
					"required": []any{},
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
						"creation_date": JSON{
							"description": "Creation date",
							"examples": []string{
								"2006-01-02T15:04:05Z07:00",
							},
							"format": "date-time",
							"type":   "string",
						},
						"modifiedd_date": JSON{
							"description": "Modification date",
							"examples": []string{
								"2006-01-02T15:04:05Z07:00",
							},
							"format": "date-time",
							"type":   "string",
						},
					},
					"required": []JSON{},
					"type":     "object",
				},
			},
		},
		Info: Info{
			Title:   "BoxOpenAPI",
			Version: "1",
		},
		Openapi: "3.1.0",
		Paths: JSON{
			"/recursive": JSON{
				"get": JSON{
					"operationId": "listRecursiveObjects",
					"responses": JSON{
						"default": JSON{
							"content": JSON{
								"application/json": JSON{
									"schema": JSON{
										"$ref": "#/components/schemas/RecursiveObject",
									},
								},
							},
							"description": "some human description",
						},
					},
				},
			},
			"/scalarValues": JSON{
				"get": JSON{
					"operationId": "getScalarValues",
					"responses": JSON{
						"default": JSON{
							"content": JSON{
								"application/json": JSON{
									"schema": JSON{
										"$ref": "#/components/schemas/ScalarsValues",
									},
								},
							},
							"description": "some human description",
						},
					},
				},
			},
			"/time.Time": JSON{
				"get": JSON{
					"operationId": "getTypeTime",
					"responses": JSON{
						"default": JSON{
							"content": JSON{
								"application/json": JSON{
									"schema": JSON{
										"$ref": "#/components/schemas/TypeTime",
									},
								},
							},
							"description": "some human description",
						},
					},
				},
			},
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
		Servers: []Server{
			{
				Url: "http://localhost:8080",
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
