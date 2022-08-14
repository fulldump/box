package main

import (
	"context"
	"fmt"
	"net/http"

	. "github.com/fulldump/box"
)

func main() {

	b := NewBox()

	b.WithInterceptors(
		func(next H) H {
			return func(ctx context.Context) {
				fmt.Println("A[")
				next(ctx)
				fmt.Println("]A")
			}
		},
		func(next H) H {
			return func(ctx context.Context) {
				fmt.Println("  B[")
				next(ctx)
				fmt.Println("  ]B")
			}
		},
	)

	b.Resource("/hello").
		WithInterceptors(
			func(next H) H {
				return func(ctx context.Context) {
					fmt.Println("    C[")
					next(ctx)
					fmt.Println("    ]C")
				}
			},
			func(next H) H {
				return func(ctx context.Context) {
					fmt.Println("      D[")
					next(ctx)
					fmt.Println("      ]D")
				}
			},
		).
		WithActions(
			Get(func() {
				fmt.Println("hello")
			}).WithInterceptors(
				func(next H) H {
					return func(ctx context.Context) {
						fmt.Println("E[")
						next(ctx)
						fmt.Println("      ]E")
					}
				},
				func(next H) H {
					return func(ctx context.Context) {
						fmt.Println("          F[")
						next(ctx)
						fmt.Println("          ]F")
					}
				},
			),
		)

	b.Resource("/users/{userId}/history").
		WithActions(
			Get(GetUserHistory),
			Post(CreateHistory),
			Action(RevertHistory),
		)

	b.ListenAndServe()
}

func CreateHistory(input string) {
	fmt.Println("Create history!!", input)
	//	fmt.Println("input:", input)
}

func GetUserHistory(ctx context.Context, w http.ResponseWriter) string {
	w.WriteHeader(555)
	userId := GetUrlParameter(ctx, "userId")
	return "Heyyy this is the history for userId " + userId
}

func RevertHistory() {
	fmt.Println("GET HISTORY HANDLER")
}
