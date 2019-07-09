package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type tCpuData [7]int

const (
	sUrlRoot         = "/"
	sUrlVars         = "/varz"
	sUrlHealth       = "/healthz"
	sUrlStatus       = "/statusz"
	sUrlHealthInvert = "/healthzInvert"
)

var (
	sParPort       = flag.String("listen-port", "8080", "The port to listen on for HTTP requests.")
	iParRefresh    = flag.Int("refresh-interval", 30, "Interval for front page refreshes (0=disable)")
	iParRequestLog = flag.Int("request-log", 1, "Log request to console (0=disable)")
)

var stCpu tCpuData
var saUrls = []string{sUrlVars, sUrlHealth, sUrlStatus, sUrlHealthInvert}
var pageTop = "<html><head>%s<title>Hello World by Go lang</title></head><body><h1>Hello World from Go!</h1><p>Available end points: "
var pageBottom = "</body></html>"
var tiStarted time.Time
var iStatHealth = true
var iStatReady = false
var sHelp = `Application end points:
- /                     main front page
- /varz                 CPU usage
- /healthz              application health state
- /statusz              application ready state
- /healthzInvert        inverts health state`

func main() {
	tiStarted = time.Now()
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Simple http server that exposes CPU usage to world\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\n"+sHelp)
	}
	flag.Parse()
	if *iParRefresh > 0 {
		pageTop = fmt.Sprintf(pageTop, "<meta http-equiv=refresh content="+strconv.Itoa(*iParRefresh)+">")
	}
	go getCpuLoad()

	http.HandleFunc(sUrlHealth, func(w http.ResponseWriter, r *http.Request) {
		iHdr, sStr := 200, "ok"
		RequestLog(r)
		if !iStatHealth {
			iHdr, sStr = 500, "error"
		}
		w.WriteHeader(iHdr)
		w.Write([]byte(sStr))
	})
	http.HandleFunc(sUrlHealthInvert, func(w http.ResponseWriter, r *http.Request) {
		RequestLog(r)
		iStatHealth = !iStatHealth
		w.Write([]byte(fmt.Sprintf("Health status changed from %v to %v", !iStatHealth, iStatHealth)))
	})

	http.HandleFunc(sUrlStatus, func(w http.ResponseWriter, r *http.Request) {
		iHdr, sStr := 200, "ready"
		RequestLog(r)
		if !iStatReady {
			iHdr, sStr = 500, "not ready"
		}
		w.WriteHeader(iHdr)
		w.Write([]byte(sStr))
	})

	http.HandleFunc(sUrlVars, func(w http.ResponseWriter, r *http.Request) {
		tUp := time.Now().Unix() - tiStarted.Unix()
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("hww_user %d\nhww_nice %d\nhww_system %d\nhww_idle %d\nhww_iowait %d\nhww_uptime %d\nhww_health %v\nhww_ready %v\n",
			stCpu[0], stCpu[1], stCpu[2], stCpu[3], stCpu[4], tUp, Bool2Int(iStatHealth), Bool2Int(iStatReady))))
	})

	fmt.Printf("Listening on port %s\n", *sParPort)
	http.HandleFunc(sUrlRoot, PageIndex)
	iStatReady = true
	if err := http.ListenAndServe(":"+*sParPort, nil); err != nil {
		log.Fatal("Failed to start server", err)
		iStatHealth, iStatReady = false, false
	}
}

func RunCommand(sCmd string) string {
	sOut, err := exec.Command(sCmd).Output()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		iStatHealth, iStatReady = false, false
	}
	return string(sOut)
}

func RequestLog(request *http.Request) {
	if *iParRequestLog == 1 {
		fmt.Printf("%v %v %v\n", time.Now(), request.RemoteAddr, request.URL.Path)
	}
}

func PageIndex(writer http.ResponseWriter, request *http.Request) {
	RequestLog(request)

	fmt.Fprint(writer, pageTop)
	for _, s := range saUrls {
		fmt.Fprintf(writer, "| <a href=\"%s\">%s</a> ", s, s)
	}
	sPage := fmt.Sprintf("date:<br>%s\nuptime:<br>%s\nps:<br>%s\n", RunCommand("date"), RunCommand("uptime"), RunCommand("ps"))
	fmt.Fprintf(writer, "<pre>%s</pre>", sPage)
	fmt.Fprint(writer, pageBottom)
}

func getCpuData() tCpuData {
	rawBytes, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		log.Fatal(err)
		iStatHealth, iStatReady = false, false
		return tCpuData{}
	}
	sLine := strings.Split(string(rawBytes), "\n")[0]
	iCpu := strings.Count(string(rawBytes), "cpu") - 1
	var aiCpuData tCpuData

	reSpace := regexp.MustCompile(`\s+`)
	s := reSpace.ReplaceAllString(sLine, " ")

	as := strings.Split(s, " ")
	for i, v := range as[1:8] {
		aiCpuData[i], err = strconv.Atoi(v)
		aiCpuData[i] /= iCpu
	}
	return aiCpuData
}

func getCpuLoad() {
	for {
		data0 := getCpuData()
		time.Sleep(time.Second)
		data1 := getCpuData()

		for i, v := range data1 {
			stCpu[i] = v - data0[i]
		}
	}
}

func Bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}
