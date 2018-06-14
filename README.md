# go-echo-gae-sample
Sample project for GoLang, echo, Google Appengine with clean architecture

# Notes
app-engine.yml is not contained.

# How to run
### local
```bash
go run app.go routes.go app-standalone.go
```

[Build Tags](https://golang.org/pkg/go/build/#hdr-Build_Constraints) doesn't work with go run command, so you should manually pass required files as arguments to run locally

# Related Contents
- [echo/cookbook/google-app-engine](https://echo.labstack.com/cookbook/google-app-engine)
