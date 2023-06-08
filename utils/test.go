package utils

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Test(filename, kode string) {
	pass, fail, rte := 0, 0, 0
	cmd := exec.Command("g++", "-o", "YOUR_AWFUL_CODE", filename)
	err := cmd.Run()
	if err != nil {
		println("\033[31mCompilation error:\033[0m", err)
		return
	}
	testcode := map[string]TestCase{
		"1":       sample,
		"lemon":   T1,
		"skylar":  T2,
		"jimin":   T3,
		"tehyung": T4,
		"ganjar":  T5,
		"annies":  T6,
		"jihyo":   T7,
		"ryujin":  T8,
	}
	testcase, ok := testcode[kode]
	if !ok {
		fmt.Println("\033[31m[-] Paket not found\033[0m")
		return
	}

	fmt.Printf("\n🧪 RUN \033[33m%d\033[0m TEST CASE TO %s\n", len(testcase), filename)
	fmt.Print("⌛ Testing program ....")
	time.Sleep(3 * time.Second)
	fmt.Println("\r" + "                         ")
	for i, test := range testcase {
		startTime := time.Now()
		cmd := exec.Command("./YOUR_AWFUL_CODE")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		if err := cmd.Start(); err != nil {
			panic(err)
		}
		fmt.Fprintf(stdin, "%s\n", test.input)
		stdin.Close()
		var got []byte
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 {
				got = append(got, buf[:n]...)
			}
			if err != nil {
				break
			}
		}
		if err := cmd.Wait(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				// check for runtime error
				rte++
				fmt.Printf("\033[33m \u00D7 runtime error: %s\033[0m\n", exitErr.Error())
			} else {
				// some other error occurred
				fmt.Printf("\033[33m \u00D7 Error: %s\033[0m\n", err.Error())
			}
			continue
		}
		endTime := time.Now()
		if endTime.Sub(startTime) > time.Second*2 {
			fmt.Printf("\033[33m  \u00D7 TLE  Test %d \t\033[0m%v\n", i+1, endTime.Sub(startTime))
			continue
		}
		expected := strings.TrimSpace(string(got))
		expected = strings.ReplaceAll(expected, "\r", "")
		if expected != test.output {
			fail++
			fmt.Println([]byte(expected))
			fmt.Println(expected)
			fmt.Println([]byte(test.output))
			fmt.Println(test.output)
			fmt.Printf("\033[31m  \u00D7 FAIL  Test %d \t\033[0m%v\n", i+1, endTime.Sub(startTime))
			continue
		}
		pass++
		fmt.Printf("\033[32m  ✓ PASS  Test %d \t\033[0m%v\n", i+1, endTime.Sub(startTime))
	}
	fmt.Printf("\n\033[1mResult - \033[32m%d Passed |\033[31m %d Failed |\033[33m %d Error\033[0m\n", pass, fail, rte)
}
