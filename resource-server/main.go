package main
import (
  "net/http"
)

//main function (entry point to the app)
func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
          w.Write([]byte("Hello World"))  
      })
  http.ListenAndServe(":8000", nil)
}
