package data_test_engine

import "fmt"

// get GTP Header from UE id(function used for multiple attach ue).
func GenerateGtpHeader(gtpTeid int) string {
	var auxGtp string

	// generates some GTP-TEIDs for the RAN-UPF tunnels(uplink) in order to make the GTP header.
	// valorGtp = 1 + 2*(ranUeId-1)
	if gtpTeid < 16 {
		auxGtp = "32ff00340000000" + fmt.Sprintf("%x", gtpTeid) + "00000000"
	} else if gtpTeid < 256 {
		auxGtp = "32ff0034000000" + fmt.Sprintf("%x", gtpTeid) + "00000000"
	} else {
		auxGtp = "32ff003400000" + fmt.Sprintf("%x", gtpTeid) + "00000000"
	}

	return auxGtp
}
