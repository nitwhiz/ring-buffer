# ring-buffer

A **generic** ring buffer implementation in Golang. Implements `io.Writer` and `io.Reader`.

## Usage

```shell
go get -u github.com/nitwhiz/ring-buffer
```

```go
buf := ring.NewBuffer[int](10)

// write to buffer
n, err := buf.Write([]int{1, 2, 3, 4, 5})
err = buf.WriteOne(7)

// read from buffer
data := make([]int, 2)
n, err = buf.Read(data)
three, err := buf.ReadOne()
```

There is also a concurrent-safe version with `ring.NewBlockingBuffer`, which only allows a single `Read`/`ReadOne` or `Write`/`WriteOne` at a time.
