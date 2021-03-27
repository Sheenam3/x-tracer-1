package main

import (
	"context"
	"fmt"
	"github.com/ITRI-ICL-Peregrine/x-tracer/pkg"
	"github.com/docker/docker/client"
	"log"
	"os"
	"strings"
	"time"

)

func main() {

	//var ucmd []string
	log.Println("Start api...")

	containerId := os.Getenv("containerId")

	if containerId == "" {
		containerId = "ec9515bb14a2"
	}

	serverIp := os.Getenv("masterIp")
	if containerId == "" {
		containerId = "ec9515bb14a2"
	}

	probeName := os.Getenv("tools")

	ucmd  := os.Getenv("userInput")

	Probe := strings.Split(probeName, ",")

	//if uretprobe == "true" { ucmd = probecmd.GetUretCmd(userinput, Probe) }


	cli, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"), client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	topResult, err := cli.ContainerTop(context.Background(), containerId, []string{"o", "pid"})
	if err != nil {
		panic(err)
	}
	fmt.Println(topResult.Processes)

	log.Printf("Start new client")
	fmt.Println("userinput-", ucmd)
	testClient := pkg.New("6666", serverIp)
	testClient.StartClient(Probe, topResult.Processes, ucmd)

	for {
		fmt.Println("x-agent - Sleeping")
		time.Sleep(10 * time.Second)
	}

}
