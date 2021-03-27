package probecmd


func GetUretCmd(userinput string, probe []string) []string {

	//cmd := make([]string,4) - direct assignment works with this declaration
	var cmd []string

        if len(probe) == 1{
                switch probe[0] {

                case "Retval":
                        //cmd[i] = getRetvalCmd(userinput) throwing index out of range error
			cmd = append(cmd,getRetvalCmd(userinput))
                        break;

                case "Count":
			cmd = append(cmd,getCountCmd(userinput))
	                break;

                case "Frequency":

                        cmd = append(cmd,getFreqCmd(userinput))
                        break;
                }


        }else if len(probe) > 1{

		cmd = append(cmd,getRetvalCmd(userinput))

                cmd = append(cmd,getCountCmd(userinput))

                cmd = append(cmd,getFreqCmd(userinput))

        }

        return cmd

}



func getRetvalCmd(userinput string) string{

        comm := `{ printf("%d %d\n", pid, retval); }`
        cmd := "uretprobe:" + userinput + comm
        return cmd

}


func getCountCmd(userinput string) string{

        cmd := "uretprobe:" + userinput + "{ @[pid] = count(); } interval:s:1 { print(@); clear(@); }"
        return cmd

}

func getFreqCmd(userinput string) string{

        cmd := "uprobe:" + userinput + "{ @start = nsecs; } uretprobe:" + userinput + "/@start/ { @time = (nsecs-@start)/1000); print(@time); delete(@start); }"
        return cmd

}

/*
func main () {

	ok := getUretCmd("/ebpfKit/main:enqeue", []string{"ount", "hello"})
	fmt.Println("Result:\n",ok[0])
	fmt.Println("Result:\n",ok[1])
	fmt.Println("Result:\n",ok[2])
}*/
