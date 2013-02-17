* Summary
  * 10k Read-only Requests (100 concurrent, MongoDB in Strong mode (very safe))
    * 1-core Go, no caching
      * 380 req/sec
    * 1-core Go, Memcache
      * 1300 req/sec (3.4x speedup from original thanks to Memcache)
    * 6-core Go, no caching
      * 2819 req/sec (7.4x speedup from original due to parallelism)
    * 6-core Go, Memcache
      * 7009 req/sec (18.4x speedup from using both combined)
  * Write-only Requests
    * TODO
  * 90% Reads, 10% Writes
    * TODO
  * 50% Reads, 50% Writes
    * TODO
  * 10% Reads, 90% Writes
    * TODO

* Technologies Used
  * Go
  * MongoDB
  * Memcache
  * ab (Apache Benchmark Tool)

* Hardware
  * CPU: AMD Phenom(tm) II X6 1090T
    * 6 cores, 3.2GHz (using 5 cores)
  * HD: SATA III Western Digital 1.5TB (_not_ SSD)
  * RAM: 8GB (only 1.8GB free but RAM was no bottleneck)

* Request types and payloads
  * Read-only payload
[{"id":"","user":{"username":"elimisteve"},"message":"New message 7","created_at":"2013-02-17T02:01:54.113-08:00"},{"id":"","user":{"username":"elimisteve"},"message":"New message 8","created_at":"2013-02-17T02:24:51.562-08:00"},{"id":"","user":{"username":"elimisteve"},"message":"New message 8","created_at":"2013-02-17T02:30:43.963-08:00"},{"id":"5120b2f73b0cad74c2000001","user":{"username":"elimisteve"},"message":"New message 9","created_at":"2013-02-17T02:37:43.709-08:00"},{"id":"5120b2fe3b0cad74c2000002","user":{"username":"elimisteve"},"message":"New message 9","created_at":"2013-02-17T02:37:50.263-08:00"},{"id":"5120b3c43b0cad762a000001","user":{"username":"elimisteve"},"message":"New message 10","created_at":"2013-02-17T02:41:08.283-08:00"}]

* Read-only requests, no caching, single-core Go

** Strong: slow, very safe

*** 10k requests (1 at a time) -- 283 req/sec, 4ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   35.297 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    283.31 [#/sec] (mean)
Time per request:       3.530 [ms] (mean)
Time per request:       3.530 [ms] (mean, across all concurrent requests)
Transfer rate:          239.32 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    3   1.0      3      15
Waiting:        1    3   1.0      3      15
Total:          1    3   1.0      3      15

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      3
  80%      3
  90%      3
  95%      4
  98%      8
  99%      8
 100%     15 (longest request)


*** 10k requests (5 at a time) -- 373 req/sec, 19ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   26.783 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    373.37 [#/sec] (mean)
Time per request:       13.392 [ms] (mean)
Time per request:       2.678 [ms] (mean, across all concurrent requests)
Transfer rate:          315.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     5   13   2.5     12      28
Waiting:        5   13   2.5     12      28
Total:          6   13   2.5     13      28

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     13
  75%     13
  80%     13
  90%     18
  95%     19
  98%     20
  99%     21
 100%     28 (longest request)


*** 10k requests (10 at a time) -- 375 req/sec, 33ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   26.623 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    375.61 [#/sec] (mean)
Time per request:       26.623 [ms] (mean)
Time per request:       2.662 [ms] (mean, across all concurrent requests)
Transfer rate:          317.29 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     8   26   3.6     25      41
Waiting:        7   26   3.6     25      41
Total:          8   27   3.6     25      41

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     26
  75%     27
  80%     31
  90%     32
  95%     33
  98%     35
  99%     37
 100%     41 (longest request)


*** 10k requests (100 at a time) -- 380 req/sec, 276ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   26.257 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    380.85 [#/sec] (mean)
Time per request:       262.571 [ms] (mean)
Time per request:       2.626 [ms] (mean, across all concurrent requests)
Transfer rate:          321.71 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:    19  262  14.0    263     329
Waiting:       19  261  14.0    263     329
Total:         21  262  13.9    263     330

Percentage of the requests served within a certain time (ms)
  50%    263
  66%    265
  75%    266
  80%    267
  90%    270
  95%    276
  98%    279
  99%    282
 100%    330 (longest request)


$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   25.591 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    390.77 [#/sec] (mean)
Time per request:       2559.075 [ms] (mean)
Time per request:       2.559 [ms] (mean, across all concurrent requests)
Transfer rate:          330.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  248 1026.9      0    7022
Processing:    31 2230 616.5   2541    3491
Waiting:       31 2230 616.5   2541    3490
Total:         44 2477 1216.1   2544    9892

Percentage of the requests served within a certain time (ms)
  50%   2544
  66%   2573
  75%   2581
  80%   2584
  90%   2593
  95%   4398
  98%   6417
  99%   9459
 100%   9892 (longest request)


*** 10k requests (1000 at a time) -- 390 req/sec, 4398ms 95%ile

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   25.591 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    390.77 [#/sec] (mean)
Time per request:       2559.075 [ms] (mean)
Time per request:       2.559 [ms] (mean, across all concurrent requests)
Transfer rate:          330.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  248 1026.9      0    7022
Processing:    31 2230 616.5   2541    3491
Waiting:       31 2230 616.5   2541    3490
Total:         44 2477 1216.1   2544    9892

Percentage of the requests served within a certain time (ms)
  50%   2544
  66%   2573
  75%   2581
  80%   2584
  90%   2593
  95%   4398
  98%   6417
  99%   9459
 100%   9892 (longest request)


** Monotonic: faster, pretty safe

*** 10k requests (1 at a time) -- 284 req/sec, 4ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   35.185 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    284.21 [#/sec] (mean)
Time per request:       3.518 [ms] (mean)
Time per request:       3.518 [ms] (mean, across all concurrent requests)
Transfer rate:          240.08 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    3   1.0      3      15
Waiting:        1    3   1.0      3      15
Total:          1    3   1.0      3      15

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      3
  80%      3
  90%      3
  95%      4
  98%      8
  99%      8
 100%     15 (longest request)

*** 10k requests (5 at a time) -- 375 req/sec, 19ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   26.608 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    375.83 [#/sec] (mean)
Time per request:       13.304 [ms] (mean)
Time per request:       2.661 [ms] (mean, across all concurrent requests)
Transfer rate:          317.47 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     4   13   2.3     12      27
Waiting:        3   13   2.3     12      27
Total:          4   13   2.3     12      27

Percentage of the requests served within a certain time (ms)
  50%     12
  66%     13
  75%     13
  80%     13
  90%     18
  95%     19
  98%     19
  99%     20
 100%     27 (longest request)

*** 10k requests (10 at a time) -- 374 req/sec, 32ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   26.674 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    374.90 [#/sec] (mean)
Time per request:       26.674 [ms] (mean)
Time per request:       2.667 [ms] (mean, across all concurrent requests)
Transfer rate:          316.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     6   26   3.2     25      41
Waiting:        6   26   3.2     25      41
Total:          6   27   3.2     25      41

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     26
  75%     27
  80%     31
  90%     32
  95%     32
  98%     33
  99%     34
 100%     41 (longest request)

*** 10k requests (100 at a time) -- 380 req/sec, 275ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   26.249 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    380.97 [#/sec] (mean)
Time per request:       262.489 [ms] (mean)
Time per request:       2.625 [ms] (mean, across all concurrent requests)
Transfer rate:          321.81 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:    18  261  13.3    262     326
Waiting:       18  261  13.3    262     325
Total:         20  261  13.2    262     326

Percentage of the requests served within a certain time (ms)
  50%    262
  66%    264
  75%    265
  80%    266
  90%    271
  95%    275
  98%    278
  99%    279
 100%    326 (longest request)

*** 10k requests (1000 at a time) -- 392 req/sec, 2650ms 95%ile

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   25.494 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    392.25 [#/sec] (mean)
Time per request:       2549.407 [ms] (mean)
Time per request:       2.549 [ms] (mean, across all concurrent requests)
Transfer rate:          331.34 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  194 889.6      0    7020
Processing:   136 2275 516.0   2526    2654
Waiting:      136 2275 516.0   2525    2654
Total:        150 2470 1017.2   2540    9665

Percentage of the requests served within a certain time (ms)
  50%   2540
  66%   2570
  75%   2576
  80%   2582
  90%   2605
  95%   2650
  98%   5382
  99%   9449
 100%   9665 (longest request)


** Eventual: fastest, least safe (eventually consistent)

*** 10k requests (1 at a time) -- ??? req/sec, 4ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   35.346 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    282.92 [#/sec] (mean)
Time per request:       3.535 [ms] (mean)
Time per request:       3.535 [ms] (mean, across all concurrent requests)
Transfer rate:          238.99 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    3   1.0      3      15
Waiting:        1    3   1.0      3      15
Total:          1    3   1.0      3      15

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      3
  80%      3
  90%      3
  95%      4
  98%      8
  99%      8
 100%     15 (longest request)

*** 10k requests (5 at a time) -- ??? req/sec, 21ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   28.057 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    356.42 [#/sec] (mean)
Time per request:       14.028 [ms] (mean)
Time per request:       2.806 [ms] (mean, across all concurrent requests)
Transfer rate:          301.08 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     4   14   3.6     13      51
Waiting:        4   14   3.6     13      51
Total:          4   14   3.6     13      51

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     15
  75%     16
  80%     17
  90%     19
  95%     21
  98%     23
  99%     24
 100%     51 (longest request)


*** 10k requests (10 at a time) -- ??? req/sec, 41ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   26.685 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    374.74 [#/sec] (mean)
Time per request:       26.685 [ms] (mean)
Time per request:       2.668 [ms] (mean, across all concurrent requests)
Transfer rate:          316.56 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     3   27   8.0     26      67
Waiting:        3   26   8.0     25      66
Total:          3   27   8.0     26      67

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     29
  75%     32
  80%     33
  90%     37
  95%     41
  98%     45
  99%     48
 100%     67 (longest request)


*** 10k requests (100 at a time) -- ??? req/sec, 294ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   27.823 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    359.41 [#/sec] (mean)
Time per request:       278.230 [ms] (mean)
Time per request:       2.782 [ms] (mean, across all concurrent requests)
Transfer rate:          303.61 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       6
Processing:    12  277  18.0    277     328
Waiting:       12  276  18.1    276     327
Total:         16  277  17.8    277     329

Percentage of the requests served within a certain time (ms)
  50%    277
  66%    281
  75%    283
  80%    284
  90%    289
  95%    294
  98%    299
  99%    303
 100%    329 (longest request)

*** 10k requests (1000 at a time) -- ??? req/sec, ???ms 95%ile (Errors)

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
apr_socket_recv: Connection reset by peer (104)
Total of 3605 requests completed

* Read-only requests, no caching, multi-core Go

** Strong: slow, very safe

*** 10k requests (1 at a time) -- 331 req/sec, 3ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   30.178 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    331.36 [#/sec] (mean)
Time per request:       3.018 [ms] (mean)
Time per request:       3.018 [ms] (mean, across all concurrent requests)
Transfer rate:          279.91 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    3   0.6      3      11
Waiting:        1    3   0.6      3      11
Total:          2    3   0.6      3      11

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      3
  80%      3
  90%      3
  95%      3
  98%      7
  99%      7
 100%     11 (longest request)


*** 10k requests (5 at a time) -- 1245 req/sec, 8ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   8.030 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    1245.39 [#/sec] (mean)
Time per request:       4.015 [ms] (mean)
Time per request:       0.803 [ms] (mean, across all concurrent requests)
Transfer rate:          1052.01 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:     1    4   1.8      3      16
Waiting:        1    4   1.8      3      16
Total:          1    4   1.8      3      16

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      4
  80%      5
  90%      6
  95%      8
  98%      9
  99%     11
 100%     16 (longest request)

*** 10k requests (10 at a time) -- 2337 req/sec, 8ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   4.278 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    2337.66 [#/sec] (mean)
Time per request:       4.278 [ms] (mean)
Time per request:       0.428 [ms] (mean, across all concurrent requests)
Transfer rate:          1974.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       2
Processing:     1    4   1.8      4      15
Waiting:        1    4   1.7      3      15
Total:          1    4   1.8      4      15

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      5
  80%      6
  90%      7
  95%      8
  98%      9
  99%     10
 100%     15 (longest request)

*** 10k requests (100 at a time) -- 2819 req/sec, 46ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   3.547 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    2819.28 [#/sec] (mean)
Time per request:       35.470 [ms] (mean)
Time per request:       0.355 [ms] (mean, across all concurrent requests)
Transfer rate:          2381.52 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       7
Processing:    14   35   5.5     35      54
Waiting:       14   35   5.5     35      54
Total:         15   35   5.5     35      54

Percentage of the requests served within a certain time (ms)
  50%     35
  66%     37
  75%     38
  80%     39
  90%     42
  95%     46
  98%     48
  99%     50
 100%     54 (longest request)

*** 10k requests (1000 at a time) -- 3601 req/sec, 1130ms 95%ile

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   2.776 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    3601.72 [#/sec] (mean)
Time per request:       277.645 [ms] (mean)
Time per request:       0.278 [ms] (mean, across all concurrent requests)
Transfer rate:          3042.47 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   59 235.4      0    1004
Processing:    46  195 120.1    205    1715
Waiting:       46  195 120.1    205    1715
Total:         59  254 299.0    209    2718

Percentage of the requests served within a certain time (ms)
  50%    209
  66%    235
  75%    241
  80%    249
  90%    367
  95%   1130
  98%   1473
  99%   1489
 100%   2718 (longest request)


** Monotonic: faster, pretty safe

*** 10k requests (1 at a time) -- 330 req/sec, 3ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   30.298 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    330.05 [#/sec] (mean)
Time per request:       3.030 [ms] (mean)
Time per request:       3.030 [ms] (mean, across all concurrent requests)
Transfer rate:          278.81 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    3   0.7      3      10
Waiting:        1    3   0.7      3       9
Total:          1    3   0.7      3      10

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      3
  80%      3
  90%      3
  95%      3
  98%      6
  99%      7
 100%     10 (longest request)

*** 10k requests (5 at a time) -- 1340 req/sec, 8ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   7.458 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    1340.79 [#/sec] (mean)
Time per request:       3.729 [ms] (mean)
Time per request:       0.746 [ms] (mean, across all concurrent requests)
Transfer rate:          1132.60 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     1    4   1.8      3      17
Waiting:        1    4   1.8      3      17
Total:          1    4   1.9      3      17

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      4
  80%      5
  90%      6
  95%      8
  98%      9
  99%     10
 100%     17 (longest request)

*** 10k requests (10 at a time) -- 2443 req/sec, 7ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   4.092 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    2443.52 [#/sec] (mean)
Time per request:       4.092 [ms] (mean)
Time per request:       0.409 [ms] (mean, across all concurrent requests)
Transfer rate:          2064.11 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1    4   1.6      3      13
Waiting:        1    4   1.6      3      13
Total:          1    4   1.6      4      13

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      5
  80%      5
  90%      7
  95%      7
  98%      8
  99%      9
 100%     13 (longest request)

*** 10k requests (100 at a time) -- 2802 req/sec, 44ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   3.569 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    2802.03 [#/sec] (mean)
Time per request:       35.688 [ms] (mean)
Time per request:       0.357 [ms] (mean, across all concurrent requests)
Transfer rate:          2366.95 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       7
Processing:     8   35  21.3     34     257
Waiting:        8   35  21.3     34     257
Total:         15   36  21.3     34     257

Percentage of the requests served within a certain time (ms)
  50%     34
  66%     38
  75%     39
  80%     39
  90%     41
  95%     44
  98%     47
  99%    217
 100%    257 (longest request)

*** 10k requests (1000 at a time) -- 3339 req/sec, 1136ms 95%ile

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   2.994 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    3339.48 [#/sec] (mean)
Time per request:       299.448 [ms] (mean)
Time per request:       0.299 [ms] (mean, across all concurrent requests)
Transfer rate:          2820.95 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   56 229.5      0    1004
Processing:    23  199 127.0    211    1684
Waiting:       22  199 127.0    211    1684
Total:         38  255 305.1    215    2687

Percentage of the requests served within a certain time (ms)
  50%    215
  66%    240
  75%    272
  80%    286
  90%    316
  95%   1136
  98%   1509
  99%   1526
 100%   2687 (longest request)


** Eventual: fastest, least safe (eventually consistent)

*** 10k requests (1 at a time) -- 327 req/sec, 3ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   30.513 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    327.73 [#/sec] (mean)
Time per request:       3.051 [ms] (mean)
Time per request:       3.051 [ms] (mean, across all concurrent requests)
Transfer rate:          276.84 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    3   0.7      3       9
Waiting:        1    3   0.7      3       9
Total:          1    3   0.7      3       9

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      3
  80%      3
  90%      3
  95%      3
  98%      6
  99%      6
 100%      9 (longest request)

*** 10k requests (5 at a time) -- 1646 req/sec, 6ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   6.072 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    1646.79 [#/sec] (mean)
Time per request:       3.036 [ms] (mean)
Time per request:       0.607 [ms] (mean, across all concurrent requests)
Transfer rate:          1391.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1    3   1.7      2      15
Waiting:        1    3   1.6      2      15
Total:          1    3   1.7      3      15

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      4
  80%      4
  90%      5
  95%      6
  98%      8
  99%      9
 100%     15 (longest request)

*** 10k requests (10 at a time) -- 2735 req/sec, 7ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   3.655 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    2735.95 [#/sec] (mean)
Time per request:       3.655 [ms] (mean)
Time per request:       0.366 [ms] (mean, across all concurrent requests)
Transfer rate:          2311.13 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1    4   1.6      3      13
Waiting:        1    3   1.5      3      13
Total:          1    4   1.6      3      13

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      4
  75%      4
  80%      5
  90%      6
  95%      7
  98%      8
  99%      9
 100%     13 (longest request)

*** 10k requests (100 at a time) -- 2984 req/sec, 43ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   3.351 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    2984.62 [#/sec] (mean)
Time per request:       33.505 [ms] (mean)
Time per request:       0.335 [ms] (mean, across all concurrent requests)
Transfer rate:          2521.19 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       2
Processing:     7   33   8.1     32     114
Waiting:        7   33   8.0     32     114
Total:          8   33   8.1     32     114

Percentage of the requests served within a certain time (ms)
  50%     32
  66%     34
  75%     36
  80%     37
  90%     40
  95%     43
  98%     48
  99%     86
 100%    114 (longest request)

*** 10k requests (1000 at a time) -- 1459 req/sec, 1471ms 95%ile (RUN #1)

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   6.852 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    1459.36 [#/sec] (mean)
Time per request:       685.233 [ms] (mean)
Time per request:       0.685 [ms] (mean, across all concurrent requests)
Transfer rate:          1232.76 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  137 500.3      0    3008
Processing:    11  121 413.9     45    4401
Waiting:       10  120 413.9     45    4401
Total:         26  257 761.3     46    5403

Percentage of the requests served within a certain time (ms)
  50%     46
  66%     48
  75%     51
  80%     53
  90%    101
  95%   1471
  98%   3950
  99%   4001
 100%   5403 (longest request)

*** 10k requests (1000 at a time) -- 1602 req/sec, 2488ms 95%ile (RUN #2)

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   6.240 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    1602.64 [#/sec] (mean)
Time per request:       623.970 [ms] (mean)
Time per request:       0.624 [ms] (mean, across all concurrent requests)
Transfer rate:          1353.79 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  183 611.3      0    3009
Processing:    15  128 425.9     46    6219
Waiting:       14  128 425.9     46    6219
Total:         27  311 888.2     46    6233

Percentage of the requests served within a certain time (ms)
  50%     46
  66%     50
  75%     54
  80%     59
  90%   1044
  95%   2488
  98%   3749
  99%   4192
 100%   6233 (longest request)


* Read-only requests, memcached, multi-core Go

** Strong: slow, very safe

*** 10k requests (1 at a time) -- 965 req/sec, 1ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   10.360 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    965.27 [#/sec] (mean)
Time per request:       1.036 [ms] (mean)
Time per request:       1.036 [ms] (mean, across all concurrent requests)
Transfer rate:          815.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    1   0.4      1       6
Waiting:        0    1   0.4      1       6
Total:          0    1   0.4      1       6

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      4
 100%      6 (longest request)

*** 10k requests (5 at a time) -- 4499 req/sec, 3ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   2.223 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    4499.29 [#/sec] (mean)
Time per request:       1.111 [ms] (mean)
Time per request:       0.222 [ms] (mean, across all concurrent requests)
Transfer rate:          3800.67 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:     0    1   0.8      1      11
Waiting:        0    1   0.8      1      10
Total:          0    1   0.8      1      11

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      2
  95%      3
  98%      4
  99%      5
 100%     11 (longest request)

*** 10k requests (10 at a time) -- 6246 req/sec, 4ms 95%ile [RUN #1]

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   1.601 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    6246.55 [#/sec] (mean)
Time per request:       1.601 [ms] (mean)
Time per request:       0.160 [ms] (mean, across all concurrent requests)
Transfer rate:          5276.63 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       4
Processing:     0    1   1.1      1      14
Waiting:        0    1   1.0      1      14
Total:          0    2   1.1      1      14

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      2
  75%      2
  80%      2
  90%      3
  95%      4
  98%      5
  99%      5
 100%     14 (longest request)

*** 10k requests (100 at a time) -- 7009 req/sec, 28ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   1.427 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    7009.51 [#/sec] (mean)
Time per request:       14.266 [ms] (mean)
Time per request:       0.143 [ms] (mean, across all concurrent requests)
Transfer rate:          5921.12 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       3
Processing:     0   14   8.2     14      67
Waiting:        0   14   8.2     13      67
Total:          0   14   8.2     14      68

Percentage of the requests served within a certain time (ms)
  50%     14
  66%     17
  75%     19
  80%     21
  90%     24
  95%     28
  98%     33
  99%     37
 100%     68 (longest request)

*** 10k requests (1000 at a time) -- 6045 req/sec, 112ms 95%ile

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   1.654 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    6045.69 [#/sec] (mean)
Time per request:       165.407 [ms] (mean)
Time per request:       0.165 [ms] (mean, across all concurrent requests)
Transfer rate:          5106.95 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   46 208.3      0    1004
Processing:     0   46  76.4     35     640
Waiting:        0   46  76.4     34     640
Total:          0   92 266.4     35    1642

Percentage of the requests served within a certain time (ms)
  50%     35
  66%     42
  75%     48
  80%     56
  90%     73
  95%    112
  98%   1264
  99%   1632
 100%   1642 (longest request)


** Monotonic: faster, pretty safe

*** 10k requests (1 at a time) -- 968 req/sec, 1ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   10.324 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    968.58 [#/sec] (mean)
Time per request:       1.032 [ms] (mean)
Time per request:       1.032 [ms] (mean, across all concurrent requests)
Transfer rate:          818.19 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    1   0.3      1       5
Waiting:        0    1   0.3      1       5
Total:          0    1   0.3      1       5

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      4
 100%      5 (longest request)

*** 10k requests (5 at a time) -- 3536 req/sec, 4ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   2.828 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    3536.41 [#/sec] (mean)
Time per request:       1.414 [ms] (mean)
Time per request:       0.283 [ms] (mean, across all concurrent requests)
Transfer rate:          2987.30 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     0    1   0.9      1      10
Waiting:        0    1   0.9      1      10
Total:          0    1   0.9      1      10

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      2
  90%      2
  95%      4
  98%      5
  99%      5
 100%     10 (longest request)

*** 10k requests (10 at a time) -- 6553 req/sec, 4ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   1.526 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    6553.36 [#/sec] (mean)
Time per request:       1.526 [ms] (mean)
Time per request:       0.153 [ms] (mean, across all concurrent requests)
Transfer rate:          5535.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     0    1   1.0      1      16
Waiting:        0    1   1.0      1      15
Total:          0    1   1.0      1      16

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      2
  80%      2
  90%      3
  95%      4
  98%      4
  99%      5
 100%     16 (longest request)

*** 10k requests (100 at a time) -- 7817 req/sec, 25ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   1.279 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    7817.97 [#/sec] (mean)
Time per request:       12.791 [ms] (mean)
Time per request:       0.128 [ms] (mean, across all concurrent requests)
Transfer rate:          6604.04 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       2
Processing:     0   13   7.1     12      61
Waiting:        0   12   7.1     12      61
Total:          0   13   7.1     12      61

Percentage of the requests served within a certain time (ms)
  50%     12
  66%     15
  75%     17
  80%     18
  90%     22
  95%     25
  98%     29
  99%     32
 100%     61 (longest request)

*** 10k requests (1000 at a time) -- 6063 req/sec, 1005ms 95%ile

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   1.649 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    6063.81 [#/sec] (mean)
Time per request:       164.913 [ms] (mean)
Time per request:       0.165 [ms] (mean, across all concurrent requests)
Transfer rate:          5122.26 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   51 218.6      0    1004
Processing:     0   52  77.0     39     632
Waiting:        0   52  77.0     38     632
Total:          0  103 275.2     40    1635

Percentage of the requests served within a certain time (ms)
  50%     40
  66%     49
  75%     55
  80%     67
  90%     94
  95%   1005
  98%   1274
  99%   1624
 100%   1635 (longest request)



** Eventual: fastest, least safe (eventually consistent)

*** 10k requests (1 at a time) -- 966 req/sec, 1ms 95%ile

$ ab -n 10000 -c 1 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1
Time taken for tests:   10.349 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    966.24 [#/sec] (mean)
Time per request:       1.035 [ms] (mean)
Time per request:       1.035 [ms] (mean, across all concurrent requests)
Transfer rate:          816.21 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    1   0.4      1       5
Waiting:        0    1   0.4      1       5
Total:          0    1   0.4      1       5

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      4
 100%      5 (longest request)

*** 10k requests (5 at a time) -- 4066 req/sec, 3ms 95%ile

$ ab -n 10000 -c 5 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      5
Time taken for tests:   2.459 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    4066.19 [#/sec] (mean)
Time per request:       1.230 [ms] (mean)
Time per request:       0.246 [ms] (mean, across all concurrent requests)
Transfer rate:          3434.82 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     0    1   1.0      1      10
Waiting:        0    1   0.9      1      10
Total:          0    1   1.0      1      10

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      2
  95%      3
  98%      5
  99%      5
 100%     10 (longest request)

*** 10k requests (10 at a time) -- 4527 req/sec, 6ms 95%ile

$ ab -n 10000 -c 10 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      10
Time taken for tests:   2.209 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    4527.79 [#/sec] (mean)
Time per request:       2.209 [ms] (mean)
Time per request:       0.221 [ms] (mean, across all concurrent requests)
Transfer rate:          3824.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       2
Processing:     0    2   1.5      2      20
Waiting:        0    2   1.4      1      20
Total:          0    2   1.5      2      21

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      2
  75%      3
  80%      3
  90%      4
  95%      6
  98%      7
  99%      7
 100%     21 (longest request)

*** 10k requests (100 at a time) -- 7237 req/sec, 28ms 95%ile

$ ab -n 10000 -c 100 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      100
Time taken for tests:   1.382 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    7237.39 [#/sec] (mean)
Time per request:       13.817 [ms] (mean)
Time per request:       0.138 [ms] (mean, across all concurrent requests)
Transfer rate:          6113.61 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       2
Processing:     0   14   7.9     13      63
Waiting:        0   13   7.9     13      63
Total:          0   14   7.9     13      63

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     16
  75%     18
  80%     20
  90%     24
  95%     28
  98%     32
  99%     37
 100%     63 (longest request)

*** 10k requests (1000 at a time) -- 6107 req/sec, 100ms 95%ile

$ ab -n 10000 -c 1000 http://localhost:8080/user/elimisteve
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /user/elimisteve
Document Length:        758 bytes

Concurrency Level:      1000
Time taken for tests:   1.637 seconds
Complete requests:      10000
Failed requests:        0
Write errors:           0
Total transferred:      8650000 bytes
HTML transferred:       7580000 bytes
Requests per second:    6107.61 [#/sec] (mean)
Time per request:       163.730 [ms] (mean)
Time per request:       0.164 [ms] (mean, across all concurrent requests)
Transfer rate:          5159.26 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   27 161.0      0    1003
Processing:     0   40  35.6     35     621
Waiting:        0   39  35.6     34     621
Total:          0   67 179.5     36    1624

Percentage of the requests served within a certain time (ms)
  50%     36
  66%     45
  75%     50
  80%     55
  90%     76
  95%    100
  98%   1042
  99%   1124
 100%   1624 (longest request)
