// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// func process1(data interface{}, done chan bool) {
//     for i := 0; i < 4; i++ {
//         fmt.Println("Processing data 1:", data)
//         time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
//     }
//     done <- true
// }

// func process2(data interface{}, done chan bool) {
//     for i := 0; i < 4; i++ {
//         fmt.Println("Processing data 2:", data)
//         time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
//     }
//     done <- true
// }

// func main() {
//     done1 := make(chan bool)
//     done2 := make(chan bool)

//     data1 := "Lorem ipsum dolor sit amet"
//     data2 := 42

//     for i := 0; i < 4; i++ {
//         go process1(data1, done1)
//         go process2(data2, done2)
//     }

//     for i := 0; i < 8; i++ {
//         select {
//         case <-done1:
//             fmt.Println("Data 1 processing complete")
//         case <-done2:
//             fmt.Println("Data 2 processing complete")
//         }
//     }
// }




package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func process1(data interface{}, wg *sync.WaitGroup, lock *sync.Mutex) {
    defer wg.Done()
    for i := 0; i < 4; i++ {
        lock.Lock()
        fmt.Println("Processing data 1:", data)
        lock.Unlock()
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
}

func process2(data interface{}, wg *sync.WaitGroup, lock *sync.Mutex) {
    defer wg.Done()
    for i := 0; i < 4; i++ {
        lock.Lock()
        fmt.Println("Processing data 2:", data)
        lock.Unlock()
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
}

func main() {
    var wg sync.WaitGroup
    var lock sync.Mutex

    data1 := "Lorem ipsum dolor sit amet"
    data2 := 42

    for i := 0; i < 4; i++ {
        wg.Add(1)
        go process1(data1, &wg, &lock)
        wg.Add(1)
        go process2(data2, &wg, &lock)
    }

    wg.Wait()
}
