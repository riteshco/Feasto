# Performance Benchmark

## 1000 Concurrent Users and 100000 Requests::

### *\**The route is fetched through Authorization token*

```
This is ApacheBench, Version 2.3 <$Revision: 1923142 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            3000

Document Path:          /api/all-products
Document Length:        558 bytes

Concurrency Level:      1000
Time taken for tests:   7.286 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      68100000 bytes
HTML transferred:       55800000 bytes
Requests per second:    13724.88 [#/sec] (mean)
Time per request:       72.860 [ms] (mean)
Time per request:       0.073 [ms] (mean, across all concurrent requests)
Transfer rate:          9127.58 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   34   5.6     34      61
Processing:    11   38   8.4     37     128
Waiting:        7   27   7.6     26      99
Total:         42   72   7.9     71     147

Percentage of the requests served within a certain time (ms)
  50%     71
  66%     72
  75%     73
  80%     74
  90%     78
  95%     82
  98%    101
  99%    111
 100%    147 (longest request)
```