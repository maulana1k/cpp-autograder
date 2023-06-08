package main

import (
	"cpp_grader/utils"
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "github.com/akavel/rsrc/rsrc"
)

//go:embed paket/**/*.pdf
var paket embed.FS

//go:embed paket/**/*.cpp
var template embed.FS

func main() {
	// Define flags
	// rsrc.Embed("main.syso","amd64","main.exe","icon.ico")
	test := flag.NewFlagSet("test", flag.ExitOnError)
	submit := flag.NewFlagSet("submit", flag.ExitOnError)
	pkt := flag.NewFlagSet("paket", flag.ExitOnError)
	flag.NewFlagSet("login", flag.ExitOnError)
	flag.String("file", "", "Filename to test or submit")
	flag.String("p", "", "Paket number")
	// submit := flag.NewFlagSet("submit", flag.ExitOnError)
	// submitFile := flag.String("filename", "", "Filename to test or submit")

	flag.Usage = func() {
		fmt.Println(os.Getenv("NAME"))
		fmt.Println("Usage: grader <command> [<subcommand>] [<flags>]")
		fmt.Println("\nCommands:")
		fmt.Println("  login   Log in to autograder system")
		fmt.Println("  paket   Generate problem case file (PDF only)")
		fmt.Println("  test    Test program using autograder")
		fmt.Println("  submit  Submit your program report to zip file")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
	}

	// Parse command-line arguments
	flag.Parse()

	// Get command and subcommand
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(0)
	}
	cmd := args[0]
	if cmd == "login" {
		utils.Login()
		return
	}
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		// file doesn't exist
		println("\033[31m[-] You must login first\033[0m")
		return
	}
	// Process command
	switch cmd {
	case "test":
		filename := test.String("file", "", "Programs filename")
		kode := test.String("p", "", "Paket number")
		test.Parse(os.Args[2:])
		if *filename == "" || *kode == "" {
			fmt.Println("Usage: grader test -file=<filename> -p=<paket>")
			os.Exit(1)
		}
		fileInfo, err := os.Stat(*filename)
		if os.IsNotExist(err) {
			println("\033[31m[-] File does not exist\033[0m")
			return
		}
		if err != nil {
			println("Error:", err)
			return
		}
		if fileInfo.IsDir() {
			println("\033[31m[-] Filename is a directory\033[0m")
			return
		}
		utils.Test(*filename, *kode)

	case "paket":
		kode := pkt.String("p", "", "Paket number")
		pkt.Parse(os.Args[2:])
		if *kode == "" {
			fmt.Println("Usage: grader paket -p=<paket>")
			os.Exit(1)
		}
		var pkt_file []byte
		var c_file []byte
		var pktName string
		var cName string
		switch *kode {
		case "1":
			pkt_file, _ = paket.ReadFile("paket/sample/sample.pdf")
			c_file, _ = template.ReadFile("paket/sample/sample.cpp")
			cName = "sample.cpp"
			pktName = "sample.pdf"
		case "lemon":
			pkt_file, _ = paket.ReadFile("paket/Paket 1/soal_1.pdf")
			c_file, _ = template.ReadFile("paket/Paket 1/soal_1.cpp")
			cName = "main-1.c"
			pktName = "CASE_1.pdf"
		case "skylar":
			pkt_file, _ = paket.ReadFile("paket/Paket 2/soal_2.pdf")
			c_file, _ = template.ReadFile("paket/Paket 2/soal_2.cpp")
			cName = "main-2.c"
			pktName = "CASE_2.pdf"
		case "jimin":
			pkt_file, _ = paket.ReadFile("paket/Paket 3/soal_3.pdf")
			c_file, _ = template.ReadFile("paket/Paket 3/soal_3.cpp")
			cName = "main-3.cpp"
			pktName = "CASE_3.pdf"
		case "tehyung":
			pkt_file, _ = paket.ReadFile("paket/Paket 4/soal_4.pdf")
			c_file, _ = template.ReadFile("paket/Paket 4/soal_4.cpp")
			cName = "main-4.cpp"
			pktName = "CASE_4.pdf"
		case "ganjar":
			pkt_file, _ = paket.ReadFile("paket/Paket 5/soal_5.pdf")
			c_file, _ = template.ReadFile("paket/Paket 5/soal_5.cpp")
			cName = "main-5.cpp"
			pktName = "CASE_5.pdf"
		case "annies":
			pkt_file, _ = paket.ReadFile("paket/Paket 6/soal_6.pdf")
			c_file, _ = template.ReadFile("paket/Paket 6/soal_6.cpp")
			cName = "main-6.cpp"
			pktName = "CASE_6.pdf"
		case "jihyo":
			pkt_file, _ = paket.ReadFile("paket/Paket 7/soal_7.pdf")
			c_file, _ = template.ReadFile("paket/Paket 7/soal_7.cpp")
			cName = "main-7.cpp"
			pktName = "CASE_7.pdf"
		case "ryujin":
			pkt_file, _ = paket.ReadFile("paket/Paket 8/soal_8.pdf")
			c_file, _ = template.ReadFile("paket/Paket 8/soal_8.cpp")
			cName = "main-8.cpp"
			pktName = "CASE_8.pdf"
		default:
			println("\033[31m[-] Paket not found\033[0m")
			return
		}
		err := ioutil.WriteFile(pktName, pkt_file, 0644)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(cName, c_file, 0644)
		if err != nil {
			log.Fatal(err)
		}
		println("\033[32m[+] File created successfully\t\033[0m\n")
		return
	case "submit":
		filename := submit.String("file", "", "Programs filename")
		kode := submit.String("p", "", "Paket number")
		submit.Parse(os.Args[2:])
		if *filename == "" || *kode == "" {
			fmt.Println("Usage: grader submit -file=<filename> -p=<paket>")
			os.Exit(1)
		}
		utils.Submit(*filename, *kode)
	default:
		fmt.Println("Invalid command:", cmd)
		os.Exit(1)
	}
}
