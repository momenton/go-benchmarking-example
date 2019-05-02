# Go Benchmarking example

simple benchmarking of string concatenation functions
taken from: https://www.soroushjp.com/2015/01/27/beautifully-simple-benchmarking-with-go/

```
jimenezd-JSS1129:benchmarking-example jimenezd$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/runnerdave/benchmarking-example
BenchmarkSelfConcatOperator1000-12                  5000            272766 ns/op
BenchmarkSelfConcatBuffer1000-12                  200000              7214 ns/op
BenchmarkSelfConcatOperator100000-12                   1        1463489675 ns/op
BenchmarkSelfConcatBuffer100000-12                  2000            636606 ns/op
PASS
ok      github.com/runnerdave/benchmarking-example      5.728s
```


From the official Golang documentation: https://golang.org/pkg/testing/#hdr-Benchmarks

_The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably. The output_

````
BenchmarkHello    10000000    282 ns/op
````
_means that the loop ran 10000000 times at a speed of 282 ns per loop._

_If a benchmark needs some expensive setup before running, the timer may be reset:_

````
func BenchmarkBigLen(b *testing.B) {
    big := NewBig()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        big.Len()
    }
}
````

_If a benchmark needs to test performance in a parallel setting, it may use the RunParallel helper function; such benchmarks are intended to be used with the go test -cpu flag:_
````
func BenchmarkTemplateParallel(b *testing.B) {
    templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
    b.RunParallel(func(pb *testing.PB) {
        var buf bytes.Buffer
        for pb.Next() {
            buf.Reset()
            templ.Execute(&buf, "World")
        }
    })
}
````