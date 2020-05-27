package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    var s,ip string
    fmt.Printf("nhap hostname:")
    fmt.Scanf("%s",&s)
    fmt.Printf("nhap ip:"
    fmt.Scanf("%s",&ip)

    file, err := os.OpenFile("$", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    defer file.Close()
    if _, err := file.WriteString("        hostname:'" +s +"'"); err != nil {
        log.Fatal(err)
    }

    if _, err := file.WriteString("\n        ip:'" +ip+"'" ); err != nil {
        log.Fatal(err)
    }
    
    //Print the contents of the file
    data, err := ioutil.ReadFile("temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))
}
