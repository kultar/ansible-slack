package main

import (
    "fmt"
    "os/exec"
    "log"
    "bytes"
    )

var todos Todos

func ExecuteAnsible(t Todo) (s string) {
    target := t.Target
    module := t.Module
    command := t.Command
    cmd := exec.Command("ansible",target,"-m",module,"-a",command)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
            log.Fatal(err)
    }
    fmt.Printf("Executed: %q\n",out.String())
    return out.String()
}

