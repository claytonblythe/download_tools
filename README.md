## Concurrent downloading tools in Go

Here is one of my firsts projects working with the Go programming language. I wanted to write some simple programs taking advantage of goroutines, channels, and concurrent programming in Go. This is somewhat easier and more native to do in Go compared to say Python where you have to use asyncio and aiohttp and manage more of the blocking/non-blocking code yourself. 

It seems here that I/O is still very much the bottle neck, trying to write the responses from the urls to disk appears to be the bottleneck, for plain text urls in a CDN I get 20x speedup with roughly 100 workers optimal, urls such as baidu with long latency and RTT can see a 200x improvement in performance!



### Benchmarks with Slow Url (High RTT) and no File I/O

| BenchmarkDownloadFilesWorkersSlowUrls/1-12          |	       1	| 249642960054 ns/op| 	19711224 B/op	|  166919 allocs/op|

BenchmarkDownloadFilesWorkersSlowUrls/31-12 	               1	8087429436 ns/op	19658664 B/op	  163985 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/61-12 	               1	4193474715 ns/op	19532456 B/op	  161037 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/91-12 	               1	2886746090 ns/op	19436224 B/op	  158409 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/121-12         	       1	2467211415 ns/op	19419656 B/op	  155752 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/151-12         	       1	1935703262 ns/op	19489488 B/op	  157939 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/181-12         	       1	1747113468 ns/op	19511840 B/op	  157648 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/211-12         	       1	1505520783 ns/op	19435048 B/op	  156835 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/241-12         	       1	1758476666 ns/op	20086792 B/op	  159719 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/271-12         	       1	1237023498 ns/op	19472288 B/op	  156393 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/301-12         	       1	1210278315 ns/op	19533432 B/op	  156790 allocs/op
BenchmarkDownloadFilesWorkersSlowUrls/331-12         	       1	1709274561 ns/op	19464336 B/op	  156207 allocs/op

Here we can see a 200x speedup over a simple synchronous single worker