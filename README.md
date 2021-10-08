# DeepCopy Bench

```shell
$ go test -bench=. -benchmem -covermode atomic -race -v -count=1

goos: linux
goarch: amd64
pkg: github.com/jenting/deepcopy-bench
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkEngineStatusManual
BenchmarkEngineStatusManual-12             58260             17680 ns/op            5328 B/op         46 allocs/op
BenchmarkEngineStatusAuto
BenchmarkEngineStatusAuto-12               62176             18672 ns/op            5328 B/op         46 allocs/op
```

```shell
$ go test -bench=. -benchmem -covermode atomic -race -v -benchtime=10s

goos: linux
goarch: amd64
pkg: github.com/jenting/deepcopy-bench
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkEngineStatusManual
BenchmarkEngineStatusManual-12    	  654806	     17536 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto
BenchmarkEngineStatusAuto-12      	  573153	     18809 ns/op	    5328 B/op	      46 allocs/op
```

```shell
$ go test -bench=. -benchmem -covermode atomic -race -v -count=10

goos: linux
goarch: amd64
pkg: github.com/jenting/deepcopy-bench
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkEngineStatusManual
BenchmarkEngineStatusManual-12    	   64958	     17598 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   68036	     17847 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   67488	     17444 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   68929	     17400 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   68101	     17730 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   67458	     17744 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   66082	     17889 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   69452	     17876 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   68677	     17833 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusManual-12    	   66700	     17878 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto
BenchmarkEngineStatusAuto-12      	   60763	     18706 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   64200	     18674 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   63416	     18839 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   61822	     19506 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   62412	     19101 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   60321	     19315 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   60768	     19246 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   62434	     18911 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   63513	     19997 ns/op	    5328 B/op	      46 allocs/op
BenchmarkEngineStatusAuto-12      	   62738	     19454 ns/op	    5328 B/op	      46 allocs/op
```
