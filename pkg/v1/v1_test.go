package v1

import "testing"

var table map[string]uint32 = map[string]uint32{
	// http://www.webdnstools.com/dnstools/ipcalc
	"127.0.0.1":       uint32(2130706433),
	"192.168.0.1":     uint32(3232235521),
	"255.255.255.255": uint32(4294967295),
}

func TestIPv4toInt64(t *testing.T) {
	for input, output := range table {
		res, err := IPv4toUint32(input)
		if err != nil {
			t.Fatal(err)
		}
		if res != output {
			t.Errorf("expected %d, got %d", output, res)
		}
	}
}

func TestUint32toIPv4(t *testing.T) {
	for output, input := range table {
		res, err := Uint32toIPv4(input)
		if err != nil {
			t.Fatal(err)
		}
		if output != res {
			t.Errorf("expected %s, got %s", output, res)
		}
	}
}
