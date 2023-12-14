package utils

import (
  "net"
  "fmt"
  "strings"
)

// func to resolve interface name to ip addr
func ResolveInterfaceIP(iface string) (ifaceAddr string) {
  iFace, err := net.InterfaceByName(iface)
  if err != nil {
    panic(err)
  }
  addr, err := iFace.Addrs()
  if err != nil {
    fmt.Println(err)
  }
  ifaceAddr = strings.Split(addr[0].String(), "/")[0]
  return ifaceAddr
}

