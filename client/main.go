package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "time"
)

func main() {
 for {
  resp, err := http.Get("http://localhost:8080/hello")
  if err != nil {
   fmt.Println("Ошибка при выполнении запроса:", err)
  } else {
   // Чтение ответа
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
    fmt.Println("Ошибка при чтении ответа:", err)
   } else {
    fmt.Printf("Ответ сервера: %s\n", string(body))
   }
   resp.Body.Close()
  }

  // Ждем 1 секунду перед следующим запросом
  time.Sleep(1 * time.Second)
 }
}