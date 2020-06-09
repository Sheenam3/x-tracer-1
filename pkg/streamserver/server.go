package streamserver

import (
	"strings"
	"fmt"
	pb "github.com/ITRI-ICL-Peregrine/x-tracer/api"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
        "os"
)


type StreamServer struct {
	port string
}

func (s *StreamServer) RouteLog(stream pb.SentLog_RouteLogServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Res:                  "Stream closed",
			})
		}
		if err != nil {
			return err
		}

		f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0644)
   		if err != nil {
		        fmt.Println(err)
        		
    		}
//		fmt.Println("\n", r.Log)

	        parse := strings.Fields(string(r.Log))
//		fmt.Println("PID:",r.Pid)
		if r.ProbeName == "tcptracer"{

		//fmt.Println("ProbeName:",r.ProbeName)
                //fmt.Printf("{%s}\n", r.Log)
	                fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s |IP->%s | SADDR:%s | DADDR:%s | SPORT:%s | DPORT:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8],parse[9])
			_, err = fmt.Fprintln(f, "{Probe:", r.ProbeName, "|Sys_Time:", parse[0] ,"|T:",parse[1] ,"| PID:",parse[3] ,"| PNAME:",parse[4] ,"| IP:",parse[5] ,"| RADDR:",parse[6], "| RPORT:",parse[7] ,"| LADDR:",parse[8], "| LPORT:",parse[9] )
			if err != nil {
		        	fmt.Println(err)
                			f.Close()
				
			}
			err = f.Close()
			if err != nil {
     				   fmt.Println(err)
        			
    			}

                }
		if r.ProbeName == "tcpaccept"{

                //fmt.Println("ProbeName:",r.ProbeName)
		//fmt.Printf("{%s}\n", r.Log)
                	fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | RADDR:%s | RPORT:%s | LADDR:%s | LPORT:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8],parse[9])
			_, err = fmt.Fprintln(f, "{Probe:", r.ProbeName, "|Sys_Time:", parse[0], "|T:",parse[1], "| PID:",parse[3], "| PNAME:",parse[4], "| IP:",parse[5], "| RADDR:",parse[6], "| RPORT:",parse[7], "| LADDR:",parse[8], "| LPORT:",parse[9] )
                	if err != nil {
                        	fmt.Println(err)
                                	f.Close()
                        //	return
                	}
                	err = f.Close()
                	if err != nil {
                        	   fmt.Println(err)
                        	//return
               		 }


                }
		if r.ProbeName == "tcplife"{

			fmt.Printf("{Probe:%s |Sys_Time: %s |PID:%s | PNAME:%s | LADDRR:%s | LPORT:%s | RADDR:%s | RPORT:%s | TX_KB:%s | RX_KB:%s | MS: %s \n",r.ProbeName,parse[0],parse[2],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8],parse[9],parse[10])
			_, err = fmt.Fprintln(f, "{Probe:", r.ProbeName, "|Sys_Time:", parse[0] ,"|T:",parse[1] ,"| PID:",parse[3] ,"| PNAME:",parse[4] ,"| IP:",parse[5] ,"| RADDR:",parse[6], "| RPORT:",parse[7] ,"| LADDR:",parse[8], "| LPORT:",parse[9] )
			if err != nil {
		        	fmt.Println(err)
                			f.Close()
				
			}
			err = f.Close()
			if err != nil {
     				   fmt.Println(err)
        			
    			}

		
		}
		if r.ProbeName == "execsnoop"{
			fmt.Printf("{Probe:%s |Sys_Time: %s | T:%s | PNAME: %s | PID:%s | PPID:%s | RET:%s | ARGS:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7])
			_, err = fmt.Fprintln(f, "{Probe:", r.ProbeName, "| Sys_Time:", parse[0] ,"| T:",parse[1] , "| PNAME:",parse[3], "| PID:",parse[4] ,"|PPID:",parse[5] ,"| RET:",parse[6], "| ARGS:",parse[7] )
			if err != nil {
		        	fmt.Println(err)
                			f.Close()
				
			}
			err = f.Close()
			if err != nil {
     				   fmt.Println(err)
        			
    			}

		}
		if r.ProbeName == "biosnoop"{

			fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s |PNAME: %s | PID:%s | DISK:%s | R/W:%s | SECTOR:%s |BYTES: %s | Lat(ms): %s | \n",r.ProbeName,parse[0],parse[1],parse[2],parse[3],parse[4],parse[5],parse[6],parse[7],parse[9])
			_, err = fmt.Fprintln(f, "{Probe:", r.ProbeName, "|Sys_Time:", parse[0] ,"|T:",parse[1] ,"| PID:",parse[3] ,"| PNAME:",parse[2] ,"| DISK:",parse[4] ,"| R/W:",parse[5], "| SECTOR:",parse[6] ,"| BYTES:",parse[7], "| LAT(ms):",parse[9] )
			if err != nil {
		        	fmt.Println(err)
                			f.Close()
				
			}
			err = f.Close()
			if err != nil {
     				   fmt.Println(err)
        			
    			}

		}
		if r.ProbeName == "cachetop"{

			fmt.Printf("{Probe:%s |Sys_Time: %s | PID:%s | UID:%s | CMD:%s | HITS:%s | MISS:%s | DIRTIES: %s| READ_HIT%:%s | W_HIT%: %s | \n",r.ProbeName,parse[0],parse[1],parse[2],parse[3],parse[5],parse[6],parse[7],parse[8], parse[9])
			_, err = fmt.Fprintln(f, "{Probe:", r.ProbeName, "|Sys_Time:", parse[0] ,"| PID:",parse[1] ,"| UID:",parse[2] ,"| CMD:",parse[3] ,"| HITS:",parse[5], "| MISS:",parse[6] ,"| DIRTIES:",parse[7], "| READ_HIT%:",parse[8], "|W_HIT%%:",parse[9] )
			if err != nil {
		        	fmt.Println(err)
                			f.Close()
				
			}
			err = f.Close()
			if err != nil {
     				   fmt.Println(err)
        			
    			}

		}
		if r.ProbeName == "tcpconnect"{
	                fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8])
			_, err = fmt.Fprintln(f, "{Probe:", r.ProbeName, "|Sys_Time:", parse[0] ,"|T:",parse[1] ,"| PID:",parse[3] ,"| PNAME:",parse[4] ,"| IP:",parse[5] ,"| SADDR:",parse[6], "| DADDR:",parse[7] ,"| DPORT:",parse[8] )
			if err != nil {
		        	fmt.Println(err)
                			f.Close()
				
			}
			err = f.Close()
			if err != nil {
     				   fmt.Println(err)
        			
    			}


                }
		//fmt.Println(r.TimeStamp, "\n")
	}
}

func New(servicePort string) *StreamServer{
	return &StreamServer{
		servicePort}
}

func (s *StreamServer) StartServer(){
	server := grpc.NewServer()
	pb.RegisterSentLogServer(server, &StreamServer{})

	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		log.Fatalln("net.Listen error:", err)
	}

	_ = server.Serve(lis)
}

