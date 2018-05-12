package main

import (
	"context"
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

	client := github.New(context.Background(), token)

	http.HandleFunc("/pullrequests", func(w http.ResponseWriter, r *http.Request) {
		ownerParam := r.URL.Query().Get("owner")
		reposParam := r.URL.Query().Get("repos")

		if ownerParam == "" || reposParam == "" {
			fmt.Fprintln(w, "Missing owner or repos query parmaeter")
			return
		}

		repos := strings.Split(reposParam, ",")

		w.Header().Add("Content-Type", "application/json")
		rr := make([]github.RepositoryData, len(repos))
		wg := sync.WaitGroup{}
		for k, r := range repos {
			wg.Add(1)
			go func(owner, repo string, index int) {
				defer wg.Done()
				rr[index] = client.RepositoryData(owner, repo)
			}(ownerParam, r, k)
		}
		wg.Wait()
		json.NewEncoder(w).Encode(rr)
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "./static/dashboard.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
