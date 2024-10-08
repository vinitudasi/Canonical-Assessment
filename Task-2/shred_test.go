package main

import (
    "os"
    "testing"
    "path/filepath"
)

func TestShred_Success(t *testing.T) {
    randomFile := "randomfile.txt"

    err := Shred(randomFile)
    if err != nil {
        t.Fatalf("No success in shredding, got error: %v", err)
    }

    if _, err := os.Stat(randomFile); !os.IsNotExist(err) {
        t.Errorf("File still in directory!")
    }
}

func TestShred_FileNotFound(t *testing.T) {
    fakeFile := filepath.Join(os.TempDir(), "nonexistentfile.txt")

    err := Shred(fakeFile)
    if err == nil {
        t.Errorf("hmm, nonexistent file e")
    }
}

func TestShred_PermissionDenied(t *testing.T) {
    readOnlyFile := "readonlyfile.txt"

    err := os.Chmod(readOnlyFile, 0400)
    if err != nil {
        t.Fatalf("Couldn't change file permissions: %v", err)
    }

    err = Shred(readOnlyFile)
    if err == nil {
        t.Errorf("Expected a permission error, but got no error")
    }

    os.Chmod(readOnlyFile, 0644)
    os.Remove(readOnlyFile)
}

func TestShred_EmptyFile(t *testing.T) {
    emptyFile := "emptyfile.txt"

    err := Shred(emptyFile)
    if err != nil {
        t.Fatalf("Expected no error for empty file, got: %v", err)
    }

    if _, err := os.Stat(emptyFile); !os.IsNotExist(err) {
        t.Errorf("Expected the empty file to be deleted, but it still exists")
    }
}

func TestShred_LargeFile(t *testing.T) {
    largeFile := "largefile.txt"

    err := Shred(largeFile)
    if err != nil {
        t.Fatalf("Expected no error for large file, but got: %v", err)
    }

    if _, err := os.Stat(largeFile); !os.IsNotExist(err) {
        t.Errorf("Expected large file to be deleted, but it still exists")
    }
}
