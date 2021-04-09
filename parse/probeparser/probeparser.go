package probeparser

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Log struct {
	Fulllog string
	Pid     int64
	Time    float64
	Probe   string
}

const (
	timestamp int = 0
)

func GetNS(pid string) string {
	cmdName := "ls"
	out, err := exec.Command(cmdName, fmt.Sprintf("/proc/%s/ns/net", pid), "-al").Output()
	if err != nil {
		println(err)
	}
	ns := string(out)
	parse := strings.Fields(string(ns))
	s := strings.SplitN(parse[10], "[", 10)
	sep := strings.Split(s[1], "]")
	return sep[0]

}


func RunUretprobeFreq(tool string, loguretprobe chan Log, pid string, filepath string, funcname string) {

	filepath = strings.Replace(filepath, "\n", "", -1)
        funcname = strings.Replace(funcname, "\n", "", -1)
        pid = strings.Replace(pid, "\n", "", -1)
	path := "/proc/" + pid + "/root/" + filepath
        path = strings.Replace(path, "\n","", -1)
	//command := `uprobe:` + filepath + `:` + funcname + `{ @start = nsecs; }` + `uretprobe:` + filepath + `:` + funcname + `/@start/ { @time = ((nsecs - @start)/1000);  print(@time); delete(@start); }`
	//cmd := exec.Command("bpftrace", "-p", pid , "-e", command)
	cmd := exec.Command("./functime", "-s", path, "-fn", funcname)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	fmt.Println("cmd--",cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))
		if (len(parsedLine) > 0) && parsedLine[0] != "TIME"{
			//fmt.Println("Freq",parsedLine[0])
			//fmt.Println("Freq",parsedLine[1])
			timest := 0.00
			fmt.Println("Freq",string(line))
//			log := parsedLine[0] + " " + parsedLine[1] + " " + parsedLine[2] + " " + parsedLine[3]
			n := Log{Fulllog: string(line), Pid: 1234, Time: timest, Probe: tool}
			loguretprobe <- n

		}
	}
}


func RunUretprobeCount(tool string, loguretprobe chan Log, pid string, filepath string, funcname string) {

	filepath = strings.Replace(filepath, "\n", "", -1)
        funcname = strings.Replace(funcname, "\n", "", -1)
        pid = strings.Replace(pid, "\n", "", -1)
	path := "/proc/" + pid + "/root/" + filepath
        path = strings.Replace(path, "\n", "", -1)
	path = path + ":" + funcname
	//command  := `uretprobe:` + filepath + `:` + funcname + `{ @[pid] = count(); }` +`interval:s:1` + `{ print(@); clear(@); }`
	//cmd := exec.Command("bpftrace", "-p", pid , "-e", command)

	cmd := exec.Command("funccount", "-i", "1" , path)
	// testing for now the below command- it worked
	//cmd := exec.Command("bpftrace", "-e" , filepath)
	//	fmt.Println(funcname, pid)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		fmt.Println("trrr:", cmd)
		parsedLine := strings.Fields(string(line))
                if (len(parsedLine) > 0) && parsedLine[0] != "Tracing"{
			if(parsedLine[0] != "FUNC"){
			//s := strings.Split(parsedLine[0], "[")
                        //sep := strings.Split(s[1], "]")
//				log := parsedLine[0] + " " + parsedLine[1]
				timest := 0.00
//				fmt.Println("count--", log)
				n := Log{Fulllog: string(line), Pid: 1234, Time: timest, Probe: tool}
				loguretprobe <- n
			}

		}
	}
}


func RunUretprobe(tool string, loguretprobe chan Log, pid string, filepath string, funcname string) {

	filepath = strings.Replace(filepath, "\n", "", -1)
	funcname = strings.Replace(funcname, "\n", "", -1)
	pid = strings.Replace(pid, "\n", "", -1)
	path := "/proc/" + pid + "/root/" + filepath
        path = strings.Replace(path, "\n", "", -1)

	//command := `uretprobe:` +filepath+ `:` + funcname + `{ printf("%d %d\n", pid, retval); }`
	cmd := exec.Command("./funcretval", "-s", path , "-fn", funcname)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))
	fmt.Println("trrr:", cmd)

                if (len(parsedLine) > 0) && parsedLine[0] != "TIME"{

			timest := 0.00
//			log := parsedLine[0] + " " + parsedLine[1] + " " + parsedLine[2]
			n := Log{Fulllog: string(line), Pid: 1234, Time: timest, Probe: tool}
			loguretprobe <- n

		}
	}
}

func RunTcptracer(tool string, logtcptracer chan Log, pid string) {

	sep := GetNS(pid)

	cmd := exec.Command("./tcptracer.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))
		fmt.Println(cmd)
		if parsedLine[0] != "Tracing" {
			if parsedLine[0] != "TIME(s)" {
				ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
				if err != nil {
					println("Tcptracer PID Error")
				}

				timest := 0.00
				n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
				logtcptracer <- n

			}
		}
	}
}

func RunTcpconnect(tool string, logtcpconnect chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./tcpconnect.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))
		if parsedLine[0] != "TIME(s)" {
			ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
			if err != nil {
				println("TCPConnect PID Error")
			}

			timest := 0.00
			n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
			logtcpconnect <- n

		}
	}
}

func RunTcpaccept(tool string, logtcpaccept chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./tcpaccept.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		if parsedLine[0] != "TIME(s)" {
			ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
			if err != nil {
				println("TCPaccept PID Error")
			}

			timest := 0.00

			n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
			logtcpaccept <- n

		}
	}
}

func RunTcplife(tool string, logtcplife chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./tcplife.py", "-T", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))
		println(parsedLine[0])
		if parsedLine[0] != "TIME(s)" {
			ppid, err := strconv.ParseInt(parsedLine[2], 10, 64)
			if err != nil {
				println("TCPlife PID Error")
			}

			timest := 0.00

			n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
			logtcplife <- n

		}
	}
}

func RunExecsnoop(tool string, logexecsnoop chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./execsnoop.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		ppid, err := strconv.ParseInt(parsedLine[4], 10, 64)
		if err != nil {
			println("Execsnoop PID Error")
		}

		timest := 0.00

		n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
		logexecsnoop <- n

	}
}

func RunBiosnoop(tool string, logbiosnoop chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./biosnoop.py", "-T", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
		if err != nil {
			println("Biosnoop PID Error")
		}
		timest := 0.00

		n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
		logbiosnoop <- n

	}
}

func RunCachetop(tool string, logcachetop chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./Cachetop.py", "-T", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		ppid, err := strconv.ParseInt(parsedLine[1], 10, 64)
		if err != nil {
			println("Cachetop PID Error")
		}
		timest := 0.00

		n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
		logcachetop <- n

	}
}
