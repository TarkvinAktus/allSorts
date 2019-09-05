## Some basic sorts

### Benchmarks

Array[100]

    Benchmark_bubble-4                 30000             46744 ns/op               0 B/op          0 allocs/op
    Benchmark_Quick-4                 100000             13735 ns/op               0 B/op          0 allocs/op
    Benchmark_Quick_Goroutines-4       50000             34511 ns/op               0 B/op          0 allocs/op

Array[1000]

    Benchmark_bubble-4                   500           2438519 ns/op              16 B/op          0 allocs/op
    Benchmark_Quick-4                  20000             94920 ns/op               0 B/op          0 allocs/op
    Benchmark_Quick_Goroutines-4       10000            107744 ns/op               1 B/op          0 allocs/op

Array[10.000]

    Benchmark_bubble-4                     5         322999620 ns/op           16384 B/op          0 allocs/op
    Benchmark_Quick-4                   2000            983623 ns/op              40 B/op          0 allocs/op
    Benchmark_Quick_Goroutines-4        2000            658131 ns/op              44 B/op          0 allocs/op

Array[100.000]

    Benchmark_bubble-4                     1        39458179600 ns/op         802816 B/op          1 allocs/op
    Benchmark_Quick-4                    200           9422404 ns/op            4014 B/op          0 allocs/op
    Benchmark_Quick_Goroutines-4         300           5237069 ns/op            2677 B/op          0 allocs/op

Array[1.000.000]

    Benchmark_Quick-4                     20          95472870 ns/op          400179 B/op          0 allocs/op
    Benchmark_Quick_Goroutines-4          30          48626596 ns/op          267072 B/op          0 allocs/op

Array[10.000.000]
    Benchmark_Quick-4                      1        1005124600 ns/op        80003168 B/op          2 allocs/op
    Benchmark_Quick_Goroutines-4           2         519285800 ns/op        40001848 B/op          3 allocs/op

