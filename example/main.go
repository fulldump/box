package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fulldump/box"
	"github.com/fulldump/box/boxopenapi"
	"github.com/fulldump/box/boxutil"
)

func main() {

	b := box.NewBox()

	b.Use(
		box.AccessLog,
		box.RecoverFromPanic,
	)

	orgs := b.Group("/orgs")
	orgs.HandleFunc("GET", "/orgs/{org}/repos", ListOrganizationRepositories)
	orgs.HandleFunc("POST", "/orgs/{org}/repos", CreateOrganizationRepository)

	repos := b.Group("/repos/{owner}/{repo}")
	repos.HandleFunc("POST", "", GetRepository)
	repos.HandleFunc("PATCH", "", UpdateRepository)
	repos.HandleFunc("DELETE", "", DeleteRepository)
	repos.HandleFunc("PUT", "/automated-security-fixes", EnableAutomatedSecurityFixes)
	repos.HandleFunc("DELETE", "/automated-security-fixes", DisableAutomatedSecurityFixes)
	repos.HandleFunc("GET", "/codeowners/errors", ListCodeownersErrors)
	repos.HandleFunc("GET", "/contributors", ListRepositoryContributors)
	repos.HandleFunc("GET", "/languages", ListRepositoryLanguages)
	repos.HandleFunc("GET", "/tags", ListRepositoryTags)
	repos.HandleFunc("GET", "/teams", ListRepositoryTeams)
	repos.HandleFunc("GET", "/topics", ListRepositoryTopics)

	b.HandleFunc("GET", "/repositories", ListPublicRepositories)
	b.HandleFunc("GET", "/users/{username}/repos", ListRepositoriesForAUser)

	user := b.Group("/user").Use(CheckAuthorization)
	user.HandleFunc("GET", "/repos", ListRepositoriesForTheAuthenticatedUser)
	user.HandleFunc("POST", "/repos", CreateRepositoryForTheAuthenticatedUser)

	fmt.Println(boxutil.Tree(b.R))

	spec := boxopenapi.Spec(b)
	spec.Info.Title = "My service"
	spec.Info.Version = "1.0"
	spec.Servers = []boxopenapi.Server{
		{
			Url: "http://localhost:8080",
		},
	}
	b.Handle("GET", "/openapi.json", func() any {
		return spec
	})

	b.ListenAndServe()
}

func ListRepositoryTopics(w http.ResponseWriter, r *http.Request)                               {}
func ListRepositoryTeams(w http.ResponseWriter, r *http.Request)                                {}
func ListRepositoryTags(w http.ResponseWriter, r *http.Request)                                 {}
func ListRepositoryLanguages(w http.ResponseWriter, r *http.Request)                            {}
func ListRepositoryContributors(w http.ResponseWriter, r *http.Request)                         {}
func ListCodeownersErrors(w http.ResponseWriter, r *http.Request)                               {}
func DisableAutomatedSecurityFixes(w http.ResponseWriter, r *http.Request)                      {}
func EnableAutomatedSecurityFixes(w http.ResponseWriter, r *http.Request)                       {}
func ListRepositoriesForAUser(w http.ResponseWriter, r *http.Request)                           {}
func CreateRepositoryForTheAuthenticatedUser(w http.ResponseWriter, r *http.Request)            {}
func ListRepositoriesForTheAuthenticatedUser(writer http.ResponseWriter, request *http.Request) {}
func DeleteRepository(w http.ResponseWriter, r *http.Request)                                   {}
func UpdateRepository(w http.ResponseWriter, r *http.Request)                                   {}
func ListPublicRepositories(w http.ResponseWriter, r *http.Request)                             {}
func CreateOrganizationRepository(w http.ResponseWriter, r *http.Request)                       {}
func GetRepository(w http.ResponseWriter, r *http.Request)                                      {}
func ListOrganizationRepositories(w http.ResponseWriter, r *http.Request)                       {}

func CheckAuthorization(next box.H) box.H {
	return func(ctx context.Context) {
		user, pass, ok := box.GetRequest(ctx).BasicAuth()
		if !ok {
			return
		}
		if user != "admin" {
			return
		}
		if pass != "123456" {
			return
		}
		next(ctx)
	}
}
