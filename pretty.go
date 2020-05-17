package main

import (
        "fmt"
        "strings"
        "strconv"
        "os"
        "flag"
        "io/ioutil"
        "encoding/json"
)

const (
        indentMaxLength = 10
)

var (
        i      = flag.Int("i", 4, "")
        t      = flag.String("t", "", "")
)

var usage = `Usage:  pretty [OPTIONS] <JSON STRING>
        curl -s http://date.jsontest.com | pretty [OPTIONS]

Options:
  -i  Set indentation level to n spaces. this number is capped at 10 (if it's greater,
      the value is just 10). Values less than 1 indicate that no space should be used.
      (default 4)
  -t  specify an indentation by text (the first 10 characters of the text, if it's longer than that).
`

func main() {
        flag.Usage = func() {
                fmt.Fprint(os.Stderr, usage)
        }
        flag.Parse()
        var s string
        if flag.NArg() != 0 {
                s = flag.Args()[0]
        }
        stat, _ := os.Stdin.Stat()
        if stat.Mode()&os.ModeCharDevice == 0 {
                if b, err := ioutil.ReadAll(os.Stdin); err == nil {
                        s = string(b)
                }
        }
        if len(s) == 0 {
                flag.Usage()
                os.Exit(2)
        }
        h := json.RawMessage(s)
        b, _ := json.MarshalIndent(&h, "", resolveIndent())
        os.Stdout.Write(b)
}

func resolveIndent() string {
        if s := resolveIndentFromText(); len(s) != 0 {
                return s
        }

        i := *i
        if i > indentMaxLength {
                i = indentMaxLength
        }

        var b strings.Builder
        for n := 0; n < i; n++ {
                b.WriteString(" ")
        }

        return b.String()
}

func resolveIndentFromText() string {
        t := *t
        if len(t) > indentMaxLength {
                t = t[:indentMaxLength]
        }
        s, err := strconv.Unquote(`"`+t+`"`)
        if err != nil {
                return ""
        }
        return s
}
