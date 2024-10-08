package main

import (
    "crypto/rand"
    "fmt"
    "os"
)

func Shred(path string) error {
    file, err := os.OpenFile(path, os.O_WRONLY, 0)
    if err != nil {
        return err
    }
    defer file.Close()

    info, err := file.Stat()
    if err != nil {
        return err
    }
    size := info.Size()

    // 3 times random data
    for i := 0; i < 3; i++ {
        if _, err := file.Seek(0, 0); err != nil {
            return err
        }
        randomData := make([]byte, size)
        if _, err := rand.Read(randomData); err != nil {
            return err
        }
        if _, err := file.Write(randomData); err != nil {
            return err
        }
        fmt.Printf("Overwrite pass %d completed.\n", i+1)
    }

   
    file.Close()

    // Delete file
    if err := os.Remove(path); err != nil {
        return err
    }
    fmt.Println("Mission Accomplished")
    return nil
}


