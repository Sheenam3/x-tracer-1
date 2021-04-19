package getval


var NAMESPACE string = "default"


func GetProbeNames() []string {

        pn := []string{"tcptracer", "tcpconnect", "tcpaccept", "tcplife", "execsnoop", "biosnoop", "cachestat", "All TCP Probes", "All Probes"}
        return pn

}

func GetTcpProbeNames() []string {

	pn := []string{"tcptracer", "tcpconnect", "tcpaccept"}
	return pn

}
