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
        tests := []struct {
                name     string
                args     []string
                expected string
        }{
                {
                        "specify indent text",
                        []string{"cmd", "-t", "text"},
                        "text",
                },
                {
                        "specify indent text that contains unprintable character",
                        []string{"cmd", "-t", "\t"},
                        "\t",
                },
                {
                        "specify indent text that exceeds maximum length",
                        []string{"cmd", "-t", "0123456789X"},
                        "0123456789",
                },
        }
        for _, tc := range tests {
                t.Run(tc.name, func(t *testing.T) {
                        oldArgs := os.Args
                        defer func() { os.Args = oldArgs }()
                        os.Args = tc.args
                        flag.Parse()
                        if got := resolveIndent(); got != tc.expected {
                                t.Errorf("expected '%s', but got '%s'", tc.expected, got)
                        }
                })
        }
}
