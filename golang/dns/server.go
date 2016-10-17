//  dns server

package main

import (
	"github.com/miekg/dns"
	"log"
	"net"
	"runtime"
	"time"
)

type conDns struct {
	tenant string
	epg    string
}

type dnsReader struct {
	id string
}

type dnsWriter struct {
	id string
}

func markLog() {
	if pc, _, l, ok := runtime.Caller(1); ok {
		log.Printf("==== marker %+v@:%d ====", runtime.FuncForPC(pc).Name(), l)
	} else {
		log.Printf("======== unknwon ========")
	}
}

func main() {
	cDns := conDns{}
	dReader := dnsReader{}
        dWriter := dnsWriter{}
	// loop thru
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:55134")
	if err != nil {
		log.Printf("failed to resolve udp address")
		return
	}

	dConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Printf("failed to start udp listener")

	}

	s1 := dns.Server{
		Net:               "UDP",
		UDPSize:           1024,
                ReadTimeout:       5 * time.Second,
		PacketConn:        dConn,
		Handler:           &cDns,
		NotifyStartedFunc: func() { log.Printf("started") },
		DecorateReader:    func(dns.Reader) dns.Reader { return &dReader },
                DecorateWriter:    func(dns.Writer) dns.Writer { return &dWriter },
	}

	log.Printf("=== starting ====")
	if err := s1.ActivateAndServe(); err != nil {
		log.Printf("failed to start %s", err)
	}

}

func (cDns *conDns) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	markLog()
        m := &dns.Msg{}
	m.SetReply(r)

	answers := []dns.RR{}

	for _, question := range r.Question {
		switch question.Qtype {
		case dns.TypeA:
		case dns.TypeSRV:
		}
	}

	// If we have no answers, that means we found nothing or didn't get a query
	// we can reply to. Reply with no answers so we ensure the query moves on to
	// the next server.
	if len(answers) == 0 {
		m.SetRcode(r, dns.RcodeSuccess)
		w.WriteMsg(m)
		return
	}

	// Without these the glibc resolver gets very angry.
	m.Authoritative = true
	m.RecursionAvailable = true
	m.Answer = answers

	w.WriteMsg(m)
}

// ReadTCP reads a raw message from a TCP connection. Implementations may alter
// connection properties, for example the read-deadline.
func (dr *dnsReader) ReadTCP(conn net.Conn, timeout time.Duration) ([]byte, error) {
	markLog()
	return []byte{}, nil
}

func dnsReq(q string, t uint16, c uint16) (dns.Msg, error) {
	m := dns.Msg{Question: []dns.Question{{}}}
	m.Question[0] = dns.Question{Name: q, Qtype: t, Qclass: c}
	return m, nil
}

// ReadUDP reads a raw message from a UDP connection. Implementations may alter
// connection properties, for example the read-deadline.
func (dr *dnsReader) ReadUDP(conn *net.UDPConn, timeout time.Duration) ([]byte, *dns.SessionUDP, error) {
        s := &dns.SessionUDP{}

	markLog()
        time.Sleep(timeout)
	m, err := dnsReq("abc.example.", dns.TypeA, dns.ClassINET)
        if err != nil {
                log.Printf("error failed to create dns req")
                return []byte{}, nil, dns.ErrShortRead
        }
        if b, e := m.Pack(); e != nil {
                log.Printf("fail %s", e)
                return b, s, e
        }else {
                 log.Printf("success")
                return b, s, e
        }
}

func (dr *dnsWriter) Write(p []byte) (n int, err error) {
        markLog()

        log.Printf("==write()=== %s", string(p))
        return 0, nil
}
