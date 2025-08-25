# Preface
I think with the prevalence of _vibe coding_ and AI generated slop 🤮 it's vital to deeply understand the technologies we use. This means not only applying modern tools and frameworks, but understanding how the tools themselves work.

I wanted to understand backend frameworks better, so what better way to do that then to __CREATE__ one!

# Whiz
Whiz is a lighting-fast and minimal backend framework.
Whiz ships with a custom server, router and logger, I built from the ground up.

## Quickstart
Here is a simple demo to show off Whiz's most powerful features!

Start by creating a Whiz logger and Whiz router.
```go
func main() {
	logger := logging.NewLogger(
		logging.Info,
		"LOG_",
		logging.OneKB,
	)
	err := logger.Open()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	router := routing.NewRouter(logger)
    // ...
}
```

Now, lets define a route. Whiz supports dynamic routes by using the symbol ":" followed by the variable name

**_NOTE: Whiz handlers return a highly abstracted response model, making it easy for beginners to get into backend!_**

```go
func main() {
    router.Get("/hello/:name", func(params map[string]string) *routing.Response {
		// Dynamic parameters can be accessed through the params map
		name := params["name"]
		greeting := fmt.Sprintf("<h1>Hello there, %s! Nice to meet you.</h1>", name)

		res := routing.Response{
			Status:  200,
			Data:    greeting,
			Headers: map[string]string{"Content-Type": "text/html; charset=utf-8"},
		}
		res.Headers["Content-Length"] = strconv.Itoa(len(res.Data))

		return &res
	})
	router.Get("/hellos/:numhellos/:name", func(params map[string]string) *routing.Response {
		name := params["name"]
		numHellosParam := params["numhellos"]
		numHellos, err := strconv.Atoi(numHellosParam)
		if err != nil {
			logger.Error(fmt.Sprintf("Could not convert string to int (%s)", numHellosParam))
		}

		hellos := strings.Repeat(fmt.Sprintf("<h3>Hello %s!<h3>", name), numHellos)
		res := routing.Response{
			Status:  200,
			Data:    hellos,
			Headers: map[string]string{"Content-Type": "text/html; charset=utf-8"},
		}
		res.Headers["Content-Length"] = strconv.Itoa(len(res.Data))
		return &res
	})
}
```

As you can see, Whiz supports dynamic routing. A Whiz router will store the routes as a trie data structure. To match a path to the correct handler, the router will traverse this data structure and map the requested path onto the trie. To make these routes accessible, attach it to a whiz server.

```go
func main() {
    // ...
    server := server.NewServer(logger, 8000)
	server.AddRouter(router)
	server.Serve()
}
```

We can compile this code and run the server with:

```zsh
go build -o whiz
./whiz
```

When you hit http://localhost:8000/hello/Andrew you will see the following rendered to the browser.

```html
<h1>Hello there, Andrew! Nice to meet you.</h1>
```

And http://localhost:8000/3/Andrew renders:

```html
<h3>Hello Andrew!<h3>
<h3>Hello Andrew!<h3>
<h3>Hello Andrew!<h3>
```

## Logging

Every time you hit an endpoint from the server, it will be logged using a custom build Whiz logger. Here is an example of a log file: 

```log
[2025-08-22T02:01:05Z] INFO .../whiz/logger/logger.go:66 (GET /hellos/3/Andrew)
[2025-08-22T02:01:06Z] INFO .../whiz/logger/logger.go:66 (GET /hellos/3/Andrew)
[2025-08-22T02:01:14Z] INFO .../whiz/logger/logger.go:66 (GET /hellos/3/Andrew)
```
