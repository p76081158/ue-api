package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"strings"
	"strconv"

	"gonum.org/v1/gonum/stat/distuv"
	"golang.org/x/exp/rand"
)

var Server = "google.com"
var Request_pattern = "500:50"
var Interval_delay = 2
var TimeWindow = 5
var Resource_ratio = 10
var TotalRequestSend = 0

// convert string to int
func StringToInt(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
	return i
}

// sending requests to server
func SendRequest() {
	input_cmd := "curl " + Server
	cmd := exec.Command("/bin/sh", "-c", input_cmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	TotalRequestSend++
	return
}

// sending requests with poisson distribution during the timewindow
func RequestPoisson(lambda float64, request_num int) {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	p := distuv.Poisson{lambda, r}
	for i := 0; i < request_num; i++ {
		request_delay := int(p.Rand())
		go SendRequest()
		time.Sleep(time.Duration(request_delay) * time.Millisecond)
	}
}

// generate lambda of requests number and requests delay within a timewindow
// number of request is generate by poisson distribution
// e.g. num_lambda   = 250  (average number requests between timewindow is 250)
//      delay_lambda = 20   (average delay between requests is 20ms)
func RequestPatternGenerator(resource int, duration int) {
	timeWindowNum := duration / TimeWindow
	num_lambda := float64((resource / Resource_ratio) * TimeWindow)
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	num_poisson := distuv.Poisson{num_lambda, r}
	for i := 0; i < timeWindowNum; i++ {
		request_nums := int(num_poisson.Rand())
		delay_lambda := float64(1000 / request_nums / TimeWindow)
		go RequestPoisson(delay_lambda ,request_nums)
		time.Sleep(time.Duration(TimeWindow) * time.Second)
	}
}

// schedule between request interval
func RequestScheduler(pattern string) {
	interval := strings.Split(pattern, ",")
	for i := 0; i < len(interval); i++ {
		temp := strings.Split(pattern, ":")
		resource := StringToInt(temp[0])
		duration := StringToInt(temp[1])
		go RequestPatternGenerator(resource, duration)
		time.Sleep(time.Duration(duration + Interval_delay) * time.Second)
	}
}

func main() {
	if (os.Args[1]!="") {
		Server = string(os.Args[1])
	}
    if (os.Args[2]!="") {
		Request_pattern = string(os.Args[2])
	}
	if (os.Args[3]!="") {
		Resource_ratio = StringToInt(string(os.Args[3]))
	}

	start := time.Now()
	RequestScheduler(Request_pattern)

	fmt.Println("")
	fmt.Println("Duration of Sending Requests: ", time.Since(start))
	fmt.Println("Total Requests Sended: ", TotalRequestSend)
}