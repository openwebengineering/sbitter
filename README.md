# SBitter: Santa Barbara Twitter Clone

Demo app created to show off [OpenWeb](http://www.openweb.com)'s
ability to build scalable web applications.

Thank you to [Santa Barbara Makerspace](http://sbhackerspace.com) for
creating this competition.


## Benchmark Summary (see
   [BENCHMARKS.md](https://github.com/openwebengineering/sbitter/blob/master/BENCHMARKS.md)
   for details)

Setup: 10,000 read-only requests, 100 concurrently, with MongoDB in
Strong mode (very safe).

### Results

* 1-core Go, no caching
  * *380* req/sec
* 1-core Go, Memcache
  * 1300 req/sec (3.4x speedup from original thanks to Memcache)
* 6-core Go, no caching
  * 2819 req/sec (7.4x speedup from original due to parallelism)
* 6-core Go, Memcache
  * *7009* req/sec (18.4x speedup from using both combined)
