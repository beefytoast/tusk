package initproject

import (
    "fmt"
    "os"
    "path/filepath"
)

var tuskDir = ".tusk"
var objectsDir = filepath.Join(tuskDir, "objects")
var activeIssueFile = filepath.Join(tuskDir, "active_issue")

func Init() {
    if _, err := os.Stat(tuskDir); os.IsNotExist(err) {
        os.Mkdir(tuskDir, os.ModePerm)
        os.Mkdir(objectsDir, os.ModePerm)
        file, err := os.Create(activeIssueFile)
        if err == nil {
            file.Close()
        }
        fmt.Println("Initialized empty Tusk project in", tuskDir)
    } else {
        fmt.Println("Tusk project already initialized.")
    }
}

