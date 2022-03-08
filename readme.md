Plotting the speed of read and write on the disk on my Macbook Air M1 Chip (L1 Cahce: 128K, 196K) w.r.t to different block sizes.

Bash scrip does the io and stores the result.
Go code generates a graph.

Assuming linux based O.S. and go environment is set.

`bash runme.bash`
`go build`
`go run plot.go`

bar.html file holds the graph for the output

Observation: There is a sharp decrease in speeds once the data croses 256K block size.

Conclusions: There is dip in read speeds drastically after using more size of data then the L1 Cache in the system.