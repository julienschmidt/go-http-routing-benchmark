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
 * [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter)

## Results

Benchmark System:
 * Intel Core i5 M 580 (4x 2.67GHz)
 * 2x 4 GiB DDR3-1066 RAM
 * go1.2.1 linux/amd64
 * Arch Linux amd64 (Linux Kernel 3.13.5)

```
#GithubAPI Routes: 203
#GPlusAPI Routes: 13
#ParseAPI Routes: 26
#Static Routes: 157

BenchmarkHttpServeMux_StaticAll             2000           1206800 ns/op             104 B/op          8 allocs/op
BenchmarkGocraftWeb_StaticAll              10000            215578 ns/op           49164 B/op        951 allocs/op
BenchmarkGorillaMux_StaticAll                200           8634939 ns/op           72242 B/op        965 allocs/op
BenchmarkHttpRouter_StaticAll              50000             39715 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll                   500           4380765 ns/op          145600 B/op       2521 allocs/op
BenchmarkPat_StaticAll                      1000           2119511 ns/op          554153 B/op      11249 allocs/op
BenchmarkTigerTonic_StaticAll              20000             76547 ns/op            7776 B/op        158 allocs/op
BenchmarkTraffic_StaticAll                   100          15082073 ns/op         3794375 B/op      27918 allocs/op
BenchmarkKocha_StaticAll                   50000             69915 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_Param                1000000              2117 ns/op             672 B/op          9 allocs/op
BenchmarkGorillaMux_Param                 500000              6304 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_Param                2000000               766 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Param                    200000              7905 ns/op            1185 B/op         13 allocs/op
BenchmarkPat_Param                        500000              3500 ns/op            1061 B/op         17 allocs/op
BenchmarkTigerTonic_Param                1000000              3222 ns/op             830 B/op         16 allocs/op
BenchmarkTraffic_Param                    200000              8085 ns/op            2025 B/op         23 allocs/op
BenchmarkKocha_Param                     5000000               397 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_ParamWrite           1000000              2249 ns/op             681 B/op         10 allocs/op
BenchmarkGorillaMux_ParamWrite            500000              6604 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite           2000000               829 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParamWrite               200000              9273 ns/op            1285 B/op         16 allocs/op
BenchmarkPat_ParamWrite                   500000              4010 ns/op            1128 B/op         19 allocs/op
BenchmarkTigerTonic_ParamWrite            500000              5436 ns/op            1288 B/op         22 allocs/op
BenchmarkTraffic_ParamWrite               200000              9619 ns/op            2457 B/op         27 allocs/op
BenchmarkKocha_ParamWrite                5000000               473 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_GithubStatic         1000000              1272 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             48685 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000                93.2 ns/op             0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21634 ns/op             859 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14488 ns/op            3788 B/op         76 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               420 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             83558 ns/op           23358 B/op        172 allocs/op
BenchmarkKocha_GithubStatic             10000000               160 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GithubParam          1000000              2640 ns/op             735 B/op         10 allocs/op
BenchmarkGorillaMux_GithubParam            50000             33128 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          1000000              1044 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GithubParam              100000             29038 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10337 ns/op            2624 B/op         56 allocs/op
BenchmarkTigerTonic_GithubParam           500000              6001 ns/op            1289 B/op         25 allocs/op
BenchmarkTraffic_GithubParam               50000             38607 ns/op            7147 B/op         60 allocs/op
BenchmarkKocha_GithubParam               2000000               878 ns/op             116 B/op          3 allocs/op

BenchmarkGocraftWeb_GithubAll               5000            512132 ns/op          136323 B/op       1914 allocs/op
BenchmarkGorillaMux_GithubAll                100          21263115 ns/op          153308 B/op       1419 allocs/op
BenchmarkHttpRouter_GithubAll              10000            176380 ns/op           57346 B/op        347 allocs/op
BenchmarkMartini_GithubAll                   100          12964238 ns/op          245285 B/op       2940 allocs/op
BenchmarkPat_GithubAll                       500           5949296 ns/op         1588119 B/op      32571 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1157693 ns/op          218176 B/op       5582 allocs/op
BenchmarkTraffic_GithubAll                   100          18409192 ns/op         3172246 B/op      24930 allocs/op
BenchmarkKocha_GithubAll                   10000            173136 ns/op           21341 B/op        508 allocs/op

BenchmarkGocraftWeb_GPlusStatic          1000000              1172 ns/op             297 B/op          6 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4342 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000                52.5 ns/op             0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              200000              8963 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               490 ns/op              99 B/op          2 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               238 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6391 ns/op            1510 B/op         19 allocs/op
BenchmarkKocha_GPlusStatic              20000000               113 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GPlusParam           1000000              2330 ns/op             672 B/op          9 allocs/op
BenchmarkGorillaMux_GPlusParam            100000             14654 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam           2000000               858 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlusParam               200000             11640 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2719 ns/op             752 B/op         14 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              3694 ns/op             907 B/op         16 allocs/op
BenchmarkTraffic_GPlusParam               200000             11012 ns/op            2039 B/op         23 allocs/op
BenchmarkKocha_GPlusParam                5000000               661 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_GPlus2Params         1000000              2664 ns/op             735 B/op         10 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             29953 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         2000000               971 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlus2Params              50000             36663 ns/op            1318 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000             11717 ns/op            2399 B/op         41 allocs/op
BenchmarkTigerTonic_GPlus2Params          200000              9033 ns/op            1391 B/op         25 allocs/op
BenchmarkTraffic_GPlus2Params              50000             31592 ns/op            3620 B/op         35 allocs/op
BenchmarkKocha_GPlus2Params              2000000               935 ns/op             116 B/op          3 allocs/op

BenchmarkGocraftWeb_GPlusAll              100000             30018 ns/op            8338 B/op        117 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            185822 ns/op            9724 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll              200000             10378 ns/op            3774 B/op         22 allocs/op
BenchmarkMartini_GPlusAll                  10000            206067 ns/op           15522 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             67647 ns/op           17685 B/op        346 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             80948 ns/op           13327 B/op        289 allocs/op
BenchmarkTraffic_GPlusAll                  10000            247186 ns/op           42068 B/op        446 allocs/op
BenchmarkKocha_GPlusAll                   200000              8516 ns/op             886 B/op         27 allocs/op

BenchmarkGocraftWeb_ParseStatic          1000000              1264 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8750 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000                41.4 ns/op             0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              9175 ns/op             860 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1238 ns/op             249 B/op          5 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               355 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9983 ns/op            2390 B/op         25 allocs/op
BenchmarkKocha_ParseStatic              20000000               126 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_ParseParam           1000000              2248 ns/op             688 B/op          9 allocs/op
BenchmarkGorillaMux_ParseParam            200000             10315 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam           1000000              1088 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParseParam               200000             14338 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3957 ns/op            1196 B/op         20 allocs/op
BenchmarkTigerTonic_ParseParam            500000              3547 ns/op             888 B/op         16 allocs/op
BenchmarkTraffic_ParseParam               200000             10151 ns/op            2324 B/op         25 allocs/op
BenchmarkKocha_ParseParam                5000000               474 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_Parse2Params         1000000              2551 ns/op             735 B/op         10 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10691 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params         2000000              1049 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Parse2Params             200000             11347 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3942 ns/op             907 B/op         21 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              6033 ns/op            1293 B/op         25 allocs/op
BenchmarkTraffic_Parse2Params             200000             10340 ns/op            2130 B/op         25 allocs/op
BenchmarkKocha_Parse2Params              2000000              1079 ns/op             116 B/op          3 allocs/op

BenchmarkGocraftWeb_ParseAll               50000             52036 ns/op           14297 B/op        209 allocs/op
BenchmarkGorillaMux_ParseAll                5000            385096 ns/op           17253 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              200000             15579 ns/op            5489 B/op         33 allocs/op
BenchmarkMartini_ParseAll                   5000            384920 ns/op           27677 B/op        333 allocs/op
BenchmarkPat_ParseAll                      10000            104405 ns/op           18272 B/op        385 allocs/op
BenchmarkTigerTonic_ParseAll               10000            116717 ns/op           17729 B/op        372 allocs/op
BenchmarkTraffic_ParseAll                   5000            395349 ns/op           70545 B/op        763 allocs/op
BenchmarkKocha_ParseAll                   200000             12214 ns/op            1007 B/op         35 allocs/op
```
