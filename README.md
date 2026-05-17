
<img width="500" height="500" alt="1779015396637_df86a7be" src="https://github.com/user-attachments/assets/9c218758-67bf-4abe-b05a-4ab4217df3fc" />

<br>

&rarr; Ping-Pil0t 

A simple Command and Control (C2) framework that leverages the ICMP protocol for communication, implemented in GoLang.

---

## **Features** :

- **ICMP Communication:** Uses ICMP echo requests and replies for covert communication.

- **Simple and Lightweight:** Designed to be minimalistic and easy to use.
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

This tool is intended for educational and research purposes only.
