# Go Errors

Go Errors is a package for handling errors in Go. It provides a simple and flexible way to create, handle and manage
errors in your code.

## Installation

To use Go Errors in your Go project, you need to install it using Go modules. Run the following command in your
terminal:

```go
go get github.com/developersgotech/go-error
```

## Usage

### Creating an error

To create an error, you can use the New function. It takes a message and an optional code as arguments and returns a new
error:

```go
err := goerror.New("An error occurred", "ERR-001")
```

### Handling an error

To handle an error, you can use the DeferError function. It takes a callback function as an argument and calls it with
the recovered error and stack trace:

```go
goerror.DeferError(func (err error, stackTrace string) {
    log.Printf("Error: %v\nStack Trace: %s", err, stackTrace)
})
```

### Getting the stack trace

To get the stack trace of an error, you can use the GetStackTrace function. It returns a string with the stack trace of
the error:

```go
stackTrace := err.GetStackTrace()
```

### Getting the error message and code
To get the message and code of an error, you can use the Message and Code methods respectively:

```go
message := err.Message()
code := err.Code()
```

## License
Go Errors is licensed under the BSD3 License. See [LICENSE](LICENSE) for more information.
