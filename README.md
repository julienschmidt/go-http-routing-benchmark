Go HTTP Router Benchmark
========================

This benchmark suite aims to compare the performance of HTTP request routers for [Go](https://golang.org) by implementing the routing structure of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.

Of course the tested routers can be used for any kind of HTTP request → handler function routing, not only (REST) APIs.


#### Tested routers & frameworks:

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
 * [Possum](https://github.com/mikespook/possum)
 * [R2router](https://github.com/vanng822/r2router)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)


## Motivation

Go is a great language for web applications. Since the [default *request multiplexer*](http://golang.org/pkg/net/http/#ServeMux) of Go's net/http package is very simple and limited, an accordingly high number of HTTP request routers exist.

Unfortunately, most of the (early) routers use pretty bad routing algorithms. Moreover, many of them are very wasteful with memory allocations, which can become a problem in a language with Garbage Collection like Go, since every (heap) allocation results in more work for the Garbage Collector.

Lately more and more bloated frameworks pop up, outdoing one another in the number of features. This benchmark tries to measure their overhead.

Be aware that we are comparing apples and oranges here. We compare feature-rich frameworks to packages with simple routing functionality only. But since we are only interested in decent request routing, I think this is not entirely unfair. The frameworks are configured to do as little additional work as possible.

If you care about performance, this benchmark can maybe help you find the right router, which scales with your application.

Personally, I prefer slim and optimized software, which is why I implemented [HttpRouter](https://github.com/julienschmidt/httprouter), which is also tested here. In fact, this benchmark suite started as part of the packages tests, but was then extended to a generic benchmark suite.
So keep in mind, that I am not completely unbiased :relieved:


## Results

Benchmark System:
 * Intel Core i5-2500K (4x 3,30GHz + Turbo Boost), CPU-governor: performance
 * 2x 4 GiB DDR3-1333 RAM, dual-channel
 * go version go1.3rc1 linux/amd64
 * Ubuntu 14.04 amd64 (Linux Kernel 3.13.0-29), fresh installation


### Memory Consumption

Besides the micro-benchmarks, there are 3 sets of benchmarks where we play around with clones of some real-world APIs, and one benchmark with static routes only, to allow a comparison with [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).
The following table shows the memory required only for loading the routing structure for the respective API.
The best 3 values for each test are bold. I'm pretty sure you can detect a pattern :wink:

| Router       | Static    | GitHub     | Google+   | Parse     |
|:-------------|----------:|-----------:|----------:|----------:|
| HttpServeMux |__18064 B__|         -  |        -  |        -  |
| Beego        |  79472 B  |  497248 B  |  26480 B  |  38768 B  |
| Denco        |  44752 B  |  107632 B  |  54896 B  |  36368 B  |
| Gocraft Web  |  57976 B  |   95736 B  |   8024 B  |  13120 B  |
| Goji         |  32400 B  | __58424 B__| __3392 B__| __6704 B__|
| Go-Json-Rest | 152608 B  |  148352 B  |  11696 B  |  13712 B  |
| Gorilla Mux  | 685152 B  | 1557216 B  |  80240 B  | 125480 B  |
| HttpRouter   |__26232 B__| __44344 B__| __3144 B__| __5792 B__|
| HttpTreeMux  |  75624 B  |   81408 B  |   7712 B  |   7616 B  |
| Kocha        | 130336 B  |  811744 B  | 139968 B  | 191632 B  |
| Martini      | 312592 B  |  579472 B  |  27520 B  |  50608 B  |
| Pat          |__21272 B__| __18968 B__| __1448 B__| __2360 B__|
| TigerTonic   |  85264 B  |   99392 B  |  10576 B  |  11008 B  |
| Traffic      | 649568 B  | 1124704 B  |  57984 B  |  98168 B  |

The first place goes to [Pat](https://github.com/bmizerany/pat), followed by [HttpRouter](https://github.com/julienschmidt/httprouter) and [Goji](https://github.com/zenazn/goji/). Now, before everyone starts reading the documentation of Pat, `[SPOILER]` this low memory consumption comes at the price of relatively bad routing performance. The routing structure of Pat is simple - probably too simple. `[/SPOILER]`.

Moreover main memory is cheap and usually not a scarce resource. As long as the router doesn't require Megabytes of memory, it should be no deal breaker. But it gives us a first hint how efficient or wasteful a router works.


### Static Routes

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

The only intention of this benchmark is to allow a comparison with the default router of Go's net/http package, [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is limited to static routes and does not support parameters in the route pattern.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.8% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.

```
BenchmarkHttpServeMux_StaticAll         5000     706222 ns/op          96 B/op        6 allocs/op

BenchmarkBeego_StaticAll                2000    1408954 ns/op      482433 B/op    14088 allocs/op
BenchmarkDenco_StaticAll              200000      12679 ns/op           0 B/op        0 allocs/op
BenchmarkGocraftWeb_StaticAll          10000     154142 ns/op       51468 B/op      947 allocs/op
BenchmarkGoji_StaticAll                20000      80518 ns/op           0 B/op        0 allocs/op
BenchmarkGoJsonRest_StaticAll           2000     978164 ns/op      180973 B/op     3945 allocs/op
BenchmarkGorillaMux_StaticAll           1000    1763690 ns/op       71804 B/op      956 allocs/op
BenchmarkHttpRouter_StaticAll         100000      15010 ns/op           0 B/op        0 allocs/op
BenchmarkHttpTreeMux_StaticAll        100000      15123 ns/op           0 B/op        0 allocs/op
BenchmarkKocha_StaticAll              100000      23093 ns/op           0 B/op        0 allocs/op
BenchmarkMartini_StaticAll               500    3444278 ns/op      156015 B/op     2351 allocs/op
BenchmarkPat_StaticAll                  1000    1640745 ns/op      549187 B/op    11186 allocs/op
BenchmarkTigerTonic_StaticAll          50000      58264 ns/op        7714 B/op      157 allocs/op
BenchmarkTraffic_StaticAll               500    7230129 ns/op     3763731 B/op    27453 allocs/op
```

### Micro Benchmarks

The following benchmarks measure the cost of some very basic operations.

In the first benchmark, only a single route, containing a parameter, is loaded into the routers. Then a request for a URL matching this pattern is made and the router has to call the respective registered handler function. End.
```
BenchmarkBeego_Param                  500000       5495 ns/op        1165 B/op       14 allocs/op
BenchmarkDenco_Param                 5000000        312 ns/op          50 B/op        2 allocs/op
BenchmarkGocraftWeb_Param            1000000       1440 ns/op         684 B/op        9 allocs/op
BenchmarkGoji_Param                  5000000        748 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_Param             500000       6980 ns/op        1787 B/op       29 allocs/op
BenchmarkGorillaMux_Param            1000000       2665 ns/op         780 B/op        7 allocs/op
BenchmarkHttpRouter_Param           20000000        139 ns/op          33 B/op        1 allocs/op
BenchmarkHttpTreeMux_Param           5000000        558 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_Param                 5000000        377 ns/op          58 B/op        2 allocs/op
BenchmarkMartini_Param                500000       6265 ns/op        1251 B/op       12 allocs/op
BenchmarkPat_Param                   1000000       1620 ns/op         670 B/op       11 allocs/op
BenchmarkTigerTonic_Param            1000000       2766 ns/op        1015 B/op       18 allocs/op
BenchmarkTraffic_Param                500000       4440 ns/op        2013 B/op       22 allocs/op
```

Same as before, but now with multiple parameters, all in the same single route. The intention is to see how the routers scale with the number of parameters. The values of the parameters must be passed to the handler function somehow, which requires allocations. Let's see how clever the routers solve this task with a route containing 5 and 20 parameters:
```
BenchmarkBeego_Param5                 100000      18473 ns/op        1291 B/op       14 allocs/op
BenchmarkDenco_Param5                2000000        982 ns/op         405 B/op        5 allocs/op
BenchmarkGocraftWeb_Param5           1000000       2218 ns/op         957 B/op       12 allocs/op
BenchmarkGoji_Param5                 1000000       1093 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_Param5            200000      10462 ns/op        3264 B/op       40 allocs/op
BenchmarkGorillaMux_Param5            500000       4680 ns/op         906 B/op        7 allocs/op
BenchmarkHttpRouter_Param5           5000000        319 ns/op         162 B/op        1 allocs/op
BenchmarkHttpTreeMux_Param5          2000000        898 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_Param5                1000000       1326 ns/op         448 B/op        7 allocs/op
BenchmarkMartini_Param5               200000      13027 ns/op        1376 B/op       12 allocs/op
BenchmarkPat_Param5                   500000       3416 ns/op        1435 B/op       18 allocs/op
BenchmarkTigerTonic_Param5            200000       9247 ns/op        2568 B/op       41 allocs/op
BenchmarkTraffic_Param5               500000       7206 ns/op        2312 B/op       26 allocs/op

BenchmarkBeego_Param20                 10000     106746 ns/op        3681 B/op       17 allocs/op
BenchmarkDenco_Param20               1000000       2882 ns/op        1666 B/op        7 allocs/op
BenchmarkGocraftWeb_Param20           500000       7156 ns/op        3857 B/op       16 allocs/op
BenchmarkGoji_Param20                1000000       3197 ns/op        1260 B/op        2 allocs/op
BenchmarkGoJsonRest_Param20           100000      25809 ns/op       10605 B/op       75 allocs/op
BenchmarkGorillaMux_Param20           200000       9885 ns/op        3295 B/op        9 allocs/op
BenchmarkHttpRouter_Param20          2000000        954 ns/op         646 B/op        1 allocs/op
BenchmarkHttpTreeMux_Param20          500000       5016 ns/op        2216 B/op        4 allocs/op
BenchmarkKocha_Param20                500000       4268 ns/op        1836 B/op       17 allocs/op
BenchmarkMartini_Param20               50000      55039 ns/op        3765 B/op       14 allocs/op
BenchmarkPat_Param20                  500000       3412 ns/op        1435 B/op       18 allocs/op
BenchmarkTigerTonic_Param20            50000      36825 ns/op       10710 B/op      131 allocs/op
BenchmarkTraffic_Param20              100000      22605 ns/op        8077 B/op       49 allocs/op
```

Now let's see how expensive it is to access a parameter. The handler function reads the value (by the name of the parameter, e.g. with a map lookup; depends on the router) and writes it to our [web scale storage](https://www.youtube.com/watch?v=b2F-DItXtZs) (`/dev/null`).
```
BenchmarkBeego_ParamWrite             500000       6604 ns/op        1602 B/op       18 allocs/op
BenchmarkDenco_ParamWrite            5000000        377 ns/op          50 B/op        2 allocs/op
BenchmarkGocraftWeb_ParamWrite       1000000       1590 ns/op         693 B/op        9 allocs/op
BenchmarkGoji_ParamWrite             2000000        818 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_ParamWrite        200000       8388 ns/op        2265 B/op       33 allocs/op
BenchmarkGorillaMux_ParamWrite       1000000       2913 ns/op         780 B/op        7 allocs/op
BenchmarkHttpRouter_ParamWrite      10000000        193 ns/op          33 B/op        1 allocs/op
BenchmarkHttpTreeMux_ParamWrite      5000000        649 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_ParamWrite            5000000        435 ns/op          58 B/op        2 allocs/op
BenchmarkMartini_ParamWrite           500000       7538 ns/op        1359 B/op       15 allocs/op
BenchmarkPat_ParamWrite              1000000       2940 ns/op        1109 B/op       15 allocs/op
BenchmarkTigerTonic_ParamWrite        500000       4639 ns/op        1471 B/op       23 allocs/op
BenchmarkTraffic_ParamWrite           500000       5855 ns/op        2435 B/op       25 allocs/op
```

### [Parse.com](https://parse.com/docs/rest#summary)

Enough of the micro benchmark stuff. Let's play a bit with real APIs. In the first set of benchmarks, we use a clone of the structure of [Parse](https://parse.com)'s decent medium-sized REST API, consisting of 26 routes.

The tasks are 1.) routing a static URL (no parameters), 2.) routing a URL containing 1 parameter, 3.) same with 2 parameters, 4.) route all of the routes once (like the StaticAll benchmark, but the routes now contain parameters).

Worth noting is, that the requested route might be a good case for some routing algorithms, while it is a bad case for another algorithm. The values might vary slightly depending on the selected route.

```
BenchmarkBeego_ParseStatic            500000       3461 ns/op        1247 B/op       15 allocs/op
BenchmarkDenco_ParseStatic          50000000         42.6 ns/op         0 B/op        0 allocs/op
BenchmarkGocraftWeb_ParseStatic      2000000        889 ns/op         328 B/op        6 allocs/op
BenchmarkGoji_ParseStatic            5000000        341 ns/op           0 B/op        0 allocs/op
BenchmarkGoJsonRest_ParseStatic       500000       5860 ns/op        1136 B/op       25 allocs/op
BenchmarkGorillaMux_ParseStatic      1000000       2760 ns/op         456 B/op        6 allocs/op
BenchmarkHttpRouter_ParseStatic     50000000         36.7 ns/op         0 B/op        0 allocs/op
BenchmarkHttpTreeMux_ParseStatic    50000000         62.6 ns/op         0 B/op        0 allocs/op
BenchmarkKocha_ParseStatic          50000000         72.2 ns/op         0 B/op        0 allocs/op
BenchmarkMartini_ParseStatic          500000       5528 ns/op         927 B/op       11 allocs/op
BenchmarkPat_ParseStatic             2000000        809 ns/op         246 B/op        5 allocs/op
BenchmarkTigerTonic_ParseStatic     10000000        264 ns/op          49 B/op        1 allocs/op
BenchmarkTraffic_ParseStatic          500000       5008 ns/op        2377 B/op       24 allocs/op

BenchmarkBeego_ParseParam             500000       7983 ns/op        1775 B/op       28 allocs/op
BenchmarkDenco_ParseParam            5000000        347 ns/op          50 B/op        2 allocs/op
BenchmarkGocraftWeb_ParseParam       1000000       1535 ns/op         700 B/op        9 allocs/op
BenchmarkGoji_ParseParam             2000000        983 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_ParseParam        500000       7208 ns/op        1789 B/op       29 allocs/op
BenchmarkGorillaMux_ParseParam       1000000       3186 ns/op         780 B/op        7 allocs/op
BenchmarkHttpRouter_ParseParam      10000000        178 ns/op          65 B/op        1 allocs/op
BenchmarkHttpTreeMux_ParseParam      5000000        617 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_ParseParam            5000000        413 ns/op          58 B/op        2 allocs/op
BenchmarkMartini_ParseParam           500000       7524 ns/op        1251 B/op       12 allocs/op
BenchmarkPat_ParseParam              1000000       2707 ns/op        1160 B/op       18 allocs/op
BenchmarkTigerTonic_ParseParam       1000000       3010 ns/op        1048 B/op       19 allocs/op
BenchmarkTraffic_ParseParam           500000       5228 ns/op        2314 B/op       24 allocs/op

BenchmarkBeego_Parse2Params           200000       9217 ns/op        1935 B/op       28 allocs/op
BenchmarkDenco_Parse2Params          5000000        542 ns/op         115 B/op        3 allocs/op
BenchmarkGocraftWeb_Parse2Params     1000000       1756 ns/op         750 B/op       10 allocs/op
BenchmarkGoji_Parse2Params           2000000        954 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_Parse2Params      500000       8131 ns/op        2145 B/op       32 allocs/op
BenchmarkGorillaMux_Parse2Params      500000       3623 ns/op         812 B/op        7 allocs/op
BenchmarkHttpRouter_Parse2Params    10000000        202 ns/op          65 B/op        1 allocs/op
BenchmarkHttpTreeMux_Parse2Params    5000000        708 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_Parse2Params          5000000        666 ns/op         132 B/op        4 allocs/op
BenchmarkMartini_Parse2Params         200000       7723 ns/op        1283 B/op       12 allocs/op
BenchmarkPat_Parse2Params            1000000       2687 ns/op         887 B/op       19 allocs/op
BenchmarkTigerTonic_Parse2Params      500000       4720 ns/op        1473 B/op       28 allocs/op
BenchmarkTraffic_Parse2Params         500000       5467 ns/op        2120 B/op       24 allocs/op

BenchmarkBeego_ParseAll                10000     197920 ns/op       38877 B/op      616 allocs/op
BenchmarkDenco_ParseAll               500000       7692 ns/op        1000 B/op       35 allocs/op
BenchmarkGocraftWeb_ParseAll           50000      36226 ns/op       14639 B/op      208 allocs/op
BenchmarkGoji_ParseAll                100000      19721 ns/op        5448 B/op       32 allocs/op
BenchmarkGoJsonRest_ParseAll           10000     180128 ns/op       41202 B/op      727 allocs/op
BenchmarkGorillaMux_ParseAll           10000     120929 ns/op       17138 B/op      173 allocs/op
BenchmarkHttpRouter_ParseAll          500000       3592 ns/op         660 B/op       16 allocs/op
BenchmarkHttpTreeMux_ParseAll         200000      11650 ns/op        5452 B/op       32 allocs/op
BenchmarkKocha_ParseAll               200000       9371 ns/op        1163 B/op       44 allocs/op
BenchmarkMartini_ParseAll              10000     200307 ns/op       29375 B/op      305 allocs/op
BenchmarkPat_ParseAll                  50000      53113 ns/op       18017 B/op      363 allocs/op
BenchmarkTigerTonic_ParseAll           50000      67208 ns/op       20547 B/op      419 allocs/op
BenchmarkTraffic_ParseAll              10000     164938 ns/op       70161 B/op      743 allocs/op
```


### [GitHub](http://developer.github.com/v3/)

The GitHub API is rather large, consisting of 203 routes. The tasks are basically the same as in the benchmarks before.

```
BenchmarkBeego_GithubStatic           500000       3880 ns/op        1148 B/op       31 allocs/op
BenchmarkDenco_GithubStatic         50000000         60.5 ns/op         0 B/op        0 allocs/op
BenchmarkGocraftWeb_GithubStatic     2000000        933 ns/op         328 B/op        6 allocs/op
BenchmarkGoji_GithubStatic           5000000        401 ns/op           0 B/op        0 allocs/op
BenchmarkGoJsonRest_GithubStatic      500000       6006 ns/op        1150 B/op       25 allocs/op
BenchmarkGorillaMux_GithubStatic      100000      18227 ns/op         456 B/op        6 allocs/op
BenchmarkHttpRouter_GithubStatic    50000000         63.2 ns/op         0 B/op        0 allocs/op
BenchmarkHttpTreeMux_GithubStatic   50000000         65.1 ns/op         0 B/op        0 allocs/op
BenchmarkKocha_GithubStatic         20000000         99.5 ns/op         0 B/op        0 allocs/op
BenchmarkMartini_GithubStatic         100000      18546 ns/op         927 B/op       11 allocs/op
BenchmarkPat_GithubStatic             200000      11503 ns/op        3754 B/op       76 allocs/op
BenchmarkTigerTonic_GithubStatic     5000000        308 ns/op          49 B/op        1 allocs/op
BenchmarkTraffic_GithubStatic          50000      44923 ns/op       23105 B/op      168 allocs/op

BenchmarkBeego_GithubParam             50000      44645 ns/op        2973 B/op       50 allocs/op
BenchmarkDenco_GithubParam           5000000        643 ns/op         115 B/op        3 allocs/op
BenchmarkGocraftWeb_GithubParam      1000000       1855 ns/op         750 B/op       10 allocs/op
BenchmarkGoji_GithubParam            1000000       1314 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_GithubParam       200000       8427 ns/op        2159 B/op       32 allocs/op
BenchmarkGorillaMux_GithubParam       200000      11485 ns/op         813 B/op        7 allocs/op
BenchmarkHttpRouter_GithubParam      5000000        304 ns/op          97 B/op        1 allocs/op
BenchmarkHttpTreeMux_GithubParam     2000000        770 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_GithubParam           5000000        754 ns/op         132 B/op        4 allocs/op
BenchmarkMartini_GithubParam          100000      22637 ns/op        1284 B/op       12 allocs/op
BenchmarkPat_GithubParam              500000       7319 ns/op        2538 B/op       44 allocs/op
BenchmarkTigerTonic_GithubParam       500000       4722 ns/op        1467 B/op       26 allocs/op
BenchmarkTraffic_GithubParam          100000      18700 ns/op        7076 B/op       58 allocs/op

BenchmarkBeego_GithubAll                 100   23430165 ns/op      502614 B/op     9871 allocs/op
BenchmarkDenco_GithubAll               10000     120365 ns/op       21219 B/op      506 allocs/op
BenchmarkGocraftWeb_GithubAll           5000     358982 ns/op      139091 B/op     1903 allocs/op
BenchmarkGoji_GithubAll                 5000     604522 ns/op       56896 B/op      340 allocs/op
BenchmarkGoJsonRest_GithubAll           1000    1645794 ns/op      404222 B/op     6303 allocs/op
BenchmarkGorillaMux_GithubAll            500    6634737 ns/op      152277 B/op     1402 allocs/op
BenchmarkHttpRouter_GithubAll          50000      51138 ns/op       14039 B/op      168 allocs/op
BenchmarkHttpTreeMux_GithubAll         10000     132507 ns/op       56907 B/op      340 allocs/op
BenchmarkKocha_GithubAll               10000     143398 ns/op       24117 B/op      676 allocs/op
BenchmarkMartini_GithubAll               200    9802351 ns/op      258349 B/op     2713 allocs/op
BenchmarkPat_GithubAll                   500    4154815 ns/op     1539081 B/op    24970 allocs/op
BenchmarkTigerTonic_GithubAll           2000     920839 ns/op      247085 B/op     5171 allocs/op
BenchmarkTraffic_GithubAll               200    8087393 ns/op     3143039 B/op    23958 allocs/op
```

### [Google+](https://developers.google.com/+/api/latest/)

Last but not least the Google+ API, consisting of 13 routes. In reality this is just a subset of a much larger API.

```
BenchmarkBeego_GPlusStatic           1000000       2321 ns/op         808 B/op       11 allocs/op
BenchmarkDenco_GPlusStatic          50000000         37.2 ns/op         0 B/op        0 allocs/op
BenchmarkGocraftWeb_GPlusStatic      2000000        862 ns/op         312 B/op        6 allocs/op
BenchmarkGoji_GPlusStatic           10000000        270 ns/op           0 B/op        0 allocs/op
BenchmarkGoJsonRest_GPlusStatic       500000       5827 ns/op        1136 B/op       25 allocs/op
BenchmarkGorillaMux_GPlusStatic      1000000       1793 ns/op         456 B/op        6 allocs/op
BenchmarkHttpRouter_GPlusStatic     50000000         34.6 ns/op         0 B/op        0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic    50000000         35.4 ns/op         0 B/op        0 allocs/op
BenchmarkKocha_GPlusStatic          50000000         63.8 ns/op         0 B/op        0 allocs/op
BenchmarkMartini_GPlusStatic          500000       4887 ns/op         927 B/op       11 allocs/op
BenchmarkPat_GPlusStatic             5000000        336 ns/op          98 B/op        2 allocs/op
BenchmarkTigerTonic_GPlusStatic     10000000        186 ns/op          33 B/op        1 allocs/op
BenchmarkTraffic_GPlusStatic          500000       3350 ns/op        1503 B/op       18 allocs/op

BenchmarkBeego_GPlusParam             200000       7657 ns/op        1231 B/op       16 allocs/op
BenchmarkDenco_GPlusParam            5000000        365 ns/op          50 B/op        2 allocs/op
BenchmarkGocraftWeb_GPlusParam       1000000       1519 ns/op         684 B/op        9 allocs/op
BenchmarkGoji_GPlusParam             2000000        889 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_GPlusParam        500000       7388 ns/op        1806 B/op       29 allocs/op
BenchmarkGorillaMux_GPlusParam        500000       4040 ns/op         780 B/op        7 allocs/op
BenchmarkHttpRouter_GPlusParam      10000000        203 ns/op          65 B/op        1 allocs/op
BenchmarkHttpTreeMux_GPlusParam      5000000        638 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_GPlusParam            5000000        444 ns/op          58 B/op        2 allocs/op
BenchmarkMartini_GPlusParam           200000       8672 ns/op        1251 B/op       12 allocs/op
BenchmarkPat_GPlusParam              1000000       1895 ns/op         719 B/op       12 allocs/op
BenchmarkTigerTonic_GPlusParam       1000000       3166 ns/op        1085 B/op       18 allocs/op
BenchmarkTraffic_GPlusParam           500000       5369 ns/op        2030 B/op       22 allocs/op

BenchmarkBeego_GPlus2Params           200000       9999 ns/op        1293 B/op       16 allocs/op
BenchmarkDenco_GPlus2Params          5000000        618 ns/op         115 B/op        3 allocs/op
BenchmarkGocraftWeb_GPlus2Params     1000000       1860 ns/op         750 B/op       10 allocs/op
BenchmarkGoji_GPlus2Params           1000000       1296 ns/op         340 B/op        2 allocs/op
BenchmarkGoJsonRest_GPlus2Params      200000       8516 ns/op        2178 B/op       32 allocs/op
BenchmarkGorillaMux_GPlus2Params      200000       9007 ns/op         812 B/op        7 allocs/op
BenchmarkHttpRouter_GPlus2Params    10000000        246 ns/op          65 B/op        1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params    2000000        751 ns/op         340 B/op        2 allocs/op
BenchmarkKocha_GPlus2Params          5000000        744 ns/op         132 B/op        4 allocs/op
BenchmarkMartini_GPlus2Params         100000      27700 ns/op        1382 B/op       16 allocs/op
BenchmarkPat_GPlus2Params             500000       5981 ns/op        2347 B/op       34 allocs/op
BenchmarkTigerTonic_GPlus2Params      500000       5076 ns/op        1561 B/op       27 allocs/op
BenchmarkTraffic_GPlus2Params         200000      12711 ns/op        3599 B/op       34 allocs/op

BenchmarkBeego_GPlusAll                10000     113601 ns/op       15897 B/op      217 allocs/op
BenchmarkDenco_GPlusAll               500000       5761 ns/op         880 B/op       27 allocs/op
BenchmarkGocraftWeb_GPlusAll          100000      20527 ns/op        8513 B/op      116 allocs/op
BenchmarkGoji_GPlusAll                200000      12312 ns/op        3746 B/op       22 allocs/op
BenchmarkGoJsonRest_GPlusAll           20000      99250 ns/op       23871 B/op      386 allocs/op
BenchmarkGorillaMux_GPlusAll           50000      63046 ns/op        9655 B/op       90 allocs/op
BenchmarkHttpRouter_GPlusAll         1000000       2513 ns/op         655 B/op       11 allocs/op
BenchmarkHttpTreeMux_GPlusAll         500000       7706 ns/op        3748 B/op       22 allocs/op
BenchmarkKocha_GPlusAll               500000       6858 ns/op        1017 B/op       35 allocs/op
BenchmarkMartini_GPlusAll              10000     155402 ns/op       16368 B/op      179 allocs/op
BenchmarkPat_GPlusAll                  50000      47397 ns/op       17270 B/op      302 allocs/op
BenchmarkTigerTonic_GPlusAll           50000      49864 ns/op       15160 B/op      311 allocs/op
BenchmarkTraffic_GPlusAll              10000     108007 ns/op       41779 B/op      430 allocs/op
```


## Conclusions
First of all, there is no reason to use net/http's default [ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is very limited and does not have especially good performance. There are enough alternatives coming in every flavor, choose the one you like best.

Secondly, the broad range of functions of some of the frameworks comes at a high price in terms of performance. For example Martini has great flexibility, but very bad performance. Martini has the worst performance of all tested routers in a lot of the benchmarks. Beego seems to have some scalability problems and easily defeats Martini with even worse performance, when the number of parameters or routes is high. I really hope, that the routing of these packages can be optimized. I think the Go-ecosystem needs great feature-rich frameworks like these.

Last but not least, we have to determine the performance champion.

Denco and its predecessor Kocha-urlrouter seem to have great performance, but are not convenient to use as a router for the net/http package. A lot of extra work is necessary to use it as a http.Handler. [The README of Denco claims](https://github.com/naoina/denco/blob/b03dbb499269a597afd0db715d408ebba1329d04/README.md), that the package is not intended as a replacement for [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).

[Goji](https://github.com/zenazn/goji/) looks very decent. It has great performance while also having a great range of features, more than any other router / framework in the top group.

Currently no router can beat the performance of the [HttpRouter](https://github.com/julienschmidt/httprouter) package, which currently dominates nearly all benchmarks.

In the end, performance can not be the (only) criterion for choosing a router. Play around a bit with some of the routers, and choose the one you like best.

## Usage

If you'd like to run these benchmarks locally, you'll need to install the package first:

```bash
go get github.com/julienschmidt/go-http-routing-benchmark
```
This may take a while due to the large number of dependencies that need to be downloaded. Once that command has finished you can run the full set of benchmarks like this:

```bash
cd $GOPATH/src/github.com/julienschmidt/go-http-routing-benchmark
go test -bench=.
```

> **Note:** If you run the tests and it SIGQUIT's make the go test timeout longer (#44)
>
```
go test -timeout=2h -bench=.
```


You can bench specific frameworks only by using a regular expression as the value of the `bench` parameter:
```bash
go test -bench="Martini|Gin|HttpMux"
```
