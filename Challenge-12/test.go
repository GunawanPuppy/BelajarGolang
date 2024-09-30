package main
// import (
// 	"fmt"
// 	"encoding/json"
// )
// type User struct {
// 	//kasih alias
// 	FullName string `json:"Name"`
// 	Age int
// }

// func main() {
// 	//data json
//   var jsonString = `{"Name": "john doe", "Age": 27}`

// 	//mengkonversi jsonString
//   var jsonData = []byte(jsonString)
 

//   var data User
//   //decode jsonData dan dimasukkan ke variabel penampung data yg memiliki tipe User
//   var err = json.Unmarshal(jsonData, &data)
//   if err != nil {
//     fmt.Println(err.Error())
//     return
//   }
  
//   fmt.Println("user :", data.FullName)
//   fmt.Println("age  :", data.Age)
// }