package main

import (
  "net/http"
  "encoding/xml"
)

func setupTwilioListeners() {
  http.HandleFunc("/twiml", twiml)
  http.ListenAndServe(":8080", nil)
}

func twiml(w http.ResponseWriter, r *http.Request) {
  twiml := TwiML{Message: "Hello World!"}
  x, err := xml.Marshal(twiml)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/xml")
  w.Write(x)
}

type TwiML struct {
  XMLName xml.Name `xml:"Response"`
  Message    string `xml:",omitempty"`
}
