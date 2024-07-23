package issue

import (
    "bytes"
    "compress/zlib"
    "crypto/sha1"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
)

type Issue struct {
    Description string `json:"description"`
    CreatedAt   string `json:"created_at"`
    Tasks       []Task `json:"tasks"`
}

type Task struct {
    Description string    `json:"description"`
    LogTimes    []LogTime `json:"log_times"`
}

type LogTime struct {
    Start time.Time `json:"start"`
    Stop  time.Time `json:"stop"`
}

var objectsDir = filepath.Join(".tusk", "objects")
var activeIssueFile = filepath.Join(".tusk", "active_issue")

func Add(description string) {
    issue := Issue{
        Description: description,
        CreatedAt:   time.Now().Format(time.RFC3339),
        Tasks:       []Task{},
    }
    issueHash := hashIssue(issue)
    saveIssue(issueHash, issue)
    fmt.Println("Issue added with hash:", issueHash)
}

func List() {
    files, err := ioutil.ReadDir(objectsDir)
    if err != nil {
        fmt.Println("Error reading objects directory:", err)
        return
    }

    for _, file := range files {
        issue := loadIssue(file.Name())
        fmt.Printf("Hash: %s, Description: %s, Created At: %s\n", file.Name(), issue.Description, issue.CreatedAt)
    }
}

func Switch(hash string) {
    if _, err := os.Stat(filepath.Join(objectsDir, hash)); os.IsNotExist(err) {
        fmt.Println("Issue hash not found:", hash)
        return
    }
    ioutil.WriteFile(activeIssueFile, []byte(hash), 0644)
    fmt.Println("Switched active issue to:", hash)
}

func GetActiveIssue() string {
    data, err := ioutil.ReadFile(activeIssueFile)
    if err != nil {
        return ""
    }
    return string(data)
}

func hashIssue(issue Issue) string {
    hasher := sha1.New()
    hasher.Write([]byte(issue.Description + issue.CreatedAt))
    return hex.EncodeToString(hasher.Sum(nil))
}

func saveIssue(hash string, issue Issue) {
    data, _ := json.Marshal(issue)
    compressedData := compressData(data)
    os.MkdirAll(objectsDir, os.ModePerm)
    ioutil.WriteFile(filepath.Join(objectsDir, hash), compressedData, 0644)
}

func loadIssue(hash string) Issue {
    var issue Issue
    data, err := ioutil.ReadFile(filepath.Join(objectsDir, hash))
    if err == nil {
        decompressedData := decompressData(data)
        json.Unmarshal(decompressedData, &issue)
    }
    return issue
}

func compressData(data []byte) []byte {
    var compressedData bytes.Buffer
    writer := zlib.NewWriter(&compressedData)
    writer.Write(data)
    writer.Close()
    return compressedData.Bytes()
}

func decompressData(data []byte) []byte {
    var decompressedData bytes.Buffer
    reader, err := zlib.NewReader(bytes.NewReader(data))
    if err != nil {
        return nil
    }
    io.Copy(&decompressedData, reader)
    reader.Close()
    return decompressedData.Bytes()
}

