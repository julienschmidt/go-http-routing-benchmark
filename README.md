go-http-routing-benchmark
=========================

This benchmark suite aims to compare the performance of available HTTP request routers for Go by implementing the routing of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.


Included Routers:

 * [Gocraft Web](https://github.com/gocraft/web)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [net/http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [HttpTreeMux](https://github.com/dimfeld/httptreemux)
 * [Martini](https://github.com/codegangsta/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)

## Results

Benchmark System:
 * Mid-2012 Apple Retina Macbook Pro
 * Intel Core i7 2.3 GHz
 * 2x4 GiB DDR3-1600 RAM
 * go1.2.1 darwin/amd64
 * Mac OSX 10.9.2

```
#GithubAPI Routes: 203
#GPlusAPI Routes: 13
#ParseAPI Routes: 26
#Static Routes: 157

BenchmarkGocraftWeb_Param	           1000000	      1729 ns/op	     672 B/op	       9 allocs/op
BenchmarkGorillaMux_Param	            500000	      5291 ns/op	     785 B/op	       7 allocs/op
BenchmarkHttpRouter_Param	           5000000	       629 ns/op	     343 B/op	       2 allocs/op
BenchmarkHttpTreeMux_Param	         2000000	       792 ns/op	     343 B/op	       2 allocs/op
BenchmarkMartini_Param	              500000	      7630 ns/op	    1185 B/op	      13 allocs/op
BenchmarkPat_Param	                  500000	      3213 ns/op	    1061 B/op	      17 allocs/op
BenchmarkTigerTonic_Param	            500000	      3212 ns/op	    1025 B/op	      19 allocs/op
BenchmarkTraffic_Param	              500000	      6207 ns/op	    2026 B/op	      23 allocs/op


BenchmarkGocraftWeb_ParamWrite	     1000000	      1828 ns/op	     681 B/op	      10 allocs/op
BenchmarkGorillaMux_ParamWrite	      500000	      5378 ns/op	     785 B/op	       7 allocs/op
BenchmarkHttpRouter_ParamWrite	     5000000	       612 ns/op	     343 B/op	       2 allocs/op
BenchmarkHttpTreeMux_ParamWrite  	   2000000	       757 ns/op	     343 B/op	       2 allocs/op
BenchmarkMartini_ParamWrite	          200000	      7639 ns/op	    1285 B/op	      16 allocs/op
BenchmarkPat_ParamWrite	              500000	      3405 ns/op	    1128 B/op	      19 allocs/op
BenchmarkTigerTonic_ParamWrite  	    500000	      5275 ns/op	    1483 B/op	      25 allocs/op
BenchmarkTraffic_ParamWrite  	        200000	      9421 ns/op	    2459 B/op	      27 allocs/op


BenchmarkGocraftWeb_GithubStatic	    1000000	      1275 ns/op	     313 B/op	       6 allocs/op
BenchmarkGorillaMux_GithubStatic	      50000	     46377 ns/op	     459 B/op	       6 allocs/op
BenchmarkHttpRouter_GithubStatic	   50000000	      66.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_GithubStatic	    5000000	       320 ns/op	      49 B/op	       1 allocs/op
BenchmarkMartini_GithubStatic	         100000	     21940 ns/op	     860 B/op	      12 allocs/op
BenchmarkPat_GithubStatic	             200000	     13251 ns/op	    3788 B/op	      76 allocs/op
BenchmarkTigerTonic_GithubStatic	    5000000	       345 ns/op	      49 B/op	       1 allocs/op
BenchmarkTraffic_GithubStatic	          50000	     67452 ns/op	   23361 B/op	     172 allocs/op


BenchmarkGocraftWeb_GithubParam	      1000000	      2213 ns/op	     736 B/op	      10 allocs/op
BenchmarkGorillaMux_GithubParam	       100000	     28081 ns/op	     818 B/op	       7 allocs/op
BenchmarkHttpRouter_GithubParam	      2000000	       743 ns/op	     343 B/op	       2 allocs/op
BenchmarkHttpTreeMux_GithubParam	    2000000	       881 ns/op	     343 B/op	       2 allocs/op
BenchmarkMartini_GithubParam	         100000	     24432 ns/op	    1219 B/op	      13 allocs/op
BenchmarkPat_GithubParam	             200000	      8603 ns/op	    2625 B/op	      56 allocs/op
BenchmarkTigerTonic_GithubParam	       500000	      5603 ns/op	    1483 B/op	      28 allocs/op
BenchmarkTraffic_GithubParam	          50000	     31182 ns/op	    7148 B/op	      60 allocs/op


BenchmarkGocraftWeb_GithubAll	           5000	    426058 ns/op	  136369 B/op	    1914 allocs/op
BenchmarkGorillaMux_GithubAll	            100	  16784051 ns/op	  153364 B/op	    1419 allocs/op
BenchmarkHttpRouter_GithubAll	          10000	    135000 ns/op	   57352 B/op	     347 allocs/op
BenchmarkHttpTreeMux_GithubAll	        10000	    168837 ns/op	   59145 B/op	     384 allocs/op
BenchmarkMartini_GithubAll	              100	  10733463 ns/op	  245391 B/op	    2941 allocs/op
BenchmarkPat_GithubAll	                  500	   4846657 ns/op	 1588374 B/op	   32572 allocs/op
BenchmarkTigerTonic_GithubAll	           2000	   1076119 ns/op	  250826 B/op	    6086 allocs/op
BenchmarkTraffic_GithubAll	              100	  14512012 ns/op	 3171328 B/op	   24922 allocs/op


BenchmarkGocraftWeb_GPlusStatic	      1000000	      1004 ns/op	     297 B/op	       6 allocs/op
BenchmarkGorillaMux_GPlusStatic	       500000	      3760 ns/op	     459 B/op	       6 allocs/op
BenchmarkHttpRouter_GPlusStatic	     50000000	      35.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic	   10000000	       287 ns/op	      49 B/op	       1 allocs/op
BenchmarkMartini_GPlusStatic	         500000	      5196 ns/op	     860 B/op	      12 allocs/op
BenchmarkPat_GPlusStatic	            5000000	       379 ns/op	      99 B/op	       2 allocs/op
BenchmarkTigerTonic_GPlusStatic	     10000000	       206 ns/op	      33 B/op	       1 allocs/op
BenchmarkTraffic_GPlusStatic	         500000	      5148 ns/op	    1510 B/op	      19 allocs/op


BenchmarkGocraftWeb_GPlusParam	      1000000	      1791 ns/op	     673 B/op	       9 allocs/op
BenchmarkGorillaMux_GPlusParam	       200000	      8704 ns/op	     785 B/op	       7 allocs/op
BenchmarkHttpRouter_GPlusParam	      5000000	       635 ns/op	     343 B/op	       2 allocs/op
BenchmarkHttpTreeMux_GPlusParam	      2000000	       719 ns/op	     343 B/op	       2 allocs/op
BenchmarkMartini_GPlusParam	           200000	      9330 ns/op	    1186 B/op	      13 allocs/op
BenchmarkPat_GPlusParam	              1000000	      2250 ns/op	     752 B/op	      14 allocs/op
BenchmarkTigerTonic_GPlusParam	       500000	      3634 ns/op	    1102 B/op	      19 allocs/op
BenchmarkTraffic_GPlusParam	           200000	      8286 ns/op	    2039 B/op	      23 allocs/op


BenchmarkGocraftWeb_GPlus2Params	    1000000	      2228 ns/op	     735 B/op	      10 allocs/op
BenchmarkGorillaMux_GPlus2Params	     100000	     23953 ns/op	     818 B/op	       7 allocs/op
BenchmarkHttpRouter_GPlus2Params	    5000000	       731 ns/op	     343 B/op	       2 allocs/op
BenchmarkHttpTreeMux_GPlus2Params	    2000000	       858 ns/op	     343 B/op	       2 allocs/op
BenchmarkMartini_GPlus2Params	          50000	     30463 ns/op	    1318 B/op	      17 allocs/op
BenchmarkPat_GPlus2Params	             500000	      7039 ns/op	    2400 B/op	      41 allocs/op
BenchmarkTigerTonic_GPlus2Params	     500000	      5907 ns/op	    1584 B/op	      28 allocs/op
BenchmarkTraffic_GPlus2Params	         100000	     25820 ns/op	    3620 B/op	      35 allocs/op


BenchmarkGocraftWeb_GPlusAll	         100000	     24310 ns/op	    8341 B/op	     117 allocs/op
BenchmarkGorillaMux_GPlusAll	          10000	    150843 ns/op	    9724 B/op	      91 allocs/op
BenchmarkHttpRouter_GPlusAll	         200000	      7820 ns/op	    3773 B/op	      22 allocs/op
BenchmarkHttpTreeMux_GPlusAll	         200000	      9737 ns/op	    3873 B/op	      24 allocs/op
BenchmarkMartini_GPlusAll	              10000	    170488 ns/op	   15530 B/op	     194 allocs/op
BenchmarkPat_GPlusAll	                  50000	     55544 ns/op	   17689 B/op	     346 allocs/op
BenchmarkTigerTonic_GPlusAll	          50000	     58732 ns/op	   15474 B/op	     322 allocs/op
BenchmarkTraffic_GPlusAll	              10000	    180953 ns/op	   42078 B/op	     446 allocs/op


BenchmarkGocraftWeb_ParseStatic	      1000000	      1064 ns/op	     313 B/op	       6 allocs/op
BenchmarkGorillaMux_ParseStatic	       500000	      7565 ns/op	     459 B/op	       6 allocs/op
BenchmarkHttpRouter_ParseStatic	     50000000	      36.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_ParseStatic	    5000000	       310 ns/op	      49 B/op	       1 allocs/op
BenchmarkMartini_ParseStatic	         500000	      5860 ns/op	     860 B/op	      12 allocs/op
BenchmarkPat_ParseStatic	            2000000	       906 ns/op	     249 B/op	       5 allocs/op
BenchmarkTigerTonic_ParseStatic	     10000000	       296 ns/op	      49 B/op	       1 allocs/op
BenchmarkTraffic_ParseStatic	         200000	      8140 ns/op	    2391 B/op	      25 allocs/op


BenchmarkGocraftWeb_ParseParam	      1000000	      1828 ns/op	     689 B/op	       9 allocs/op
BenchmarkGorillaMux_ParseParam	       200000	      8378 ns/op	     786 B/op	       7 allocs/op
BenchmarkHttpRouter_ParseParam	      5000000	       592 ns/op	     343 B/op	       2 allocs/op
BenchmarkHttpTreeMux_ParseParam	      5000000	       737 ns/op	     343 B/op	       2 allocs/op
BenchmarkMartini_ParseParam	           200000	      8101 ns/op	    1186 B/op	      13 allocs/op
BenchmarkPat_ParseParam	               500000	      3190 ns/op	    1196 B/op	      20 allocs/op
BenchmarkTigerTonic_ParseParam	      1000000	      3463 ns/op	    1083 B/op	      19 allocs/op
BenchmarkTraffic_ParseParam	           200000	      8042 ns/op	    2325 B/op	      25 allocs/op


BenchmarkGocraftWeb_Parse2Params	    1000000	      2076 ns/op	     736 B/op	      10 allocs/op
BenchmarkGorillaMux_Parse2Params	     200000	      8400 ns/op	     818 B/op	       7 allocs/op
BenchmarkHttpRouter_Parse2Params	    5000000	       676 ns/op	     343 B/op	       2 allocs/op
BenchmarkHttpTreeMux_Parse2Params	    2000000	       839 ns/op	     343 B/op	       2 allocs/op
BenchmarkMartini_Parse2Params	         200000	      8424 ns/op	    1219 B/op	      13 allocs/op
BenchmarkPat_Parse2Params	            1000000	      3177 ns/op	     907 B/op	      21 allocs/op
BenchmarkTigerTonic_Parse2Params	     500000	      5572 ns/op	    1487 B/op	      28 allocs/op
BenchmarkTraffic_Parse2Params	         200000	      8157 ns/op	    2130 B/op	      25 allocs/op


BenchmarkGocraftWeb_ParseAll	          50000	     42910 ns/op	   14301 B/op	     209 allocs/op
BenchmarkGorillaMux_ParseAll	           5000	    322454 ns/op	   17258 B/op	     175 allocs/op
BenchmarkHttpRouter_ParseAll	         200000	     11377 ns/op	    5489 B/op	      33 allocs/op
BenchmarkHttpTreeMux_ParseAll	         100000	     16928 ns/op	    5987 B/op	      43 allocs/op
BenchmarkMartini_ParseAll	              10000	    224285 ns/op	   27685 B/op	     333 allocs/op
BenchmarkPat_ParseAll	                  50000	     62583 ns/op	   18278 B/op	     385 allocs/op
BenchmarkTigerTonic_ParseAll	          20000	     77796 ns/op	   20851 B/op	     420 allocs/op
BenchmarkTraffic_ParseAll	              10000	    256974 ns/op	   70548 B/op	     762 allocs/op


BenchmarkHttpServeMux_StaticAll	         2000	   1184924 ns/op	     104 B/op	       8 allocs/op
BenchmarkGocraftWeb_StaticAll	          10000	    183888 ns/op	   49173 B/op	     951 allocs/op
BenchmarkGorillaMux_StaticAll	            500	   4872757 ns/op	   72314 B/op	     966 allocs/op
BenchmarkHttpRouter_StaticAll	         100000	     19081 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_StaticAll	       100000	     19020 ns/op	       0 B/op	       0 allocs/op
BenchmarkMartini_StaticAll	              500	   3765492 ns/op	  145657 B/op	    2521 allocs/op
BenchmarkPat_StaticAll	                 1000	   1834709 ns/op	  554100 B/op	   11249 allocs/op
BenchmarkTigerTonic_StaticAll	          50000	     67576 ns/op	    7777 B/op	     158 allocs/op
BenchmarkTraffic_StaticAll	              100	  11682322 ns/op	 3793165 B/op	   27907 allocs/op
```
