```
D:\Projects\Golang\mygopath\src\github.com\kataras\go-http-routing-benchmark>go test -bench=. -timeout=60m
#GithubAPI Routes: 203
   Ace: 48992 Bytes
   Bear: 161704 Bytes
   Beego: 144536 Bytes
   Bone: 97696 Bytes
   Denco: 36728 Bytes
   Echo: 76424 Bytes
   Gin: 52464 Bytes
   GocraftWeb: 95304 Bytes
   Goji: 86104 Bytes
   Gojiv2: 144408 Bytes
   GoJsonRest: 134072 Bytes
   GoRestful: 1395560 Bytes
   GorillaMux: 1494848 Bytes
   HttpRouter: 37464 Bytes
   HttpTreeMux: 78736 Bytes
   Iris: 62312 Bytes
   Kocha: 785120 Bytes
   LARS: 49016 Bytes
   Macaron: 127304 Bytes
   Martini: 556192 Bytes
   Pat: 21200 Bytes
   Possum: 82928 Bytes
   R2router: 47104 Bytes
   Revel: 141504 Bytes
   Rivet: 42840 Bytes
   Tango: 54584 Bytes
   TigerTonic: 96032 Bytes
   Traffic: 1053776 Bytes
   Vulcan: 464976 Bytes

#GPlusAPI Routes: 13
   Ace: 3600 Bytes
   Bear: 7112 Bytes
   Beego: 9712 Bytes
   Bone: 6448 Bytes
   Denco: 3256 Bytes
   Echo: 7224 Bytes
   Gin: 3856 Bytes
   GocraftWeb: 7496 Bytes
   Goji: 2912 Bytes
   Gojiv2: 7376 Bytes
   GoJsonRest: 11416 Bytes
   GoRestful: 87608 Bytes
   GorillaMux: 71072 Bytes
   HttpRouter: 2712 Bytes
   HttpTreeMux: 7376 Bytes
   Iris: 6480 Bytes
   Kocha: 128880 Bytes
   LARS: 3624 Bytes
   Macaron: 8448 Bytes
   Martini: 23936 Bytes
   Pat: 1856 Bytes
   Possum: 7248 Bytes
   R2router: 3928 Bytes
   Revel: 10768 Bytes
   Rivet: 3064 Bytes
   Tango: 4912 Bytes
   TigerTonic: 9408 Bytes
   Traffic: 49472 Bytes
   Vulcan: 25496 Bytes

#ParseAPI Routes: 26
   Ace: 6592 Bytes
   Bear: 12320 Bytes
   Beego: 18416 Bytes
   Bone: 10992 Bytes
   Denco: 4184 Bytes
   Echo: 8144 Bytes
   Gin: 6816 Bytes
   GocraftWeb: 12800 Bytes
   Goji: 5232 Bytes
   Gojiv2: 14464 Bytes
   GoJsonRest: 14088 Bytes
   GoRestful: 126216 Bytes
   GorillaMux: 122184 Bytes
   HttpRouter: 4976 Bytes
   HttpTreeMux: 7784 Bytes
   Iris: 9928 Bytes
   Kocha: 181712 Bytes
   LARS: 6616 Bytes
   Macaron: 13232 Bytes
   Martini: 45952 Bytes
   Pat: 2560 Bytes
   Possum: 10224 Bytes
   R2router: 7056 Bytes
   Revel: 15488 Bytes
   Rivet: 5680 Bytes
   Tango: 8664 Bytes
   TigerTonic: 9840 Bytes
   Traffic: 93480 Bytes
   Vulcan: 44712 Bytes

#Static Routes: 157
   HttpServeMux: 18784 Bytes
   Ace: 30080 Bytes
   Bear: 30424 Bytes
   Beego: 93768 Bytes
   Bone: 37872 Bytes
   Denco: 9696 Bytes
   Echo: 61120 Bytes
   Gin: 30400 Bytes
   GocraftWeb: 55464 Bytes
   Goji: 27200 Bytes
   Gojiv2: 104464 Bytes
   GoJsonRest: 135736 Bytes
   GoRestful: 908200 Bytes
   GorillaMux: 668496 Bytes
   HttpRouter: 21128 Bytes
   HttpTreeMux: 73384 Bytes
   Iris: 37264 Bytes
   Kocha: 114816 Bytes
   LARS: 30104 Bytes
   Macaron: 35136 Bytes
   Martini: 308784 Bytes
   Pat: 20464 Bytes
   Possum: 90848 Bytes
   R2router: 23712 Bytes
   Revel: 93392 Bytes
   Rivet: 23880 Bytes
   Tango: 28008 Bytes
   TigerTonic: 80320 Bytes
   Traffic: 624432 Bytes
   Vulcan: 369144 Bytes

PASS
BenchmarkAce_Param                       5000000               347 ns/op              32 B/op          1 allocs/op
BenchmarkBear_Param                      1000000              1313 ns/op             456 B/op          5 allocs/op
BenchmarkBeego_Param                     1000000              1173 ns/op              64 B/op          4 allocs/op
BenchmarkBone_Param                      1000000              1255 ns/op             384 B/op          3 allocs/op
BenchmarkDenco_Param                    10000000               230 ns/op              32 B/op          1 allocs/op
BenchmarkEcho_Param                     20000000                91.8 ns/op             0 B/op          0 allocs/op
BenchmarkGin_Param                      20000000                87.0 ns/op             0 B/op          0 allocs/op
BenchmarkGocraftWeb_Param                1000000              1961 ns/op             648 B/op          8 allocs/op
BenchmarkGoji_Param                      2000000               979 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_Param                    2000000               837 ns/op             176 B/op          5 allocs/op
BenchmarkGoJsonRest_Param                1000000              2129 ns/op             649 B/op         13 allocs/op
BenchmarkGoRestful_Param                  200000              8860 ns/op            2696 B/op         27 allocs/op
BenchmarkGorillaMux_Param                1000000              3152 ns/op             752 B/op          8 allocs/op
BenchmarkHttpRouter_Param               10000000               156 ns/op              32 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param               2000000               901 ns/op             352 B/op          3 allocs/op
BenchmarkIris_Param                     30000000                49.7 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_Param                     3000000               411 ns/op              56 B/op          3 allocs/op
BenchmarkLARS_Param                     20000000                92.3 ns/op             0 B/op          0 allocs/op
BenchmarkMacaron_Param                   1000000              3269 ns/op            1040 B/op          9 allocs/op
BenchmarkMartini_Param                    500000              5216 ns/op            1104 B/op         11 allocs/op
BenchmarkPat_Param                       1000000              2203 ns/op             648 B/op         12 allocs/op
BenchmarkPossum_Param                    1000000              1988 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_Param                  1000000              1235 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_Param                      300000              6060 ns/op            1632 B/op         26 allocs/op
BenchmarkRivet_Param                     5000000               263 ns/op              48 B/op          1 allocs/op
BenchmarkTango_Param                     1000000              1436 ns/op             256 B/op          9 allocs/op
BenchmarkTigerTonic_Param                1000000              3699 ns/op             976 B/op         16 allocs/op
BenchmarkTraffic_Param                    300000              5837 ns/op            1960 B/op         21 allocs/op
BenchmarkVulcan_Param                    2000000               859 ns/op              98 B/op          3 allocs/op
BenchmarkAce_Param5                      2000000               661 ns/op             160 B/op          1 allocs/op
BenchmarkBear_Param5                     1000000              1770 ns/op             501 B/op          5 allocs/op
BenchmarkBeego_Param5                    1000000              1668 ns/op             128 B/op          4 allocs/op
BenchmarkBone_Param5                     1000000              1687 ns/op             432 B/op          3 allocs/op
BenchmarkDenco_Param5                    2000000               636 ns/op             160 B/op          1 allocs/op
BenchmarkEcho_Param5                    10000000               163 ns/op               0 B/op          0 allocs/op
BenchmarkGin_Param5                     10000000               147 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_Param5               1000000              3091 ns/op             920 B/op         11 allocs/op
BenchmarkGoji_Param5                     1000000              1294 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_Param5                   1000000              1175 ns/op             240 B/op          5 allocs/op
BenchmarkGoJsonRest_Param5               1000000              3989 ns/op            1097 B/op         16 allocs/op
BenchmarkGoRestful_Param5                 200000             10665 ns/op            2872 B/op         27 allocs/op
BenchmarkGorillaMux_Param5                500000              4872 ns/op             816 B/op          8 allocs/op
BenchmarkHttpRouter_Param5               3000000               499 ns/op             160 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param5              1000000              2133 ns/op             576 B/op          6 allocs/op
BenchmarkIris_Param5                    30000000                49.3 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_Param5                    1000000              1917 ns/op             440 B/op         10 allocs/op
BenchmarkLARS_Param5                    10000000               155 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_Param5                  1000000              3742 ns/op            1040 B/op          9 allocs/op
BenchmarkMartini_Param5                   300000              7177 ns/op            1232 B/op         11 allocs/op
BenchmarkPat_Param5                       300000              5273 ns/op             964 B/op         32 allocs/op
BenchmarkPossum_Param5                   1000000              2037 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_Param5                 1000000              1569 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_Param5                     200000              7785 ns/op            1984 B/op         33 allocs/op
BenchmarkRivet_Param5                    2000000               867 ns/op             240 B/op          1 allocs/op
BenchmarkTango_Param5                    1000000              3771 ns/op             944 B/op         17 allocs/op
BenchmarkTigerTonic_Param5                200000             11565 ns/op            2471 B/op         38 allocs/op
BenchmarkTraffic_Param5                   200000              8700 ns/op            2248 B/op         25 allocs/op
BenchmarkVulcan_Param5                   1000000              1122 ns/op              98 B/op          3 allocs/op
BenchmarkAce_Param20                     1000000              1707 ns/op             640 B/op          1 allocs/op
BenchmarkBear_Param20                     300000              5340 ns/op            1665 B/op          5 allocs/op
BenchmarkBeego_Param20                    500000              4020 ns/op             192 B/op          4 allocs/op
BenchmarkBone_Param20                     300000              7937 ns/op            2540 B/op          5 allocs/op
BenchmarkDenco_Param20                   1000000              2145 ns/op             640 B/op          1 allocs/op
BenchmarkEcho_Param20                    3000000               439 ns/op               0 B/op          0 allocs/op
BenchmarkGin_Param20                     5000000               358 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_Param20               200000             10870 ns/op            3796 B/op         15 allocs/op
BenchmarkGoji_Param20                    1000000              4127 ns/op            1247 B/op          2 allocs/op
BenchmarkGojiv2_Param20                  1000000              1969 ns/op             480 B/op          5 allocs/op
BenchmarkGoJsonRest_Param20               100000             13760 ns/op            4484 B/op         20 allocs/op
BenchmarkGoRestful_Param20                100000             18921 ns/op            5443 B/op         29 allocs/op
BenchmarkGorillaMux_Param20               200000             11510 ns/op            2923 B/op         10 allocs/op
BenchmarkHttpRouter_Param20              1000000              1513 ns/op             640 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param20              200000              9770 ns/op            3194 B/op         10 allocs/op
BenchmarkKocha_Param20                    300000              5990 ns/op            1808 B/op         27 allocs/op
BenchmarkLARS_Param20                    5000000               344 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_Param20                  200000              9670 ns/op            2892 B/op         11 allocs/op
BenchmarkMartini_Param20                  100000             13810 ns/op            3597 B/op         13 allocs/op
BenchmarkPat_Param20                      100000             24021 ns/op            4687 B/op        111 allocs/op
BenchmarkPossum_Param20                  1000000              1986 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_Param20                 300000              7117 ns/op            2283 B/op          7 allocs/op
BenchmarkRevel_Param20                    100000             17220 ns/op            5511 B/op         52 allocs/op
BenchmarkRivet_Param20                   1000000              2958 ns/op            1024 B/op          1 allocs/op
BenchmarkTango_Param20                    100000             20371 ns/op            8224 B/op         47 allocs/op
BenchmarkTigerTonic_Param20                30000             45469 ns/op           10344 B/op        118 allocs/op
BenchmarkTraffic_Param20                   50000             29161 ns/op            7944 B/op         45 allocs/op
BenchmarkVulcan_Param20                  1000000              1784 ns/op              98 B/op          3 allocs/op
BenchmarkAce_ParamWrite                  3000000               447 ns/op              40 B/op          2 allocs/op
BenchmarkBear_ParamWrite                 1000000              1370 ns/op             456 B/op          5 allocs/op
BenchmarkBeego_ParamWrite                1000000              1298 ns/op              72 B/op          5 allocs/op
BenchmarkBone_ParamWrite                 1000000              1347 ns/op             384 B/op          3 allocs/op
BenchmarkDenco_ParamWrite                5000000               289 ns/op              32 B/op          1 allocs/op
BenchmarkEcho_ParamWrite                10000000               204 ns/op               8 B/op          1 allocs/op
BenchmarkGin_ParamWrite                 10000000               181 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParamWrite           1000000              2096 ns/op             656 B/op          9 allocs/op
BenchmarkGoji_ParamWrite                 1000000              1051 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_ParamWrite               1000000              1118 ns/op             208 B/op          7 allocs/op
BenchmarkGoJsonRest_ParamWrite           1000000              3461 ns/op            1128 B/op         18 allocs/op
BenchmarkGoRestful_ParamWrite             200000              9085 ns/op            2704 B/op         28 allocs/op
BenchmarkGorillaMux_ParamWrite           1000000              3290 ns/op             752 B/op          8 allocs/op
BenchmarkHttpRouter_ParamWrite          10000000               211 ns/op              32 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParamWrite          2000000               967 ns/op             352 B/op          3 allocs/op
BenchmarkIris_ParamWrite                10000000               145 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParamWrite                3000000               476 ns/op              56 B/op          3 allocs/op
BenchmarkLARS_ParamWrite                10000000               174 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_ParamWrite              1000000              3847 ns/op            1144 B/op         13 allocs/op
BenchmarkMartini_ParamWrite               300000              5983 ns/op            1208 B/op         15 allocs/op
BenchmarkPat_ParamWrite                  1000000              3637 ns/op            1072 B/op         17 allocs/op
BenchmarkPossum_ParamWrite               1000000              2011 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_ParamWrite             1000000              1335 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_ParamWrite                 200000              7270 ns/op            2096 B/op         31 allocs/op
BenchmarkRivet_ParamWrite                3000000               540 ns/op             144 B/op          3 allocs/op
BenchmarkTango_ParamWrite                2000000               775 ns/op             136 B/op          4 allocs/op
BenchmarkTigerTonic_ParamWrite            300000              5723 ns/op            1408 B/op         22 allocs/op
BenchmarkTraffic_ParamWrite               300000              7230 ns/op            2384 B/op         25 allocs/op
BenchmarkVulcan_ParamWrite               2000000               854 ns/op              98 B/op          3 allocs/op
BenchmarkAce_GithubStatic               10000000               223 ns/op               0 B/op          0 allocs/op
BenchmarkBear_GithubStatic               2000000               613 ns/op             120 B/op          3 allocs/op
BenchmarkBeego_GithubStatic              1000000              1173 ns/op              64 B/op          4 allocs/op
BenchmarkBone_GithubStatic                100000             14740 ns/op            2880 B/op         60 allocs/op
BenchmarkDenco_GithubStatic             30000000                46.8 ns/op             0 B/op          0 allocs/op
BenchmarkEcho_GithubStatic              20000000               111 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GithubStatic               20000000               114 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GithubStatic         1000000              1140 ns/op             296 B/op          5 allocs/op
BenchmarkGoji_GithubStatic               5000000               260 ns/op               0 B/op          0 allocs/op
BenchmarkGojiv2_GithubStatic             2000000               810 ns/op             160 B/op          4 allocs/op
BenchmarkGoRestful_GithubStatic            30000             43835 ns/op            3720 B/op         32 allocs/op
BenchmarkGoJsonRest_GithubStatic         1000000              1501 ns/op             329 B/op         11 allocs/op
BenchmarkGorillaMux_GithubStatic          100000             20281 ns/op             448 B/op          7 allocs/op
BenchmarkHttpRouter_GithubStatic        30000000                56.5 ns/op             0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       20000000                66.5 ns/op             0 B/op          0 allocs/op
BenchmarkIris_GithubStatic              30000000                48.5 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic             20000000                76.2 ns/op             0 B/op          0 allocs/op
BenchmarkLARS_GithubStatic              20000000               109 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_GithubStatic            1000000              2751 ns/op             752 B/op          8 allocs/op
BenchmarkMartini_GithubStatic             100000             15170 ns/op             784 B/op         10 allocs/op
BenchmarkPat_GithubStatic                 100000             12500 ns/op            3648 B/op         76 allocs/op
BenchmarkPossum_GithubStatic             1000000              1338 ns/op             416 B/op          3 allocs/op
BenchmarkR2router_GithubStatic           2000000               672 ns/op             144 B/op          4 allocs/op
BenchmarkRevel_GithubStatic               500000              4906 ns/op            1248 B/op         23 allocs/op
BenchmarkRivet_GithubStatic             10000000               121 ns/op               0 B/op          0 allocs/op
BenchmarkTango_GithubStatic              1000000              1620 ns/op             256 B/op          9 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               350 ns/op              48 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              30000             55736 ns/op           18904 B/op        148 allocs/op
BenchmarkVulcan_GithubStatic             1000000              1266 ns/op              98 B/op          3 allocs/op
BenchmarkAce_GithubParam                 2000000               584 ns/op              96 B/op          1 allocs/op
BenchmarkBear_GithubParam                1000000              1628 ns/op             496 B/op          5 allocs/op
BenchmarkBeego_GithubParam               1000000              1667 ns/op             192 B/op          4 allocs/op
BenchmarkBone_GithubParam                 300000              6553 ns/op            1456 B/op         16 allocs/op
BenchmarkDenco_GithubParam               3000000               525 ns/op             128 B/op          1 allocs/op
BenchmarkEcho_GithubParam               10000000               186 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GithubParam                10000000               184 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GithubParam          1000000              2264 ns/op             712 B/op          9 allocs/op
BenchmarkGoji_GithubParam                1000000              1355 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_GithubParam              1000000              1707 ns/op             256 B/op          7 allocs/op
BenchmarkGoJsonRest_GithubParam          1000000              2794 ns/op             713 B/op         14 allocs/op
BenchmarkGoRestful_GithubParam             10000            134707 ns/op            3016 B/op         31 allocs/op
BenchmarkGorillaMux_GithubParam           200000             11270 ns/op             768 B/op          8 allocs/op
BenchmarkHttpRouter_GithubParam          5000000               367 ns/op              96 B/op          1 allocs/op
BenchmarkHttpTreeMux_GithubParam         1000000              1198 ns/op             384 B/op          4 allocs/op
BenchmarkIris_GithubParam               30000000                49.3 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_GithubParam               2000000               872 ns/op             128 B/op          5 allocs/op
BenchmarkLARS_GithubParam               10000000               189 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_GithubParam             1000000              3210 ns/op            1040 B/op          9 allocs/op
BenchmarkMartini_GithubParam              200000             13075 ns/op            1136 B/op         11 allocs/op
BenchmarkPat_GithubParam                  200000              8415 ns/op            2464 B/op         48 allocs/op
BenchmarkPossum_GithubParam              1000000              1893 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_GithubParam            1000000              1271 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_GithubParam                200000              6570 ns/op            1744 B/op         28 allocs/op
BenchmarkRivet_GithubParam               3000000               554 ns/op              96 B/op          1 allocs/op
BenchmarkTango_GithubParam               1000000              2423 ns/op             480 B/op         12 allocs/op
BenchmarkTigerTonic_GithubParam           300000              6103 ns/op            1408 B/op         22 allocs/op
BenchmarkTraffic_GithubParam              100000             20411 ns/op            5992 B/op         52 allocs/op
BenchmarkVulcan_GithubParam              1000000              1848 ns/op              98 B/op          3 allocs/op
BenchmarkAce_GithubAll                     10000            121206 ns/op           13792 B/op        167 allocs/op
BenchmarkBear_GithubAll                    10000            348919 ns/op           86448 B/op        943 allocs/op
BenchmarkBeego_GithubAll                    5000            296816 ns/op           16608 B/op        524 allocs/op
BenchmarkBone_GithubAll                      500           2502143 ns/op          548736 B/op       7241 allocs/op
BenchmarkDenco_GithubAll                   20000             99705 ns/op           20224 B/op        167 allocs/op
BenchmarkEcho_GithubAll                    30000             45469 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GithubAll                     50000             39402 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GithubAll               5000            446025 ns/op          131656 B/op       1686 allocs/op
BenchmarkGoji_GithubAll                     3000            547698 ns/op           56112 B/op        334 allocs/op
BenchmarkGojiv2_GithubAll                   2000            763043 ns/op          118864 B/op       3103 allocs/op
BenchmarkGoJsonRest_GithubAll               5000            538030 ns/op          134371 B/op       2737 allocs/op
BenchmarkGoRestful_GithubAll                 100          14870850 ns/op          837832 B/op       6913 allocs/op
BenchmarkGorillaMux_GithubAll                200           6690383 ns/op          144464 B/op       1588 allocs/op
BenchmarkHttpRouter_GithubAll              20000             65653 ns/op           13792 B/op        167 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            215312 ns/op           65856 B/op        671 allocs/op
BenchmarkIris_GithubAll                   100000             20731 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GithubAll                   10000            167209 ns/op           23304 B/op        843 allocs/op
BenchmarkLARS_GithubAll                    30000             41069 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_GithubAll                  2000            665038 ns/op          201138 B/op       1803 allocs/op
BenchmarkMartini_GithubAll                   300           5433644 ns/op          228213 B/op       2483 allocs/op
BenchmarkPat_GithubAll                       300           4210240 ns/op         1499569 B/op      27435 allocs/op
BenchmarkPossum_GithubAll                  10000            255114 ns/op           84448 B/op        609 allocs/op
BenchmarkR2router_GithubAll                10000            237113 ns/op           77328 B/op        979 allocs/op
BenchmarkRevel_GithubAll                    2000           1150565 ns/op          337424 B/op       5512 allocs/op
BenchmarkRivet_GithubAll                   20000             96555 ns/op           16272 B/op        167 allocs/op
BenchmarkTango_GithubAll                    5000            417423 ns/op           87075 B/op       2267 allocs/op
BenchmarkTigerTonic_GithubAll               2000            994556 ns/op          233680 B/op       5035 allocs/op
BenchmarkTraffic_GithubAll                   200           7770444 ns/op         2659331 B/op      21848 allocs/op
BenchmarkVulcan_GithubAll                   5000            292216 ns/op           19894 B/op        609 allocs/op
BenchmarkAce_GPlusStatic                10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkBear_GPlusStatic                3000000               431 ns/op             104 B/op          3 allocs/op
BenchmarkBeego_GPlusStatic               2000000               943 ns/op              32 B/op          4 allocs/op
BenchmarkBone_GPlusStatic               10000000               204 ns/op              32 B/op          1 allocs/op
BenchmarkDenco_GPlusStatic              50000000                31.5 ns/op             0 B/op          0 allocs/op
BenchmarkEcho_GPlusStatic               20000000                90.1 ns/op             0 B/op          0 allocs/op
BenchmarkGin_GPlusStatic                20000000                88.8 ns/op             0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlusStatic          2000000               897 ns/op             280 B/op          5 allocs/op
BenchmarkGoji_GPlusStatic               10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkGojiv2_GPlusStatic              2000000               626 ns/op             160 B/op          4 allocs/op
BenchmarkGoJsonRest_GPlusStatic          1000000              1163 ns/op             329 B/op         11 allocs/op
BenchmarkGoRestful_GPlusStatic            200000              7915 ns/op            2360 B/op         26 allocs/op
BenchmarkGorillaMux_GPlusStatic          1000000              1892 ns/op             448 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000                33.7 ns/op             0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        30000000                42.3 ns/op             0 B/op          0 allocs/op
BenchmarkIris_GPlusStatic               30000000                51.9 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic              30000000                58.5 ns/op             0 B/op          0 allocs/op
BenchmarkLARS_GPlusStatic               20000000                88.5 ns/op             0 B/op          0 allocs/op
BenchmarkMacaron_GPlusStatic             1000000              2237 ns/op             752 B/op          8 allocs/op
BenchmarkMartini_GPlusStatic              500000              3892 ns/op             784 B/op         10 allocs/op
BenchmarkPat_GPlusStatic                 5000000               336 ns/op              96 B/op          2 allocs/op
BenchmarkPossum_GPlusStatic              1000000              1188 ns/op             416 B/op          3 allocs/op
BenchmarkR2router_GPlusStatic            3000000               560 ns/op             144 B/op          4 allocs/op
BenchmarkRevel_GPlusStatic                500000              4448 ns/op            1232 B/op         23 allocs/op
BenchmarkRivet_GPlusStatic              20000000                80.3 ns/op             0 B/op          0 allocs/op
BenchmarkTango_GPlusStatic               1000000              1153 ns/op             208 B/op          9 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               223 ns/op              32 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic             1000000              2782 ns/op            1192 B/op         15 allocs/op
BenchmarkVulcan_GPlusStatic              2000000               752 ns/op              98 B/op          3 allocs/op
BenchmarkAce_GPlusParam                  3000000               423 ns/op              64 B/op          1 allocs/op
BenchmarkBear_GPlusParam                 1000000              1094 ns/op             480 B/op          5 allocs/op
BenchmarkBeego_GPlusParam                1000000              1317 ns/op             128 B/op          4 allocs/op
BenchmarkBone_GPlusParam                 1000000              1063 ns/op             384 B/op          3 allocs/op
BenchmarkDenco_GPlusParam                5000000               314 ns/op              64 B/op          1 allocs/op
BenchmarkEcho_GPlusParam                20000000               120 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GPlusParam                 20000000               116 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlusParam           1000000              1728 ns/op             648 B/op          8 allocs/op
BenchmarkGoji_GPlusParam                 2000000               859 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_GPlusParam               2000000               935 ns/op             176 B/op          5 allocs/op
BenchmarkGoJsonRest_GPlusParam           1000000              1961 ns/op             649 B/op         13 allocs/op
BenchmarkGoRestful_GPlusParam             100000             15490 ns/op            2760 B/op         29 allocs/op
BenchmarkGorillaMux_GPlusParam            500000              3954 ns/op             752 B/op          8 allocs/op
BenchmarkHttpRouter_GPlusParam          10000000               236 ns/op              64 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2000000               821 ns/op             352 B/op          3 allocs/op
BenchmarkIris_GPlusParam                30000000                53.7 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_GPlusParam                3000000               433 ns/op              56 B/op          3 allocs/op
BenchmarkLARS_GPlusParam                10000000               121 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_GPlusParam              1000000              3153 ns/op            1040 B/op          9 allocs/op
BenchmarkMartini_GPlusParam               300000              5510 ns/op            1104 B/op         11 allocs/op
BenchmarkPat_GPlusParam                  1000000              2166 ns/op             688 B/op         12 allocs/op
BenchmarkPossum_GPlusParam               1000000              1797 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_GPlusParam             1000000              1107 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_GPlusParam                 500000              5534 ns/op            1664 B/op         26 allocs/op
BenchmarkRivet_GPlusParam                5000000               304 ns/op              48 B/op          1 allocs/op
BenchmarkTango_GPlusParam                1000000              1541 ns/op             272 B/op          9 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              3604 ns/op            1040 B/op         16 allocs/op
BenchmarkTraffic_GPlusParam               500000              5780 ns/op            1976 B/op         21 allocs/op
BenchmarkVulcan_GPlusParam               1000000              1066 ns/op              98 B/op          3 allocs/op
BenchmarkAce_GPlus2Params                3000000               465 ns/op              64 B/op          1 allocs/op
BenchmarkBear_GPlus2Params               1000000              1410 ns/op             496 B/op          5 allocs/op
BenchmarkBeego_GPlus2Params              1000000              1753 ns/op             256 B/op          4 allocs/op
BenchmarkBone_GPlus2Params               1000000              2846 ns/op             736 B/op          7 allocs/op
BenchmarkDenco_GPlus2Params              5000000               407 ns/op              64 B/op          1 allocs/op
BenchmarkEcho_GPlus2Params              10000000               166 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GPlus2Params               10000000               148 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlus2Params         1000000              2112 ns/op             712 B/op          9 allocs/op
BenchmarkGoji_GPlus2Params               1000000              1202 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_GPlus2Params             1000000              1796 ns/op             256 B/op          8 allocs/op
BenchmarkGoJsonRest_GPlus2Params         1000000              2628 ns/op             713 B/op         14 allocs/op
BenchmarkGoRestful_GPlus2Params           100000             18971 ns/op            2920 B/op         31 allocs/op
BenchmarkGorillaMux_GPlus2Params          200000              8545 ns/op             768 B/op          8 allocs/op
BenchmarkHttpRouter_GPlus2Params         5000000               269 ns/op              64 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        1000000              1108 ns/op             384 B/op          4 allocs/op
BenchmarkIris_GPlus2Params              30000000                53.3 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_GPlus2Params              2000000               818 ns/op             128 B/op          5 allocs/op
BenchmarkLARS_GPlus2Params              10000000               149 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_GPlus2Params            1000000              2937 ns/op            1040 B/op          9 allocs/op
BenchmarkMartini_GPlus2Params             200000             13015 ns/op            1232 B/op         15 allocs/op
BenchmarkPat_GPlus2Params                 200000              6500 ns/op            2256 B/op         34 allocs/op
BenchmarkPossum_GPlus2Params             1000000              1681 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_GPlus2Params           1000000              1162 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_GPlus2Params               500000              5910 ns/op            1760 B/op         28 allocs/op
BenchmarkRivet_GPlus2Params              5000000               411 ns/op              96 B/op          1 allocs/op
BenchmarkTango_GPlus2Params              1000000              1874 ns/op             448 B/op         11 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              5584 ns/op            1456 B/op         22 allocs/op
BenchmarkTraffic_GPlus2Params             200000             12500 ns/op            3272 B/op         31 allocs/op
BenchmarkVulcan_GPlus2Params             1000000              1486 ns/op              98 B/op          3 allocs/op
BenchmarkAce_GPlusAll                     300000              5583 ns/op             640 B/op         11 allocs/op
BenchmarkBear_GPlusAll                    100000             15530 ns/op            5488 B/op         61 allocs/op
BenchmarkBeego_GPlusAll                   100000             17470 ns/op            1440 B/op         44 allocs/op
BenchmarkBone_GPlusAll                    100000             19361 ns/op            4912 B/op         61 allocs/op
BenchmarkDenco_GPlusAll                   500000              4558 ns/op             672 B/op         11 allocs/op
BenchmarkEcho_GPlusAll                   1000000              2290 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GPlusAll                    1000000              1951 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlusAll              100000             23171 ns/op            8040 B/op        103 allocs/op
BenchmarkGoji_GPlusAll                    200000             12025 ns/op            3696 B/op         22 allocs/op
BenchmarkGojiv2_GPlusAll                  100000             15520 ns/op            2640 B/op         76 allocs/op
BenchmarkGoJsonRest_GPlusAll               50000             28341 ns/op            8117 B/op        170 allocs/op
BenchmarkGoRestful_GPlusAll                10000            168209 ns/op           38664 B/op        389 allocs/op
BenchmarkGorillaMux_GPlusAll               20000             63253 ns/op            9248 B/op        102 allocs/op
BenchmarkHttpRouter_GPlusAll              500000              3066 ns/op             640 B/op         11 allocs/op
BenchmarkHttpTreeMux_GPlusAll             200000             10690 ns/op            4032 B/op         38 allocs/op
BenchmarkIris_GPlusAll                   1000000              1155 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GPlusAll                   200000              7400 ns/op             976 B/op         43 allocs/op
BenchmarkLARS_GPlusAll                   1000000              2024 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_GPlusAll                  50000             36762 ns/op           12944 B/op        115 allocs/op
BenchmarkMartini_GPlusAll                  20000             88555 ns/op           14448 B/op        165 allocs/op
BenchmarkPat_GPlusAll                      30000             48602 ns/op           16576 B/op        298 allocs/op
BenchmarkPossum_GPlusAll                  100000             15940 ns/op            5408 B/op         39 allocs/op
BenchmarkR2router_GPlusAll                100000             13560 ns/op            5040 B/op         63 allocs/op
BenchmarkRevel_GPlusAll                    20000             71154 ns/op           21136 B/op        342 allocs/op
BenchmarkRivet_GPlusAll                   300000              4613 ns/op             768 B/op         11 allocs/op
BenchmarkTango_GPlusAll                   100000             21081 ns/op            4304 B/op        129 allocs/op
BenchmarkTigerTonic_GPlusAll               30000             55203 ns/op           14256 B/op        272 allocs/op
BenchmarkTraffic_GPlusAll                  10000            108306 ns/op           37360 B/op        392 allocs/op
BenchmarkVulcan_GPlusAll                  100000             15290 ns/op            1274 B/op         39 allocs/op
BenchmarkAce_ParseStatic                10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkBear_ParseStatic                3000000               497 ns/op             120 B/op          3 allocs/op
BenchmarkBeego_ParseStatic               1000000              1010 ns/op              32 B/op          4 allocs/op
BenchmarkBone_ParseStatic                2000000               685 ns/op             144 B/op          3 allocs/op
BenchmarkDenco_ParseStatic              30000000                47.2 ns/op             0 B/op          0 allocs/op
BenchmarkEcho_ParseStatic               20000000                92.5 ns/op             0 B/op          0 allocs/op
BenchmarkGin_ParseStatic                20000000                91.5 ns/op             0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParseStatic          2000000               944 ns/op             296 B/op          5 allocs/op
BenchmarkGoji_ParseStatic                5000000               251 ns/op               0 B/op          0 allocs/op
BenchmarkGojiv2_ParseStatic              2000000               638 ns/op             160 B/op          4 allocs/op
BenchmarkGoJsonRest_ParseStatic          1000000              1218 ns/op             329 B/op         11 allocs/op
BenchmarkGoRestful_ParseStatic            200000             12490 ns/op            3656 B/op         30 allocs/op
BenchmarkGorillaMux_ParseStatic           500000              2952 ns/op             448 B/op          7 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000                36.8 ns/op             0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        20000000                70.8 ns/op             0 B/op          0 allocs/op
BenchmarkIris_ParseStatic               30000000                49.2 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic              20000000                60.0 ns/op             0 B/op          0 allocs/op
BenchmarkLARS_ParseStatic               20000000                92.9 ns/op             0 B/op          0 allocs/op
BenchmarkMacaron_ParseStatic             1000000              2265 ns/op             752 B/op          8 allocs/op
BenchmarkMartini_ParseStatic              500000              4386 ns/op             784 B/op         10 allocs/op
BenchmarkPat_ParseStatic                 2000000               820 ns/op             240 B/op          5 allocs/op
BenchmarkPossum_ParseStatic              1000000              1178 ns/op             416 B/op          3 allocs/op
BenchmarkR2router_ParseStatic            2000000               608 ns/op             144 B/op          4 allocs/op
BenchmarkRevel_ParseStatic                500000              4662 ns/op            1248 B/op         23 allocs/op
BenchmarkRivet_ParseStatic              20000000                85.7 ns/op             0 B/op          0 allocs/op
BenchmarkTango_ParseStatic               1000000              1284 ns/op             256 B/op          9 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               306 ns/op              48 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              500000              4098 ns/op            1816 B/op         20 allocs/op
BenchmarkVulcan_ParseStatic              2000000               832 ns/op              98 B/op          3 allocs/op
BenchmarkAce_ParseParam                  5000000               397 ns/op              64 B/op          1 allocs/op
BenchmarkBear_ParseParam                 1000000              1084 ns/op             467 B/op          5 allocs/op
BenchmarkBeego_ParseParam                1000000              1171 ns/op              64 B/op          4 allocs/op
BenchmarkBone_ParseParam                 1000000              1320 ns/op             464 B/op          4 allocs/op
BenchmarkDenco_ParseParam                5000000               316 ns/op              64 B/op          1 allocs/op
BenchmarkEcho_ParseParam                20000000               106 ns/op               0 B/op          0 allocs/op
BenchmarkGin_ParseParam                 20000000                97.3 ns/op             0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParseParam           1000000              1634 ns/op             664 B/op          8 allocs/op
BenchmarkGoji_ParseParam                 2000000               928 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_ParseParam               2000000               986 ns/op             208 B/op          6 allocs/op
BenchmarkGoJsonRest_ParseParam           1000000              1816 ns/op             649 B/op         13 allocs/op
BenchmarkGoRestful_ParseParam             100000             13360 ns/op            4024 B/op         31 allocs/op
BenchmarkGorillaMux_ParseParam            500000              3366 ns/op             752 B/op          8 allocs/op
BenchmarkHttpRouter_ParseParam          10000000               225 ns/op              64 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParseParam          2000000               706 ns/op             352 B/op          3 allocs/op
BenchmarkIris_ParseParam                30000000                50.8 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_ParseParam                5000000               406 ns/op              56 B/op          3 allocs/op
BenchmarkLARS_ParseParam                20000000               101 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_ParseParam              1000000              2744 ns/op            1040 B/op          9 allocs/op
BenchmarkMartini_ParseParam               300000              5056 ns/op            1104 B/op         11 allocs/op
BenchmarkPat_ParseParam                  1000000              2871 ns/op            1120 B/op         17 allocs/op
BenchmarkPossum_ParseParam               1000000              1701 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_ParseParam             1000000              1033 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_ParseParam                 500000              5266 ns/op            1664 B/op         26 allocs/op
BenchmarkRivet_ParseParam                5000000               264 ns/op              48 B/op          1 allocs/op
BenchmarkTango_ParseParam                1000000              1387 ns/op             288 B/op          9 allocs/op
BenchmarkTigerTonic_ParseParam           1000000              3236 ns/op             992 B/op         16 allocs/op
BenchmarkTraffic_ParseParam               500000              5372 ns/op            2248 B/op         23 allocs/op
BenchmarkVulcan_ParseParam               2000000               949 ns/op              98 B/op          3 allocs/op
BenchmarkAce_Parse2Params                5000000               421 ns/op              64 B/op          1 allocs/op
BenchmarkBear_Parse2Params               1000000              1250 ns/op             496 B/op          5 allocs/op
BenchmarkBeego_Parse2Params              1000000              1436 ns/op             128 B/op          4 allocs/op
BenchmarkBone_Parse2Params               1000000              1192 ns/op             416 B/op          3 allocs/op
BenchmarkDenco_Parse2Params              5000000               373 ns/op              64 B/op          1 allocs/op
BenchmarkEcho_Parse2Params              10000000               131 ns/op               0 B/op          0 allocs/op
BenchmarkGin_Parse2Params               20000000               118 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_Parse2Params         1000000              2030 ns/op             712 B/op          9 allocs/op
BenchmarkGoji_Parse2Params               2000000               946 ns/op             336 B/op          2 allocs/op
BenchmarkGojiv2_Parse2Params             2000000               954 ns/op             192 B/op          5 allocs/op
BenchmarkGoJsonRest_Parse2Params         1000000              2251 ns/op             713 B/op         14 allocs/op
BenchmarkGoRestful_Parse2Params           100000             22051 ns/op            6856 B/op         39 allocs/op
BenchmarkGorillaMux_Parse2Params          500000              3574 ns/op             768 B/op          8 allocs/op
BenchmarkHttpRouter_Parse2Params        10000000               238 ns/op              64 B/op          1 allocs/op
BenchmarkHttpTreeMux_Parse2Params        1000000              1007 ns/op             384 B/op          4 allocs/op
BenchmarkIris_Parse2Params              30000000                52.0 ns/op             0 B/op          0 allocs/op
BenchmarkKocha_Parse2Params              2000000               747 ns/op             128 B/op          5 allocs/op
BenchmarkLARS_Parse2Params              20000000               119 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_Parse2Params            1000000              2984 ns/op            1040 B/op          9 allocs/op
BenchmarkMartini_Parse2Params             300000              4966 ns/op            1136 B/op         11 allocs/op
BenchmarkPat_Parse2Params                1000000              2879 ns/op             832 B/op         17 allocs/op
BenchmarkPossum_Parse2Params             1000000              1703 ns/op             560 B/op          6 allocs/op
BenchmarkR2router_Parse2Params           1000000              1187 ns/op             432 B/op          5 allocs/op
BenchmarkRevel_Parse2Params               300000              5790 ns/op            1728 B/op         28 allocs/op
BenchmarkRivet_Parse2Params              5000000               395 ns/op              96 B/op          1 allocs/op
BenchmarkTango_Parse2Params              1000000              1832 ns/op             416 B/op         11 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              5286 ns/op            1376 B/op         22 allocs/op
BenchmarkTraffic_Parse2Params             300000              5793 ns/op            2040 B/op         22 allocs/op
BenchmarkVulcan_Parse2Params             1000000              1059 ns/op              98 B/op          3 allocs/op
BenchmarkAce_ParseAll                     200000              9285 ns/op             640 B/op         16 allocs/op
BenchmarkBear_ParseAll                     50000             27281 ns/op            8928 B/op        110 allocs/op
BenchmarkBeego_ParseAll                   100000             23191 ns/op             800 B/op         36 allocs/op
BenchmarkBone_ParseAll                     50000             29121 ns/op            8048 B/op         90 allocs/op
BenchmarkDenco_ParseAll                   300000              6850 ns/op             928 B/op         16 allocs/op
BenchmarkEcho_ParseAll                    300000              4143 ns/op               0 B/op          0 allocs/op
BenchmarkGin_ParseAll                     500000              3680 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParseAll               30000             41269 ns/op           13728 B/op        181 allocs/op
BenchmarkGoji_ParseAll                    100000             19371 ns/op            5376 B/op         32 allocs/op
BenchmarkGojiv2_ParseAll                  100000             22791 ns/op            4496 B/op        121 allocs/op
BenchmarkGoJsonRest_ParseAll               30000             48336 ns/op           13866 B/op        321 allocs/op
BenchmarkGoRestful_ParseAll                 5000            447025 ns/op          125600 B/op        868 allocs/op
BenchmarkGorillaMux_ParseAll               10000            126807 ns/op           16560 B/op        198 allocs/op
BenchmarkHttpRouter_ParseAll              500000              4346 ns/op             640 B/op         16 allocs/op
BenchmarkHttpTreeMux_ParseAll             100000             15080 ns/op            5728 B/op         51 allocs/op
BenchmarkIris_ParseAll                   1000000              2233 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParseAll                   200000             10095 ns/op            1112 B/op         54 allocs/op
BenchmarkLARS_ParseAll                    500000              3672 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_ParseAll                  20000             72204 ns/op           24160 B/op        224 allocs/op
BenchmarkMartini_ParseAll                  10000            134707 ns/op           25600 B/op        276 allocs/op
BenchmarkPat_ParseAll                      30000             55703 ns/op           17264 B/op        343 allocs/op
BenchmarkPossum_ParseAll                   50000             31801 ns/op           10816 B/op         78 allocs/op
BenchmarkR2router_ParseAll                100000             25061 ns/op            8352 B/op        120 allocs/op
BenchmarkRevel_ParseAll                    10000            132607 ns/op           39424 B/op        652 allocs/op
BenchmarkRivet_ParseAll                   300000              6797 ns/op             912 B/op         16 allocs/op
BenchmarkTango_ParseAll                    50000             38662 ns/op            7664 B/op        240 allocs/op
BenchmarkTigerTonic_ParseAll               20000             72554 ns/op           19424 B/op        360 allocs/op
BenchmarkTraffic_ParseAll                  10000            147408 ns/op           57776 B/op        642 allocs/op
BenchmarkVulcan_ParseAll                   50000             28641 ns/op            2548 B/op         78 allocs/op
BenchmarkAce_StaticAll                     30000             45069 ns/op               0 B/op          0 allocs/op
BenchmarkHttpServeMux_StaticAll             2000            749542 ns/op              96 B/op          8 allocs/op
BenchmarkBeego_StaticAll                   10000            229913 ns/op           15744 B/op        628 allocs/op
BenchmarkBear_StaticAll                    10000            102805 ns/op           20336 B/op        461 allocs/op
BenchmarkBone_StaticAll                    20000             74654 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_StaticAll                  100000             12750 ns/op               0 B/op          0 allocs/op
BenchmarkEcho_StaticAll                    50000             29521 ns/op               0 B/op          0 allocs/op
BenchmarkGin_StaticAll                     50000             27861 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_StaticAll              10000            165709 ns/op           46440 B/op        785 allocs/op
BenchmarkGoji_StaticAll                    20000             60603 ns/op               0 B/op          0 allocs/op
BenchmarkGojiv2_StaticAll                  10000            155908 ns/op           25120 B/op        628 allocs/op
BenchmarkGoJsonRest_StaticAll              10000            246314 ns/op           51653 B/op       1727 allocs/op
BenchmarkGoRestful_StaticAll                 300           5300303 ns/op          392312 B/op       4694 allocs/op
BenchmarkGorillaMux_StaticAll               1000           2013115 ns/op           70432 B/op       1107 allocs/op
BenchmarkHttpRouter_StaticAll             100000             16890 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll            100000             16830 ns/op               0 B/op          0 allocs/op
BenchmarkIris_StaticAll                   100000             15580 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_StaticAll                  100000             23181 ns/op               0 B/op          0 allocs/op
BenchmarkLARS_StaticAll                    50000             27121 ns/op               0 B/op          0 allocs/op
BenchmarkMacaron_StaticAll                  5000            374421 ns/op          118065 B/op       1256 allocs/op
BenchmarkMartini_StaticAll                   500           2550146 ns/op          132818 B/op       2178 allocs/op
BenchmarkPat_StaticAll                      1000           1524087 ns/op          533904 B/op      11123 allocs/op
BenchmarkPossum_StaticAll                  10000            183510 ns/op           65312 B/op        471 allocs/op
BenchmarkR2router_StaticAll                10000            114406 ns/op           22608 B/op        628 allocs/op
BenchmarkRevel_StaticAll                    3000            698039 ns/op          198240 B/op       3611 allocs/op
BenchmarkRivet_StaticAll                   50000             32081 ns/op               0 B/op          0 allocs/op
BenchmarkTango_StaticAll                   10000            268015 ns/op           40481 B/op       1413 allocs/op
BenchmarkTigerTonic_StaticAll              20000             63353 ns/op            7504 B/op        157 allocs/op
BenchmarkTraffic_StaticAll                  1000           1976113 ns/op          729736 B/op      14287 allocs/op
BenchmarkVulcan_StaticAll                  10000            198611 ns/op           15386 B/op        471 allocs/op
ok      github.com/kataras/go-http-routing-benchmark    972.432s

D:\Projects\Golang\mygopath\src\github.com\kataras\go-http-routing-benchmark>
```