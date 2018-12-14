package main

import (
  "net/http"
  "math/rand"
)

func main() {
  http.HandleFunc("/", avatarCreator)
  http.ListenAndServe(":3000", nil)
}

func avatarCreator(w http.ResponseWriter, r *http.Request) {
  names := [10]string{"Aang", "Katara", "Sokka", "Zuko", "Iroh", "Appa", "Momo", "Pasang", "Tashi", "Yue"}
  name := names[rand.Intn(10)]
  w.Write([]byte(name))
}
