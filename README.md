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

BenchmarkGocraftWeb_Param            1000000          2102 ns/op         672 B/op          9 allocs/op
BenchmarkGorillaMux_Param             500000          5623 ns/op         786 B/op          7 allocs/op
BenchmarkHttpRouter_Param            5000000           643 ns/op         343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Param           5000000           702 ns/op         343 B/op          2 allocs/op
BenchmarkMartini_Param                500000          7276 ns/op        1186 B/op         13 allocs/op
BenchmarkPat_Param                   1000000          3033 ns/op        1061 B/op         17 allocs/op
BenchmarkTigerTonic_Param             500000          3607 ns/op        1025 B/op         19 allocs/op
BenchmarkTraffic_Param                500000          6747 ns/op        2026 B/op         23 allocs/op


BenchmarkGocraftWeb_ParamWrite       1000000          2003 ns/op         682 B/op         10 allocs/op
BenchmarkGorillaMux_ParamWrite        500000          5689 ns/op         786 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite       5000000           673 ns/op         343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParamWrite      5000000           736 ns/op         343 B/op          2 allocs/op
BenchmarkMartini_ParamWrite           200000          8291 ns/op        1285 B/op         16 allocs/op
BenchmarkPat_ParamWrite               500000          3580 ns/op        1128 B/op         19 allocs/op
BenchmarkTigerTonic_ParamWrite        500000          5605 ns/op        1483 B/op         25 allocs/op
BenchmarkTraffic_ParamWrite           200000          8182 ns/op        2459 B/op         27 allocs/op


BenchmarkGocraftWeb_GithubStatic     1000000          1170 ns/op         313 B/op          6 allocs/op
BenchmarkGorillaMux_GithubStatic       50000         43501 ns/op         459 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic    50000000          68.5 ns/op           0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic   20000000          76.6 ns/op           0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic         100000         21715 ns/op         860 B/op         12 allocs/op
BenchmarkPat_GithubStatic             200000         13604 ns/op        3788 B/op         76 allocs/op
BenchmarkTigerTonic_GithubStatic     5000000           396 ns/op          49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic          50000         73229 ns/op       23360 B/op        172 allocs/op


BenchmarkGocraftWeb_GithubParam      1000000          2340 ns/op         736 B/op         10 allocs/op
BenchmarkGorillaMux_GithubParam        50000         29444 ns/op         818 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam      2000000           803 ns/op         343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GithubParam     2000000           914 ns/op         343 B/op          2 allocs/op
BenchmarkMartini_GithubParam          100000         26148 ns/op        1219 B/op         13 allocs/op
BenchmarkPat_GithubParam              200000          9289 ns/op        2625 B/op         56 allocs/op
BenchmarkTigerTonic_GithubParam       500000          6007 ns/op        1483 B/op         28 allocs/op
BenchmarkTraffic_GithubParam           50000         33384 ns/op        7148 B/op         60 allocs/op


BenchmarkGocraftWeb_GithubAll           5000        462027 ns/op      136369 B/op       1914 allocs/op
BenchmarkGorillaMux_GithubAll            100      17838396 ns/op      153368 B/op       1419 allocs/op
BenchmarkHttpRouter_GithubAll          10000        146248 ns/op       57352 B/op        347 allocs/op
BenchmarkHttpTreeMux_GithubAll         10000        164037 ns/op       57349 B/op        347 allocs/op
BenchmarkMartini_GithubAll               100      11596404 ns/op      245412 B/op       2941 allocs/op
BenchmarkPat_GithubAll                   500       5348021 ns/op     1588380 B/op      32572 allocs/op
BenchmarkTigerTonic_GithubAll           2000       1204584 ns/op      250837 B/op       6086 allocs/op
BenchmarkTraffic_GithubAll               100      15706344 ns/op     3171229 B/op      24921 allocs/op


BenchmarkGocraftWeb_GPlusStatic      1000000          1089 ns/op         297 B/op          6 allocs/op
BenchmarkGorillaMux_GPlusStatic       500000          4133 ns/op         460 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic     50000000          37.7 ns/op           0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic    50000000          42.3 ns/op           0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic          500000          5973 ns/op         860 B/op         12 allocs/op
BenchmarkPat_GPlusStatic             5000000           419 ns/op          99 B/op          2 allocs/op
BenchmarkTigerTonic_GPlusStatic     10000000           224 ns/op          33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic          500000          5545 ns/op        1510 B/op         19 allocs/op


BenchmarkGocraftWeb_GPlusParam       1000000          1951 ns/op         673 B/op          9 allocs/op
BenchmarkGorillaMux_GPlusParam        200000          9264 ns/op         786 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam       5000000           686 ns/op         343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlusParam      2000000           754 ns/op         343 B/op          2 allocs/op
BenchmarkMartini_GPlusParam           200000         10017 ns/op        1186 B/op         13 allocs/op
BenchmarkPat_GPlusParam              1000000          2435 ns/op         752 B/op         14 allocs/op
BenchmarkTigerTonic_GPlusParam        500000          3980 ns/op        1102 B/op         19 allocs/op
BenchmarkTraffic_GPlusParam           200000          9026 ns/op        2039 B/op         23 allocs/op


BenchmarkGocraftWeb_GPlus2Params     1000000          2371 ns/op         735 B/op         10 allocs/op
BenchmarkGorillaMux_GPlus2Params      100000         25805 ns/op         818 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params     2000000           781 ns/op         343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlus2Params    2000000           917 ns/op         343 B/op          2 allocs/op
BenchmarkMartini_GPlus2Params          50000         32006 ns/op        1318 B/op         17 allocs/op
BenchmarkPat_GPlus2Params             200000          7644 ns/op        2400 B/op         41 allocs/op
BenchmarkTigerTonic_GPlus2Params      500000          6375 ns/op        1584 B/op         28 allocs/op
BenchmarkTraffic_GPlus2Params         100000         27404 ns/op        3621 B/op         35 allocs/op


BenchmarkGocraftWeb_GPlusAll          100000         26458 ns/op        8341 B/op        117 allocs/op
BenchmarkGorillaMux_GPlusAll           10000        161635 ns/op        9725 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll          200000          8348 ns/op        3773 B/op         22 allocs/op
BenchmarkHttpTreeMux_GPlusAll         200000          9325 ns/op        3773 B/op         22 allocs/op
BenchmarkMartini_GPlusAll              10000        184521 ns/op       15529 B/op        194 allocs/op
BenchmarkPat_GPlusAll                  50000         60253 ns/op       17691 B/op        346 allocs/op
BenchmarkTigerTonic_GPlusAll           50000         63487 ns/op       15475 B/op        322 allocs/op
BenchmarkTraffic_GPlusAll              10000        193173 ns/op       42077 B/op        446 allocs/op


BenchmarkGocraftWeb_ParseStatic      1000000          1139 ns/op         313 B/op          6 allocs/op
BenchmarkGorillaMux_ParseStatic       200000          8058 ns/op         459 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic     50000000          38.6 ns/op           0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic    20000000          74.1 ns/op           0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic          500000          6298 ns/op         860 B/op         12 allocs/op
BenchmarkPat_ParseStatic             2000000           977 ns/op         249 B/op          5 allocs/op
BenchmarkTigerTonic_ParseStatic      5000000           316 ns/op          49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic          200000          8655 ns/op        2391 B/op         25 allocs/op


BenchmarkGocraftWeb_ParseParam       1000000          1973 ns/op         689 B/op          9 allocs/op
BenchmarkGorillaMux_ParseParam        200000          8838 ns/op         786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam       5000000           649 ns/op         343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParseParam      5000000           730 ns/op         343 B/op          2 allocs/op
BenchmarkMartini_ParseParam           200000          8758 ns/op        1186 B/op         13 allocs/op
BenchmarkPat_ParseParam               500000          3419 ns/op        1196 B/op         20 allocs/op
BenchmarkTigerTonic_ParseParam        500000          3795 ns/op        1083 B/op         19 allocs/op
BenchmarkTraffic_ParseParam           200000          8684 ns/op        2325 B/op         25 allocs/op


BenchmarkGocraftWeb_Parse2Params     1000000          2249 ns/op         736 B/op         10 allocs/op
BenchmarkGorillaMux_Parse2Params      200000          9035 ns/op         818 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params     5000000           731 ns/op         343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Parse2Params    2000000           871 ns/op         343 B/op          2 allocs/op
BenchmarkMartini_Parse2Params         200000          9133 ns/op        1219 B/op         13 allocs/op
BenchmarkPat_Parse2Params             500000          3368 ns/op         907 B/op         21 allocs/op
BenchmarkTigerTonic_Parse2Params      500000          6018 ns/op        1486 B/op         28 allocs/op
BenchmarkTraffic_Parse2Params         200000          8780 ns/op        2130 B/op         25 allocs/op


BenchmarkGocraftWeb_ParseAll           50000         46535 ns/op       14300 B/op        209 allocs/op
BenchmarkGorillaMux_ParseAll            5000        341061 ns/op       17258 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll          200000         12386 ns/op        5489 B/op         33 allocs/op
BenchmarkHttpTreeMux_ParseAll         200000         14633 ns/op        5489 B/op         33 allocs/op
BenchmarkMartini_ParseAll              10000        245201 ns/op       27684 B/op        333 allocs/op
BenchmarkPat_ParseAll                  50000         74691 ns/op       18278 B/op        385 allocs/op
BenchmarkTigerTonic_ParseAll           10000        101282 ns/op       20849 B/op        420 allocs/op
BenchmarkTraffic_ParseAll               5000        333251 ns/op       70549 B/op        762 allocs/op


BenchmarkHttpServeMux_StaticAll         2000       1302656 ns/op         104 B/op          8 allocs/op
BenchmarkGocraftWeb_StaticAll          10000        205982 ns/op       49171 B/op        951 allocs/op
BenchmarkGorillaMux_StaticAll            500       5407232 ns/op       72319 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll         100000         20377 ns/op           0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll        100000         20282 ns/op           0 B/op          0 allocs/op
BenchmarkMartini_StaticAll               500       4322012 ns/op      145650 B/op       2521 allocs/op
BenchmarkPat_StaticAll                  1000       2113952 ns/op      554086 B/op      11249 allocs/op
BenchmarkTigerTonic_StaticAll          20000         74334 ns/op        7777 B/op        158 allocs/op
BenchmarkTraffic_StaticAll               100      13198546 ns/op     3793052 B/op      27906 allocs/op
```
