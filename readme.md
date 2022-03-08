Plotting the speed of read and write on the disk on my Mac w.r.t to different block sizes.

Bash scrip does the io and stores the result.
Go code generates a graph.

Assuming linux based O.S. and go environment is set.

`bash runme.bash`
`go build`
`go run plot.go`

bar.html file holds the graph for the output

Observation: There is a sharp decrease in speeds once the data croses 256K block size.