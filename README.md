
&rarr; Ping-Pil0t 

A simple Command and Control (C2) framework that leverages the ICMP protocol for communication, implemented in GoLang.

---

## **Features** :

- **ICMP Communication:** Uses ICMP echo requests and replies for covert communication.

- **Simple and Lightweight:** Designed to be minimalistic and easy to use.

- **Command Execution:** Allows remote execution of commands on target systems.
---

## **Usage** :

&nbsp; **Build** :

```
go build -o pingpil0t main.go
```

&nbsp; **Initiate Server** :

```
./pingpilot server --iface <interface_name>
```

&nbsp; **Initiate Client** :

```
./pingpilot client ---iface <interface_name> [--size <chunk_size>]
```

---

## **Disclaimer** :

This tool is intended for educational and research purposes only. Use it responsibly and only on systems you own.
