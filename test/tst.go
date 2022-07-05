package main


import (
       "fmt"
       "os"
       cap "github.com/maccoy54/EchaQueue" 
       )


func main() {

        // Parametres
        if len(os.Args) != 3 {
            fmt.Printf("usage : %s tranche entre \"AAAA-MM-DD\" et \"AAAA-MM-DD\"",os.Args[0])
            os.Exit(1)
        }
        q := cap.CreateQueue(os.Args[1],os.Args[2])
        var k string
        depile := len(q)
        for depile > 0 {
                k,depile = cap.GetKey() 
                fmt.Println(k)
                //depile = len(cap.Cle)
        }
}

     
