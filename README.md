### GithubAPI Routes: 203

```
   Ace: 48992 Bytes
   Bear: 161704 Bytes
   Beego: 144536 Bytes
   Bone: 97696 Bytes
   Denco: 36728 Bytes
   Echo: 76312 Bytes
   Gin: 52464 Bytes
   GocraftWeb: 95512 Bytes
   Goji: 86104 Bytes
   Gojiv2: 144408 Bytes
   GoJsonRest: 133832 Bytes
   GoRestful: 1395576 Bytes
   GorillaMux: 1494848 Bytes
   HttpRouter: 37464 Bytes
   HttpTreeMux: 78736 Bytes
   Iris: 62312 Bytes
   Kocha: 785120 Bytes
   LARS: 49016 Bytes
   Macaron: 128408 Bytes
   Martini: 556192 Bytes
   Pat: 21200 Bytes
   Possum: 83888 Bytes
   R2router: 47104 Bytes
   Revel: 141088 Bytes
   Rivet: 42840 Bytes
   Tango: 54584 Bytes
   TigerTonic: 96032 Bytes
   Traffic: 1053696 Bytes
   Vulcan: 465392 Bytes
```

### GPlusAPI Routes: 13

```
   Ace: 3600 Bytes
   Bear: 7112 Bytes
   Beego: 9712 Bytes
   Bone: 6448 Bytes
   Denco: 3256 Bytes
   Echo: 7112 Bytes
   Gin: 3856 Bytes
   GocraftWeb: 7496 Bytes
   Goji: 2912 Bytes
   Gojiv2: 7376 Bytes
   GoJsonRest: 11400 Bytes
   GoRestful: 87608 Bytes
   GorillaMux: 71072 Bytes
   HttpRouter: 2712 Bytes
   HttpTreeMux: 7376 Bytes
   Iris: 6352 Bytes
   Kocha: 128944 Bytes
   LARS: 3624 Bytes
   Macaron: 8448 Bytes
   Martini: 23936 Bytes
   Pat: 1856 Bytes
   Possum: 7728 Bytes
   R2router: 3928 Bytes
   Revel: 10768 Bytes
   Rivet: 3064 Bytes
   Tango: 4912 Bytes
   TigerTonic: 9408 Bytes
   Traffic: 49472 Bytes
   Vulcan: 25496 Bytes
```


### ParseAPI Routes: 26

```
   Ace: 6592 Bytes
   Bear: 12320 Bytes
   Beego: 18416 Bytes
   Bone: 10992 Bytes
   Denco: 4184 Bytes
   Echo: 8032 Bytes
   Gin: 6816 Bytes
   GocraftWeb: 12800 Bytes
   Goji: 5232 Bytes
   Gojiv2: 14464 Bytes
   GoJsonRest: 14248 Bytes
   GoRestful: 126216 Bytes
   GorillaMux: 122184 Bytes
   HttpRouter: 4976 Bytes
   HttpTreeMux: 7784 Bytes
   Iris: 9864 Bytes
   Kocha: 181776 Bytes
   LARS: 6616 Bytes
   Macaron: 13232 Bytes
   Martini: 45952 Bytes
   Pat: 2560 Bytes
   Possum: 9200 Bytes
   R2router: 7056 Bytes
   Revel: 15488 Bytes
   Rivet: 5680 Bytes
   Tango: 8664 Bytes
   TigerTonic: 9840 Bytes
   Traffic: 93480 Bytes
   Vulcan: 44504 Bytes
```


### Static Routes: 157

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

The only intention of this benchmark is to allow a comparison with the default router of Go's net/http package, [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is limited to static routes and does not support parameters in the route pattern.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.8% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.
```   
   HttpServeMux: 17824 Bytes
   Ace: 30080 Bytes
   Bear: 30424 Bytes
   Beego: 93768 Bytes
   Bone: 37872 Bytes
   Denco: 9984 Bytes
   Echo: 61008 Bytes
   Gin: 30400 Bytes
   GocraftWeb: 55464 Bytes
   Goji: 27200 Bytes
   Gojiv2: 104464 Bytes
   GoJsonRest: 135816 Bytes
   GoRestful: 908200 Bytes
   GorillaMux: 668496 Bytes
   HttpRouter: 21128 Bytes
   HttpTreeMux: 73384 Bytes
   Iris: 37200 Bytes
   Kocha: 114880 Bytes
   LARS: 30104 Bytes
   Macaron: 34928 Bytes
   Martini: 308784 Bytes
   Pat: 20464 Bytes
   Possum: 91328 Bytes
   R2router: 23712 Bytes
   Revel: 94016 Bytes
   Rivet: 23880 Bytes
   Tango: 28008 Bytes
   TigerTonic: 80320 Bytes
   Traffic: 624432 Bytes
   Vulcan: 368728 Bytes
```


### Micro Benchmarks

The following benchmarks measure the cost of some very basic operations.

In the first benchmark, only a single route, containing a parameter, is loaded into the routers. Then a request for a URL matching this pattern is made and the router has to call the respective registered handler function. End.
```
BenchmarkAce_Param               	 5000000	       359 ns/op	      32 B/op	       1 allocs/op
BenchmarkBear_Param              	 1000000	      1389 ns/op	     456 B/op	       5 allocs/op
BenchmarkBeego_Param             	 2000000	       698 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_Param              	 1000000	      1427 ns/op	     384 B/op	       3 allocs/op
BenchmarkDenco_Param             	 5000000	       258 ns/op	      32 B/op	       1 allocs/op
BenchmarkEcho_Param              	20000000	       101 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param               	20000000	        91.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param        	 1000000	      2074 ns/op	     648 B/op	       8 allocs/op
BenchmarkGoji_Param              	 2000000	       983 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_Param            	 2000000	       846 ns/op	     176 B/op	       5 allocs/op
BenchmarkGoJsonRest_Param        	 1000000	      2112 ns/op	     649 B/op	      13 allocs/op
BenchmarkGoRestful_Param         	  200000	      8894 ns/op	    2696 B/op	      27 allocs/op
BenchmarkGorillaMux_Param        	 1000000	      3456 ns/op	     752 B/op	       8 allocs/op
BenchmarkHttpRouter_Param        	10000000	       188 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param       	 1000000	      1011 ns/op	     352 B/op	       3 allocs/op
BenchmarkIris_Param              	30000000	        54.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_Param             	 3000000	       459 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_Param              	20000000	        93.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Param           	 1000000	      3827 ns/op	    1040 B/op	       9 allocs/op
BenchmarkMartini_Param           	  500000	      5591 ns/op	    1104 B/op	      11 allocs/op
BenchmarkPat_Param               	 1000000	      2898 ns/op	     648 B/op	      12 allocs/op
BenchmarkPossum_Param            	 1000000	      2287 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_Param          	 1000000	      1409 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_Param             	  300000	      6641 ns/op	    1632 B/op	      26 allocs/op
BenchmarkRivet_Param             	 5000000	       333 ns/op	      48 B/op	       1 allocs/op
BenchmarkTango_Param             	 1000000	      1679 ns/op	     256 B/op	       9 allocs/op
BenchmarkTigerTonic_Param        	  500000	      4152 ns/op	     976 B/op	      16 allocs/op
BenchmarkTraffic_Param           	  200000	      6667 ns/op	    1960 B/op	      21 allocs/op
BenchmarkVulcan_Param            	 2000000	       973 ns/op	      98 B/op	       3 allocs/op
```

Same as before, but now with multiple parameters, all in the same single route. The intention is to see how the routers scale with the number of parameters. The values of the parameters must be passed to the handler function somehow, which requires allocations. Let's see how clever the routers solve this task with a route containing 5 and 20 parameters:
```
BenchmarkAce_Param5              	 2000000	       644 ns/op	     160 B/op	       1 allocs/op
BenchmarkBear_Param5             	 1000000	      1867 ns/op	     501 B/op	       5 allocs/op
BenchmarkBeego_Param5            	 2000000	       860 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_Param5             	 1000000	      1767 ns/op	     432 B/op	       3 allocs/op
BenchmarkDenco_Param5            	 2000000	       686 ns/op	     160 B/op	       1 allocs/op
BenchmarkEcho_Param5             	10000000	       172 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param5              	10000000	       162 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param5       	 1000000	      3280 ns/op	     920 B/op	      11 allocs/op
BenchmarkGoji_Param5             	 1000000	      1286 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_Param5           	 1000000	      1195 ns/op	     240 B/op	       5 allocs/op
BenchmarkGoJsonRest_Param5       	  500000	      4241 ns/op	    1097 B/op	      16 allocs/op
BenchmarkGoRestful_Param5        	  200000	     11228 ns/op	    2872 B/op	      27 allocs/op
BenchmarkGorillaMux_Param5       	  500000	      5092 ns/op	     816 B/op	       8 allocs/op
BenchmarkHttpRouter_Param5       	 3000000	       479 ns/op	     160 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param5      	 1000000	      2032 ns/op	     576 B/op	       6 allocs/op
BenchmarkIris_Param5             	30000000	        50.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_Param5            	 1000000	      1898 ns/op	     440 B/op	      10 allocs/op
BenchmarkLARS_Param5             	10000000	       156 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Param5          	  500000	      3848 ns/op	    1040 B/op	       9 allocs/op
BenchmarkMartini_Param5          	  300000	      6557 ns/op	    1232 B/op	      11 allocs/op
BenchmarkPat_Param5              	  300000	      5738 ns/op	     964 B/op	      32 allocs/op
BenchmarkPossum_Param5           	 1000000	      2232 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_Param5         	 1000000	      1650 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_Param5            	  200000	      7602 ns/op	    1984 B/op	      33 allocs/op
BenchmarkRivet_Param5            	 2000000	       860 ns/op	     240 B/op	       1 allocs/op
BenchmarkTango_Param5            	 1000000	      3625 ns/op	     944 B/op	      17 allocs/op
BenchmarkTigerTonic_Param5       	  200000	     12190 ns/op	    2471 B/op	      38 allocs/op
BenchmarkTraffic_Param5          	  200000	      8879 ns/op	    2248 B/op	      25 allocs/op
BenchmarkVulcan_Param5           	 1000000	      1151 ns/op	      98 B/op	       3 allocs/op
 
BenchmarkAce_Param20             	 1000000	      1732 ns/op	     640 B/op	       1 allocs/op
BenchmarkBear_Param20            	  300000	      5446 ns/op	    1665 B/op	       5 allocs/op
BenchmarkBeego_Param20           	 1000000	      2469 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_Param20            	  200000	      8268 ns/op	    2538 B/op	       5 allocs/op
BenchmarkDenco_Param20           	 1000000	      2204 ns/op	     640 B/op	       1 allocs/op
BenchmarkEcho_Param20            	 3000000	       472 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param20             	 5000000	       370 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param20      	  200000	     11590 ns/op	    3795 B/op	      15 allocs/op
BenchmarkGoji_Param20            	  500000	      4677 ns/op	    1247 B/op	       2 allocs/op
BenchmarkGojiv2_Param20          	 1000000	      2260 ns/op	     480 B/op	       5 allocs/op
BenchmarkGoJsonRest_Param20      	  100000	     15627 ns/op	    4484 B/op	      20 allocs/op
BenchmarkGoRestful_Param20       	  100000	     21221 ns/op	    5444 B/op	      29 allocs/op
BenchmarkGorillaMux_Param20      	  100000	     12336 ns/op	    2924 B/op	      10 allocs/op
BenchmarkHttpRouter_Param20      	 1000000	      1628 ns/op	     640 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param20     	  200000	     10204 ns/op	    3195 B/op	      10 allocs/op
BenchmarkKocha_Param20           	  300000	      6263 ns/op	    1808 B/op	      27 allocs/op
BenchmarkLARS_Param20            	 5000000	       370 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Param20         	  200000	     10388 ns/op	    2892 B/op	      11 allocs/op
BenchmarkMartini_Param20         	  100000	     15105 ns/op	    3596 B/op	      13 allocs/op
BenchmarkPat_Param20             	   50000	     27182 ns/op	    4687 B/op	     111 allocs/op
BenchmarkPossum_Param20          	 1000000	      2244 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_Param20        	  200000	      7788 ns/op	    2284 B/op	       7 allocs/op
BenchmarkRevel_Param20           	  100000	     17886 ns/op	    5512 B/op	      52 allocs/op
BenchmarkRivet_Param20           	 1000000	      3172 ns/op	    1024 B/op	       1 allocs/op
BenchmarkTango_Param20           	  100000	     21303 ns/op	    8224 B/op	      47 allocs/op
BenchmarkTigerTonic_Param20      	   30000	     48806 ns/op	   10343 B/op	     118 allocs/op
BenchmarkTraffic_Param20         	   50000	     30890 ns/op	    7944 B/op	      45 allocs/op
BenchmarkVulcan_Param20          	 1000000	      1906 ns/op	      98 B/op	       3 allocs/op
```


Now let's see how expensive it is to access a parameter. The handler function reads the value (by the name of the parameter, e.g. with a map lookup; depends on the router) and writes it to our [web scale storage](https://www.youtube.com/watch?v=b2F-DItXtZs) (`/dev/null`).
```
BenchmarkAce_ParamWrite          	 3000000	       486 ns/op	      40 B/op	       2 allocs/op
BenchmarkBear_ParamWrite         	 1000000	      1430 ns/op	     456 B/op	       5 allocs/op
BenchmarkBeego_ParamWrite        	 2000000	       780 ns/op	       8 B/op	       1 allocs/op
BenchmarkBone_ParamWrite         	 1000000	      1445 ns/op	     384 B/op	       3 allocs/op
BenchmarkDenco_ParamWrite        	 5000000	       315 ns/op	      32 B/op	       1 allocs/op
BenchmarkEcho_ParamWrite         	10000000	       212 ns/op	       8 B/op	       1 allocs/op
BenchmarkGin_ParamWrite          	10000000	       188 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParamWrite   	 1000000	      2238 ns/op	     656 B/op	       9 allocs/op
BenchmarkGoji_ParamWrite         	 1000000	      1133 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_ParamWrite       	 1000000	      1173 ns/op	     208 B/op	       7 allocs/op
BenchmarkGoJsonRest_ParamWrite   	 1000000	      3670 ns/op	    1128 B/op	      18 allocs/op
BenchmarkGoRestful_ParamWrite    	  200000	      9800 ns/op	    2704 B/op	      28 allocs/op
BenchmarkGorillaMux_ParamWrite   	 1000000	      3599 ns/op	     752 B/op	       8 allocs/op
BenchmarkHttpRouter_ParamWrite   	10000000	       233 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_ParamWrite  	 1000000	      1087 ns/op	     352 B/op	       3 allocs/op
BenchmarkIris_ParamWrite         	10000000	       169 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_ParamWrite        	 3000000	       538 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_ParamWrite         	10000000	       187 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParamWrite      	  500000	      4285 ns/op	    1144 B/op	      13 allocs/op
BenchmarkMartini_ParamWrite      	  300000	      6633 ns/op	    1208 B/op	      15 allocs/op
BenchmarkPat_ParamWrite          	  500000	      4156 ns/op	    1072 B/o	      17 allocs/op
BenchmarkPossum_ParamWrite       	 1000000	      2226 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_ParamWrite     	 1000000	      1464 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_ParamWrite        	  200000	      8246 ns/op	    2096 B/op	      31 allocs/op
BenchmarkRivet_ParamWrite        	 3000000	       614 ns/op	     144 B/op	       3 allocs/op
BenchmarkTango_ParamWrite        	 2000000	       848 ns/op	     136 B/op	       4 allocs/op
BenchmarkTigerTonic_ParamWrite   	  300000	      6257 ns/op	    1408 B/op	      22 allocs/op
BenchmarkTraffic_ParamWrite      	  200000	      7946 ns/op	    2384 B/op	      25 allocs/op
BenchmarkVulcan_ParamWrite       	 2000000	       947 ns/op	      98 B/op	       3 allocs/op
```


### [GitHub](http://developer.github.com/v3/)

The GitHub API is rather large, consisting of 203 routes. The tasks are basically the same as in the benchmarks before.
```
BenchmarkAce_GithubStatic        	10000000	       241 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GithubStatic       	 2000000	       692 ns/op	     120 B/op	       3 allocs/op
BenchmarkBeego_GithubStatic      	 2000000	       731 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_GithubStatic       	  100000	     18364 ns/op	    2880 B/op	      60 allocs/op
BenchmarkDenco_GithubStatic      	20000000	        63.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubStatic       	10000000	       116 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubStatic        	10000000	       127 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubStatic 	 1000000	      1347 ns/op	     296 B/op	       5 allocs/op
BenchmarkGoji_GithubStatic       	 5000000	       296 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_GithubStatic     	 2000000	       912 ns/op	     160 B/op	       4 allocs/op
BenchmarkGoRestful_GithubStatic  	   30000	     53375 ns/op	    3720 B/op	      32 allocs/op
BenchmarkGoJsonRest_GithubStatic 	 1000000	      1688 ns/op	     329 B/op	      11 allocs/op
BenchmarkGorillaMux_GithubStatic 	  100000	     20838 ns/op	     448 B/op	       7 allocs/op
BenchmarkHttpRouter_GithubStatic 	20000000	        64.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_GithubStatic	20000000	        75.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIris_GithubStatic       	30000000	        56.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GithubStatic      	20000000	        90.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkLARS_GithubStatic       	10000000	       117 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GithubStatic    	 1000000	      3145 ns/op	     752 B/op	       8 allocs/op
BenchmarkMartini_GithubStatic    	  100000	     17773 ns/op	     784 B/op	      10 allocs/op
BenchmarkPat_GithubStatic        	  100000	     13798 ns/op	    3648 B/op	      76 allocs/op
BenchmarkPossum_GithubStatic     	 1000000	      1382 ns/op	     416 B/op	       3 allocs/op
BenchmarkR2router_GithubStatic   	 2000000	       758 ns/op	     144 B/op	       4 allocs/op
BenchmarkRevel_GithubStatic      	  500000	      5504 ns/op	    1248 B/op	      23 allocs/op
BenchmarkRivet_GithubStatic      	10000000	       136 ns/op	       0 B/op	       0 allocs/op
BenchmarkTango_GithubStatic      	 1000000	      1800 ns/op	     256 B/op	       9 allocs/op
BenchmarkTigerTonic_GithubStatic 	 5000000	       390 ns/op	      48 B/op	       1 allocs/op
BenchmarkTraffic_GithubStatic    	   20000	     59225 ns/op	   18904 B/op	     148 allocs/op
BenchmarkVulcan_GithubStatic     	 1000000	      1386 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_GithubParam         	 2000000	       648 ns/op	      96 B/op	       1 allocs/op
BenchmarkBear_GithubParam        	 1000000	      1695 ns/op	     496 B/op	       5 allocs/op
BenchmarkBeego_GithubParam       	 2000000	       916 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_GithubParam        	  300000	      7186 ns/op	    1456 B/op	      16 allocs/op
BenchmarkDenco_GithubParam       	 3000000	       606 ns/op	     128 B/op	       1 allocs/op
BenchmarkEcho_GithubParam        	10000000	       202 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubParam         	10000000	       208 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubParam  	 1000000	      2597 ns/op	     712 B/op	       9 allocs/op
BenchmarkGoji_GithubParam        	 1000000	      1490 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_GithubParam      	 1000000	      1990 ns/op	     256 B/op	       7 allocs/op
BenchmarkGoJsonRest_GithubParam  	 1000000	      3075 ns/op	     713 B/op	      14 allocs/op
BenchmarkGoRestful_GithubParam   	   10000	    142378 ns/op	    3016 B/op	      31 allocs/op
BenchmarkGorillaMux_GithubParam  	  200000	     13113 ns/op	     768 B/op	       8 allocs/op
BenchmarkHttpRouter_GithubParam  	 3000000	       444 ns/op	      96 B/op	       1 allocs/op
BenchmarkHttpTreeMux_GithubParam 	 1000000	      1335 ns/op	     384 B/op	       4 allocs/op
BenchmarkIris_GithubParam        	30000000	        55.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GithubParam       	 1000000	      1001 ns/op	     128 B/op	       5 allocs/op
BenchmarkLARS_GithubParam        	10000000	       190 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GithubParam     	 1000000	      3569 ns/op	    1040 B/op	       9 allocs/op
BenchmarkMartini_GithubParam     	  100000	     13832 ns/op	    1136 B/op	      11 allocs/op
BenchmarkPat_GithubParam         	  200000	      9163 ns/op	    2464 B/op	      48 allocs/op
BenchmarkPossum_GithubParam      	 1000000	      2124 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_GithubParam    	 1000000	      1440 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_GithubParam       	  200000	      6924 ns/op	    1744 B/op	      28 allocs/op
BenchmarkRivet_GithubParam       	 3000000	       600 ns/op	      96 B/op	       1 allocs/op
BenchmarkTango_GithubParam       	 1000000	      2719 ns/op	     480 B/op	      12 allocs/op
BenchmarkTigerTonic_GithubParam  	  300000	      6585 ns/op	    1408 B/op	      22 allocs/op
BenchmarkTraffic_GithubParam     	  100000	     22477 ns/op	    5992 B/op	      52 allocs/op
BenchmarkVulcan_GithubParam      	 1000000	      1988 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_GithubAll           	   10000	    115472 ns/op	   13792 B/op	     167 allocs/op
BenchmarkBear_GithubAll          	    5000	    356255 ns/op	   86448 B/op	     943 allocs/op
BenchmarkBeego_GithubAll         	   10000	    188006 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_GithubAll          	     500	   2719314 ns/op	  548736 B/op	    7241 allocs/op
BenchmarkDenco_GithubAll         	   10000	    113264 ns/op	   20224 B/op	     167 allocs/op
BenchmarkEcho_GithubAll          	   30000	     49096 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll           	   30000	     44661 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubAll    	    5000	    501882 ns/op	  131656 B/op	    1686 allocs/op
BenchmarkGoji_GithubAll          	    3000	    660196 ns/op	   56112 B/op	     334 allocs/op
BenchmarkGojiv2_GithubAll        	    2000	    838054 ns/op	  118864 B/op	    3103 allocs/op
BenchmarkGoJsonRest_GithubAll    	    3000	    649118 ns/op	  134371 B/op	    2737 allocs/op
BenchmarkGoRestful_GithubAll     	     100	  19901761 ns/op	  837832 B/op	    6913 allocs/op
BenchmarkGorillaMux_GithubAll    	     200	   7417761 ns/op	  144464 B/op	    1588 allocs/op
BenchmarkHttpRouter_GithubAll    	   20000	     72545 ns/op	   13792 B/op	     167 allocs/op
BenchmarkHttpTreeMux_GithubAll   	   10000	    241502 ns/op	   65856 B/op	     671 allocs/op
BenchmarkIris_GithubAll          	  100000	     22663 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GithubAll         	   10000	    188043 ns/op	   23304 B/op	     843 allocs/op
BenchmarkLARS_GithubAll          	   30000	     43820 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GithubAll       	    2000	    810354 ns/op	  201138 B/op	    1803 allocs/op
BenchmarkMartini_GithubAll       	     200	   6001410 ns/op	  228214 B/op	    2483 allocs/op
BenchmarkPat_GithubAll           	     300	   4316839 ns/op	 1499569 B/op	   27435 allocs/op
BenchmarkPossum_GithubAll        	   10000	    291630 ns/op	   84448 B/op	     609 allocs/op
BenchmarkR2router_GithubAll      	   10000	    254326 ns/op	   77328 B/op	     979 allocs/op
BenchmarkRevel_GithubAll         	    2000	   1219235 ns/op	  337424 B/op	    5512 allocs/op
BenchmarkRivet_GithubAll         	   10000	    108660 ns/op	   16272 B/op	     167 allocs/op
BenchmarkTango_GithubAll         	    5000	    449599 ns/op	   87075 B/op	    2267 allocs/op
BenchmarkTigerTonic_GithubAll    	    2000	   1140471 ns/op	  233680 B/op	    5035 allocs/op
BenchmarkTraffic_GithubAll       	     200	   8731615 ns/op	 2659331 B/op	   21848 allocs/op
BenchmarkVulcan_GithubAll        	    5000	    329104 ns/op	   19894 B/op	     609 allocs/op
```


### [Google+](https://developers.google.com/+/api/latest/)

Last but not least the Google+ API, consisting of 13 routes. In reality this is just a subset of a much larger API.
```
BenchmarkAce_GPlusStatic         	10000000	       204 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GPlusStatic        	 3000000	       476 ns/op	     104 B/op	       3 allocs/op
BenchmarkBeego_GPlusStatic       	 2000000	       636 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_GPlusStatic        	10000000	       219 ns/op	      32 B/op	       1 allocs/op
BenchmarkDenco_GPlusStatic       	50000000	        35.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GPlusStatic        	20000000	        95.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusStatic         	20000000	        90.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlusStatic  	 2000000	       949 ns/op	     280 B/op	       5 allocs/op
BenchmarkGoji_GPlusStatic        	10000000	       201 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_GPlusStatic      	 2000000	       694 ns/op	     160 B/op	       4 allocs/op
BenchmarkGoJsonRest_GPlusStatic  	 1000000	      1239 ns/op	     329 B/op	      11 allocs/op
BenchmarkGoRestful_GPlusStatic   	  200000	      8815 ns/op	    2360 B/op	      26 allocs/op
BenchmarkGorillaMux_GPlusStatic  	 1000000	      2157 ns/op	     448 B/op	       7 allocs/op
BenchmarkHttpRouter_GPlusStatic  	30000000	        38.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic 	30000000	        46.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkIris_GPlusStatic        	30000000	        53.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GPlusStatic       	20000000	        61.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkLARS_GPlusStatic        	20000000	        95.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlusStatic     	 1000000	      2390 ns/op	     752 B/op	       8 allocs/op
BenchmarkMartini_GPlusStatic     	  500000	      4118 ns/op	     784 B/op	      10 allocs/op
BenchmarkPat_GPlusStatic         	 5000000	       374 ns/op	      96 B/op	       2 allocs/op
BenchmarkPossum_GPlusStatic      	 1000000	      1296 ns/op	     416 B/op	       3 allocs/op
BenchmarkR2router_GPlusStatic    	 2000000	       613 ns/op	     144 B/op	       4 allocs/op
BenchmarkRevel_GPlusStatic       	  500000	      5128 ns/op	    1232 B/op	      23 allocs/op
BenchmarkRivet_GPlusStatic       	20000000	        86.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkTango_GPlusStatic       	 1000000	      1186 ns/op	     208 B/op	       9 allocs/op
BenchmarkTigerTonic_GPlusStatic  	10000000	       244 ns/op	      32 B/op	       1 allocs/op
BenchmarkTraffic_GPlusStatic     	 1000000	      3131 ns/op	    1192 B/op	      15 allocs/op
BenchmarkVulcan_GPlusStatic      	 2000000	       815 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_GPlusParam          	 3000000	       479 ns/op	      64 B/op	       1 allocs/op
BenchmarkBear_GPlusParam         	 1000000	      1228 ns/op	     480 B/op	       5 allocs/op
BenchmarkBeego_GPlusParam        	 2000000	       799 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_GPlusParam         	 1000000	      1132 ns/op	     384 B/op	       3 allocs/op
BenchmarkDenco_GPlusParam        	 5000000	       344 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_GPlusParam         	10000000	       134 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusParam          	10000000	       124 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlusParam   	 1000000	      1833 ns/op	     648 B/op	       8 allocs/op
BenchmarkGoji_GPlusParam         	 1000000	      1077 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_GPlusParam       	 2000000	       960 ns/op	     176 B/op	       5 allocs/op
BenchmarkGoJsonRest_GPlusParam   	 1000000	      2187 ns/op	     649 B/op	      13 allocs/op
BenchmarkGoRestful_GPlusParam    	  100000	     16238 ns/op	    2760 B/op	      29 allocs/op
BenchmarkGorillaMux_GPlusParam   	  500000	      4075 ns/op	     752 B/op	       8 allocs/op
BenchmarkHttpRouter_GPlusParam   	 5000000	       275 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_GPlusParam  	 2000000	       887 ns/op	     352 B/op	       3 allocs/op
BenchmarkIris_GPlusParam         	30000000	        55.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GPlusParam        	 3000000	       479 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_GPlusParam         	10000000	       132 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlusParam      	 1000000	      3093 ns/op	    1040 B/op	       9 allocs/op
BenchmarkMartini_GPlusParam      	  300000	      5693 ns/op	    1104 B/op	      11 allocs/op
BenchmarkPat_GPlusParam          	 1000000	      2085 ns/op	     688 B/op	      12 allocs/op
BenchmarkPossum_GPlusParam       	 1000000	      1724 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_GPlusParam     	 1000000	      1090 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_GPlusParam        	  500000	      5519 ns/op	    1664 B/op	      26 allocs/op
BenchmarkRivet_GPlusParam        	 5000000	       302 ns/op	      48 B/op	       1 allocs/op
BenchmarkTango_GPlusParam        	 1000000	      1505 ns/op	     272 B/op	       9 allocs/op
BenchmarkTigerTonic_GPlusParam   	  500000	      3583 ns/op	    1040 B/op	      16 allocs/op
BenchmarkTraffic_GPlusParam      	  500000	      5883 ns/op	    1976 B/op	      21 allocs/op
BenchmarkVulcan_GPlusParam       	 1000000	      1089 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_GPlus2Params        	 3000000	       474 ns/op	      64 B/op	       1 allocs/op
BenchmarkBear_GPlus2Params       	 1000000	      1395 ns/op	     496 B/op	       5 allocs/op
BenchmarkBeego_GPlus2Params      	 2000000	       910 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_GPlus2Params       	 1000000	      3012 ns/op	     736 B/op	       7 allocs/op
BenchmarkDenco_GPlus2Params      	 5000000	       421 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_GPlus2Params       	10000000	       170 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlus2Params        	10000000	       151 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlus2Params 	 1000000	      2200 ns/op	     712 B/op	       9 allocs/op
BenchmarkGoji_GPlus2Params       	 1000000	      1212 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_GPlus2Params     	 1000000	      1737 ns/op	     256 B/op	       8 allocs/op
BenchmarkGoJsonRest_GPlus2Params 	 1000000	      2497 ns/op	     713 B/op	      14 allocs/op
BenchmarkGoRestful_GPlus2Params  	  100000	     18489 ns/op	    2920 B/op	      31 allocs/op
BenchmarkGorillaMux_GPlus2Params 	  200000	      8099 ns/op	     768 B/op	       8 allocs/op
BenchmarkHttpRouter_GPlus2Params 	 5000000	       278 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params	 1000000	      1017 ns/op	     384 B/op	       4 allocs/op
BenchmarkIris_GPlus2Params       	30000000	        51.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GPlus2Params      	 2000000	       797 ns/op	     128 B/op	       5 allocs/op
BenchmarkLARS_GPlus2Params       	10000000	    358127 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlus2Params    	 1000000	      2815 ns/op	    1040 B/op	       9 allocs/op
BenchmarkMartini_GPlus2Params    	  200000	     22391 ns/op	    1232 B/op	      15 allocs/op
BenchmarkPat_GPlus2Params        	   50000	     39427 ns/op	    2256 B/op	      34 allocs/op
BenchmarkPossum_GPlus2Params     	  300000	      8552 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_GPlus2Params   	  500000	      5151 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_GPlus2Params      	  100000	     24954 ns/op	    1760 B/op	      28 allocs/op
BenchmarkRivet_GPlus2Params      	 1000000	      1886 ns/op	      96 B/op	       1 allocs/op
BenchmarkTango_GPlus2Params      	  300000	      6960 ns/op	     448 B/op	      11 allocs/op
BenchmarkTigerTonic_GPlus2Params 	  100000	     20905 ns/op	    1456 B/op	      22 allocs/op
BenchmarkTraffic_GPlus2Params    	   30000	     45991 ns/op	    3272 B/op	      31 allocs/op
BenchmarkVulcan_GPlus2Params     	  300000	  11270106 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_GPlusAll            	  300000	      6247 ns/op	     640 B/op	      11 allocs/op
BenchmarkBear_GPlusAll           	  100000	     15683 ns/op	    5488 B/op	      61 allocs/op
BenchmarkBeego_GPlusAll          	  200000	      9879 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_GPlusAll           	  100000	     20070 ns/op	    4912 B/op	      61 allocs/op
BenchmarkDenco_GPlusAll          	  500000	      4430 ns/op	     672 B/op	      11 allocs/op
BenchmarkEcho_GPlusAll           	 1000000	      2235 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusAll            	 1000000	      1848 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlusAll     	  100000	     22867 ns/op	    8040 B/op	     103 allocs/op
BenchmarkGoji_GPlusAll           	  200000	     11873 ns/op	    3696 B/op	      22 allocs/op
BenchmarkGojiv2_GPlusAll         	  100000	     14798 ns/op	    2640 B/op	      76 allocs/op
BenchmarkGoJsonRest_GPlusAll     	   50000	     26591 ns/op	    8117 B/op	     170 allocs/op
BenchmarkGoRestful_GPlusAll      	   10000	    165798 ns/op	   38664 B/op	     389 allocs/op
BenchmarkGorillaMux_GPlusAll     	   20000	     61564 ns/op	    9248 B/op	     102 allocs/op
BenchmarkHttpRouter_GPlusAll     	 1000000	      3013 ns/op	     640 B/op	      11 allocs/op
BenchmarkHttpTreeMux_GPlusAll    	  200000	     10172 ns/op	    4032 B/op	      38 allocs/op
BenchmarkIris_GPlusAll           	 1000000	      1175 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GPlusAll          	  200000	      7258 ns/op	     976 B/op	      43 allocs/op
BenchmarkLARS_GPlusAll           	 1000000	      1901 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlusAll        	   50000	     35269 ns/op	   12944 B/op	     115 allocs/op
BenchmarkMartini_GPlusAll        	   20000	     87743 ns/op	   14448 B/op	     165 allocs/op
BenchmarkPat_GPlusAll            	   30000	     49012 ns/op	   16576 B/op	     298 allocs/op
BenchmarkPossum_GPlusAll         	  100000	     15444 ns/op	    5408 B/op	      39 allocs/op
BenchmarkR2router_GPlusAll       	  100000	     13569 ns/op	    5040 B/op	      63 allocs/op
BenchmarkRevel_GPlusAll          	   20000	     68250 ns/op	   21136 B/op	     342 allocs/op
BenchmarkRivet_GPlusAll          	  500000	      4351 ns/op	     768 B/op	      11 allocs/op
BenchmarkTango_GPlusAll          	  100000	     19997 ns/op	    4304 B/op	     129 allocs/op
BenchmarkTigerTonic_GPlusAll     	   30000	     51951 ns/op	   14256 B/op	     272 allocs/op
BenchmarkTraffic_GPlusAll        	   10000	    100447 ns/op	   37360 B/op	     392 allocs/op
BenchmarkVulcan_GPlusAll         	  100000	     14706 ns/op	    1274 B/op	      39 allocs/op
```


### [Parse.com](https://parse.com/docs/rest#summary)

Enough of the micro benchmark stuff. Let's play a bit with real APIs. In the first set of benchmarks, we use a clone of the structure of [Parse](https://parse.com)'s decent medium-sized REST API, consisting of 26 routes.

The tasks are 1.) routing a static URL (no parameters), 2.) routing a URL containing 1 parameter, 3.) same with 2 parameters, 4.) route all of the routes once (like the StaticAll benchmark, but the routes now contain parameters).

Worth noting is, that the requested route might be a good case for some routing algorithms, while it is a bad case for another algorithm. The values might vary slightly depending on the selected route.

```
BenchmarkAce_ParseStatic         	10000000	       199 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_ParseStatic        	 3000000	       502 ns/op	     120 B/op	       3 allocs/op
BenchmarkBeego_ParseStatic       	 2000000	       610 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_ParseStatic        	 2000000	       702 ns/op	     144 B/op	       3 allocs/op
BenchmarkDenco_ParseStatic       	30000000	        43.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_ParseStatic        	20000000	        92.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseStatic         	20000000	        87.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParseStatic  	 2000000	       955 ns/op	     296 B/op	       5 allocs/op
BenchmarkGoji_ParseStatic        	 5000000	       246 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_ParseStatic      	 2000000	       630 ns/op	     160 B/op	       4 allocs/op
BenchmarkGoJsonRest_ParseStatic  	 1000000	      1213 ns/op	     329 B/op	      11 allocs/op
BenchmarkGoRestful_ParseStatic   	  200000	     12025 ns/op	    3656 B/op	      30 allocs/op
BenchmarkGorillaMux_ParseStatic  	  500000	      2867 ns/op	     448 B/op	       7 allocs/op
BenchmarkHttpRouter_ParseStatic  	50000000	        37.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_ParseStatic 	20000000	        72.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkIris_ParseStatic        	30000000	        51.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_ParseStatic       	30000000	        59.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkLARS_ParseStatic        	20000000	        85.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParseStatic     	 1000000	      2275 ns/op	     752 B/op	       8 allocs/op
BenchmarkMartini_ParseStatic     	  500000	      4318 ns/op	     784 B/op	      10 allocs/op
BenchmarkPat_ParseStatic         	 2000000	       796 ns/op	     240 B/op	       5 allocs/op
BenchmarkPossum_ParseStatic      	 1000000	      1159 ns/op	     416 B/op	       3 allocs/op
BenchmarkR2router_ParseStatic    	 2000000	       609 ns/op	     144 B/op	       4 allocs/op
BenchmarkRevel_ParseStatic       	  500000	      4432 ns/op	    1248 B/op	      23 allocs/op
BenchmarkRivet_ParseStatic       	20000000	        85.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkTango_ParseStatic       	 1000000	      1227 ns/op	     256 B/op	       9 allocs/op
BenchmarkTigerTonic_ParseStatic  	 5000000	       306 ns/op	      48 B/op	       1 allocs/op
BenchmarkTraffic_ParseStatic     	 1000000	      4133 ns/op	    1816 B/op	      20 allocs/op
BenchmarkVulcan_ParseStatic      	 2000000	       829 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_ParseParam          	 5000000	       380 ns/op	      64 B/op	       1 allocs/op
BenchmarkBear_ParseParam         	 1000000	      1065 ns/op	     467 B/op	       5 allocs/op
BenchmarkBeego_ParseParam        	 2000000	       673 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_ParseParam         	 1000000	      1300 ns/op	     464 B/op	       4 allocs/op
BenchmarkDenco_ParseParam        	 5000000	       305 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_ParseParam         	20000000	       106 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseParam          	20000000	        92.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParseParam   	 1000000	      1639 ns/op	     664 B/op	       8 allocs/op
BenchmarkGoji_ParseParam         	 2000000	       934 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_ParseParam       	 2000000	       968 ns/op	     208 B/op	       6 allocs/op
BenchmarkGoJsonRest_ParseParam   	 1000000	      1803 ns/op	     649 B/op	      13 allocs/op
BenchmarkGoRestful_ParseParam    	  200000	     13474 ns/op	    4024 B/op	      31 allocs/op
BenchmarkGorillaMux_ParseParam   	  500000	      3169 ns/op	     752 B/op	       8 allocs/op
BenchmarkHttpRouter_ParseParam   	10000000	       215 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_ParseParam  	 2000000	       737 ns/op	     352 B/op	       3 allocs/op
BenchmarkIris_ParseParam         	30000000	        53.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_ParseParam        	 3000000	       441 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_ParseParam         	20000000	       108 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParseParam      	 1000000	      3065 ns/op	    1040 B/op	       9 allocs/op
BenchmarkMartini_ParseParam      	  500000	      5376 ns/op	    1104 B/op	      11 allocs/op
BenchmarkPat_ParseParam          	 1000000	      3171 ns/op	    1120 B/op	      17 allocs/op
BenchmarkPossum_ParseParam       	 1000000	      1840 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_ParseParam     	 1000000	      1131 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_ParseParam        	  500000	      5826 ns/op	    1664 B/op	      26 allocs/op
BenchmarkRivet_ParseParam        	 5000000	       280 ns/op	      48 B/op	       1 allocs/op
BenchmarkTango_ParseParam        	 1000000	      1412 ns/op	     288 B/op	       9 allocs/op
BenchmarkTigerTonic_ParseParam   	 1000000	      3558 ns/op	     992 B/op	      16 allocs/op
BenchmarkTraffic_ParseParam      	  500000	      6026 ns/op	    2248 B/op	      23 allocs/op
BenchmarkVulcan_ParseParam       	 1000000	      1034 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_Parse2Params        	 3000000	       460 ns/op	      64 B/op	       1 allocs/op
BenchmarkBear_Parse2Params       	 1000000	      1384 ns/op	     496 B/op	       5 allocs/op
BenchmarkBeego_Parse2Params      	 2000000	       811 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_Parse2Params       	 1000000	      1329 ns/op	     416 B/op	       3 allocs/op
BenchmarkDenco_Parse2Params      	 5000000	       397 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_Parse2Params       	10000000	       143 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Parse2Params        	10000000	       127 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Parse2Params 	 1000000	      2177 ns/op	     712 B/op	       9 allocs/op
BenchmarkGoji_Parse2Params       	 1000000	      1052 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_Parse2Params     	 1000000	      1040 ns/op	     192 B/op	       5 allocs/op
BenchmarkGoJsonRest_Parse2Params 	 1000000	      2567 ns/op	     713 B/op	      14 allocs/op
BenchmarkGoRestful_Parse2Params  	  100000	     24052 ns/op	    6856 B/op	      39 allocs/op
BenchmarkGorillaMux_Parse2Params 	  500000	      3918 ns/op	     768 B/op	       8 allocs/op
BenchmarkHttpRouter_Parse2Params 	 5000000	       258 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Parse2Params	 1000000	      1047 ns/op	     384 B/op	       4 allocs/op
BenchmarkIris_Parse2Params       	30000000	        50.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_Parse2Params      	 2000000	       744 ns/op	     128 B/op	       5 allocs/op
BenchmarkLARS_Parse2Params       	20000000	       115 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Parse2Params    	 1000000	      2794 ns/op	    1040 B/op	       9 allocs/op
BenchmarkMartini_Parse2Params    	  500000	      4938 ns/op	    1136 B/op	      11 allocs/op
BenchmarkPat_Parse2Params        	 1000000	      2746 ns/op	     832 B/op	      17 allocs/op
BenchmarkPossum_Parse2Params     	 1000000	      1635 ns/op	     560 B/op	       6 allocs/op
BenchmarkR2router_Parse2Params   	 1000000	      1174 ns/op	     432 B/op	       5 allocs/op
BenchmarkRevel_Parse2Params      	  500000	      6121 ns/op	    1728 B/op	      28 allocs/op
BenchmarkRivet_Parse2Params      	 3000000	       433 ns/op	      96 B/op	       1 allocs/op
BenchmarkTango_Parse2Params      	 1000000	      1947 ns/op	     416 B/op	      11 allocs/op
BenchmarkTigerTonic_Parse2Params 	  500000	      5573 ns/op	    1376 B/op	      22 allocs/op
BenchmarkTraffic_Parse2Params    	  500000	      6020 ns/op	    2040 B/op	      22 allocs/op
BenchmarkVulcan_Parse2Params     	 1000000	      1215 ns/op	      98 B/op	       3 allocs/op
```

```
BenchmarkAce_ParseAll            	  200000	     10065 ns/op	     640 B/op	      16 allocs/op
BenchmarkBear_ParseAll           	   50000	     28476 ns/op	    8928 B/op	     110 allocs/op
BenchmarkBeego_ParseAll          	  100000	     18675 ns/op	       0 B/op	       0 allocs/op
BenchmarkBone_ParseAll           	   50000	     31888 ns/op	    8048 B/op	      90 allocs/op
BenchmarkDenco_ParseAll          	  200000	      7168 ns/op	     928 B/op	      16 allocs/op
BenchmarkEcho_ParseAll           	  300000	      4464 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseAll            	  500000	      3691 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParseAll     	   50000	     39247 ns/op	   13728 B/op	     181 allocs/op
BenchmarkGoji_ParseAll           	  100000	     20836 ns/op	    5376 B/op	      32 allocs/op
BenchmarkGojiv2_ParseAll         	  100000	     23901 ns/op	    4496 B/op	     121 allocs/op
BenchmarkGoJsonRest_ParseAll     	   30000	     50044 ns/op	   13866 B/op	     321 allocs/op
BenchmarkGoRestful_ParseAll      	    5000	    461934 ns/op	  125600 B/op	     868 allocs/op
BenchmarkGorillaMux_ParseAll     	   10000	    132990 ns/op	   16560 B/op	     198 allocs/op
BenchmarkHttpRouter_ParseAll     	  300000	      4675 ns/op	     640 B/op	      16 allocs/op
BenchmarkHttpTreeMux_ParseAll    	  100000	     15765 ns/op	    5728 B/op	      51 allocs/op
BenchmarkIris_ParseAll           	 1000000	      2208 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_ParseAll          	  200000	      9909 ns/op	    1112 B/op	      54 allocs/op
BenchmarkLARS_ParseAll           	  500000	      3476 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParseAll        	   20000	     68841 ns/op	   24160 B/op	     224 allocs/op
BenchmarkMartini_ParseAll        	   10000	    133464 ns/op	   25600 B/op	     276 allocs/op
BenchmarkPat_ParseAll            	   30000	     53668 ns/op	   17264 B/op	     343 allocs/op
BenchmarkPossum_ParseAll         	   50000	     30886 ns/op	   10816 B/op	      78 allocs/op
BenchmarkR2router_ParseAll       	  100000	     25608 ns/op	    8352 B/op	     120 allocs/op
BenchmarkRevel_ParseAll          	   10000	    142645 ns/op	   39424 B/op	     652 allocs/op
BenchmarkRivet_ParseAll          	  200000	      7406 ns/op	     912 B/op	      16 allocs/op
BenchmarkTango_ParseAll          	   30000	     40996 ns/op	    7664 B/op	     240 allocs/op
BenchmarkTigerTonic_ParseAll     	   20000	     77872 ns/op	   19424 B/op	     360 allocs/op
BenchmarkTraffic_ParseAll        	   10000	    162703 ns/op	   57776 B/op	     642 allocs/op
BenchmarkVulcan_ParseAll         	   50000	     31148 ns/op	    2548 B/op	      78 allocs/op
```



### Static Routes

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

The only intention of this benchmark is to allow a comparison with the default router of Go's net/http package, [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is limited to static routes and does not support parameters in the route pattern.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.8% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.
```
BenchmarkAce_StaticAll           	   30000	     47612 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpServeMux_StaticAll  	    2000	    808918 ns/op	      96 B/op	       8 allocs/op
BenchmarkBeego_StaticAll         	   10000	    140125 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_StaticAll          	   10000	    111635 ns/op	   20336 B/op	     461 allocs/op
BenchmarkBone_StaticAll          	   20000	     79318 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_StaticAll         	  100000	     14002 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_StaticAll          	   50000	     32832 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_StaticAll           	   50000	     29305 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_StaticAll    	   10000	    180748 ns/op	   46440 B/op	     785 allocs/op
BenchmarkGoji_StaticAll          	   20000	     61934 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_StaticAll        	   10000	    153009 ns/op	   25120 B/op	     628 allocs/op
BenchmarkGoJsonRest_StaticAll    	   10000	    264210 ns/op	   51653 B/op	    1727 allocs/op
BenchmarkGoRestful_StaticAll     	     300	   5580313 ns/op	  392312 B/op	    4694 allocs/op
BenchmarkGorillaMux_StaticAll    	    1000	   2029901 ns/op	   70432 B/op	    1107 allocs/op
BenchmarkHttpRouter_StaticAll    	  100000	     18088 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_StaticAll   	  100000	     17547 ns/op	       0 B/op	       0 allocs/op
BenchmarkIris_StaticAll          	  100000	     18780 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_StaticAll         	   50000	     27329 ns/op	       0 B/op	       0 allocs/op
BenchmarkLARS_StaticAll          	   50000	     30731 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_StaticAll       	    5000	    419889 ns/op	  118065 B/op	    1256 allocs/op
BenchmarkMartini_StaticAll       	     500	   2774058 ns/op	  132818 B/op	    2178 allocs/op
BenchmarkPat_StaticAll           	    1000	   1736504 ns/op	  533904 B/op	   11123 allocs/op
BenchmarkPossum_StaticAll        	   10000	    190275 ns/op	   65312 B/op	     471 allocs/op
BenchmarkR2router_StaticAll      	   10000	    123840 ns/op	   22608 B/op	     628 allocs/op
BenchmarkRevel_StaticAll         	    3000	    751825 ns/op	  198240 B/op	    3611 allocs/op
BenchmarkRivet_StaticAll         	   50000	     36155 ns/op	       0 B/op	       0 allocs/op
BenchmarkTango_StaticAll         	   10000	    282255 ns/op	   40481 B/op	    1413 allocs/op
BenchmarkTigerTonic_StaticAll    	   20000	     72426 ns/op	    7504 B/op	     157 allocs/op
BenchmarkTraffic_StaticAll       	    1000	   2265895 ns/op	  729736 B/op	   14287 allocs/op
BenchmarkVulcan_StaticAll        	   10000	    219291 ns/op	   15386 B/op	     471 allocs/op
```
