package main

import (
        "os"
        "flag"
        "io/ioutil"
        "encoding/json"
)

func main() {
        flag.Parse()
        stat, _ := os.Stdin.Stat()
        var s string
        if stat.Mode()&os.ModeCharDevice == 0 {
                if b, err := ioutil.ReadAll(os.Stdin); err == nil {
                        s = string(b)
                }
        } else {
                args := flag.Args()
                if len(args) != 0 {
                        s = args[0]
                }
        }

        h := json.RawMessage(s)
        b, _ := json.MarshalIndent(&h, "", "    ")
        os.Stdout.Write(b)
}
