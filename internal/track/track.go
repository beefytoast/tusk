package track

import (
    "bytes"
    "compress/zlib"
    "encoding/json"
    "fmt"
    "io"
    "os"
    "io/ioutil"
    "path/filepath"
    "time"
    "github.com/beefytoast/tusk/internal/issue"
)

var objectsDir = filepath.Join(".tusk", "objects")
var activeIssueFile = filepath.Join(".tusk", "active_issue")

func Add(description string) {
    issueHash := issue.GetActiveIssue()
    if issueHash == "" {
        fmt.Println("No active issue set. Please switch to an active issue using 'tusk issue --switch <hash>'")
        return
    }
    iss := loadIssue(issueHash)
    newTask := issue.Task{
        Description: description,
        LogTimes:    []issue.LogTime{},
    }
    iss.Tasks = append(iss.Tasks, newTask)
    saveIssue(issueHash, iss)
    fmt.Println("Task added to issue", issueHash)
}

func Start() {
    issueHash := issue.GetActiveIssue()
    if issueHash == "" {
        fmt.Println("No active issue set. Please switch to an active issue using 'tusk issue --switch <hash>'")
        return
    }
    iss := loadIssue(issueHash)
    if len(iss.Tasks) == 0 {
        fmt.Println("No tasks found in issue", issueHash)
        return
    }
    iss.Tasks[len(iss.Tasks)-1].LogTimes = append(iss.Tasks[len(iss.Tasks)-1].LogTimes, issue.LogTime{Start: time.Now()})
    saveIssue(issueHash, iss)
    fmt.Println("Started logging time for the last task in issue", issueHash)
}

func Stop() {
    issueHash := issue.GetActiveIssue()
    if issueHash == "" {
        fmt.Println("No active issue set. Please switch to an active issue using 'tusk issue --switch <hash>'")
        return
    }
    iss := loadIssue(issueHash)
    if len(iss.Tasks) == 0 || len(iss.Tasks[len(iss.Tasks)-1].LogTimes) == 0 {
        fmt.Println("No active time log found for the last task in issue", issueHash)
        return
    }
    iss.Tasks[len(iss.Tasks)-1].LogTimes[len(iss.Tasks[len(iss.Tasks)-1].LogTimes)-1].Stop = time.Now()
    saveIssue(issueHash, iss)
    fmt.Println("Stopped logging time for the last task in issue", issueHash)
}

func loadIssue(hash string) issue.Issue {
    var iss issue.Issue
    data, err := ioutil.ReadFile(filepath.Join(objectsDir, hash))
    if err == nil {
        decompressedData := decompressData(data)
        json.Unmarshal(decompressedData, &iss)
    }
    return iss
}

func saveIssue(hash string, iss issue.Issue) {
    data, _ := json.Marshal(iss)
    compressedData := compressData(data)
    os.MkdirAll(objectsDir, os.ModePerm)
    ioutil.WriteFile(filepath.Join(objectsDir, hash), compressedData, 0644)
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

