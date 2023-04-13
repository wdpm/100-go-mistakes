# notes

## 16

linters: 
- https://golang.org/cmd/vet/—A standard Go analyzer
- https://github.com/kisielk/errcheck—An error checker
- https://github.com/fzipp/gocyclo—A cyclomatic complexity analyzer
- https://github.com/jgautheron/goconst—A repeated string constants analyzer

formatters:
- https://golang.org/cmd/gofmt/—A standard Go code formatter
- https://godoc.org/golang.org/x/tools/cmd/goimports—A standard Go imports formatter


值 or pointer receiver, see #42

（goroutine）同步用mutex 锁，（goroutine）协作用channel 交流。

(A context allows you to carry a deadline, a cancellation signal, and/or a list of keys-values)

```
let’s remember that a channel without data should be
expressed with a chan struct{} type. This way, it clarifies for receivers that they
shouldn’t expect any meaning from a message’s content—only the fact that they have
received a message. In Go, such channels are called notification channels.
```