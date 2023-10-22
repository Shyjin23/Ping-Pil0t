#include <stdio.h>
#include <stdlib.h>
#include <unistd.h> 
#include <netinet/ip.h>
#include <netinet/tcp.h> 
#include <sys/socket.h>
#include <arpa/inet.h> 

int main() {
  int sock_raw;
  printf("[+] Creating a raw socket.. ");
  sock_raw = socket(AF_INET, SOCK_RAW, IPPROTO_TCP);
  if(sock_raw == -1) {
    printf("\n[-] Error while creating socket!\n"); 
    return 1;
  }
  printf(" => 0k!\n"); 
  close(sock_raw);
  return 0; 
}
