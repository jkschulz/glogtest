package main

/*
 * Implements a simple program that outputs a random log statement at a regular interval.
 */
import (
	"fmt"
  "flag"
	"math/rand"
	"os"
	"time"

	"github.com/golang/glog"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// Take the log interval and logging header from the environment
	logInterval := convertToInt(os.Getenv("LOG_INTERVAL"), 2)
	header := os.Getenv("HEADER")
	whenToStart := rand.Intn(10)
	time.Sleep(time.Duration(whenToStart) * time.Second)
	for true {
		time.Sleep(time.Duration(logInterval) * time.Second)
		glog.Info("!an info message")
		glog.Info("@#another info")
		glog.Error("weewoo errah")
		glog.V(2).Infoln("&this is a v2 info line ok")
		glog.Warning("^WARN^")
	}
}
