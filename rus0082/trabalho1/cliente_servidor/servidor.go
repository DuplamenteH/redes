/*****************************************************************************
 * servidor.go                                                                 
 * Nome: Carlos Matheus Rodrigues Martins
 * Matrícula: 403481
 * Nome: Lauana Maria Cartaxo de Oliveira
 * Matrícula: 399428
 *****************************************************************************/

package main

import (
  "fmt"
  "io/ioutil"
  "net"
  "os"
  "log"
)

const RECV_BUFFER_SIZE = 2048

/* TODO: server()
 * Abra socket e espere o cliente conectar
 * Imprima a mensagem recebida em stdout
*/
func handlerServer(link net.Conn){
  info, _:=ioutil.ReadAll(link)
  infoString := string(info)
  fmt.Print(infoString)
}


func server(server_port string) {
  service := ":"+server_port
  addr,_:= net.ResolveTCPAddr("tcp", service)
  l,err := net.ListenTCP("tcp", addr)
  for{
      connection,_:= l.AcceptTCP()
      if err != nil{
        continue
      }

     handlerServer(connection)
  }
  l.Close()

}


// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
  if len(os.Args) != 2 {
    log.Fatal("Uso: ./servidor [porta servidor]")
  }
  server_port := os.Args[1]
  server(server_port)
}
