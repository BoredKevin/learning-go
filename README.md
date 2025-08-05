# learning-go

This is the repository that I use to learn golang and go functions, each cmd dir contains a different learning material.

## Project Structure

This project follows a standard Go project layout:

-   `/cmd`: Contains the main applications for the project. Each subdirectory is a separate executable.
-   `/internal`: Contains private application code. Code in here is not importable by other projects.

## How to Build and Run

To run an application, use `go run` from the project root. For example, to run the `calculator` application:

```sh
go run ./cmd/calculator
```

To build the binary into the `build/` directory (using the VS Code task or manually):
```sh
go build -o build/calculator ./cmd/calculator
```
