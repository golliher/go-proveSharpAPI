package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Pull out the characters up to the first \r
func parseResult(resultstring []byte) string {
	parsed := strings.Split(string(resultstring), "\r")
	return parsed[0]
}

// Send transmits Sharp Aquos API commands to the Television over the network
func Send(sharpCommand string, sharpParameter string, ip string, port string) string {
	cmdString := fmt.Sprintf("%4s%-4s\r", sharpCommand, sharpParameter)

	connectString := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.DialTimeout("tcp", connectString, time.Duration(100*time.Millisecond))

	if err != nil {
		return ("Error connecting to TV")
	}

	fmt.Fprintf(conn, cmdString)
	if err != nil {
		fmt.Println("An error occured.")
		fmt.Println(err.Error())
	}

	apiResult := make([]byte, 32)
	var resultString string

	bytesRead, err := conn.Read(apiResult)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Only read in %d bytes:", bytesRead)

	} else {
		resultString = parseResult(apiResult)
	}

	if ip == "192.168.4.21" {
		bytesRead, err := conn.Read(apiResult)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("Only read in %d bytes:", bytesRead)

		} else {
			resultString = parseResult(apiResult)
		}

	}
	conn.Close()
	return resultString

}

func printRow(cmd1, arg1, result1, cmd2, arg2, result2 string) {
	fmt.Printf(" %4s%-4s      %-17s               %4s%-4s      %-17s\n", cmd1, arg1, result1, cmd2, arg2, result2)
}

func testCMD(cmd string, arg string) {
	ip1 := "192.168.4.11"
	port1 := "10002"

	cmd1 := cmd
	arg1 := arg
	result1 := Send(cmd1, arg1, ip1, port1)

	ip2 := "192.168.4.21"
	port2 := "10002"

	cmd2 := cmd
	arg2 := arg
	result2 := Send(cmd2, arg2, ip2, port2)

	printRow(cmd1, arg1, result1, cmd2, arg2, result2)
}

func main() {

	fmt.Println("")
	fmt.Println("----------OLD TV---------------               ----------NEW TV---------------")
	fmt.Println("")
	fmt.Println(" SENT          RECEIVED                        SENT          RECEIVED")
	fmt.Println(" ========      ================                ========      ================")

	testCMD("POWR", "1")
	testCMD("MUTE", "0")
	testCMD("SWVN", "1")
	testCMD("MNRD", "1")
	testCMD("POWR", "?")
	testCMD("MUTE", "?")
	testCMD("MUTE", "1")
	testCMD("VOLM", "?")
	testCMD("IAVD", "?")
	testCMD("MUTE", "2")
	testCMD("IAVD", "?")
	testCMD("VOLM", "1")
	testCMD("RCKY", "33")
	testCMD("RCKY", "32")
	testCMD("RCKY", "39")
	testCMD("RCKY", "46")
	testCMD("RCKY", "36")

	fmt.Println("")

}
