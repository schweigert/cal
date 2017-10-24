package main

import(
  "./rsa"
  "fmt"
)

func main () {
  private, public := rsa.NewCert()
  fmt.Println(public)
  fmt.Println(private)
}
