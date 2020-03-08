## Concurrent downloading tools in Go

Here is one of my firsts projects working with the Go programming language. I wanted to write some simple programs taking advantage of goroutines, channels, and concurrent programming in Go. This is somewhat easier and more native to do in Go compared to say Python where you have to use asyncio and aiohttp and manage more of the blocking/non-blocking code yourself. 

It seems here that I/O is still very much the bottle neck, trying to write the responses from the urls to disk appears to be the bottleneck, for plain text urls in a CDN I get 20x speedup with roughly 100 workers optimal, urls such as baidu with long latency and RTT can see a 200x improvement in performance!



### Benchmarks with Slow Url (High RTT) and no File I/O

| Benchmark | Performance |
| --- | --- | 
| BenchmarkDownloadFilesWorkersSlowUrls/1         |	   249642960054 ns/op | 
| BenchmarkDownloadFilesWorkersSlowUrls/31 |	          8087429436 ns/op |
BenchmarkDownloadFilesWorkersSlowUrls/61	     |         4193474715 ns/op	
BenchmarkDownloadFilesWorkersSlowUrls/91	      |        2886746090 ns/op	
BenchmarkDownloadFilesWorkersSlowUrls/121         |	      2467211415 ns/op	
BenchmarkDownloadFilesWorkersSlowUrls/151        | 	      1935703262 ns/op
BenchmarkDownloadFilesWorkersSlowUrls/181        | 	      1747113468 ns/op
BenchmarkDownloadFilesWorkersSlowUrls/211        | 	      1505520783 ns/op	
BenchmarkDownloadFilesWorkersSlowUrls/241         |	      1758476666 ns/op
BenchmarkDownloadFilesWorkersSlowUrls/271        	|       1237023498 ns/op	
BenchmarkDownloadFilesWorkersSlowUrls/301        | 	      	1210278315 ns/op	|
BenchmarkDownloadFilesWorkersSlowUrls/331        | 	      1709274561 ns/op |

Here we can see a 200x speedup over a simple synchronous single worker