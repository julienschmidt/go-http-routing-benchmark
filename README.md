go-http-routing-benchmark
=========================

This benchmark suite aims to compare the performance of available HTTP request routers for Go by implementing the routing of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.


Included Routers:

 * [Gocraft Web](https://github.com/gocraft/web)
 * [Goji](https://github.com/zenazn/goji/)
 * [Go-Json-Rest](https://github.com/ant0ine/go-json-rest/)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [net/http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [httptreemux](https://github.com/dimfeld/httptreemux)
 * [Martini](https://github.com/codegangsta/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)

## Results

Benchmark System:
 * Intel Core i7-4771 (4x 3.50GHz + Turbo Boost, Hyper-Threaded)
 * 12GB DDR3 1600MHz RAM + 128GB SSD
 * go1.2.2 darwin/amd64
 * Mac OS X 10.9.3
 
```
#GithubAPI Routes: 203
#GPlusAPI Routes: 13
#ParseAPI Routes: 26
#Static Routes: 157


benchmarkgocraftweb_param          1000000     1323     ns/op       672     B/op        9     allocs/op
benchmarkgoji_param                5000000     682      ns/op       343     B/op        2     allocs/op
benchmarkgojsonrest_param          500000      4981     ns/op       1793    B/op        30    allocs/op
benchmarkgorillamux_param          500000      3445     ns/op       785     B/op        7     allocs/op
benchmarkhttprouter_param          5000000     443      ns/op       343     B/op        2     allocs/op
benchmarkhttptreemux_param         5000000     485      ns/op       343     B/op        2     allocs/op
benchmarkmartini_param             500000      4574     ns/op       1185    B/op        13    allocs/op
benchmarkpat_param                 1000000     1523     ns/op       687     B/op        14    allocs/op
benchmarktigertonic_param          1000000     2442     ns/op       1025    B/op        19    allocs/op
benchmarktraffic_param             500000      4440     ns/op       2024    B/op        23    allocs/op

benchmarkgocraftweb_param20        500000      6300     ns/op       3863    B/op        17    allocs/op
benchmarkgoji_param20              1000000     2714     ns/op       1260    B/op        2     allocs/op
benchmarkgojsonrest_param20        100000      21353    ns/op       10628   B/op        77    allocs/op
benchmarkgorillamux_param20        200000      12867    ns/op       3311    B/op        10    allocs/op
benchmarkhttprouter_param20        500000      3522     ns/op       2218    B/op        4     allocs/op
benchmarkhttptreemux_param20       500000      4529     ns/op       2219    B/op        4     allocs/op
benchmarkmartini_param20           50000       46866    ns/op       3712    B/op        16    allocs/op
benchmarkpat_param20               500000      3156     ns/op       1492    B/op        25    allocs/op
benchmarktigertonic_param20        50000       33852    ns/op       11259   B/op        179   allocs/op
benchmarktraffic_param20           100000      23082    ns/op       8233    B/op        67    allocs/op

benchmarkgocraftweb_paramwrite     1000000     1421     ns/op       681     B/op        10    allocs/op
benchmarkgoji_paramwrite           2000000     732      ns/op       343     B/op        2     allocs/op
benchmarkgojsonrest_paramwrite     500000      5925     ns/op       2271    B/op        35    allocs/op
benchmarkgorillamux_paramwrite     500000      3632     ns/op       785     B/op        7     allocs/op
benchmarkhttprouter_paramwrite     5000000     478      ns/op       343     B/op        2     allocs/op
benchmarkhttptreemux_paramwrite    5000000     506      ns/op       343     B/op        2     allocs/op
benchmarkmartini_paramwrite        500000      5133     ns/op       1285    B/op        16    allocs/op
benchmarkpat_paramwrite            1000000     2433     ns/op       1128    B/op        19    allocs/op
benchmarktigertonic_paramwrite     500000      3878     ns/op       1482    B/op        25    allocs/op
benchmarktraffic_paramwrite        500000      5566     ns/op       2457    B/op        27    allocs/op

benchmarkgocraftweb_githubstatic   2000000     841      ns/op       313     B/op        6     allocs/op
benchmarkgoji_githubstatic         5000000     345      ns/op       0       B/op        0     allocs/op
benchmarkgojsonrest_githubstatic   500000      4285     ns/op       1155    B/op        26    allocs/op
benchmarkgorillamux_githubstatic   50000       30427    ns/op       459     B/op        6     allocs/op
benchmarkhttprouter_githubstatic   50000000    44.0     ns/op       0       B/op        0     allocs/op
benchmarkhttptreemux_githubstatic  50000000    51.7     ns/op       0       B/op        0     allocs/op
benchmarkmartini_githubstatic      100000      15542    ns/op       859     B/op        12    allocs/op
benchmarkpat_githubstatic          200000      10239    ns/op       3788    B/op        76    allocs/op
benchmarktigertonic_githubstatic   10000000    300      ns/op       49      B/op        1     allocs/op
benchmarktraffic_githubstatic      50000       55501    ns/op       23358   B/op        172   allocs/op

benchmarkgocraftweb_githubparam    1000000     1705     ns/op       735     B/op        10    allocs/op
benchmarkgoji_githubparam          1000000     1168     ns/op       343     B/op        2     allocs/op
benchmarkgojsonrest_githubparam    500000      6368     ns/op       2180    B/op        33    allocs/op
benchmarkgorillamux_githubparam    100000      20148    ns/op       818     B/op        7     allocs/op
benchmarkhttprouter_githubparam    5000000     570      ns/op       343     B/op        2     allocs/op
benchmarkhttptreemux_githubparam   5000000     644      ns/op       343     B/op        2     allocs/op
benchmarkmartini_githubparam       100000      17711    ns/op       1219    B/op        13    allocs/op
benchmarkpat_githubparam           500000      6704     ns/op       2624    B/op        56    allocs/op
benchmarktigertonic_githubparam    500000      4328     ns/op       1482    B/op        28    allocs/op
benchmarktraffic_githubparam       100000      23398    ns/op       7147    B/op        60    allocs/op

benchmarkgocraftweb_githuball      5000        336929   ns/op       136317  B/op        1914  allocs/op
benchmarkgoji_githuball            5000        540415   ns/op       57341   B/op        347   allocs/op
benchmarkgojsonrest_githuball      2000        1257656  ns/op       405606  B/op        6553  allocs/op
benchmarkgorillamux_githuball      100         13067727 ns/op       153329  B/op        1419  allocs/op
benchmarkhttprouter_githuball      10000       112636   ns/op       57345   B/op        347   allocs/op
benchmarkhttptreemux_githuball     10000       123952   ns/op       57342   B/op        347   allocs/op
benchmarkmartini_githuball         200         8073963  ns/op       245198  B/op        2940  allocs/op
benchmarkpat_githuball             500         4418522  ns/op       1587926 B/op        32571 allocs/op
benchmarktigertonic_githuball      2000        899634   ns/op       250656  B/op        6085  allocs/op
benchmarktraffic_githuball         100         11295696 ns/op       3172088 B/op        24928 allocs/op

benchmarkgocraftweb_gplusstatic    2000000     810      ns/op       297     B/op        6     allocs/op
benchmarkgoji_gplusstatic          10000000    242      ns/op       0       B/op        0     allocs/op
benchmarkgojsonrest_gplusstatic    500000      4335     ns/op       1143    B/op        26    allocs/op
benchmarkgorillamux_gplusstatic    1000000     2836     ns/op       459     B/op        6     allocs/op
benchmarkhttprouter_gplusstatic    100000000   26.7     ns/op       0       B/op        0     allocs/op
benchmarkhttptreemux_gplusstatic   100000000   27.7     ns/op       0       B/op        0     allocs/op
benchmarkmartini_gplusstatic       500000      3699     ns/op       860     B/op        12    allocs/op
benchmarkpat_gplusstatic           10000000    294      ns/op       99      B/op        2     allocs/op
benchmarktigertonic_gplusstatic    10000000    160      ns/op       33      B/op        1     allocs/op
benchmarktraffic_gplusstatic       500000      3904     ns/op       1509    B/op        19    allocs/op

benchmarkgocraftweb_gplusparam     1000000     1377     ns/op       672     B/op        9     allocs/op
benchmarkgoji_gplusparam           5000000     757      ns/op       343     B/op        2     allocs/op
benchmarkgojsonrest_gplusparam     500000      5367     ns/op       1815    B/op        30    allocs/op
benchmarkgorillamux_gplusparam     500000      6009     ns/op       785     B/op        7     allocs/op
benchmarkhttprouter_gplusparam     5000000     490      ns/op       343     B/op        2     allocs/op
benchmarkhttptreemux_gplusparam    5000000     532      ns/op       343     B/op        2     allocs/op
benchmarkmartini_gplusparam        500000      6775     ns/op       1186    B/op        13    allocs/op
benchmarkpat_gplusparam            1000000     1742     ns/op       752     B/op        14    allocs/op
benchmarktigertonic_gplusparam     1000000     2934     ns/op       1102    B/op        19    allocs/op
benchmarktraffic_gplusparam        500000      6271     ns/op       2038    B/op        23    allocs/op

benchmarkgocraftweb_gplus2params   1000000     1722     ns/op       735     B/op        10    allocs/op
benchmarkgoji_gplus2params         1000000     1132     ns/op       343     B/op        2     allocs/op
benchmarkgojsonrest_gplus2params   500000      6749     ns/op       2189    B/op        33    allocs/op
benchmarkgorillamux_gplus2params   100000      17109    ns/op       818     B/op        7     allocs/op
benchmarkhttprouter_gplus2params   5000000     568      ns/op       343     B/op        2     allocs/op
benchmarkhttptreemux_gplus2params  5000000     629      ns/op       343     B/op        2     allocs/op
benchmarkmartini_gplus2params      100000      21411    ns/op       1318    B/op        17    allocs/op
benchmarkpat_gplus2params          500000      5351     ns/op       2400    B/op        41    allocs/op
benchmarktigertonic_gplus2params   500000      4619     ns/op       1584    B/op        28    allocs/op
benchmarktraffic_gplus2params      100000      18866    ns/op       3620    B/op        35    allocs/op

benchmarkgocraftweb_gplusall       100000      18978    ns/op       8338    B/op        117   allocs/op
benchmarkgoji_gplusall             200000      11164    ns/op       3773    B/op        22    allocs/op
benchmarkgojsonrest_gplusall       50000       74203    ns/op       24033   B/op        402   allocs/op
benchmarkgorillamux_gplusall       10000       107060   ns/op       9724    B/op        91    allocs/op
benchmarkhttprouter_gplusall       500000      6688     ns/op       3774    B/op        22    allocs/op
benchmarkhttptreemux_gplusall      200000      7832     ns/op       3774    B/op        22    allocs/op
benchmarkmartini_gplusall          10000       125425   ns/op       15520   B/op        194   allocs/op
benchmarkpat_gplusall              50000       42806    ns/op       17685   B/op        346   allocs/op
benchmarktigertonic_gplusall       50000       44188    ns/op       15470   B/op        322   allocs/op
benchmarktraffic_gplusall          10000       129847   ns/op       42065   B/op        446   allocs/op

benchmarkgocraftweb_parsestatic    2000000     829      ns/op       313     B/op        6     allocs/op
benchmarkgoji_parsestatic          5000000     312      ns/op       0       B/op        0     allocs/op
benchmarkgojsonrest_parsestatic    500000      4349     ns/op       1143    B/op        26    allocs/op
benchmarkgorillamux_parsestatic    500000      5420     ns/op       459     B/op        6     allocs/op
benchmarkhttprouter_parsestatic    100000000   24.4     ns/op       0       B/op        0     allocs/op
benchmarkhttptreemux_parsestatic   50000000    48.6     ns/op       0       B/op        0     allocs/op
benchmarkmartini_parsestatic       500000      4180     ns/op       860     B/op        12    allocs/op
benchmarkpat_parsestatic           5000000     700      ns/op       249     B/op        5     allocs/op
benchmarktigertonic_parsestatic    10000000    234      ns/op       49      B/op        1     allocs/op
benchmarktraffic_parsestatic       500000      6817     ns/op       2390    B/op        25    allocs/op

benchmarkgocraftweb_parseparam     1000000     1713     ns/op       688     B/op        9     allocs/op
benchmarkgoji_parseparam           1000000     1005     ns/op       343     B/op        2     allocs/op
benchmarkgojsonrest_parseparam     500000      6093     ns/op       1803    B/op        30    allocs/op
benchmarkgorillamux_parseparam     500000      5874     ns/op       786     B/op        7     allocs/op
benchmarkhttprouter_parseparam     5000000     498      ns/op       343     B/op        2     allocs/op
benchmarkhttptreemux_parseparam    5000000     561      ns/op       343     B/op        2     allocs/op
benchmarkmartini_parseparam        500000      5814     ns/op       1186    B/op        13    allocs/op
benchmarkpat_parseparam            1000000     2551     ns/op       1196    B/op        20    allocs/op
benchmarktigertonic_parseparam     1000000     2772     ns/op       1083    B/op        19    allocs/op
benchmarktraffic_parseparam        500000      6293     ns/op       2324    B/op        25    allocs/op

benchmarkgocraftweb_parse2params   1000000     1634     ns/op       735     B/op        10    allocs/op
benchmarkgoji_parse2params         2000000     909      ns/op       343     B/op        2     allocs/op
benchmarkgojsonrest_parse2params   500000      6700     ns/op       2161    B/op        33    allocs/op
benchmarkgorillamux_parse2params   500000      6396     ns/op       818     B/op        7     allocs/op
benchmarkhttprouter_parse2params   5000000     562      ns/op       343     B/op        2     allocs/op
benchmarkhttptreemux_parse2params  5000000     645      ns/op       343     B/op        2     allocs/op
benchmarkmartini_parse2params      500000      6453     ns/op       1219    B/op        13    allocs/op
benchmarkpat_parse2params          1000000     2600     ns/op       907     B/op        21    allocs/op
benchmarktigertonic_parse2params   500000      4598     ns/op       1486    B/op        28    allocs/op
benchmarktraffic_parse2params      500000      6412     ns/op       2129    B/op        25    allocs/op

benchmarkgocraftweb_parseall       50000       35884    ns/op       14294   B/op        209   allocs/op
benchmarkgoji_parseall             100000      19034    ns/op       5489    B/op        33    allocs/op
benchmarkgojsonrest_parseall       10000       139099   ns/op       41514   B/op        758   allocs/op
benchmarkgorillamux_parseall       10000       227415   ns/op       17254   B/op        175   allocs/op
benchmarkhttprouter_parseall       200000      9141     ns/op       5489    B/op        33    allocs/op
benchmarkhttptreemux_parseall      200000      10346    ns/op       5489    B/op        33    allocs/op
benchmarkmartini_parseall          10000       157143   ns/op       27673   B/op        333   allocs/op
benchmarkpat_parseall              50000       47640    ns/op       18273   B/op        385   allocs/op
benchmarktigertonic_parseall       50000       59743    ns/op       20843   B/op        420   allocs/op
benchmarktraffic_parseall          10000       195778   ns/op       70533   B/op        762   allocs/op

benchmarkhttpservemux_staticall    2000        930269   ns/op       104     B/op        8     allocs/op
benchmarkgocraftweb_staticall      10000       147355   ns/op       49153   B/op        951   allocs/op
benchmarkgoji_staticall            50000       73174    ns/op       0       B/op        0     allocs/op
benchmarkgojsonrest_staticall      5000        741010   ns/op       182597  B/op        4128  allocs/op
benchmarkgorillamux_staticall      500         3706206  ns/op       72310   B/op        966   allocs/op
benchmarkhttprouter_staticall      200000      13843    ns/op       0       B/op        0     allocs/op
benchmarkhttptreemux_staticall     200000      13919    ns/op       0       B/op        0     allocs/op
benchmarkmartini_staticall         1000        2874761  ns/op       145575  B/op        2521  allocs/op
benchmarkpat_staticall             1000        1462445  ns/op       554141  B/op        11249 allocs/op
benchmarktigertonic_staticall      50000       53678    ns/op       7776    B/op        158   allocs/op
benchmarktraffic_staticall         200         9174624  ns/op       3794337 B/op        27918 allocs/op
```
