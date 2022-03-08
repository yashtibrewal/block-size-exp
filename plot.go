package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getXAndYData(path string) ([]int64, []opts.BarData) {
	file, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	y_axis := make([]opts.BarData, 0)
	x_axis := make([]int64, 0)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")
		y_value, err := strconv.ParseInt(lines[1], 10, 64)
		check(err)
		y_value = y_value / 1024 // Bytes to KB
		y_value = y_value / 1024 // KB to MB
		y_axis = append(y_axis, opts.BarData{Value: y_value})
		x_value, err := strconv.ParseInt(lines[0], 10, 64)
		check(err)
		x_axis = append(x_axis, x_value)
	}

	file.Close()
	return x_axis, y_axis
}

func main() {
	fmt.Println("Plotting the graphs for read and write speed of the SSD w.r.t to changing block sizes.")

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Reads and Write Speed(MB/s) in bytes per second to with different block sizes(in K blocks)",
		Subtitle: "",
	}))
	x_axis, r_y_axis := getXAndYData("reads.txt")
	bar.SetXAxis(x_axis)
	bar.AddSeries("Reads", r_y_axis)
	_, w_y_axis := getXAndYData("writes.txt")
	bar.AddSeries("Writes", w_y_axis)
	figure, _ := os.Create("bar.html")
	bar.Render(figure)
}
