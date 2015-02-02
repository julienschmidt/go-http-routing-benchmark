Go HTTP Router Benchmark
========================

This benchmark suite aims to compare the performance of HTTP request routers for [Go](https://golang.org) by implementing the routing structure of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.

Of course the tested routers can be used for any kind of HTTP request → handler function routing, not only (REST) APIs.


#### Tested routers & frameworks:

 * [Ace](https://github.com/plimble/ace)
 * [Bear](https://github.com/ursiform/bear)
 * [Beego](http://beego.me/)
 * [go-json-rest](https://github.com/ant0ine/go-json-rest)
 * [Denco](https://github.com/naoina/denco)
 * [Gocraft Web](https://github.com/gocraft/web)
 * [Goji](https://github.com/zenazn/goji/)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [HttpTreeMux](https://github.com/dimfeld/httptreemux)
 * [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter)
 * [Martini](https://github.com/go-martini/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)


## Motivation

Go is a great language for web applications. Since the [default *request multiplexer*](http://golang.org/pkg/net/http/#ServeMux) of Go's net/http package is very simple and limited, an accordingly high number of HTTP request routers exist.

Unfortunately, most of the (early) routers use pretty bad routing algorithms. Moreover, many of them are very wasteful with memory allocations, which can become a problem in a language with Garbage Collection like Go, since every (heap) allocation results in more work for the Garbage Collector.

Lately more and more bloated frameworks pop up, outdoing one another in the number of features. This benchmark tries to measure their overhead.

Beware that we are comparing apples to oranges here, we compare feature-rich frameworks to packages with simple routing functionality only. But since we are only interested in decent request routing, I think this is not entirely unfair. The frameworks are configured to do as little additional work as possible.

If you care about performance, this benchmark can maybe help you find the right router, which scales with your application.

Personally, I prefer slim and optimized software, which is why I implemented [HttpRouter](https://github.com/julienschmidt/httprouter), which is also tested here. In fact, this benchmark suite started as part of the packages tests, but was then extended to a generic benchmark suite.
So keep in mind, that I am not completely unbiased :relieved:


## Results

Benchmark System:
 * Intel Core i7 (4x 2.7GHz)
 * 2x 8 GiB DDR3-1600 RAM
 * go version go1.4 darwin/amd64
 * OS X, 10.10.2


### Memory Consumption

Besides the micro-benchmarks, there are 3 sets of benchmarks where we play around with clones of some real-world APIs, and one benchmark with static routes only, to allow a comparison with [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).
The following table shows the memory required only for loading the routing structure for the respective API.
The best 3 values for each test are bold. I'm pretty sure you can detect a pattern :wink:

| Router       | Static    | GitHub     | Google+   | Parse     |
|:-------------|----------:|-----------:|----------:|----------:|
| HttpServeMux |  17728 B  |         _  |        _  |        _  |
| Ace          |  37720 B  |   60520 B  |   4280 B  |   7968 B  |
| Bear         |  36760 B  |   92008 B  |   7672 B  |  13344 B  |
| Beego        | 103032 B  |  164984 B  |  11368 B  |  21040 B  |
| Bone         |  40560 B  |   92960 B  |   6160 B  |  10672 B  |
| Denco        | __9408 B__|   36440 B  |   3256 B  | __4184 B__|
| Gin          |  41760 B  |   66896 B  |   4800 B  |   8712 B  |
| GocraftWeb   |  51288 B  |   88840 B  |   7032 B  |  12128 B  |
| Goji         |  29728 B  |   50568 B  | __3136 B__|   5664 B  |
| GoJsonRest   | 142256 B  |  140640 B  |  17696 B  |  20368 B  |
| GoRestful    | 886200 B  | 1369208 B  |  89176 B  | 142856 B  |
| GorillaMux   | 668512 B  | 1493136 B  |  71056 B  | 122136 B  |
| HttpRouter   |  25064 B  |   44184 B  |   3144 B  |   5792 B  |
| HttpTreeMux  |  73288 B  |   74992 B  |   7120 B  |   7616 B  |
| Kocha        | 114816 B  |  785120 B  | 128880 B  | 181712 B  |
| Macaron      |  52168 B  |  102664 B  |   8264 B  |  13384 B  |
| Martini      | 309040 B  |  476992 B  |  23904 B  |  45952 B  |
| Pat          |__15704 B__| __15832 B__| __1448 B__| __1976 B__|
| Revel        |  93600 B  |  141296 B  |  10768 B  |  15488 B  |
| Rivet        |  42872 B  |  110080 B  |   7896 B  |  13496 B  |
| Tango        | 256032 B  | __22864 B__|   3168 B  |   4432 B  |
| TigerTonic   |  80736 B  |   96240 B  |   9408 B  |   9840 B  |
| Traffic      | 624416 B  | 1053488 B  |  49488 B  |  93496 B  |
| Vulcan       | 624416 B  |  426208 B  |  25496 B  |  44504 B  |
| Zeus         |__14560 B__| __15840 B__| __1456 B__| __1984 B__|




Main memory is cheap and usually not a scarce resource. As long as the router doesn't require Megabytes of memory, it should be no deal breaker. But it gives us a first hint how efficient or wasteful a router works.


### Static Routes

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

The only intention of this benchmark is to allow a comparison with the default router of Go's net/http package, [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is limited to static routes and does not support parameters in the route pattern.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.5% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.

```
BenchmarkHttpServeMux_StaticAll  2000       845916 ns/op          96 B/op          8 allocs/op

BenchmarkAce_StaticAll           50000       26512 ns/op           0 B/op          0 allocs/op
BenchmarkBear_StaticAll          10000      129529 ns/op       26520 B/op        619 allocs/op
BenchmarkBeego_StaticAll         10000      214671 ns/op       57744 B/op       1099 allocs/op
BenchmarkBone_StaticAll          30000       42939 ns/op           0 B/op          0 allocs/op
BenchmarkDenco_StaticAll         100000      12685 ns/op           0 B/op          0 allocs/op
BenchmarkGin_StaticAll           50000       28448 ns/op           0 B/op          0 allocs/op
BenchmarkGocraftWeb_StaticAll    10000      152033 ns/op       47696 B/op        942 allocs/op
BenchmarkGoji_StaticAll          20000       63010 ns/op           0 B/op          0 allocs/op
BenchmarkGoJsonRest_StaticAll    300       5017869 ns/op      610135 B/op      15066 allocs/op
BenchmarkGoRestful_StaticAll     300       4752691 ns/op      294345 B/op       5007 allocs/op
BenchmarkGorillaMux_StaticAll    1000      1745932 ns/op       72944 B/op       1264 allocs/op
BenchmarkHttpRouter_StaticAll    100000      14535 ns/op           0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll   100000      14595 ns/op           0 B/op          0 allocs/op
BenchmarkKocha_StaticAll         100000      22398 ns/op           0 B/op          0 allocs/op
BenchmarkMacaron_StaticAll       5000       340507 ns/op      118032 B/op       1256 allocs/op
BenchmarkMartini_StaticAll       500       3258465 ns/op      140352 B/op       2335 allocs/op
BenchmarkPat_StaticAll           1000      1488026 ns/op      533904 B/op      11123 allocs/op
BenchmarkRevel_StaticAll         2000       804102 ns/op      204528 B/op       3925 allocs/op
BenchmarkRivet_StaticAll         20000       69601 ns/op       17584 B/op        314 allocs/op
BenchmarkTango_StaticAll         1000      2360343 ns/op       54848 B/op       1544 allocs/op
BenchmarkTigerTonic_StaticAll    30000       54839 ns/op        7504 B/op        157 allocs/op
BenchmarkTraffic_StaticAll       1000      1902986 ns/op      732248 B/op      14444 allocs/op
BenchmarkVulcan_StaticAll        1000      1926207 ns/op      732248 B/op      14444 allocs/op
BenchmarkZeus_StaticAll          30000       53468 ns/op        4736 B/op         48 allocs/op
```

### Micro Benchmarks

The following benchmarks measure the cost of some very basic operations.

In the first benchmark, only a single route, containing a parameter, is loaded into the routers. Then a request for a URL matching this pattern is made and the router has to call the respective registered handler function. End.
```
BenchmarkAce_Param          10000000    188 ns/op          32 B/op          1 allocs/op
BenchmarkBear_Param         1000000    1174 ns/op         496 B/op          6 allocs/op
BenchmarkBeego_Param        1000000    1854 ns/op         720 B/op         10 allocs/op
BenchmarkBone_Param         2000000     831 ns/op         384 B/op          3 allocs/op
BenchmarkDenco_Param        10000000    196 ns/op          32 B/op          1 allocs/op
BenchmarkGin_Param          10000000    201 ns/op          32 B/op          1 allocs/op
BenchmarkGocraftWeb_Param   1000000    1455 ns/op         656 B/op          9 allocs/op
BenchmarkGoji_Param         2000000     736 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_Param   50000     31624 ns/op        4208 B/op         98 allocs/op
BenchmarkGoRestful_Param    500000     2812 ns/op         568 B/op         16 allocs/op
BenchmarkGorillaMux_Param   1000000    2859 ns/op         784 B/op          9 allocs/op
BenchmarkHttpRouter_Param   10000000    131 ns/op          32 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param  3000000     586 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_Param        5000000     352 ns/op          56 B/op          3 allocs/op
BenchmarkMacaron_Param      1000000    3072 ns/op        1144 B/op         13 allocs/op
BenchmarkMartini_Param      300000     5988 ns/op        1152 B/op         12 allocs/op
BenchmarkPat_Param          1000000    1704 ns/op         656 B/op         14 allocs/op
BenchmarkRevel_Param        300000     5789 ns/op        1672 B/op         28 allocs/op
BenchmarkRivet_Param        2000000     932 ns/op         464 B/op          5 allocs/op
BenchmarkTango_Param        1000000    1825 ns/op         712 B/op         13 allocs/op
BenchmarkTigerTonic_Param   1000000    2793 ns/op         992 B/op         19 allocs/op
BenchmarkTraffic_Param      300000     4558 ns/op        1984 B/op         23 allocs/op
BenchmarkVulcan_Param       2000000     758 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_Param         1000000    1027 ns/op         421 B/op          6 allocs/op
```

Same as before, but now with multiple parameters, all in the same single route. The intention is to see how the routers scale with the number of parameters. The values of the parameters must be passed to the handler function somehow, which requires allocations. Let's see how clever the routers solve this task with a route containing 5 and 20 parameters:
```
BenchmarkAce_Param5             5000000           404 ns/op         160 B/op          1 allocs/op
BenchmarkBear_Param5            1000000          1837 ns/op         560 B/op          6 allocs/op
BenchmarkBeego_Param5           1000000          2724 ns/op         992 B/op         13 allocs/op
BenchmarkBone_Param5            1000000          1196 ns/op         432 B/op          3 allocs/op
BenchmarkDenco_Param5           3000000           476 ns/op         160 B/op          1 allocs/op
BenchmarkGin_Param5             3000000           414 ns/op         160 B/op          1 allocs/op
BenchmarkGocraftWeb_Param5      1000000          2232 ns/op         928 B/op         12 allocs/op
BenchmarkGoji_Param5            1000000          1072 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_Param5      50000           33555 ns/op        4688 B/op        101 allocs/op
BenchmarkGoRestful_Param5       500000           3327 ns/op         568 B/op         16 allocs/op
BenchmarkGorillaMux_Param5      300000           4755 ns/op         912 B/op          9 allocs/op
BenchmarkHttpRouter_Param5      5000000           337 ns/op         160 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param5     2000000           878 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_Param5           1000000          1360 ns/op         440 B/op         10 allocs/op
BenchmarkMacaron_Param5         500000           3876 ns/op        1416 B/op         16 allocs/op
BenchmarkMartini_Param5         200000          12523 ns/op        1280 B/op         12 allocs/op
BenchmarkPat_Param5             500000           4322 ns/op        1008 B/op         42 allocs/op
BenchmarkRevel_Param5           200000           6823 ns/op        2024 B/op         35 allocs/op
BenchmarkRivet_Param5           1000000          1545 ns/op         528 B/op          9 allocs/op
BenchmarkTango_Param5           500000           3063 ns/op         856 B/op         29 allocs/op
BenchmarkTigerTonic_Param5      200000           9510 ns/op        2519 B/op         53 allocs/op
BenchmarkTraffic_Param5         200000           7277 ns/op        2280 B/op         31 allocs/op
BenchmarkVulcan_Param5          1000000          1023 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_Param5            1000000          1795 ns/op         592 B/op          9 allocs/op

BenchmarkAce_Param20            1000000          1138 ns/op         640 B/op          1 allocs/op
BenchmarkBear_Param20           200000           6731 ns/op        2667 B/op          8 allocs/op
BenchmarkBeego_Param20          200000           8081 ns/op        3867 B/op         17 allocs/op
BenchmarkBone_Param20           300000           5565 ns/op        2540 B/op          5 allocs/op
BenchmarkDenco_Param20          1000000          1503 ns/op         640 B/op          1 allocs/op
BenchmarkGin_Param20            1000000          1149 ns/op         640 B/op          1 allocs/op
BenchmarkGocraftWeb_Param20     200000           7407 ns/op        3803 B/op         16 allocs/op
BenchmarkGoji_Param20           1000000          3160 ns/op        1247 B/op          2 allocs/op
BenchmarkGoJsonRest_Param20     30000           42759 ns/op        8114 B/op        105 allocs/op
BenchmarkGoRestful_Param20      500000           4000 ns/op         568 B/op         16 allocs/op
BenchmarkGorillaMux_Param20     200000          10464 ns/op        3276 B/op         11 allocs/op
BenchmarkHttpRouter_Param20     1000000          1069 ns/op         640 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param20    300000           5063 ns/op        2187 B/op          4 allocs/op
BenchmarkKocha_Param20          300000           4323 ns/op        1808 B/op         27 allocs/op
BenchmarkMacaron_Param20        200000           9183 ns/op        4292 B/op         20 allocs/op
BenchmarkMartini_Param20        30000           54214 ns/op        3645 B/op         14 allocs/op
BenchmarkPat_Param20            100000          19926 ns/op        4885 B/op        151 allocs/op
BenchmarkRevel_Param20          100000          13015 ns/op        5551 B/op         54 allocs/op
BenchmarkRivet_Param20          300000           6634 ns/op        2620 B/op         26 allocs/op
BenchmarkTango_Param20          200000          11355 ns/op        3655 B/op         91 allocs/op
BenchmarkTigerTonic_Param20     50000           38843 ns/op       10542 B/op        178 allocs/op
BenchmarkTraffic_Param20        100000          23000 ns/op        7997 B/op         66 allocs/op
BenchmarkVulcan_Param20         1000000          1774 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_Param20           200000           7830 ns/op        3356 B/op         26 allocs/op
```

Now let's see how expensive it is to access a parameter. The handler function reads the value (by the name of the parameter, e.g. with a map lookup; depends on the router) and writes it to our [web scale storage](https://www.youtube.com/watch?v=b2F-DItXtZs) (`/dev/null`).
```
BenchmarkAce_ParamWrite  5000000           297 ns/op          40 B/op          2 allocs/op
BenchmarkBear_ParamWrite     1000000          1240 ns/op         504 B/op          7 allocs/op
BenchmarkBeego_ParamWrite    1000000          1973 ns/op         728 B/op         11 allocs/op
BenchmarkBone_ParamWrite     1000000          1030 ns/op         432 B/op          4 allocs/op
BenchmarkDenco_ParamWrite    5000000           256 ns/op          32 B/op          1 allocs/op
BenchmarkGin_ParamWrite  5000000           308 ns/op          40 B/op          2 allocs/op
BenchmarkGocraftWeb_ParamWrite   1000000          1595 ns/op         664 B/op         10 allocs/op
BenchmarkGoji_ParamWrite     2000000           794 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_ParamWrite     50000         33423 ns/op        4680 B/op        103 allocs/op
BenchmarkGoRestful_ParamWrite     500000          2830 ns/op         568 B/op         16 allocs/op
BenchmarkGorillaMux_ParamWrite    500000          3126 ns/op         792 B/op         10 allocs/op
BenchmarkHttpRouter_ParamWrite  10000000           184 ns/op          32 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParamWrite  2000000           638 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_ParamWrite    3000000           413 ns/op          56 B/op          3 allocs/op
BenchmarkMacaron_ParamWrite   500000          3551 ns/op        1208 B/op         15 allocs/op
BenchmarkMartini_ParamWrite   200000          6896 ns/op        1256 B/op         16 allocs/op
BenchmarkPat_ParamWrite  1000000          2875 ns/op        1088 B/op         19 allocs/op
BenchmarkRevel_ParamWrite     200000          6750 ns/op        2128 B/op         33 allocs/op
BenchmarkRivet_ParamWrite    1000000          1067 ns/op         472 B/op          6 allocs/op
BenchmarkTango_ParamWrite    1000000          1859 ns/op         712 B/op         13 allocs/op
BenchmarkTigerTonic_ParamWrite    300000          4566 ns/op        1440 B/op         25 allocs/op
BenchmarkTraffic_ParamWrite   300000          5846 ns/op        2400 B/op         27 allocs/op
BenchmarkVulcan_ParamWrite   2000000           758 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_ParamWrite     1000000          1253 ns/op         469 B/op          7 allocs/op
```

### [Parse.com](https://parse.com/docs/rest#summary)

Enough of the micro benchmark stuff. Let's play a bit with real APIs. In the first set of benchmarks, we use a clone of the structure of [Parse](https://parse.com)'s decent medium-sized REST API, consisting of 26 routes.

The tasks are 1.) routing a static URL (no parameters), 2.) routing a URL containing 1 parameter, 3.) same with 2 parameters, 4.) route all of the routes once (like the StaticAll benchmark, but the routes now contain parameters).

Worth noting is, that the requested route might be a good case for some routing algorithms, while it is a bad case for another algorithm. The values might vary slightly depending on the selected route.

```
BenchmarkAce_ParseStatic    20000000            95.3 ns/op         0 B/op          0 allocs/op
BenchmarkBear_ParseStatic    2000000           647 ns/op         160 B/op          4 allocs/op
BenchmarkBeego_ParseStatic   1000000          1190 ns/op         368 B/op          7 allocs/op
BenchmarkBone_ParseStatic    3000000           480 ns/op         144 B/op          3 allocs/op
BenchmarkDenco_ParseStatic  30000000            44.6 ns/op         0 B/op          0 allocs/op
BenchmarkGin_ParseStatic    20000000           110 ns/op           0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParseStatic  2000000           880 ns/op         304 B/op          6 allocs/op
BenchmarkGoji_ParseStatic    5000000           284 ns/op           0 B/op          0 allocs/op
BenchmarkGoJsonRest_ParseStatic    50000         31277 ns/op        3649 B/op         95 allocs/op
BenchmarkGoRestful_ParseStatic    200000         10855 ns/op        2904 B/op         32 allocs/op
BenchmarkGorillaMux_ParseStatic   500000          2826 ns/op         464 B/op          8 allocs/op
BenchmarkHttpRouter_ParseStatic 50000000            36.3 ns/op         0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic    20000000            60.2 ns/op         0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic  20000000            66.3 ns/op         0 B/op          0 allocs/op
BenchmarkMacaron_ParseStatic     1000000          2099 ns/op         752 B/op          8 allocs/op
BenchmarkMartini_ParseStatic      300000          4974 ns/op         832 B/op         11 allocs/op
BenchmarkPat_ParseStatic     2000000           757 ns/op         240 B/op          5 allocs/op
BenchmarkRevel_ParseStatic    500000          5028 ns/op        1288 B/op         25 allocs/op
BenchmarkRivet_ParseStatic   5000000           371 ns/op         112 B/op          2 allocs/op
BenchmarkTango_ParseStatic   1000000          1018 ns/op         320 B/op          8 allocs/op
BenchmarkTigerTonic_ParseStatic  5000000           247 ns/op          48 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic      500000          3762 ns/op        1832 B/op         21 allocs/op
BenchmarkVulcan_ParseStatic  2000000           853 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_ParseStatic    1000000          1002 ns/op         256 B/op          6 allocs/op

BenchmarkAce_ParseParam 10000000           229 ns/op          64 B/op          1 allocs/op
BenchmarkBear_ParseParam     1000000          1234 ns/op         512 B/op          6 allocs/op
BenchmarkBeego_ParseParam    1000000          1915 ns/op         736 B/op         10 allocs/op
BenchmarkBone_ParseParam     1000000          1087 ns/op         464 B/op          4 allocs/op
BenchmarkDenco_ParseParam    5000000           249 ns/op          64 B/op          1 allocs/op
BenchmarkGin_ParseParam  5000000           247 ns/op          64 B/op          1 allocs/op
BenchmarkGocraftWeb_ParseParam   1000000          1506 ns/op         672 B/op          9 allocs/op
BenchmarkGoji_ParseParam     2000000           933 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_ParseParam     50000         32233 ns/op        4208 B/op         98 allocs/op
BenchmarkGoRestful_ParseParam     300000          6147 ns/op         568 B/op         16 allocs/op
BenchmarkGorillaMux_ParseParam    500000          3244 ns/op         784 B/op          9 allocs/op
BenchmarkHttpRouter_ParseParam  10000000           167 ns/op          64 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParseParam  3000000           608 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_ParseParam    5000000           379 ns/op          56 B/op          3 allocs/op
BenchmarkMacaron_ParseParam  1000000          2779 ns/op        1120 B/op         11 allocs/op
BenchmarkMartini_ParseParam   200000          7132 ns/op        1152 B/op         12 allocs/op
BenchmarkPat_ParseParam  1000000          2659 ns/op        1136 B/op         20 allocs/op
BenchmarkRevel_ParseParam     300000          5883 ns/op        1704 B/op         28 allocs/op
BenchmarkRivet_ParseParam    2000000           964 ns/op         464 B/op          5 allocs/op
BenchmarkTango_ParseParam    1000000          2760 ns/op        1176 B/op         19 allocs/op
BenchmarkTigerTonic_ParseParam   1000000          2966 ns/op        1024 B/op         19 allocs/op
BenchmarkTraffic_ParseParam   300000          5139 ns/op        2280 B/op         25 allocs/op
BenchmarkVulcan_ParseParam   2000000           870 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_ParseParam     1000000          1626 ns/op         576 B/op          9 allocs/op

BenchmarkAce_Parse2Params    5000000           252 ns/op          64 B/op          1 allocs/op
BenchmarkBear_Parse2Params   1000000          1474 ns/op         544 B/op          6 allocs/op
BenchmarkBeego_Parse2Params  1000000          2201 ns/op         784 B/op         11 allocs/op
BenchmarkBone_Parse2Params   2000000           980 ns/op         416 B/op          3 allocs/op
BenchmarkDenco_Parse2Params  5000000           297 ns/op          64 B/op          1 allocs/op
BenchmarkGin_Parse2Params    5000000           276 ns/op          64 B/op          1 allocs/op
BenchmarkGocraftWeb_Parse2Params     1000000          1740 ns/op         720 B/op         10 allocs/op
BenchmarkGoji_Parse2Params   2000000           896 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_Parse2Params       50000         32619 ns/op        4272 B/op         99 allocs/op
BenchmarkGoRestful_Parse2Params   200000          6579 ns/op         568 B/op         16 allocs/op
BenchmarkGorillaMux_Parse2Params      500000          3696 ns/op         816 B/op          9 allocs/op
BenchmarkHttpRouter_Parse2Params    10000000           193 ns/op          64 B/op          1 allocs/op
BenchmarkHttpTreeMux_Parse2Params    2000000           674 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_Parse2Params  2000000           629 ns/op         128 B/op          5 allocs/op
BenchmarkMacaron_Parse2Params    1000000          3009 ns/op        1168 B/op         12 allocs/op
BenchmarkMartini_Parse2Params     200000          7441 ns/op        1184 B/op         12 allocs/op
BenchmarkPat_Parse2Params    1000000          2618 ns/op         864 B/op         21 allocs/op
BenchmarkRevel_Parse2Params   300000          6126 ns/op        1768 B/op         30 allocs/op
BenchmarkRivet_Parse2Params  1000000          1106 ns/op         480 B/op          6 allocs/op
BenchmarkTango_Parse2Params  1000000          2218 ns/op         776 B/op         17 allocs/op
BenchmarkTigerTonic_Parse2Params      500000          4870 ns/op        1440 B/op         28 allocs/op
BenchmarkTraffic_Parse2Params     300000          5415 ns/op        2088 B/op         25 allocs/op
BenchmarkVulcan_Parse2Params     1000000          1051 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_Parse2Params   1000000          1481 ns/op         528 B/op          8 allocs/op

BenchmarkAce_ParseAll     300000          5341 ns/op         640 B/op         16 allocs/op
BenchmarkBear_ParseAll     30000         49576 ns/op       15600 B/op        233 allocs/op
BenchmarkBeego_ParseAll    30000         47858 ns/op       15600 B/op        233 allocs/op
BenchmarkBone_ParseAll     50000         27212 ns/op        9296 B/op         99 allocs/op
BenchmarkDenco_ParseAll   300000          5624 ns/op         928 B/op         16 allocs/op
BenchmarkGin_ParseAll     300000          6433 ns/op         640 B/op         16 allocs/op
BenchmarkGocraftWeb_ParseAll       50000         37088 ns/op       13936 B/op        207 allocs/op
BenchmarkGoji_ParseAll    100000         18588 ns/op        5376 B/op         32 allocs/op
BenchmarkGoji_ParseAll    100000             18588 ns/op            5376 B/op         32 allocs/op
BenchmarkGoJsonRest_ParseAll        2000        871680 ns/op      105643 B/op       2527 allocs/op
BenchmarkGoRestful_ParseAll     5000        343915 ns/op       97312 B/op        904 allocs/op
BenchmarkGorillaMux_ParseAll       10000        124983 ns/op       17280 B/op        224 allocs/op
BenchmarkHttpRouter_ParseAll      500000          3670 ns/op         640 B/op         16 allocs/op
BenchmarkHttpTreeMux_ParseAll     200000         11484 ns/op        5376 B/op         32 allocs/op
BenchmarkKocha_ParseAll   200000          9038 ns/op        1112 B/op         54 allocs/op
BenchmarkMacaron_ParseAll      20000         73140 ns/op       25584 B/op        259 allocs/op
BenchmarkMartini_ParseAll      10000        200863 ns/op       26848 B/op        302 allocs/op
BenchmarkPat_ParseAll      30000         51573 ns/op       17568 B/op        382 allocs/op
BenchmarkRevel_ParseAll    10000        152626 ns/op       40464 B/op        704 allocs/op
BenchmarkRivet_ParseAll   100000         22562 ns/op        8592 B/op        103 allocs/op
BenchmarkTango_ParseAll    30000         48243 ns/op       16560 B/op        330 allocs/op
BenchmarkTigerTonic_ParseAll       20000         68300 ns/op       20032 B/op        417 allocs/op
BenchmarkTraffic_ParseAll      10000        139429 ns/op       58496 B/op        687 allocs/op
BenchmarkVulcan_ParseAll       50000         27106 ns/op        2548 B/op         78 allocs/op
BenchmarkZeus_ParseAll     30000         56752 ns/op       18336 B/op        292 allocs/op
```


### [GitHub](http://developer.github.com/v3/)

The GitHub API is rather large, consisting of 203 routes. The tasks are basically the same as in the benchmarks before.

```
BenchmarkAce_GithubStatic   20000000           116 ns/op           0 B/op          0 allocs/op
BenchmarkBear_GithubStatic   2000000           712 ns/op         160 B/op          4 allocs/op
BenchmarkBeego_GithubStatic  1000000          1233 ns/op         368 B/op          7 allocs/op
BenchmarkBone_GithubStatic    200000          9707 ns/op        2880 B/op         60 allocs/op
BenchmarkDenco_GithubStatic 20000000            64.5 ns/op         0 B/op          0 allocs/op
BenchmarkGin_GithubStatic   10000000           134 ns/op           0 B/op          0 allocs/op
BenchmarkGocraftWeb_GithubStatic     2000000           893 ns/op         304 B/op          6 allocs/op
BenchmarkGoji_GithubStatic   5000000           297 ns/op           0 B/op          0 allocs/op
BenchmarkGoRestful_GithubStatic    50000         37217 ns/op        2968 B/op         34 allocs/op
BenchmarkGoJsonRest_GithubStatic       50000         31305 ns/op        3888 B/op         96 allocs/op
BenchmarkGorillaMux_GithubStatic      100000         19121 ns/op         464 B/op          8 allocs/op
BenchmarkHttpRouter_GithubStatic    30000000            57.8 ns/op         0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic   20000000            64.2 ns/op         0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic 20000000            93.1 ns/op         0 B/op          0 allocs/op
BenchmarkMacaron_GithubStatic    1000000          2105 ns/op         752 B/op          8 allocs/op
BenchmarkMartini_GithubStatic     100000         18156 ns/op         832 B/op         11 allocs/op
BenchmarkPat_GithubStatic     200000         10629 ns/op        3648 B/op         76 allocs/op
BenchmarkRevel_GithubStatic   300000          4987 ns/op        1288 B/op         25 allocs/op
BenchmarkRivet_GithubStatic  5000000           393 ns/op         112 B/op          2 allocs/op
BenchmarkTango_GithubStatic  1000000          1042 ns/op         320 B/op          8 allocs/op
BenchmarkTigerTonic_GithubStatic     5000000           281 ns/op          48 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic      50000         37404 ns/op       18920 B/op        149 allocs/op
BenchmarkVulcan_GithubStatic     1000000          1131 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_GithubStatic     50000         26059 ns/op        7344 B/op        136 allocs/op

BenchmarkAce_GithubParam     5000000           356 ns/op          96 B/op          1 allocs/op
BenchmarkBear_GithubParam    1000000          1724 ns/op         560 B/op          6 allocs/op
BenchmarkBeego_GithubParam   1000000          2457 ns/op         784 B/op         11 allocs/op
BenchmarkBone_GithubParam     300000          5368 ns/op        1456 B/op         16 allocs/op
BenchmarkDenco_GithubParam   3000000           447 ns/op         128 B/op          1 allocs/op
BenchmarkGin_GithubParam     5000000           374 ns/op          96 B/op          1 allocs/op
BenchmarkGocraftWeb_GithubParam  1000000          1843 ns/op         720 B/op         10 allocs/op
BenchmarkGoji_GithubParam    1000000          1229 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_GithubParam    50000         33013 ns/op        4368 B/op         99 allocs/op
BenchmarkGoRestful_GithubParam     50000         33528 ns/op         568 B/op         16 allocs/op
BenchmarkGorillaMux_GithubParam   200000         11515 ns/op         816 B/op          9 allocs/op
BenchmarkHttpRouter_GithubParam  5000000           291 ns/op          96 B/op          1 allocs/op
BenchmarkHttpTreeMux_GithubParam     2000000           760 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_GithubParam   2000000           734 ns/op         128 B/op          5 allocs/op
BenchmarkMacaron_GithubParam     1000000          3211 ns/op        1168 B/op         12 allocs/op
BenchmarkMartini_GithubParam      100000         23790 ns/op        1184 B/op         12 allocs/op
BenchmarkPat_GithubParam      200000          7109 ns/op        2480 B/op         56 allocs/op
BenchmarkRevel_GithubParam    200000          6261 ns/op        1784 B/op         30 allocs/op
BenchmarkRivet_GithubParam   1000000          1273 ns/op         480 B/op          6 allocs/op
BenchmarkTango_GithubParam    300000          6220 ns/op        2240 B/op         48 allocs/op
BenchmarkTigerTonic_GithubParam   300000          4912 ns/op        1440 B/op         28 allocs/op
BenchmarkTraffic_GithubParam      100000         16620 ns/op        6024 B/op         55 allocs/op
BenchmarkVulcan_GithubParam  1000000          1845 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_GithubParam     100000         21205 ns/op        7496 B/op        100 allocs/op

BenchmarkAce_GithubAll     20000         63741 ns/op       13792 B/op        167 allocs/op
BenchmarkBear_GithubAll     5000        323455 ns/op       97840 B/op       1146 allocs/op
BenchmarkBeego_GithubAll        3000        468911 ns/op      146272 B/op       2092 allocs/op
BenchmarkBone_GithubAll     1000       2172496 ns/op      648016 B/op       8119 allocs/op
BenchmarkDenco_GithubAll       20000         78482 ns/op       20224 B/op        167 allocs/op
BenchmarkGin_GithubAll     20000         68553 ns/op       13792 B/op        167 allocs/op
BenchmarkGocraftWeb_GithubAll       5000        362107 ns/op      133280 B/op       1889 allocs/op
BenchmarkGoji_GithubAll     3000        557240 ns/op       56113 B/op        334 allocs/op
BenchmarkGoJsonRest_GithubAll        200       7028115 ns/op      860715 B/op      19983 allocs/op
BenchmarkGoRestful_GithubAll         200       8203040 ns/op      598107 B/op       7152 allocs/op
BenchmarkGorillaMux_GithubAll        200       6705367 ns/op      153137 B/op       1791 allocs/op
BenchmarkHttpRouter_GithubAll      30000         51046 ns/op       13792 B/op        167 allocs/op
BenchmarkHttpTreeMux_GithubAll     10000        130105 ns/op       56112 B/op        334 allocs/op
BenchmarkKocha_GithubAll       10000        138286 ns/op       23304 B/op        843 allocs/op
BenchmarkMacaron_GithubAll      2000        692800 ns/op      224960 B/op       2315 allocs/op
BenchmarkMartini_GithubAll       100      11379839 ns/op      237952 B/op       2686 allocs/op
BenchmarkPat_GithubAll       300       3915465 ns/op     1504101 B/op      32222 allocs/op
BenchmarkRevel_GithubAll        2000       1222445 ns/op      345553 B/op       5918 allocs/op
BenchmarkRivet_GithubAll       10000        235504 ns/op       84272 B/op       1079 allocs/op
BenchmarkTango_GithubAll         500       3395497 ns/op     1338664 B/op      27736 allocs/op
BenchmarkTigerTonic_GithubAll       2000        938041 ns/op      241088 B/op       6052 allocs/op
BenchmarkTraffic_GithubAll       200       7263268 ns/op     2664762 B/op      22390 allocs/op
BenchmarkVulcan_GithubAll       5000        285455 ns/op       19894 B/op        609 allocs/op
BenchmarkZeus_GithubAll      200       8170547 ns/op     2755680 B/op      39483 allocs/op
```

### [Google+](https://developers.google.com/+/api/latest/)

Last but not least the Google+ API, consisting of 13 routes. In reality this is just a subset of a much larger API.

```
BenchmarkAce_GPlusStatic    20000000            94.0 ns/op         0 B/op          0 allocs/op
BenchmarkBear_GPlusStatic    3000000           562 ns/op         136 B/op          4 allocs/op
BenchmarkBeego_GPlusStatic   1000000          1123 ns/op         352 B/op          7 allocs/op
BenchmarkBone_GPlusStatic   10000000           154 ns/op          32 B/op          1 allocs/op
BenchmarkDenco_GPlusStatic  30000000            38.0 ns/op         0 B/op          0 allocs/op
BenchmarkGin_GPlusStatic    20000000           107 ns/op           0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlusStatic  2000000           824 ns/op         288 B/op          6 allocs/op
BenchmarkGoji_GPlusStatic   10000000           220 ns/op           0 B/op          0 allocs/op
BenchmarkGoJsonRest_GPlusStatic    50000         30572 ns/op        3648 B/op         95 allocs/op
BenchmarkGoRestful_GPlusStatic    200000          6973 ns/op        1736 B/op         28 allocs/op
BenchmarkGorillaMux_GPlusStatic  1000000          1911 ns/op         464 B/op          8 allocs/op
BenchmarkHttpRouter_GPlusStatic 50000000            34.2 ns/op         0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic    50000000            36.8 ns/op         0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic  20000000            61.1 ns/op         0 B/op          0 allocs/op
BenchmarkMacaron_GPlusStatic     1000000          2040 ns/op         736 B/op          8 allocs/op
BenchmarkMartini_GPlusStatic      500000          4339 ns/op         832 B/op         11 allocs/op
BenchmarkPat_GPlusStatic     5000000           323 ns/op          96 B/op          2 allocs/op
BenchmarkRevel_GPlusStatic    500000          4959 ns/op        1272 B/op         25 allocs/op
BenchmarkRivet_GPlusStatic   5000000           351 ns/op         112 B/op          2 allocs/op
BenchmarkTango_GPlusStatic   1000000          1004 ns/op         320 B/op          8 allocs/op
BenchmarkTigerTonic_GPlusStatic 10000000           173 ns/op          32 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic     1000000          2546 ns/op        1208 B/op         16 allocs/op
BenchmarkVulcan_GPlusStatic  2000000           710 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_GPlusStatic    5000000           266 ns/op          48 B/op          2 allocs/op

BenchmarkAce_GPlusParam  5000000           251 ns/op          64 B/op          1 allocs/op
BenchmarkBear_GPlusParam     1000000          1301 ns/op         512 B/op          6 allocs/op
BenchmarkBeego_GPlusParam    1000000          2068 ns/op         720 B/op         10 allocs/op
BenchmarkBone_GPlusParam     2000000           874 ns/op         384 B/op          3 allocs/op
BenchmarkDenco_GPlusParam    5000000           260 ns/op          64 B/op          1 allocs/op
BenchmarkGin_GPlusParam  5000000           272 ns/op          64 B/op          1 allocs/op
BenchmarkGocraftWeb_GPlusParam   1000000          1507 ns/op         656 B/op          9 allocs/op
BenchmarkGoji_GPlusParam     2000000           835 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlusParam     50000         33086 ns/op        4240 B/op         98 allocs/op
BenchmarkGoRestful_GPlusParam     300000          6108 ns/op         632 B/op         18 allocs/op
BenchmarkGorillaMux_GPlusParam    500000          4104 ns/op         784 B/op          9 allocs/op
BenchmarkHttpRouter_GPlusParam  10000000           192 ns/op          64 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlusParam  2000000           611 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_GPlusParam    5000000           405 ns/op          56 B/op          3 allocs/op
BenchmarkMacaron_GPlusParam  1000000          2800 ns/op        1104 B/op         11 allocs/op
BenchmarkMartini_GPlusParam   200000          8656 ns/op        1152 B/op         12 allocs/op
BenchmarkPat_GPlusParam  1000000          1870 ns/op         704 B/op         14 allocs/op
BenchmarkRevel_GPlusParam     300000          5753 ns/op        1704 B/op         28 allocs/op
BenchmarkRivet_GPlusParam    2000000          1008 ns/op         464 B/op          5 allocs/op
BenchmarkTango_GPlusParam    1000000          1880 ns/op         712 B/op         13 allocs/op
BenchmarkTigerTonic_GPlusParam   1000000          3177 ns/op        1064 B/op         19 allocs/op
BenchmarkTraffic_GPlusParam   300000          5503 ns/op        2000 B/op         23 allocs/op
BenchmarkVulcan_GPlusParam   1000000          1029 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_GPlusParam     1000000          1178 ns/op         440 B/op          6 allocs/op

BenchmarkAce_GPlus2Params    5000000           297 ns/op          64 B/op          1 allocs/op
BenchmarkBear_GPlus2Params   1000000          1701 ns/op         576 B/op          6 allocs/op
BenchmarkBeego_GPlus2Params  1000000          2504 ns/op         784 B/op         11 allocs/op
BenchmarkBone_GPlus2Params   1000000          2552 ns/op         736 B/op          7 allocs/op
BenchmarkDenco_GPlus2Params  5000000           345 ns/op          64 B/op          1 allocs/op
BenchmarkGin_GPlus2Params    5000000           302 ns/op          64 B/op          1 allocs/op
BenchmarkGocraftWeb_GPlus2Params     1000000          1831 ns/op         720 B/op         10 allocs/op
BenchmarkGoji_GPlus2Params   1000000          1190 ns/op         336 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlus2Params       50000         32878 ns/op        4384 B/op         99 allocs/op
BenchmarkGoRestful_GPlus2Params   200000          7730 ns/op         632 B/op         18 allocs/op
BenchmarkGorillaMux_GPlus2Params      200000          9201 ns/op         816 B/op          9 allocs/op
BenchmarkHttpRouter_GPlus2Params    10000000           232 ns/op          64 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params    2000000           739 ns/op         336 B/op          2 allocs/op
BenchmarkKocha_GPlus2Params  2000000           693 ns/op         128 B/op          5 allocs/op
BenchmarkMacaron_GPlus2Params    1000000          3209 ns/op        1168 B/op         12 allocs/op
BenchmarkMartini_GPlus2Params      50000         28563 ns/op        1280 B/op         16 allocs/op
BenchmarkPat_GPlus2Params     300000          5871 ns/op        2304 B/op         41 allocs/op
BenchmarkRevel_GPlus2Params   300000          6282 ns/op        1800 B/op         30 allocs/op
BenchmarkRivet_GPlus2Params  1000000          1159 ns/op         480 B/op          6 allocs/op
BenchmarkTango_GPlus2Params   300000          5111 ns/op        2136 B/op         36 allocs/op
BenchmarkTigerTonic_GPlus2Params      300000          5196 ns/op        1528 B/op         28 allocs/op
BenchmarkTraffic_GPlus2Params     200000         12093 ns/op        3312 B/op         34 allocs/op
BenchmarkVulcan_GPlus2Params     1000000          1483 ns/op          98 B/op          3 allocs/op
BenchmarkZeus_GPlus2Params    200000          7680 ns/op        2864 B/op         37 allocs/op

BenchmarkAce_GPlusAll     500000          3312 ns/op         640 B/op         11 allocs/op
BenchmarkBear_GPlusAll    100000         16969 ns/op        6168 B/op         74 allocs/op
BenchmarkBeego_GPlusAll    50000         26293 ns/op        8976 B/op        129 allocs/op
BenchmarkBone_GPlusAll    100000         21740 ns/op        6992 B/op         76 allocs/op
BenchmarkDenco_GPlusAll   500000          3718 ns/op         672 B/op         11 allocs/op
BenchmarkGin_GPlusAll     500000          3484 ns/op         640 B/op         11 allocs/op
BenchmarkGocraftWeb_GPlusAll      100000         19972 ns/op        8144 B/op        116 allocs/op
BenchmarkGoji_GPlusAll    200000         11510 ns/op        3696 B/op         22 allocs/op
BenchmarkGoJsonRest_GPlusAll        5000        421075 ns/op       54499 B/op       1274 allocs/op
BenchmarkGoRestful_GPlusAll    10000        120138 ns/op       26264 B/op        404 allocs/op
BenchmarkGorillaMux_GPlusAll       20000         65963 ns/op        9712 B/op        115 allocs/op
BenchmarkHttpRouter_GPlusAll     1000000          2441 ns/op         640 B/op         11 allocs/op
BenchmarkHttpTreeMux_GPlusAll     200000          7449 ns/op        3696 B/op         22 allocs/op
BenchmarkKocha_GPlusAll   300000          6425 ns/op         976 B/op         43 allocs/op
BenchmarkMacaron_GPlusAll      50000         36082 ns/op       13968 B/op        142 allocs/op
BenchmarkMartini_GPlusAll      10000        155220 ns/op       15072 B/op        178 allocs/op
BenchmarkPat_GPlusAll      30000         46035 ns/op       16880 B/op        343 allocs/op
BenchmarkRevel_GPlusAll    20000         77560 ns/op       21656 B/op        368 allocs/op
BenchmarkRivet_GPlusAll   200000         13012 ns/op        5408 B/op         64 allocs/op
BenchmarkTango_GPlusAll    50000         40046 ns/op       15624 B/op        280 allocs/op
BenchmarkTigerTonic_GPlusAll       30000         50101 ns/op       14800 B/op        320 allocs/op
BenchmarkTraffic_GPlusAll      10000        100134 ns/op       37760 B/op        421 allocs/op
BenchmarkVulcan_GPlusAll      100000         14525 ns/op        1274 B/op         39 allocs/op
BenchmarkZeus_GPlusAll     30000         53868 ns/op       18440 B/op        268 allocs/op
```


## Conclusions
First of all, there is no reason to use net/http's default [ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is very limited and does not have especially good performance. There are enough alternatives coming in every flavor, choose the one you like best.

Secondly, the broad range of functions of some of the frameworks comes at a high price in terms of performance. For example Martini has great flexibility, but very bad performance. Martini has the worst performance of all tested routers in a lot of the benchmarks. Beego seems to have some scalability problems and easily defeats Martini with even worse performance, when the number of parameters or routes is high. I really hope, that the routing of these packages can be optimized. I think the Go-ecosystem needs great feature-rich frameworks like these.

Last but not least, we have to determine the performance champion.

Denco and its predecessor Kocha-urlrouter seem to have great performance, but are not convenient to use as a router for the net/http package. A lot of extra work is necessary to use it as a http.Handler. [The README of Denco claims](https://github.com/naoina/denco/blob/b03dbb499269a597afd0db715d408ebba1329d04/README.md), that the package is not intended as a replacement for [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).

[Goji](https://github.com/zenazn/goji/) looks very decent. It has great performance while also having a great range of features, more than any other router / framework in the top group.

Currently no router can beat the performance of the [HttpRouter](https://github.com/julienschmidt/httprouter) package, which currently dominates nearly all benchmarks.

In the end, performance can not be the (only) criterion for choosing a router. Play around a bit with some of the routers, and choose the one you like best.
