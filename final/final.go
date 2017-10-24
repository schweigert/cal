package main

import(
  "./rsa"
  "fmt"
)

func main () {
  private, public := rsa.NewCert()
  fmt.Println(public.N)
  fmt.Println(private.N)
}
