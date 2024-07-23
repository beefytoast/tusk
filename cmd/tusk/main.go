package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/beefytoast/tusk/internal/issue"
    "github.com/beefytoast/tusk/internal/track"
    "github.com/beefytoast/tusk/internal/init"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("expected 'issue' or 'track' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "issue":
        issueCmd := flag.NewFlagSet("issue", flag.ExitOnError)
        add := issueCmd.Bool("add", false, "Add a new issue")
        desc := issueCmd.String("desc", "", "Description of the issue")
        list := issueCmd.Bool("list", false, "List all issues")
        switchCmd := issueCmd.String("switch", "", "Switch active issue")

        issueCmd.Parse(os.Args[2:])
        if *add {
            issue.Add(*desc)
        } else if *list {
            issue.List()
        } else if *switchCmd != "" {
            issue.Switch(*switchCmd)
        }

    case "track":
        trackCmd := flag.NewFlagSet("track", flag.ExitOnError)
        add := trackCmd.Bool("add", false, "Add a new task")
        desc := trackCmd.String("desc", "", "Description of the task")
        start := trackCmd.Bool("start", false, "Start logging time for task")
        stop := trackCmd.Bool("stop", false, "Stop logging time for task")

        trackCmd.Parse(os.Args[2:])
        if *add {
            track.Add(*desc)
        } else if *start {
            track.Start()
        } else if *stop {
            track.Stop()
        }
    
    case "init":
        initCmd := flag.NewFlagSet("init", flag.ExitOnError)
        initCmd.Parse(os.Args[2:])
        initproject.Init()

    default:
        fmt.Println("expected 'issue' or 'track' subcommands")
        os.Exit(1)
    }
}

