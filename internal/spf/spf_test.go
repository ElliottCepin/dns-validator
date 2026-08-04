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

func TestParseSPF(t *testing.T) {
	record, _ := dns.NewRR(`example.com.       TXT     "v=SPF1 +mx -all"`)
	parsed := parseRecord(record.(*dns.TXT).Txt)

	if parsed[0] != "v=SPF1" {
		t.Errorf("Expected 'v=SPF1', got '%s'", parsed[0])
	}

	if len(parsed) != 3 {
		t.Errorf("Expected length 3, got length %v", len(parsed))
		t.Logf("%v", parsed)
	}

}

func TestResolutionError (t *testing.T) {
	// cannot resolve host + rType
	// lookup error : *dns.Msg
}

func TestNoAnswer (t *testing.T) {
	// lookup error : no answer

}
