package scan

import (
	"encoding/json"
	"fmt"
	"github.com/praetorian-inc/fingerprintx/pkg/plugins"
	"log"
	"net/netip"
	"testing"
	"time"
)

func TestConfig_SimpleScanTarget(t *testing.T) {
	fxConfig := Config{
		DefaultTimeout: time.Duration(2) * time.Second,
		FastMode:       false,
		Verbose:        true,
		UDP:            false,
		Proxy:          "",
	}

	testTargets := make(map[string][]uint16)
	testTargets["160.20.55.34"] = []uint16{33890}
	testTargets["200.229.30.193"] = []uint16{3389}
	testTargets["185.216.178.7"] = []uint16{3306}
	testTargets["210.243.16.155"] = []uint16{3307} //mysql; test for proxy fail,why?

	// create a target list to scan
	for ip, ports := range testTargets {
		for _, port := range ports {
			target := plugins.Target{
				Address: netip.AddrPortFrom(netip.MustParseAddr(ip), port),
				Host:    ip,
			}
			targets := make([]plugins.Target, 0)
			targets = append(targets, target)

			// run the scan
			fmt.Printf("Scanning %s:%d\n", ip, port)
			results, err := ScanTargets(targets, fxConfig)
			if err != nil {
				log.Fatalf("error: %s\n", err)
			}

			// process the results
			for _, result := range results {
				fmt.Printf("%s:%d (%s/%s)\n", result.Host, result.Port, result.Transport, result.Protocol)
				data, jerr := json.Marshal(result)
				if jerr != nil {
					continue
				}
				log.Println(string(data))
			}
		}
	}
}
