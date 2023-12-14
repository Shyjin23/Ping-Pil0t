package main 

// echo "1" > /proc/sys/net/ipv4/icmp_echo_ignore_all

import (
     "os"
     "strconv"
  C2 "PingPil0t/handlers"
     "github.com/urfave/cli/v2"
)

// Initiate Server 
func PingPil0tServer(c *cli.Context) error {
  iface := c.String("iface") 
  C2.StartServer(iface) 
  return nil
}

// Initiate Client 
func PingPil0tClient(c *cli.Context) error {
  iface := c.String("iface") 
  chunkSize, _ := strconv.Atoi(c.String("chunkSize")) 
  C2.StartClient(iface, chunkSize)
  return nil
}

// Driver code
func main() {
  app := &cli.App {
    Name: "PingPil0t",
    Usage: "An ICMP Tunnel In Go!", 
    Commands: []*cli.Command {
      {
        Name: "server",
        Usage: "Use PingPil0t in server mode.",
        Flags: []cli.Flag {
          &cli.StringFlag {
            Name: "iface",
            Aliases: []string{"i"},
            Usage: "Interface to use.",
            Value: "eth0", 
          }, 
        },
        Action: PingPil0tServer,
      },
      {
        Name: "client",
        Usage: "Use PingPil0t in client mode.",
        Flags: []cli.Flag {
          &cli.StringFlag {
            Name: "iface", 
            Aliases: []string{"i"},
            Usage: "Interface to use.",
            Value: "eth0", 
          },
          &cli.StringFlag {
            Name: "chunkSize", 
            Aliases: []string{"s"},
            Usage: "Max size of payload.",
            Value: "1024",
          },
        },
        Action: PingPil0tClient,
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    panic(err)
  }
}

