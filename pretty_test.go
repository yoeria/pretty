package main

import (
        "os"
        "flag"
        "testing"
)

func TestResolveIndent(t *testing.T) {
        indent := resolveIndent()
        if indent != "    " {
                t.Errorf("expected 4 spaces, but got '%s'", indent)
        }
}

func TestResolveIdentFromLevel(t *testing.T) {
        oldArgs := os.Args
        defer func() { os.Args = oldArgs }()
        os.Args = []string{"cmd", "-i", "2"}
        flag.Parse()
        indent := resolveIndent()
        if indent != "  " {
                t.Errorf("expected 2 spaces, but got '%s'", indent)
        }
}

func TestResolveIndentExceedsMaximumLevel(t *testing.T) {
        oldArgs := os.Args
        defer func() { os.Args = oldArgs }()
        os.Args = []string{"cmd", "-i", "11"}
        flag.Parse()
        indent := resolveIndent()
        if indent != "          " {
                t.Errorf("expected 10 spaces, but got '%s'", indent)
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
