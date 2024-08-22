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

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func banner() {
	fmt.Println()
	color.Yellow("   █████████  ████                                 █████████   ████  ████ ")
	color.Yellow("  ███░░░░░███░░███                                ███░░░░░███ ░░███ ░░███ ")
	color.Yellow(" ███     ░░░  ░███   ██████  ████████    ██████  ░███    ░███  ░███  ░███ ")
	color.Yellow("░███          ░███  ███░░███░░███░░███  ███░░███ ░███████████  ░███  ░███ ")
	color.Yellow("░███          ░███ ░███ ░███ ░███ ░███ ░███████  ░███░░░░░███  ░███  ░███ ")
	color.Yellow("░░███     ███ ░███ ░███ ░███ ░███ ░███ ░███░░░   ░███    ░███  ░███  ░███ ")
	color.Yellow(" ░░█████████  █████░░██████  ████ █████░░██████  █████   █████ █████ █████")
	color.Yellow("  ░░░░░░░░░  ░░░░░  ░░░░░░  ░░░░ ░░░░░  ░░░░░░  ░░░░░   ░░░░░ ░░░░░ ░░░░░ ")
	fmt.Println()
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatalf("GitHub token not found in .env file")
	}

	banner()

	// Set up OAuth2 authentication
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	// Create a GitHub client
	client := github.NewClient(tc)

	for {
		// Get the GitHub username from the user
		reader := bufio.NewReader(os.Stdin)
		color.Yellow("Enter GitHub username (or type 'exit' to quit): ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username) // remove newline character and any surrounding spaces

		if strings.ToLower(username) == "exit" {
			color.Yellow("Exiting the program.")
			break
		}

		// Initialize an empty slice to store all repositories
		var allRepos []*github.Repository

		// Fetch all repositories for the user with pagination
		opts := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 50}} // Adjust PerPage if needed
		for {
			repos, resp, err := client.Repositories.List(ctx, username, opts)
			if err != nil {
				log.Fatalf("Error fetching repositories: %v", err)
			}

			// Append the current page of repos to allRepos
			allRepos = append(allRepos, repos...)

			// Check if there are more pages to fetch
			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}

		// List repositories and ask for user selection
		color.Yellow("Select repository to clone:")
		color.Red("0. Exit") // Option to exit
		color.Green("1. All")

		printRepoList(allRepos)

		color.Yellow("Enter your choice (number): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		choiceNum, _ := strconv.Atoi(choice)

		if choiceNum == 0 {
			color.Yellow("Exiting to username input.")
			continue
		}

		// Create a directory with the GitHub username
		os.Mkdir(fmt.Sprintf("%s's Repo", username), 0755)

		// Clone the repositories based on user choice
		if choiceNum == 1 {
			// Clone all repositories
			for _, repo := range allRepos {
				cloneRepo(username, *repo.CloneURL)
			}
		} else if choiceNum >= 2 && choiceNum < len(allRepos)+2 {
			// Clone the selected repository
			cloneRepo(username, *allRepos[choiceNum-2].CloneURL)
		} else {
			color.Red("Invalid choice")
		}
	}
}

// Function to print the repository list in multiple columns
func printRepoList(repos []*github.Repository) {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		for i, repo := range repos {
			color.Green("%d. %s", i+2, *repo.Name)
		}
		return
	}

	// Get terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default to 80 if terminal size cannot be determined
	}

	// Determine column width (longest repo name plus padding)
	maxLength := 0
	for _, repo := range repos {
		if len(*repo.Name) > maxLength {
			maxLength = len(*repo.Name)
		}
	}
	columnWidth := maxLength + 5 // Adjust the padding as needed
	columns := width / columnWidth

	// Print the repo names in columns with color
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

func cloneRepo(username, cloneURL string) {
	color.Blue("Cloning %s...", cloneURL)

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		cmd := exec.Command("git", "-c", "http.postBuffer=524288000", "clone", cloneURL)
		cmd.Dir = fmt.Sprintf("%s's Repo", username)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err == nil {
			color.Magenta("Clone completed.")
			return
		}

		color.Red("Failed to clone repository (attempt %d/%d): %v", i+1, maxRetries, err)
	}

	log.Fatalf("Failed to clone repository after %d attempts", maxRetries)
}
