package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/alexmullins/zip"
)

func Submit(filename, kode string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	config, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	code, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data, err := io.ReadAll(config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	config.Close()
	var student Student
	err = json.Unmarshal(data, &student)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("\033[33m[+] Generating report file [ REPORT.txt ]\033[0m\n")
	report, err := os.OpenFile(kode+"REPORT.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer report.Close()
	// Write some text to the file
	report.WriteString(fmt.Sprintf("%s %s %s %s\n%s\n", student.Name, student.NIM, student.Shift, kode, string(code)))
	report.WriteString(fmt.Sprintf("CPU model: %v\nCPU cores: %v\nOS type: %v\nStarted at: %v\nSubmitted at: %v\n", runtime.GOARCH, runtime.NumCPU(), runtime.GOOS, student.Session, time.Now()))
	// Origin stdout
	fmt.Printf("\033[33m[+] Running final test to program [ %s ] \033[0m\n", filename)
	originalStdout := os.Stdout
	// Redirect stdout to the file
	os.Stdout = report
	// Test the program
	Test(filename, kode)
	// Restore stdout to the terminal
	os.Stdout = originalStdout
	// Create zip file
	zipname := fmt.Sprintf("%s_%s_RESULT.zip", student.Name, kode)
	fmt.Printf("\033[33m[+] Generating grading result file [ %s ] \033[0m\n", zipname)
	fzip, err := os.OpenFile(zipname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()

	w, err := zipw.Encrypt(kode+"REPORT.txt", "asprakganteng")
	if err != nil {
		log.Fatal(err)
	}
	rfile, err := os.Open(kode + "REPORT.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	result, err := io.ReadAll(rfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	_, err = io.Copy(w, bytes.NewReader(result))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Flush()
	fmt.Printf("\033[33m[+] Clearing local workspace\033[0m\n")
	pdf, _ := filepath.Glob("*.pdf")
	exe, _ := filepath.Glob("*.exe")
	// json, _ := filepath.Glob("*.json")
	for _, file := range pdf {
		os.Remove(file)
	}
	for _, file := range exe {
		os.Remove(file)
	}
	// for _, file := range json {
	// 	os.Remove(file)
	// }
	fmt.Printf("\033[32m✓ Done \t\033[0m\n")
}
