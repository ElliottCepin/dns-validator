package spf

import (
	"testing"
	"github.com/miekg/dns"
)

func TestNoSPF(t *testing.T) {
	// txt records exist
	// no spf
	// should return none
	record, _ := dns.NewRR(`dmarc.example.com.       TXT     "v=DMARC1; p=quarantine; rua=mailto:dmarc@example.com"`)
	records := []string{record.String()}
	if s := checkHost(records); s != "none" {
		t.Errorf("Expected 'none', got '%s'", s)
	}

}

func TestResolutionError (t *testing.T) {
	// cannot resolve host + rType
	// lookup error : *dns.Msg
}

func TestNoAnswer (t *testing.T) {
	// lookup error : no answer

}
