# CloneAllRepo

[![Go Version](https://img.shields.io/badge/go-1.17%20%7C%201.18%20%7C%201.19%20%7C%201.20-blue)](https://golang.org/dl/)
[![GitHub Issues](https://img.shields.io/github/issues/mamad-1999/CloneAllRepo)](https://github.com/mamad-1999/CloneAllRepo/issues)
[![GitHub Stars](https://img.shields.io/github/stars/mamad-1999/CloneAllRepo)](https://github.com/mamad-1999/CloneAllRepo/stargazers)
[![GitHub License](https://img.shields.io/github/license/mamad-1999/CloneAllRepo)](https://github.com/mamad-1999/CloneAllRepo/blob/master/LICENSE)

<p>
    <a href="https://skillicons.dev">
      <img src="https://github.com/tandpfun/skill-icons/blob/main/icons/GoLang.svg" width="48" title="Go">
      <img src="https://github.com/tandpfun/skill-icons/blob/main/icons/Github-Dark.svg" width="48" title="github">
    </a>
</p>

CloneAllRepo is a script that allows you to clone repositories from a specified GitHub user.
It provides an interactive and colorful terminal interface to choose and clone repositories.

![clone](https://github.com/user-attachments/assets/d0caa7cc-6de2-480d-bad5-cd5aafac6351)


## Requirements

- Go (version 1.18 or later)
- Git
- GitHub Personal Access Token (for authentication and rate limit issues)

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/mamad-1999/CloneAllRepo.git
   ```

2. **Navigate to the Project Directory**:
    ```bash
   cd CloneAllRepo
    ```

> [!IMPORTANT]
> before this point you should install the go

[![Go Version](https://img.shields.io/badge/go-1.17%20%7C%201.18%20%7C%201.19%20%7C%201.20-blue)](https://golang.org/dl/)

3. **Install Dependencies**:
    ```bash
    go mod tidy
    ```

## Usage

1. **Create a .env file in the project directory with your GitHub token**:
   ```bash
   GITHUB_TOKEN=your_github_token
  ``

2. **Run the Script**:
  ```bash
go run cloneAllRepo.go
```

Follow the Prompts:

- Enter the GitHub username.
- Choose which repositories to clone (All or specific numbers).

## License

This project is licensed under the MIT License.

## Contributing

Feel free to submit issues if you have improvements or bug fixes.
