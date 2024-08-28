package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/v49/github"
	"github.com/joho/godotenv"
	"github.com/mattn/go-isatty"
	"golang.org/x/oauth2"
	"golang.org/x/term"
)

// init function to load environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// displayBanner prints the banner on the console
func displayBanner() {
	fmt.Println()
	color.Yellow(`
	   █████████  ████                                 █████████   ████  ████ 
	  ███░░░░░███░░███                                ███░░░░░███ ░░███ ░░███ 
	 ███     ░░░  ░███   ██████  ████████    ██████  ░███    ░███  ░███  ░███ 
	░███          ░███  ███░░███░░███░░███  ███░░███ ░███████████  ░███  ░███ 
	░███          ░███ ░███ ░███ ░███ ░███ ░███████  ░███░░░░░███  ░███  ░███ 
	░░███     ███ ░███ ░███ ░███ ░███ ░███ ░███░░░   ░███    ░███  ░███  ░███ 
	 ░░█████████  █████░░██████  ████ █████░░██████  █████   █████ █████ █████
	  ░░░░░░░░░  ░░░░░  ░░░░░░  ░░░░ ░░░░░  ░░░░░░  ░░░░░   ░░░░░ ░░░░░ ░░░░░ 
	`)
	fmt.Println()
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatalf("GitHub token not found in .env file")
	}

	displayBanner()

	ctx := context.Background()
	client := createGitHubClient(ctx, token)

	for {
		username := promptUsername()
		if username == "" {
			break
		}

		repos := fetchRepositories(ctx, client, username)
		if len(repos) == 0 {
			color.Red("No repositories found for user %s", username)
			continue
		}

		choice := promptRepositorySelection(repos)
		if choice == 0 {
			continue
		}

		cloneSelectedRepositories(username, repos, choice)
	}
}

// createGitHubClient initializes and returns a GitHub client using OAuth2
func createGitHubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

// promptUsername asks the user for a GitHub username and returns it
func promptUsername() string {
	reader := bufio.NewReader(os.Stdin)
	color.Yellow("Enter GitHub username (or type 'exit' to quit): ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if strings.ToLower(username) == "exit" {
		color.Yellow("Exiting the program.")
		return ""
	}

	return username
}

// fetchRepositories retrieves all repositories for the given GitHub username
func fetchRepositories(ctx context.Context, client *github.Client, username string) []*github.Repository {
	var allRepos []*github.Repository
	opts := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 50}}

	for {
		repos, resp, err := client.Repositories.List(ctx, username, opts)
		if err != nil {
			log.Fatalf("Error fetching repositories: %v", err)
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allRepos
}

// promptRepositorySelection presents the repository list and captures the user's choice
func promptRepositorySelection(repos []*github.Repository) int {
	color.Yellow("Select repository to clone:")
	color.Red("0. Exit")
	color.Green("1. All")

	printRepoList(repos)

	reader := bufio.NewReader(os.Stdin)
	color.Yellow("Enter your choice (number): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choiceNum, _ := strconv.Atoi(choice)

	if choiceNum < 0 || choiceNum > len(repos)+1 {
		color.Red("Invalid choice")
		return 0
	}

	return choiceNum
}

// cloneSelectedRepositories clones repositories based on the user's choice
func cloneSelectedRepositories(username string, repos []*github.Repository, choice int) {
	dirName := fmt.Sprintf("%s's Repo", username)
	if err := os.Mkdir(dirName, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("Error creating directory: %v", err)
	}

	switch choice {
	case 1:
		for _, repo := range repos {
			cloneRepository(dirName, *repo.CloneURL)
		}
	default:
		cloneRepository(dirName, *repos[choice-2].CloneURL)
	}
}

// printRepoList prints the repository list in a column format
func printRepoList(repos []*github.Repository) {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		for i, repo := range repos {
			color.Green("%d. %s", i+2, *repo.Name)
		}
		return
	}

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	maxLength := 0
	for _, repo := range repos {
		if len(*repo.Name) > maxLength {
			maxLength = len(*repo.Name)
		}
	}
	columnWidth := maxLength + 5
	columns := width / columnWidth

	for i, repo := range repos {
		fmt.Printf("%s. %s",
			color.GreenString("%d", i+2),
			color.GreenString("%-*s", maxLength, *repo.Name))

		if (i+1)%columns == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// cloneRepository executes the git clone command for the given repository
func cloneRepository(dirName, cloneURL string) {
	color.Blue("Cloning %s...", cloneURL)

	cmd := exec.Command("git", "-c", "http.postBuffer=524288000", "clone", cloneURL)
	cmd.Dir = dirName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := retryOnError(cmd.Run, 5); err != nil {
		log.Fatalf("Failed to clone repository after 5 attempts: %v", err)
	}

	color.Magenta("Clone completed.")
}

// retryOnError retries the given function up to maxRetries times if it returns an error.
// It returns the last error if all retries fail.
func retryOnError(f func() error, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		if err = f(); err == nil {
			return nil
		}
		color.Red("Attempt %d/%d failed: %v", i+1, maxRetries, err)
	}
	return err
}
