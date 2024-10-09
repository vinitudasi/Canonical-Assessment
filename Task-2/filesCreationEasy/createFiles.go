package main

import (
    "crypto/rand"
    "fmt"
    "os"
)

func createFile(fileName string, size int64) error {
    file, err := os.Create(fileName)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    if size > 0 {
        data := make([]byte, size)
        if _, err := rand.Read(data); err != nil {
            return fmt.Errorf("failed to generate random data: %w", err)
        }

        if _, err := file.Write(data); err != nil {
            return fmt.Errorf("failed to write data: %w", err)
        }
    }
    return nil
}

func createEmptyFile(fileName string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return fmt.Errorf("failed to create empty file: %w", err)
    }
    defer file.Close()
    return nil
}

func createReadOnlyFile(fileName string, size int64) error {
    if err := createFile(fileName, size); err != nil {
        return err
    }
    return os.Chmod(fileName, 0400)
}

func main() {
    if err := createFile("randomfile.txt", 1024); err != nil {
        fmt.Printf("Error creating randomfile.txt: %v\n", err)
    } else {
        fmt.Println("Created randomfile.txt (1 KB)")
    }

    if err := createFile("largefile.txt", 10*1024*1024); err != nil {
        fmt.Printf("Error creating largefile.txt: %v\n", err)
    } else {
        fmt.Println("Created largefile.txt (10 MB)")
    }

    if err := createReadOnlyFile("readonlyfile.txt", 1024); err != nil {
        fmt.Printf("Error creating readonlyfile.txt: %v\n", err)
    } else {
        fmt.Println("Created readonlyfile.txt (1 KB, read-only)")
    }

    if err := createEmptyFile("emptyfile.txt"); err != nil {
        fmt.Printf("Error creating emptyfile.txt: %v\n", err)
    } else {
        fmt.Println("Created emptyfile.txt (empty)")
    }
}
