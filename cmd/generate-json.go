package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/zelenin/go-tdlib/tlparser"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var outputFilePath string

	flag.StringVar(&outputFilePath, "output", "./td_api.json", "json schema file")

	flag.Parse()

	tlFile, err := os.Open("data/td_api.tl")
	if err != nil {
		log.Fatalf("open td_api.tl error: %s", err)
		return
	}
	defer tlFile.Close()

	schema, err := tlparser.Parse(bufio.NewReader(tlFile))
	if err != nil {
		log.Fatalf("schema parse error: %s", err)
		return
	}

	cppFile, err := os.Open("data/Td.cpp")
	if err != nil {
		log.Fatalf("open Td.cpp error: %s", err)
		return
	}
	defer cppFile.Close()

	err = tlparser.ParseCode(bufio.NewReader(cppFile), schema)
	if err != nil {
		log.Fatalf("parse code error: %s", err)
		return
	}

	err = os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm)
	if err != nil {
		log.Fatalf("make dir error: %s", filepath.Dir(outputFilePath))
	}

	file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %s", err)
		return
	}

	data, err := json.MarshalIndent(schema, "", strings.Repeat(" ", 4))
	if err != nil {
		log.Fatalf("json marshal error: %s", err)
		return
	}
	bufio.NewWriter(file).Write(data)
}
