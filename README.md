# Go Practice Repository

This repository is a collection of small Go experiments and learning exercises.

## Current Projects

| Folder | Description |
| ------ | ----------- |
| `hello-go` | A minimal "Hello, World!" program that demonstrates how to set up a simple Go module and print to the console. |

## Prerequisites

* Go 1.22 (or any relatively recent version of Go) is installed and in your `PATH`.
* A POSIX-compatible shell (macOS Terminal, zsh, bash, etc.).

## Running the example

1. Open a terminal and clone this repository (if you haven't already):
   ```bash
   git clone git@github.com:Sudibya/Go_practise.git
   cd Go_practise
   ```

2. Navigate to the project you want to run—in this case `hello-go`—and execute it:
   ```bash
   cd hello-go
   go run .
   # or build a binary
   go build -o hello
   ./hello
   ```
   You should see:
   ```
   Hello, World!
   ```

## Adding new examples

Feel free to create additional folders (e.g., `fib`, `http-server`) under the repository root. Each folder can be its own Go module (`go mod init`) or share the root module—whichever works best for your learning goals.

When you add new code:

```bash
# from the repository root
git add <new_folder>
git commit -m "Add <description> example"
git push origin master
```

## License

This repository is licensed under the MIT License—see the [LICENSE](LICENSE) file for details.
