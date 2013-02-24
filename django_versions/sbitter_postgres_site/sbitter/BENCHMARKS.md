##Basic Django setup (no optimizations, running locally)


##Django with basic Memcached setup
This is ApacheBench, Version 2.3 <$Revision: 655654 $>

Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/

Licensed to The Apache Software Foundation, http://www.apache.org/


Benchmarking 127.0.0.1 (be patient)


Server Software:        WSGIServer/0.1

Server Hostname:        127.0.0.1

Server Port:            8000


Document Path:          /

Document Length:        4060 bytes

Concurrency Level:      1

Time taken for tests:   43.107 seconds

Complete requests:      10000

Failed requests:        0

Write errors:           0

Total transferred:      43070000 bytes

HTML transferred:       40600000 bytes

###Requests per second:    231.98 [#/sec] (mean)

Time per request:       4.311 [ms] (mean)

Time per request:       4.311 [ms] (mean, across all concurrent requests)

Transfer rate:          975.72 [Kbytes/sec] received

Connection Times (ms)

              min  mean[+/-sd] median   max

Connect:        0    0   0.0      0       2

Processing:     3    4   1.3      4      23

Waiting:        1    4   1.2      3      22

Total:          3    4   1.3      4      23



Percentage of the requests served within a certain time (ms)

  50%      4


  66%      4

  75%      4

  80%      5

  90%      6

  95%      6

  98%      7

  99%      7

 100%     23 (longest request)
