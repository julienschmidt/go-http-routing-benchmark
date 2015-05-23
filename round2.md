#Go Webframework Benchmarks #2

May 23, 2015.

After one year of development we have released [Gin v1.0rc1](https://github.com/gin-gonic/gin/releases/tag/v1.0rc1) and we wanted to run the benchmark tests one more time.

**Machine:** intel i7 ivy bridge quad-core. 8GB RAM.  
**Notes:**
	1. echo and zeus where removed in some tests because they continuously crashed.

##Test Frameworks

Ace  
Bear  
Beego  
Bone  
Denco  
Gin  
GocraftWeb  
Goji  
GoJsonRest  
GoRestful  
GorillaMux  
HttpRouter  
HttpTreeMux  
Kocha  
Macaron  
Martini  
Pat  
Possum  
R2router  
Revel  
Rivet  
Tango  
TigerTonic  
Traffic  
Vulcan  
Zeus  

Fastest: **BenchmarkGin_Param  10000000           124 ns/op           0 B/op          0 allocs/op**  
Slowest: **BenchmarkGoRestful_Param      200000          8397 ns/op        2496 B/op         31 allocs/op**

```
BenchmarkAce_Param	 5000000	       346 ns/op	      32 B/op	       1 allocs/op
BenchmarkBear_Param	 1000000	      1003 ns/op	     424 B/op	       5 allocs/op
BenchmarkBeego_Param	 1000000	      2240 ns/op	     720 B/op	      10 allocs/op
BenchmarkBone_Param	 2000000	       932 ns/op	     384 B/op	       3 allocs/op
BenchmarkDenco_Param	10000000	       219 ns/op	      32 B/op	       1 allocs/op
BenchmarkGin_Param	10000000	       124 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param	 1000000	      1627 ns/op	     656 B/op	       9 allocs/op
BenchmarkGoji_Param	 2000000	       848 ns/op	     336 B/op	       2 allocs/op
BenchmarkGoJsonRest_Param	 1000000	      1814 ns/op	     657 B/op	      14 allocs/op
BenchmarkGoRestful_Param	  200000	      8397 ns/op	    2496 B/op	      31 allocs/op
BenchmarkGorillaMux_Param	  500000	      3423 ns/op	     784 B/op	       9 allocs/op
BenchmarkHttpRouter_Param	10000000	       142 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param	 2000000	       644 ns/op	     336 B/op	       2 allocs/op
BenchmarkKocha_Param	 5000000	       390 ns/op	      56 B/op	       3 allocs/op
BenchmarkMacaron_Param	 1000000	      3158 ns/op	    1104 B/op	      11 allocs/op
BenchmarkMartini_Param	  300000	      6609 ns/op	    1152 B/op	      12 allocs/op
BenchmarkPat_Param	 1000000	      1868 ns/op	     656 B/op	      14 allocs/op
BenchmarkPossum_Param	 1000000	      1853 ns/op	     624 B/op	       7 allocs/op
BenchmarkR2router_Param	 1000000	      1035 ns/op	     432 B/op	       6 allocs/op
BenchmarkRevel_Param	  200000	      6654 ns/op	    1672 B/op	      28 allocs/op
BenchmarkRivet_Param	 1000000	      1078 ns/op	     464 B/op	       5 allocs/op
BenchmarkTango_Param	 1000000	      1368 ns/op	     256 B/op	      10 allocs/op
BenchmarkTigerTonic_Param	  500000	      3135 ns/op	     992 B/op	      19 allocs/op
BenchmarkTraffic_Param	  300000	      5450 ns/op	    1984 B/op	      23 allocs/op
BenchmarkVulcan_Param	 2000000	       938 ns/op	      98 B/op	       3 allocs/op
BenchmarkZeus_Param	 1000000	      1100 ns/op	     368 B/op	       3 allocs/op
```

Fastest: **BenchmarkGin_Param5	10000000 215 ns/op	       0 B/op	       0 allocs/op**  
Slowest: **BenchmarkMartini_Param5   100000         15237 ns/op        1280 B/op         12 allocs/op**

```
BenchmarkAce_Param5	 2000000	       579 ns/op	     160 B/op	       1 allocs/op
BenchmarkBear_Param5	 1000000	      1599 ns/op	     469 B/op	       5 allocs/op
BenchmarkBeego_Param5	 1000000	      3472 ns/op	     992 B/op	      13 allocs/op
BenchmarkBone_Param5	 1000000	      1558 ns/op	     432 B/op	       3 allocs/op
BenchmarkDenco_Param5	 3000000	       554 ns/op	     160 B/op	       1 allocs/op
BenchmarkGin_Param5	10000000	       215 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param5	 1000000	      2689 ns/op	     928 B/op	      12 allocs/op
BenchmarkGoji_Param5	 1000000	      1194 ns/op	     336 B/op	       2 allocs/op
BenchmarkGoJsonRest_Param5	  500000	      3765 ns/op	    1105 B/op	      17 allocs/op
BenchmarkGoRestful_Param5	  200000	     11263 ns/op	    2672 B/op	      31 allocs/op
BenchmarkGorillaMux_Param5	  300000	      6050 ns/op	     912 B/op	       9 allocs/op
BenchmarkHttpRouter_Param5	 5000000	       397 ns/op	     160 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param5	 1000000	      1088 ns/op	     336 B/op	       2 allocs/op
BenchmarkKocha_Param5	 1000000	      1608 ns/op	     440 B/op	      10 allocs/op
BenchmarkMacaron_Param5	  300000	      4506 ns/op	    1376 B/op	      14 allocs/op
BenchmarkMartini_Param5	  100000	     15237 ns/op	    1280 B/op	      12 allocs/op
BenchmarkPat_Param5	  300000	      4988 ns/op	    1008 B/op	      42 allocs/op
BenchmarkPossum_Param5	 1000000	      2012 ns/op	     624 B/op	       7 allocs/op
BenchmarkR2router_Param5	 1000000	      1531 ns/op	     432 B/op	       6 allocs/op
BenchmarkRevel_Param5	  200000	      7964 ns/op	    2024 B/op	      35 allocs/op
BenchmarkRivet_Param5	 1000000	      1895 ns/op	     528 B/op	       9 allocs/op
BenchmarkTango_Param5	 1000000	      3093 ns/op	     944 B/op	      18 allocs/op
BenchmarkTigerTonic_Param5	  200000	     11992 ns/op	    2519 B/op	      53 allocs/op
BenchmarkTraffic_Param5	  200000	      8537 ns/op	    2280 B/op	      31 allocs/op
BenchmarkVulcan_Param5	 1000000	      1290 ns/op	      98 B/op	       3 allocs/op
BenchmarkZeus_Param5	 1000000	      1537 ns/op	     416 B/op	       3 allocs/op
```

Fastest: **BenchmarkGin_Param20     3000000           522 ns/op           0 B/op          0 allocs/op**  
Slowest: **BenchmarkMartini_Param20       30000         57617 ns/op        3643 B/op         14 allocs/op**

```
BenchmarkAce_Param20	 1000000	      1402 ns/op	     640 B/op	       1 allocs/op
BenchmarkBear_Param20	  300000	      4518 ns/op	    1633 B/op	       5 allocs/op
BenchmarkBeego_Param20	  200000	      9903 ns/op	    3867 B/op	      17 allocs/op
BenchmarkBone_Param20	  300000	      6532 ns/op	    2540 B/op	       5 allocs/op
BenchmarkDenco_Param20	 1000000	      1670 ns/op	     640 B/op	       1 allocs/op
BenchmarkGin_Param20	 3000000	       522 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param20	  200000	      8511 ns/op	    3804 B/op	      16 allocs/op
BenchmarkGoji_Param20	  500000	      3587 ns/op	    1247 B/op	       2 allocs/op
BenchmarkGoJsonRest_Param20	  200000	     11216 ns/op	    4492 B/op	      21 allocs/op
BenchmarkGoRestful_Param20	  100000	     17779 ns/op	    5243 B/op	      33 allocs/op
BenchmarkGorillaMux_Param20	  200000	     11616 ns/op	    3276 B/op	      11 allocs/op
BenchmarkHttpRouter_Param20	 1000000	      1156 ns/op	     640 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param20	  300000	      5690 ns/op	    2187 B/op	       4 allocs/op
BenchmarkKocha_Param20	  300000	      4750 ns/op	    1808 B/op	      27 allocs/op
BenchmarkMacaron_Param20	  200000	      9921 ns/op	    4251 B/op	      18 allocs/op
BenchmarkMartini_Param20	   30000	     57617 ns/op	    3643 B/op	      14 allocs/op
BenchmarkPat_Param20	  100000	     22388 ns/op	    4886 B/op	     151 allocs/op
BenchmarkPossum_Param20	 1000000	      1879 ns/op	     624 B/op	       7 allocs/op
BenchmarkR2router_Param20	  300000	      5827 ns/op	    2283 B/op	       8 allocs/op
BenchmarkRevel_Param20	  100000	     14821 ns/op	    5551 B/op	      54 allocs/op
BenchmarkRivet_Param20	  200000	      7539 ns/op	    2619 B/op	      26 allocs/op
BenchmarkTango_Param20	  100000	     14399 ns/op	    8224 B/op	      48 allocs/op
BenchmarkTigerTonic_Param20	   30000	     41867 ns/op	   10546 B/op	     178 allocs/op
BenchmarkTraffic_Param20	   50000	     25688 ns/op	    7999 B/op	      66 allocs/op
BenchmarkVulcan_Param20	 1000000	      1994 ns/op	      98 B/op	       3 allocs/op
BenchmarkZeus_Param20	  300000	      5882 ns/op	    2507 B/op	       5 allocs/op
```

Fastest: **BenchmarkGin_ParamWrite  5000000           261 ns/op           8 B/op          1 allocs/op**  
Slowest: **BenchmarkMartini_ParamWrite   200000          8313 ns/op        1256 B/op         16 allocs/op**

```
BenchmarkAce_ParamWrite	 3000000	       458 ns/op	      40 B/op	       2 allocs/op
BenchmarkBear_ParamWrite	 1000000	      1272 ns/op	     424 B/op	       5 allocs/op
BenchmarkBeego_ParamWrite	 1000000	      2541 ns/op	     728 B/op	      11 allocs/op
BenchmarkBone_ParamWrite	 1000000	      1124 ns/op	     384 B/op	       3 allocs/op
BenchmarkDenco_ParamWrite	 5000000	       316 ns/op	      32 B/op	       1 allocs/op
BenchmarkGin_ParamWrite	 5000000	       261 ns/op	       8 B/op	       1 allocs/op
BenchmarkGocraftWeb_ParamWrite	 1000000	      2009 ns/op	     664 B/op	      10 allocs/op
BenchmarkGoji_ParamWrite	 1000000	      1017 ns/op	     336 B/op	       2 allocs/op
BenchmarkGoJsonRest_ParamWrite	 1000000	      3440 ns/op	    1136 B/op	      19 allocs/op
BenchmarkGoRestful_ParamWrite	  200000	      9578 ns/op	    2504 B/op	      32 allocs/op
BenchmarkGorillaMux_ParamWrite	  500000	      4123 ns/op	     792 B/op	      10 allocs/op
BenchmarkHttpRouter_ParamWrite	10000000	       244 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_ParamWrite	 2000000	       759 ns/op	     336 B/op	       2 allocs/op
BenchmarkKocha_ParamWrite	 3000000	       551 ns/op	      56 B/op	       3 allocs/op
BenchmarkMacaron_ParamWrite	  500000	      4654 ns/op	    1216 B/op	      16 allocs/op
BenchmarkMartini_ParamWrite	  200000	      8313 ns/op	    1256 B/op	      16 allocs/op
BenchmarkPat_ParamWrite	 1000000	      3561 ns/op	    1088 B/op	      19 allocs/op
BenchmarkPossum_ParamWrite	 1000000	      2037 ns/op	     624 B/op	       7 allocs/op
BenchmarkR2router_ParamWrite	 1000000	      1200 ns/op	     432 B/op	       6 allocs/op
BenchmarkRevel_ParamWrite	  200000	      7716 ns/op	    2128 B/op	      33 allocs/op
BenchmarkRivet_ParamWrite	 1000000	      1177 ns/op	     472 B/op	       6 allocs/op
BenchmarkTango_ParamWrite	 2000000	       751 ns/op	     136 B/op	       5 allocs/op
BenchmarkTigerTonic_ParamWrite	  300000	      5267 ns/op	    1440 B/op	      25 allocs/op
BenchmarkTraffic_ParamWrite	  200000	      6779 ns/op	    2400 B/op	      27 allocs/op
BenchmarkVulcan_ParamWrite	 2000000	       855 ns/op	      98 B/op	       3 allocs/op
BenchmarkZeus_ParamWrite	 1000000	      1105 ns/op	     368 B/op	       3 allocs/op
```

Fastest: **BenchmarkHttpRouter_GithubStatic    20000000            69.5 ns/op         0 B/op          0 allocs/op**  
Slowest: **BenchmarkGoRestful_GithubStatic    30000         48449 ns/op        3520 B/op         36 allocs/op**

```
BenchmarkAce_GithubStatic	 5000000	       251 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GithubStatic	 3000000	       509 ns/op	      88 B/op	       3 allocs/op
BenchmarkBeego_GithubStatic	 1000000	      1426 ns/op	     368 B/op	       7 allocs/op
BenchmarkBone_GithubStatic	  200000	     10622 ns/op	    2880 B/op	      60 allocs/op
BenchmarkDenco_GithubStatic	20000000	        72.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubStatic	 1000000	      1181 ns/op	     464 B/op	       5 allocs/op
BenchmarkGin_GithubStatic	10000000	       147 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubStatic	 1000000	      1006 ns/op	     304 B/op	       6 allocs/op
BenchmarkGoji_GithubStatic	 5000000	       364 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoRestful_GithubStatic	   30000	     48449 ns/op	    3520 B/op	      36 allocs/op
BenchmarkGoJsonRest_GithubStatic	 1000000	      1353 ns/op	     337 B/op	      12 allocs/op
BenchmarkGorillaMux_GithubStatic	  100000	     21445 ns/op	     464 B/op	       8 allocs/op
BenchmarkHttpRouter_GithubStatic	20000000	        69.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_GithubStatic	20000000	        77.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GithubStatic	20000000	       100.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GithubStatic	 1000000	      2374 ns/op	     752 B/op	       8 allocs/op
BenchmarkMartini_GithubStatic	  100000	     19826 ns/op	     832 B/op	      11 allocs/op
BenchmarkPat_GithubStatic	  200000	     11476 ns/op	    3648 B/op	      76 allocs/op
BenchmarkPossum_GithubStatic	 1000000	      1459 ns/op	     480 B/op	       4 allocs/op
BenchmarkR2router_GithubStatic	 2000000	       665 ns/op	     144 B/op	       5 allocs/op
BenchmarkRevel_GithubStatic	  300000	      5813 ns/op	    1288 B/op	      25 allocs/op
BenchmarkRivet_GithubStatic	 3000000	       430 ns/op	     112 B/op	       2 allocs/op
BenchmarkTango_GithubStatic	 1000000	      1610 ns/op	     256 B/op	      10 allocs/op
BenchmarkTigerTonic_GithubStatic	 5000000	       312 ns/op	      48 B/op	       1 allocs/op
BenchmarkTraffic_GithubStatic	   30000	     40732 ns/op	   18920 B/op	     149 allocs/op
BenchmarkVulcan_GithubStatic	 1000000	      1236 ns/op	      98 B/op	       3 allocs/op
BenchmarkZeus_GithubStatic	 1000000	      2315 ns/op	     512 B/op	      11 allocs/op
```

Fastest: **BenchmarkGin_GithubParam     5000000           235 ns/op           0 B/op          0 allocs/op**  
Slowest: **BenchmarkGoRestful_GithubParam     10000        166324 ns/op        2816 B/op         35 allocs/op**

```
BenchmarkAce_GithubParam	 3000000	       550 ns/op	      96 B/op	       1 allocs/op
BenchmarkBear_GithubParam	 1000000	      1478 ns/op	     464 B/op	       5 allocs/op
BenchmarkBeego_GithubParam	 1000000	      2777 ns/op	     784 B/op	      11 allocs/op
BenchmarkBone_GithubParam	  300000	      5795 ns/op	    1456 B/op	      16 allocs/op
BenchmarkDenco_GithubParam	 3000000	       504 ns/op	     128 B/op	       1 allocs/op
BenchmarkEcho_GithubParam	 1000000	      1242 ns/op	     464 B/op	       5 allocs/op
BenchmarkGin_GithubParam	 5000000	       235 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubParam	 1000000	      2090 ns/op	     720 B/op	      10 allocs/op
BenchmarkGoji_GithubParam	 1000000	      1383 ns/op	     336 B/op	       2 allocs/op
BenchmarkGoJsonRest_GithubParam	 1000000	      2572 ns/op	     721 B/op	      15 allocs/op
BenchmarkGoRestful_GithubParam	   10000	    166324 ns/op	    2816 B/op	      35 allocs/op
BenchmarkGorillaMux_GithubParam	  100000	     12676 ns/op	     816 B/op	       9 allocs/op
BenchmarkHttpRouter_GithubParam	 5000000	       314 ns/op	      96 B/op	       1 allocs/op
BenchmarkHttpTreeMux_GithubParam	 2000000	       848 ns/op	     336 B/op	       2 allocs/op
BenchmarkKocha_GithubParam	 2000000	       819 ns/op	     128 B/op	       5 allocs/op
BenchmarkMacaron_GithubParam	  500000	      3739 ns/op	    1168 B/op	      12 allocs/op
BenchmarkMartini_GithubParam	   50000	     25915 ns/op	    1184 B/op	      12 allocs/op
BenchmarkPat_GithubParam	  200000	      7792 ns/op	    2480 B/op	      56 allocs/op
BenchmarkPossum_GithubParam	 1000000	      1883 ns/op	     624 B/op	       7 allocs/op
BenchmarkR2router_GithubParam	 1000000	      1259 ns/op	     432 B/op	       6 allocs/op
BenchmarkRevel_GithubParam	  200000	      6917 ns/op	    1784 B/op	      30 allocs/op
BenchmarkRivet_GithubParam	 1000000	      1431 ns/op	     480 B/op	       6 allocs/op
BenchmarkTango_GithubParam	 1000000	      2112 ns/op	     480 B/op	      13 allocs/op
BenchmarkTigerTonic_GithubParam	  300000	      5433 ns/op	    1440 B/op	      28 allocs/op
BenchmarkTraffic_GithubParam	  100000	     18452 ns/op	    6024 B/op	      55 allocs/op
BenchmarkVulcan_GithubParam	 1000000	      1987 ns/op	      98 B/op	       3 allocs/op
BenchmarkZeus_GithubParam	  500000	      3713 ns/op	    1312 B/op	      12 allocs/op
```

Fastest: **BenchmarkGin_GithubAll     30000         48336 ns/op           0 B/op          0 allocs/op**  
Slowest: **BenchmarkGoRestful_GithubAll         100      16682676 ns/op      797236 B/op       7725 allocs/op**

```
BenchmarkAce_GithubAll	   20000	     97572 ns/op	   13792 B/op	     167 allocs/op
BenchmarkBear_GithubAll	   10000	    265268 ns/op	   79952 B/op	     943 allocs/op
BenchmarkBeego_GithubAll	    3000	    511191 ns/op	  146272 B/op	    2092 allocs/op
BenchmarkBone_GithubAll	    1000	   2268579 ns/op	  648016 B/op	    8119 allocs/op
BenchmarkDenco_GithubAll	   20000	     87831 ns/op	   20224 B/op	     167 allocs/op
BenchmarkEcho_GithubAll	   10000	    242361 ns/op	   94196 B/op	    1015 allocs/op
BenchmarkGin_GithubAll	   30000	     48336 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubAll	    5000	    399830 ns/op	  133280 B/op	    1889 allocs/op
BenchmarkGoji_GithubAll	    3000	    614849 ns/op	   56113 B/op	     334 allocs/op
BenchmarkGoJsonRest_GithubAll	    5000	    488553 ns/op	  135995 B/op	    2940 allocs/op
BenchmarkGoRestful_GithubAll	     100	  16682676 ns/op	  797236 B/op	    7725 allocs/op
BenchmarkGorillaMux_GithubAll	     200	   7328008 ns/op	  153137 B/op	    1791 allocs/op
BenchmarkHttpRouter_GithubAll	   30000	     55733 ns/op	   13792 B/op	     167 allocs/op
BenchmarkHttpTreeMux_GithubAll	   10000	    155918 ns/op	   56112 B/op	     334 allocs/op
BenchmarkKocha_GithubAll	   10000	    152172 ns/op	   23304 B/op	     843 allocs/op
BenchmarkMacaron_GithubAll	    2000	    751826 ns/op	  224960 B/op	    2315 allocs/op
BenchmarkMartini_GithubAll	     100	  11173691 ns/op	  237952 B/op	    2686 allocs/op
BenchmarkPat_GithubAll	     300	   4345614 ns/op	 1504101 B/op	   32222 allocs/op
BenchmarkPossum_GithubAll	   10000	    271678 ns/op	   97440 B/op	     812 allocs/op
BenchmarkR2router_GithubAll	   10000	    262801 ns/op	   77328 B/op	    1182 allocs/op
BenchmarkRevel_GithubAll	    1000	   1515212 ns/op	  345553 B/op	    5918 allocs/op
BenchmarkRivet_GithubAll	   10000	    256614 ns/op	   84272 B/op	    1079 allocs/op
BenchmarkTango_GithubAll	    5000	    418542 ns/op	   87078 B/op	    2470 allocs/op
BenchmarkTigerTonic_GithubAll	    2000	   1053725 ns/op	  241088 B/op	    6052 allocs/op
BenchmarkTraffic_GithubAll	     200	   8064773 ns/op	 2664762 B/op	   22390 allocs/op
BenchmarkVulcan_GithubAll	    5000	    312481 ns/op	   19894 B/op	     609 allocs/op
BenchmarkZeus_GithubAll	    2000	    832438 ns/op	  300688 B/op	    2648 allocs/op
```
