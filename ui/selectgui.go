package ui


func getProbeNames() []string {

        pn := []string{"uretprobe","tcptracer", "tcpconnect", "tcpaccept", "tcplife", "execsnoop", "biosnoop", "cachestat", "All TCP Probes", "All Probes"}
        return pn

}


func getTcpProbeNames() []string {

        pn := []string{"tcptracer", "tcpconnect", "tcpaccept"}
        return pn

}

func getUProbeFuncType() []string {

        pt := []string{"Retval","Count", "Frequency", "All"}

        return pt

}


