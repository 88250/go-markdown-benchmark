* [Golang markdown 引擎性能基准测试](https://hacpai.com/article/1574570835061)

​    run benchmark:

```
go test -bench .
```



result:

macOS 10.15.3

macbook pro 15 ( 2.8Ghz i7 CPU / 16G mem)

```
......../go-markdown-benchmark   go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/88250/go-markdown-benchmark
BenchmarkLute-8               	     340	   3418156 ns/op	 5969818 B/op	   38354 allocs/op
BenchmarkGolangCommonMark-8   	     796	   1606813 ns/op	 2901345 B/op	   20119 allocs/op
BenchmarkGoldMark-8           	     579	   2058027 ns/op	 2449778 B/op	   14943 allocs/op
BenchmarkBlackFriday-8        	     829	   1455494 ns/op	 3230938 B/op	   19855 allocs/op
BenchmarkGoMarkdown-8         	      40	  29939149 ns/op	 2173578 B/op	   22173 allocs/op
PASS
ok  	github.com/88250/go-markdown-benchmark	9.689s
```

