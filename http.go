package main

import (
    "io"
    "log"
    "fmt"
    "net/http"
    "flag"
    "reflect"
)

func main() {
    numbPrt := flag.Int("port", 8080, "port number")
    flag.Parse()

    http.HandleFunc("/", ExampleHandler)

    log.Println("** Service Started on Port ", *numbPrt, " **")

    // Use ListenAndServeTLS() instead of ListenAndServe() which accepts two extra parameters. 
    // We need to specify both the certificate file and the key file (which we've named 
    // https-server.crt and https-server.key).
    //err := http.ListenAndServeTLS(":4567", "https-server.crt", "https-server.key", nil);
    err := http.ListenAndServe(fmt.Sprintf(":%+v", *numbPrt), nil)
    if err != nil {
        log.Fatal(err)
    }
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, string(r.Method))


    log.Println(w)
    log.Println(r)
    log.Println(r.Method)
    log.Println(reflect.ValueOf(r).Elem())
    err := r.ParseForm()
    if err != nil {
        // in case of any error
        return
    }
    val := reflect.ValueOf(r).Elem()
    for i:=0; i<val.NumField();i++{
        //fmt.Println(val.Type().Field(i).Value)
        // fmt.Print(val.Type().Field(i).Name)
        fmt.Printf("FIELD:  %+v     VALUE:  %+v\n", val.Type().Field(i).Name, val.Field(i))
        io.WriteString(w, fmt.Sprintf("FIELD:  %+v     VALUE:  %+v\n", val.Type().Field(i).Name, val.Field(i)))
    }
    //for key, values := range r {   // range over map
    //    fmt.Println(key)
        //for _, value := range values {    // range over []string
        //    fmt.Println(key, value)
        //}
    //}
}
