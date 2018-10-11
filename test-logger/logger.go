package main

/*
 * Implements a simple program that outputs a random log statement at a regular interval.
 */
import (
	"fmt"
	"os"

	"github.com/golang/glog"
)

func main() {
	glog.Info("!!!!AN INFO MESSAGE>!?")
	glog.Info("!!!!ANOTHER INFO MESSAGE>!?")
	glog.Error("!!!!This is an ERROR message!@$ hurrah")
	glog.V(2).Infoln("!!! this is a v2 info line ok")
	glog.Warning("!!!!WARNWARN MESSAGE>!?")
	glog.Flush()
}
