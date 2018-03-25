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
		result := make([]github.Repository, len(allRepos))

		var wg sync.WaitGroup
		for i, repo := range allRepos {
			r, o, i := repo, owner, i
			wg.Add(1)
			go func(string, string, int) {
				defer wg.Done()
				repoData, err := client.GetRepositoryData(o, r)
				if err != nil {
					fmt.Println(w, err)
				}
				result[i] = *repoData
			}(o, r, i)
		}
		wg.Wait()

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "./static/dashboard.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
