package main

type Todo struct {
    Target    string    `json:"target"`
    Module    string    `json:"module"`
    Command   string    `json:"command"`
}

type Todos []Todo
