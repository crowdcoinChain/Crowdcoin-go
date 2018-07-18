package main

import (
	"runtime"
	"runtime/debug"
	"fmt"
	"os"
	"limits"

)

const CROWDCOIN_CONF_FILENAME = "crowdcoin.conf"
const CROWDCOIN_PID_FILENAME = "crowdcoind.pid"


//winServiceMain inwoked only in windows detact if its running in a service
var winServiceMain func()(bool, error)

func SetupEnvironment(){

	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Limit Garbage Collection
	debug.SetGCPercent(10)

	//Up some limits.
	if err := limits.SetLimits(); err!= nil {
		fmt.Fprintf(os.Stderr, "failed to set limits: %v\n",err)
		os.Exit(1)
	}

	//Windows handle running as service
	if runtime.GOOS == "windows" {
		isService, err := winServiceMain()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if isService {
			os.Exit(0)
		}
	}

}
