I have tried to plot the speed of read and write on my Mac w.r.t to different block sizes.

Bash scrip does the io and stores the result.
Go code generates a graph.

# Assuming linux based O.S. and go environment is set.

`bash runme.bash`
`go build`
`go run plot.go`

bar.html file holds the graph for the output

Conclusion: There is a sharp decrease in speeds once the data croses 512K block size.