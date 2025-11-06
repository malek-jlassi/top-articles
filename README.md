# 📰 Top Articles

This is a small Go project that fetches article data from a public API and prints the top N article titles sorted by number of comments.

## 🎯 What you will learn
- How to structure a tiny Go project
- How to make HTTP requests and decode JSON
- How to sort data and print results

## 🧱 Project Structure
```
top-articles/
├── Dockerfile                     # Container build
├── cmd/
│   └── app/
│       └── main.go                # Program entry: calls TopArticles and prints
├── business_logic/
│   └── articles/
│       ├── model.go               # Data structures (Article, APIResponse)
│       ├── repository.go          # HTTP fetch helpers
│       ├── service.go             # TopArticles business logic
│       └── service_test.go        # Unit tests for TopArticles
├── go.mod                         # Go module + version
└── README.md                
```

## ✅ Prerequisites
- Install Go (1.21+ recommended): https://go.dev/dl/
- Windows, macOS, or Linux terminal

## ▶️ Run the app
1) Open a terminal in the project root (folder with `go.mod`).
2) Run:
```sh
go run ./cmd/app
```

You should see something like:
```
Top Articles:
1. Article A
2. Article B
3. Article C
4. Article D
5. Article E
```

To change how many titles are shown, open `cmd/app/main.go` and change the number passed to `TopArticles(...)`.

## 🧠 How the code works (in simple steps)
1) `service.go` asks the API for page 1 to learn how many total pages exist.
2) It loops over all pages one by one using `fetchPage(page)` from `repository.go`.
3) It collects all articles, picks a title (`title` or `story_title`), and sets missing comment counts to 0.
4) It sorts articles by comments (descending) and returns the top N names.

This project uses a sequential approach (no goroutines/channels) to make the flow easy to follow.

## 🧪 Run tests (optional)
From the project root:
```sh
go test ./...
```

**Author:** Malek JELASSI

