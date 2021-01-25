/*****************************************************************************
 * cliente.go
 * Nome: Carlos Matheus Rodrigues Martins
 * Matrícula: 403481
 * Nome: Lauana Maria Cartaxo de Oliveira
 * Matrícula: 399428
 *****************************************************************************/

package main

import (
  "io/ioutil"
  "log"
  "net"
  "os"
)

const SEND_BUFFER_SIZE = 2048

/* TODO: client()
 * Abrir socket e enviar mensagem de stdin.
*/

func action(link net.Conn){
  data,_:= ioutil.ReadAll(os.Stdin)
  link.Write(data)
  link.Close()

}
func client(server_ip string, server_port string) {

  service:= server_ip+":"+server_port

  addr,_:= net.ResolveTCPAddr("tcp", service)

  link,_:= net.DialTCP("tcp", nil, addr)

  action(link)
}

// Main obtém argumentos da linha de comando e chama função client
func main() {
  if len(os.Args) != 3 {
    log.Fatal("Uso: ./cliente [IP servidor] [porta servidor] < [arquivo mensagem]")
  }
  server_ip := os.Args[1]
  server_port := os.Args[2]
  client(server_ip, server_port)
}
