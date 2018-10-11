package main

import(
  "encoding/json"
  "fmt"
)

type App struct {
    Id string `json:"id, omitempty"`
    Title string `json:"-"`
    Age int `json:"age"`
}

func main(){

  data := []byte(`[{ "id": "k34rAT4", "title": "My Awesome App", "age": 23}, { "id": "k34rAT1", "title": "My Awesome App", "age": 31 }, { "id": "k00Rat4", "title": "My Teritary App", "age": 0 } ] `)

  var app []App
  err := json.Unmarshal(data, &app)
  fmt.Printf("%+v", app)

  if err != nil {
      fmt.Printf("%v", app[0])
}

  fmt.Println("*****")
  // data, err = json.Marshal(app)
  // fmt.Printf("%s", data)
  fmt.Println()
  fmt.Println()

for _, val := range app {
  fmt.Println(val)
  
}

}
