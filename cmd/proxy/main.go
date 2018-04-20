package main

import (
	"path"
	"strconv"

	"github.com/magenta9/zcdis/utils"

	"github.com/docopt/docopt-go"
	log "github.com/ngaut/logging"
)

var (
	cpus       = 2
	addr       = ":9000"
	httpAddr   = ":9001"
	configFile = "config.ini"
)

var usage = `usage: proxy [-c <config_file>] [-L <log_file>] [--log-level=<loglevel>] [--cpu=<cpu_num>] [--addr=<proxy_listen_addr>] [--http-addr=<debug_http_server_addr>]

options:
   -c	set config file
   -L	set output log file, default is stdout
   --log-level=<loglevel>	set log level: info, warn, error, debug [default: info]
   --cpu=<cpu_num>		num of cpu cores that proxy can use
   --addr=<proxy_listen_addr>		proxy listen address, example: 0.0.0.0:9000
   --http-addr=<debug_http_server_addr>		debug vars http server
`

func main() {
	log.SetLevelByString("info")

	args, err := docopt.Parse(usage, nil, true, "v1.0", false)
	if err != nil {
		log.Error(err)
	}

	//config file
	if args["-c"] != nil {
		configFile = args["-c"].(string)
	}

	//output log file
	if args["-L"] != nil {
		log.SetOutputByName(args["-L"].(string))
	}

	//log level: fatal, error, warn, debug, info
	if args["--log-level"] != nil {
		log.SetLevelByString(args["--log-level"].(string))
	}

	if args["--cpu"] != nil {
		cpus, err = strconv.Atoi(args["--cpu"].(string))
		if err != nil {
			log.Fatal(err)
		}
	}

	// addr
	if args["--addr"] != nil {
		addr = args["--addr"].(string)
	}

	// http addr
	if args["--http-addr"] != nil {
		httpAddr = args["--http-addr"].(string)
	}

	dumpPath := utils.GetExecPath()
	log.Info("dump file path:", dumpPath)
	log.CrashLog(path.Join(dumpPath, "proxy.dump"))



}
