/*
Learning Golang
Just Kalkulator Sederhana bjir


Aku Mah Masih Pemula T.T
*/

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
//  "strings"
)

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
      fmt.Println("Opsi Tidak Valid!")
      continue
      }

    // Input Ke-1 Bjir
    fmt.Println("Input Angka Ke-1: ")
    scanner.Scan()
    num1Str := scanner.Text()
    num1, err := strconv.ParseFloat(num1Str, 64)
    if err != nil {
     fmt.Println("Masukin angka jir lah")
      continue
    }

    // Input Ke-2
    fmt.Println("Input Angka Ke-1: ")
    scanner.Scan()
    num2Str := scanner.Text()
    num2, err := strconv.ParseFloat(num2Str, 64)
    if err != nil {
     fmt.Println("Masukin angka jir lah")
      continue
    }

    // Operasi Melalui pilihan dengan float & strings
    var result float64
    var operator string

    switch choice {
    case "1":
      result = num1 + num2
      operator = "+"
    case "2":
      result = num1 - num2
      operator = "-"
    case "3":
      result = num1 * num2
      operator = "*"
    case "4":
    if num2 == 0 {
      fmt.Println("Tidak bisa membagi noll")
      continue
      }
    result = num1 / num2
      operator = "/"
      }

// Test Ombak jir
    fmt.Println("\nHasil: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
    
  }
}
