go-http-routing-benchmark
=========================

This benchmark suite aims to compare the performance of available HTTP request routers for Go by implementing the routing of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.


Included Routers:

 * [Gocraft Web](https://github.com/gocraft/web)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [net/http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [Martini](https://github.com/codegangsta/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)

## Results

Benchmark System:
 * Intel Core i5-2500K (4x 3,30GHz + Turbo Boost, no Hyper-Threading)
 * 2x 4 GiB DDR3-1333 RAM
 * go1.2 linux/amd64
 * Ubuntu 13.10 Desktop amd64 (Linux Kernel 3.11.0)

```
#GithubAPI Routes: 203
#GPlusAPI Routes: 13
#ParseAPI Routes: 26
#Static Routes: 157

BenchmarkHttpServeMux_StaticAll         2000    1052206 ns/op         104 B/op        8 allocs/op
BenchmarkGocraftWeb_StaticAll          10000     166399 ns/op       49157 B/op      951 allocs/op
BenchmarkGorillaMux_StaticAll            500    4697626 ns/op       72291 B/op      965 allocs/op
BenchmarkHttpRouter_StaticAll         100000      16391 ns/op           0 B/op        0 allocs/op
BenchmarkMartini_StaticAll               500    3679323 ns/op      161525 B/op     3153 allocs/op
BenchmarkPat_StaticAll                  1000    1715771 ns/op      554033 B/op    11248 allocs/op
BenchmarkTigerTonic_StaticAll          50000      61535 ns/op        7775 B/op      158 allocs/op
BenchmarkTraffic_StaticAll               100   10972450 ns/op     3792598 B/op    27902 allocs/op


BenchmarkGocraftWeb_Param            1000000       1603 ns/op         672 B/op        9 allocs/op
BenchmarkGorillaMux_Param             500000       4703 ns/op         785 B/op        7 allocs/op
BenchmarkHttpRouter_Param            5000000        537 ns/op         343 B/op        2 allocs/op
BenchmarkMartini_Param                200000       7603 ns/op        1285 B/op       17 allocs/op
BenchmarkPat_Param                   1000000       2647 ns/op        1061 B/op       17 allocs/op
BenchmarkTigerTonic_Param            1000000       2505 ns/op         830 B/op       16 allocs/op
BenchmarkTraffic_Param                500000       6129 ns/op        2025 B/op       23 allocs/op

BenchmarkGocraftWeb_ParamWrite       1000000       1731 ns/op         681 B/op       10 allocs/op
BenchmarkGorillaMux_ParamWrite        500000       5050 ns/op         785 B/op        7 allocs/op
BenchmarkHttpRouter_ParamWrite       5000000        575 ns/op         343 B/op        2 allocs/op
BenchmarkMartini_ParamWrite           200000       8386 ns/op        1385 B/op       20 allocs/op
BenchmarkPat_ParamWrite              1000000       3185 ns/op        1128 B/op       19 allocs/op
BenchmarkTigerTonic_ParamWrite        500000       4368 ns/op        1288 B/op       22 allocs/op
BenchmarkTraffic_ParamWrite           500000       7157 ns/op        2457 B/op       27 allocs/op


BenchmarkGocraftWeb_GithubStatic     1000000       1032 ns/op         313 B/op        6 allocs/op
BenchmarkGorillaMux_GithubStatic       50000      39107 ns/op         459 B/op        6 allocs/op
BenchmarkHttpRouter_GithubStatic    50000000         58.5 ns/op         0 B/op        0 allocs/op
BenchmarkMartini_GithubStatic         100000      20292 ns/op         961 B/op       16 allocs/op
BenchmarkPat_GithubStatic             200000      12377 ns/op        3788 B/op       76 allocs/op
BenchmarkTigerTonic_GithubStatic     5000000        339 ns/op          49 B/op        1 allocs/op
BenchmarkTraffic_GithubStatic          50000      65180 ns/op       23359 B/op      172 allocs/op

BenchmarkGocraftWeb_GithubParam      1000000       2033 ns/op         736 B/op       10 allocs/op
BenchmarkGorillaMux_GithubParam       100000      26645 ns/op         818 B/op        7 allocs/op
BenchmarkHttpRouter_GithubParam      5000000        692 ns/op         343 B/op        2 allocs/op
BenchmarkMartini_GithubParam          100000      24937 ns/op        1318 B/op       17 allocs/op
BenchmarkPat_GithubParam              200000       8105 ns/op        2624 B/op       56 allocs/op
BenchmarkTigerTonic_GithubParam       500000       4732 ns/op        1289 B/op       25 allocs/op
BenchmarkTraffic_GithubParam          100000      29873 ns/op        7147 B/op       60 allocs/op

BenchmarkGocraftWeb_GithubAll           5000     395016 ns/op      136345 B/op     1914 allocs/op
BenchmarkGorillaMux_GithubAll            100   16305457 ns/op      153308 B/op     1419 allocs/op
BenchmarkHttpRouter_GithubAll          10000     123328 ns/op       57345 B/op      347 allocs/op
BenchmarkMartini_GithubAll               100   10639735 ns/op      265753 B/op     3757 allocs/op
BenchmarkPat_GithubAll                   500    4747315 ns/op     1588126 B/op    32571 allocs/op
BenchmarkTigerTonic_GithubAll           2000     910842 ns/op      218170 B/op     5582 allocs/op
BenchmarkTraffic_GithubAll               100   14018032 ns/op     3170857 B/op    24918 allocs/op


BenchmarkGocraftWeb_GPlusStatic      2000000        914 ns/op         297 B/op        6 allocs/op
BenchmarkGorillaMux_GPlusStatic       500000       3613 ns/op         459 B/op        6 allocs/op
BenchmarkHttpRouter_GPlusStatic     50000000         32.2 ns/op         0 B/op        0 allocs/op
BenchmarkMartini_GPlusStatic          500000       6128 ns/op         961 B/op       16 allocs/op
BenchmarkPat_GPlusStatic             5000000        357 ns/op          99 B/op        2 allocs/op
BenchmarkTigerTonic_GPlusStatic     10000000        192 ns/op          33 B/op        1 allocs/op
BenchmarkTraffic_GPlusStatic          500000       4734 ns/op        1509 B/op       19 allocs/op

BenchmarkGocraftWeb_GPlusParam       1000000       1692 ns/op         672 B/op        9 allocs/op
BenchmarkGorillaMux_GPlusParam        200000       9108 ns/op         785 B/op        7 allocs/op
BenchmarkHttpRouter_GPlusParam       5000000        584 ns/op         342 B/op        2 allocs/op
BenchmarkMartini_GPlusParam           200000      10352 ns/op        1286 B/op       17 allocs/op
BenchmarkPat_GPlusParam              1000000       2132 ns/op         752 B/op       14 allocs/op
BenchmarkTigerTonic_GPlusParam       1000000       2919 ns/op         907 B/op       16 allocs/op
BenchmarkTraffic_GPlusParam           200000       7812 ns/op        2038 B/op       23 allocs/op

BenchmarkGocraftWeb_GPlus2Params     1000000       2050 ns/op         735 B/op       10 allocs/op
BenchmarkGorillaMux_GPlus2Params      100000      23288 ns/op         818 B/op        7 allocs/op
BenchmarkHttpRouter_GPlus2Params     5000000        670 ns/op         342 B/op        2 allocs/op
BenchmarkMartini_GPlus2Params          50000      29990 ns/op        1419 B/op       21 allocs/op
BenchmarkPat_GPlus2Params             500000       6880 ns/op        2399 B/op       41 allocs/op
BenchmarkTigerTonic_GPlus2Params      500000       4936 ns/op        1391 B/op       25 allocs/op
BenchmarkTraffic_GPlus2Params         100000      24330 ns/op        3619 B/op       35 allocs/op

BenchmarkGocraftWeb_GPlusAll          100000      23006 ns/op        8338 B/op      117 allocs/op
BenchmarkGorillaMux_GPlusAll           10000     147360 ns/op        9722 B/op       91 allocs/op
BenchmarkHttpRouter_GPlusAll          500000       7276 ns/op        3772 B/op       22 allocs/op
BenchmarkMartini_GPlusAll              10000     182433 ns/op       16842 B/op      246 allocs/op
BenchmarkPat_GPlusAll                  50000      52997 ns/op       17685 B/op      346 allocs/op
BenchmarkTigerTonic_GPlusAll           50000      48984 ns/op       13327 B/op      289 allocs/op
BenchmarkTraffic_GPlusAll              10000     165602 ns/op       42062 B/op      446 allocs/op


BenchmarkGocraftWeb_ParseStatic      1000000       1008 ns/op         313 B/op        6 allocs/op
BenchmarkGorillaMux_ParseStatic       500000       7243 ns/op         459 B/op        6 allocs/op
BenchmarkHttpRouter_ParseStatic     50000000         30.8 ns/op         0 B/op        0 allocs/op
BenchmarkMartini_ParseStatic          500000       6841 ns/op         961 B/op       16 allocs/op
BenchmarkPat_ParseStatic             2000000        862 ns/op         249 B/op        5 allocs/op
BenchmarkTigerTonic_ParseStatic     10000000        281 ns/op          49 B/op        1 allocs/op
BenchmarkTraffic_ParseStatic          500000       7526 ns/op        2389 B/op       25 allocs/op

BenchmarkGocraftWeb_ParseParam       1000000       1688 ns/op         688 B/op        9 allocs/op
BenchmarkGorillaMux_ParseParam        500000       7632 ns/op         785 B/op        7 allocs/op
BenchmarkHttpRouter_ParseParam       5000000        562 ns/op         343 B/op        2 allocs/op
BenchmarkMartini_ParseParam           200000       8900 ns/op        1286 B/op       17 allocs/op
BenchmarkPat_ParseParam              1000000       2888 ns/op        1196 B/op       20 allocs/op
BenchmarkTigerTonic_ParseParam       1000000       2621 ns/op         888 B/op       16 allocs/op
BenchmarkTraffic_ParseParam           500000       7349 ns/op        2323 B/op       25 allocs/op

BenchmarkGocraftWeb_Parse2Params     1000000       1848 ns/op         735 B/op       10 allocs/op
BenchmarkGorillaMux_Parse2Params      500000       7559 ns/op         818 B/op        7 allocs/op
BenchmarkHttpRouter_Parse2Params     5000000        631 ns/op         343 B/op        2 allocs/op
BenchmarkMartini_Parse2Params         200000       9059 ns/op        1318 B/op       17 allocs/op
BenchmarkPat_Parse2Params            1000000       2862 ns/op         907 B/op       21 allocs/op
BenchmarkTigerTonic_Parse2Params      500000       4415 ns/op        1293 B/op       25 allocs/op
BenchmarkTraffic_Parse2Params         500000       7363 ns/op        2129 B/op       25 allocs/op

BenchmarkGocraftWeb_ParseAll           50000      38023 ns/op       14293 B/op      209 allocs/op
BenchmarkGorillaMux_ParseAll           10000     296433 ns/op       17248 B/op      175 allocs/op
BenchmarkHttpRouter_ParseAll          200000      10597 ns/op        5488 B/op       33 allocs/op
BenchmarkMartini_ParseAll              10000     236616 ns/op       30296 B/op      438 allocs/op
BenchmarkPat_ParseAll                  50000      56646 ns/op       18265 B/op      385 allocs/op
BenchmarkTigerTonic_ParseAll           50000      61418 ns/op       17723 B/op      372 allocs/op
BenchmarkTraffic_ParseAll              10000     236261 ns/op       70505 B/op      762 allocs/op
```
