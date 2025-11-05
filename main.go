/*
Learning Golang

Just Kalkulator Sederhana bjir


*/

package main

import {
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
}

func main() {
 scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Println("===Kalkulator Sederhana===")
    fmt.Println("\n1. Penjumlahan")
    fmt.Println("2. Pengurangan")
    fmt.Println("3. Perkalian")
    fmt.Println("4. Pembagian")
    fmt.Println("0. Keluar Dari Program")
    fmt.Println("Input Operasi (1-5): ")

    // Melakukan scaning pada input
    scanner.Scan()
    choice := scanner.Text()

    // Jika Memilih 0 = exit
    if choice == "0" {
      fmt.Println("Terimakasih sudah mencoba kalkulator by Pemula")
      break
  }

    // Jika Memilih bukan 0 - 4 Maka Invalid 
    if choice < "0" || choice > "4" {


    
  }
