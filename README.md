<p align="center">
  <a href="https://golang.org/dl/"><img src="https://img.shields.io/badge/go-1.17%20%7C%201.18%20%7C%201.19%20%7C%201.20-blue"></a>
  <a href="https://github.com/mamad-1999/CloneAllRepo/issues"><img src="https://img.shields.io/github/issues/mamad-1999/CloneAllRepo"></a>
  <a href="https://github.com/mamad-1999/CloneAllRepo/stargazers"><img src="https://img.shields.io/github/stars/mamad-1999/CloneAllRepo"></a>
  <a href="https://github.com/mamad-1999/CloneAllRepo/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mamad-1999/CloneAllRepo"></a>
</p>
<h4 align="center">CloneAllRepo is a script that allows you to clone repositories from a specified GitHub user.
It provides an interactive and colorful terminal interface to choose and clone repositories.</h4>
<p align="center">
  <a href="#installation"><img src="https://img.shields.io/badge/Install-blue?style=for-the-badge" alt="Install"></a>
  <a href="#usage"><img src="https://img.shields.io/badge/Usage-green?style=for-the-badge" alt="Usage"></a>
  <a href="#preview"><img src="https://img.shields.io/badge/Preview-red?style=for-the-badge" alt="Preview"></a>
  <a href="#contributing"><img src="https://img.shields.io/badge/Contributing-yellow?style=for-the-badge" alt="Contributing"></a>
</p>
<p align="center">
    <a href="https://skillicons.dev">
      <img src="https://github.com/tandpfun/skill-icons/blob/main/icons/GoLang.svg" width="48" title="Go">
      <img src="https://github.com/tandpfun/skill-icons/blob/main/icons/Github-Dark.svg" width="48" title="github">
    </a>
</p>

### Preview
![clone](https://github.com/user-attachments/assets/d0caa7cc-6de2-480d-bad5-cd5aafac6351)

> [!IMPORTANT]
> ### Requirements
> - Go (version 1.18 or later) [![Go Version](https://img.shields.io/badge/go-1.17%20%7C%201.18%20%7C%201.19%20%7C%201.20-blue)](https://golang.org/dl/)
> - Git
> - GitHub Personal Access Token (for authentication and rate limit issues)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/mamad-1999/CloneAllRepo.git
   ```
2. **Navigate to the Project Directory**:
    ```bash
   cd CloneAllRepo
    ```
3. **Install Dependencies**:
    ```bash
    go mod tidy
    ```

## Usage

1. **Create a .env file in the project directory with your GitHub token**:
   ```bash
   GITHUB_TOKEN=your_github_token

2. **Run the Script**:
      ```bash
      go run cloneAllRepo.go
      ```
> [!TIP]
> Follow the Prompts:
> - Enter the GitHub username.
> - Choose which repositories to clone (All or specific numbers).

## Contributing
Feel free to submit issues if you have improvements or bug fixes.
