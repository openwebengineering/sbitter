# SBitter: Santa Barbara Makerspace Twitter Clone

Demo app created to show off
[OpenWeb](http://www.openwebengineering.com)'s ability to build
scalable web applications as part of
[Santa Barbara Makerspace](http://sbhackerspace.com)'s
[SBitter](https://github.com/sbhackerspace/sbhx-sbitter#sbitter)
programming competition.

Go version: [![Build Status](https://drone.io/github.com/openwebengineering/sbitter/status.png)](https://drone.io/github.com/openwebengineering/sbitter/latest)

## Benchmark Summary

Setup: 10,000 read-only requests, 100 concurrently, with MongoDB in
Strong mode (very safe).


### Results

* 1-core Go, no caching
  * __380__ req/sec
* 1-core Go, Memcache
  * 1300 req/sec => 3.4x speedup from original thanks to Memcache
* 6-core Go, no caching
  * 2819 req/sec => 7.4x speedup from original due to parallelism
* 6-core Go, Memcache
  * __7009__ req/sec => __18.4x__ speedup total

See [BENCHMARKS.md](https://github.com/openwebengineering/sbitter/blob/master/BENCHMARKS.md) for details.
