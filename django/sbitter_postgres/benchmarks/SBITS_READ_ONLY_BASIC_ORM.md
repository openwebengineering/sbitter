This is the result of a read only 'ab' test on '/sbits/' using what I am terming 'Basic ORM', this means the view was setup like so.

		::python
		def view_all_sbits(request):
    		sbits = Sbit.objects.all()
    		return render(request, 'sbitter/sbits_list.html', locals())
		
The command used to was

		::bash
		ab -n 2000 -c 1 http://localhost:8000/sbits/


Memcached has been installed with no configuration, and caching setup in Django with no modifications (as layed out in https://docs.djangoproject.com/en/dev/topics/cache/)

Next will be testing better configuration of Memcached as well as Django ORM tricks.


###ab Output


		This is ApacheBench, Version 2.3 <$Revision: 655654 $>
		Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
		Licensed to The Apache Software Foundation, http://www.apache.org/

		Benchmarking localhost (be patient)


		Server Software:        WSGIServer/0.1
		Server Hostname:        localhost
		Server Port:            8000

		Document Path:          /sbits/
		Document Length:        2989 bytes

		Concurrency Level:      1
		Time taken for tests:   8.322 seconds
		Complete requests:      2000
		Failed requests:        0
		Write errors:           0
		Total transferred:      6500000 bytes
		HTML transferred:       5978000 bytes
		Requests per second:    240.33 [#/sec] (mean)
		Time per request:       4.161 [ms] (mean)
		Time per request:       4.161 [ms] (mean, across all concurrent requests)
		Transfer rate:          762.77 [Kbytes/sec] received

		Connection Times (ms)
					  min  mean[+/-sd] median   max
		Connect:        0    0   0.0      0       0
		Processing:     2    4   1.6      3      24
		Waiting:        1    3   1.4      3      24
		Total:          2    4   1.6      3      24

		Percentage of the requests served within a certain time (ms)
		  50%      3
		  66%      4
		  75%      5
		  80%      5
		  90%      6
		  95%      7
		  98%      8
		  99%     10
		 100%     24 (longest request)
