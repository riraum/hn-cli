# Specifications

This doc should guide you through making this new project.
Since it's going to explore a couple more topics, I'll get a bit more involved in the setup.

## Code

- Create a type `Item` that's a struct containing: `title`, `author`, `score`, `time`, and `URL`.

  Try to figure out the best type for each.
  
- Create a type `Items` that is `[]Item`.
- Create a method on the type `Items` that formats all the items in a table based on the passed width.

  E.g.
  ```
   0 1234 authorname    title is very long, and might get truncated if it ovela...  1d ago
   1 5564 anotherauthor a shorter title                                             3h ago
  ...
  45 4568 name          this title is very long, but it doesn't matter because ...  1m ago
  ```
- Create a method that converts a [`time.Time`](https://pkg.go.dev/time#Time) into a relative with the possible values:
  - `<1h` if it was less than 1h ago
  - `Nh` N is the number of hours, if it was less than 24 hours ago
  - `Nd` N is the number of days, if it was less than 30 days ago
  - `Nm` N is the number of months, if it was less than 12 months ago
  - `Ny` N is the number of years
  Make sure that the returned string is exactly 3 characters long, for easier padding.

- Use https://pkg.go.dev/net/http to make HTTP GET request to the API.
- Use https://pkg.go.dev/encoding/json to unmarshal the data into the `Item` struct.
- Use https://hacker-news.firebaseio.com/v0/topstories.json to get the latest stories.
- Use https://hacker-news.firebaseio.com/v0/item/<ID>.json to get the title, author, score, time, and URL for the story.
- Use https://pkg.go.dev/golang.org/x/term#GetSize to get the current terminal size.
- UI should be the following:

  ```
   0 1234 authorname    title is very long, and might get truncated if it ovela...  1d ago
   1 5564 anotherauthor a shorter title                                             3h ago
  ...
  45 4568 name          this title is very long, but it doesn't matter because ...  1m ago
  >
  ```
  The `>` expect a command:
  - `next`: gets the next page of items
  - `open X`: opens the item with index/id `X` in the browser
  - `quit`: quits the program
  - `refresh`: reload the top items
  - `comments`: open the comments page in the browser

## CI

- Create a GitHub workflow that checks for code linting and run tests
- Use https://golangci-lint.run/ for additional linting
- Create a GitHub workflow that build and runs the code
