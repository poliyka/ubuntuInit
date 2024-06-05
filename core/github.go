package core

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"log"
)

func GetGithubRepLatestRelease(owner string, repo string) string {
	ctx := context.Background()
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		log.Fatalf("Error fetching latest release: %v", err)
	}

	fmt.Printf("The latest nvm release is %s\n", *release.TagName)
	return *release.TagName
}
