// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// func process1(data interface{}) {
//     for i := 1; i <= 4; i++ {
//         time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
//         fmt.Println("Bisa 1", "Bisa 2", "Bisa 3" , data)
//     }
// }

// func process2(data interface{}) {
//     for i := 1; i <= 4; i++ {
//         time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
//         fmt.Println("Coba 1", "Coba 2", "Coba 3", data)
//     }
// }

// func main() {
//     rand.Seed(time.Now().UnixNano())

//     for i := 1; i <= 4; i++ {
//         go process1(fmt.Sprintf(" %d", i))
//         go process2(fmt.Sprintf(" %d", i))
//     }

//     time.Sleep(5 * time.Second)
// }
