package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/nenadstojanovikj/ghremind/pkg/github"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Println("Missing GITHUB_TOKEN, can only serve public repositories")
	}
	client := github.NewClient(token)

	http.HandleFunc("/pullrequests", func(w http.ResponseWriter, r *http.Request) {
		owner := r.URL.Query().Get("owner")
		repos := r.URL.Query().Get("repos")

		if owner == "" || repos == "" {
			fmt.Fprintln(w, "Missing owner or repos query parmaeter")
			return
		}

		allRepos := strings.Split(repos, ",")
		allPRs := make(map[string][]github.PullRequest)

		var wg sync.WaitGroup
		for _, repo := range allRepos {
			r, o := repo, owner
			wg.Add(1)
			go func(string, string) {
				defer wg.Done()
				prs, err := client.GetPullRequests(o, r)
				if err != nil {
					fmt.Println(w, err)
				}
				allPRs[r] = prs
			}(o, r)
		}
		wg.Wait()

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(allPRs)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
