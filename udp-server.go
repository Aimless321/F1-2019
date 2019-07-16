package f12019

import (
	"log"
	"net"
)

func startUdpServer() {
	//Start udp server
	pc, err := net.ListenPacket("udp", ":20777")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		//Largest packet is 1347 bytes
		buf := make([]byte, 1400)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}

		go serve(pc, addr, buf[:n])
	}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	handleData(buf)
}