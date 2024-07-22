package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	name, numRepos, err := getGithubInfo("sergiolpf")
	fmt.Printf("Name: %v, Num_repos: %v, err: %v", name, numRepos, err)
}

func getGithubInfo(username string) (string, int, error) {
	resp, err := http.Get("https://api.github.com/users/" + url.PathEscape(username))

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	r := Reply{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	return r.Name, r.Public_Repos, nil
}

type Reply struct {
	Name         string
	Public_Repos int
}
