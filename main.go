func firstEndPointHandler(w http.ResponseWriter, r *http.Request) {
   message := "this is the first endpoint"
   _, err := w.Write([]byte(message))
   if err != nil {
      log.Fatal(err)
   }
}


func secondEndPointHandler(w http.ResponseWriter, r *http.Request) {
   message := "second endpoint"
   _, err := w.Write([]byte(message))
   if err != nil {
      log.Fatal(err)
   }
}


func thirdEndPointHandler(w http.ResponseWriter, r *http.Request) {
   message := "second endpoint"
   _, err := w.Write([]byte(message))
   if err != nil {
      log.Fatal(err)
   }
}


func main() {
   http.HandleFunc("/first", firstEndPointHandler)
   http.HandleFunc("/second", secondEndPointHandler)
   http.HandleFunc("/third", thirdEndPointHandler)
   err := http.ListenAndServe(":8083", nil)
   log.Fatal(err)
}
