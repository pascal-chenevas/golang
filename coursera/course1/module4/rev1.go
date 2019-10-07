package main

import (
    "fmt"
    "encoding/json"
    "os"
    "bufio"
    "strings"
)

func main() {
    m := make(map[string]string)
    name := readInput("Enter a name:")
    address := readInput("Enter an address:")
    m["address"] =  address
    m["name"] = name
    jsonObject, err := json.Marshal(m)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(jsonObject))
    }
}
func readInput(prompt string) string {
    fmt.Println(prompt)
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Invalid input")
    }
    return strings.TrimSpace(input)
}
