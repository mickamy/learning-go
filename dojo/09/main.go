package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

type RuneScanner struct {
	reader io.Reader
	err    error
	char   rune
}

func NewScanner(r io.Reader) RuneScanner {
	scanner := RuneScanner{
		reader: r,
	}
	return scanner
}

func (s *RuneScanner) Scan() bool {
	if s.err != nil {
		return false
	}

	buf := [16]byte{}
	n, err := s.reader.Read(buf[:])
	if err != nil {
		s.err = err
		return false
	}

	r, size := utf8.DecodeRune(buf[:n])
	if r == utf8.RuneError {
		s.err = fmt.Errorf("invalid utf-8")
		return false
	}
	s.char = r
	s.reader = io.MultiReader(bytes.NewReader(buf[size:n]), s.reader)

	return true
}

func (s *RuneScanner) Char() rune {
	return s.char
}

func (s *RuneScanner) Err() error {
	return s.err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := NewScanner(reader)
	for s.Scan() {
		fmt.Println(string(s.Char()))
	}
	if s.Err() != nil {
		fmt.Println("err:", s.Err())
		os.Exit(1)
	}
}
