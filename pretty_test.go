package main

import (
        "os"
        "flag"
        "testing"
)

func TestResolveIndent(t *testing.T) {
        tests := []struct {
                name     string
                args     []string
                expected string
        }{
                {
                        "default indent level",
                        []string{"cmd"},
                        "    ",
                },
                {
                        "specify indent level",
                        []string{"cmd", "-i", "2"},
                        "  ",
                },
                {
                        "specify indent level that exceeds maximum number",
                        []string{"cmd", "-i", "11"},
                        "          ",
                },
        }
        for _, tc := range tests {
                t.Run(tc.name, func(t *testing.T) {
                        oldArgs := os.Args
                        defer func() { os.Args = oldArgs }()
                        os.Args = tc.args
                        flag.Parse()
                        if got := resolveIndent(); got != tc.expected {
                                t.Errorf("expected %d spaces, but got %d", len(tc.expected), len(got))
                        }
                })
        }
}

func TestResolveIndentFromText(t *testing.T) {
        oldArgs := os.Args
        defer func() { os.Args = oldArgs }()
        os.Args = []string{"cmd", "-t", "text"}
        flag.Parse()
        indent := resolveIndent()
        if indent != "text" {
                t.Errorf("expected 'test', but got '%s'", indent)
        }
}

func TestResolveIndentFromUnprintableText(t *testing.T) {
        oldArgs := os.Args
        defer func() { os.Args = oldArgs }()
        os.Args = []string{"cmd", "-t", "\t"}
        flag.Parse()
        indent := resolveIndent()
        if indent != "\t" {
                t.Errorf("expected <tab>, but got '%s'", indent)
        }
}

func TestResolveIndentExceedsMaximumText(t *testing.T) {
        oldArgs := os.Args
        defer func() { os.Args = oldArgs }()
        os.Args = []string{"cmd", "-t", "0123456789X"}
        flag.Parse()
        indent := resolveIndent()
        if indent != "0123456789" {
                t.Errorf("expected '0123456789', but got '%s'", indent)
        }
}
