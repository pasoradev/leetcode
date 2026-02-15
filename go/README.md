# LeetCode Go Solutions

This folder contains my LeetCode solutions written in Go.

## Running the Code

Each problem has its own folder containing a `main.go` file. To run a solution:

```bash
cd "go/<Problem Name>"
go run main.go
```

For example, to run the "Maximum Depth of N-ary Tree" solution:

```bash
cd "go/559. Maximum Depth of N-ary Tree"
go run main.go
```

## Code Structure

- Each problem is in its own folder (e.g., `559. Maximum Depth of N-ary Tree/`)
- Each folder contains a `main.go` file with:
  - The solution function(s)
  - A `main()` function with embedded test cases
  - Output via `fmt.Println()` or `println()`

## Test Cases

Test cases are embedded directly in the `main()` function of each file. Modify the test inputs in `main()` to test different cases.

## New Problem Setup

You don't need to run `go mod init` for these problems - the code uses only the standard library, so you can simply:

1. Create a new folder with the problem name (e.g., `"123. Problem Name/"`)
2. Create a `main.go` file with your solution
3. Run it with `go run main.go`
