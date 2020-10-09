package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/felipecaputo/github-changelog-extractor/pkg/githubclient"
)

func main() {
	client := githubclient.NewClient(os.Getenv("GITHUB_TOKEN"))
	if query, err := client.GetClosedMilestones("felipecaputo", "git-project-manager", ""); err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		data, _ := json.MarshalIndent(query, "", "\t")
		fmt.Printf("Found: \n %s", string(data))
	}
}
