package core

//
// import (
// 	"fmt"
//
// 	"code.google.com/p/goauth2/oauth"
// 	"github.com/google/go-github/github"
// 	"gopkg.in/libgit2/git2go.v22"
// )
//
// type Repository struct {
// 	client *github.Client
// }
//
// func NewRepository() *Repository {
// 	return new(Repository)
// }
//
// func (r *Repository) Connect(token string) {
//
// 	t := &oauth.Transport{
// 		Token: &oauth.Token{AccessToken: token},
// 	}
//
// 	r.client = github.NewClient(t.Client())
// }
//
// func (r *Repository) Clone(repository string) {
// 	_, err := git.Clone(repository, Directory(), &git.CloneOptions{Bare: true})
// 	if err != nil {
// 		fmt.Println("test")
// 	}
// }
//
// func (r *Repository) Create(name string) {
// 	/*repo, err */ git.InitRepository(Directory(), true)
//
// }
//
// func (r *Repository) Push() {
//
// }
