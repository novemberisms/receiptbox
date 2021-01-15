package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// const address = "http://localhost:3885/"
// const address = "http://192.168.254.129:3885/"

const port = "3885"

var address string

func main() {

	ip := "localhost"

	if len(os.Args) > 1 {
		ip = os.Args[1]
	}

	address = fmt.Sprintf("http://%s:%s/", ip, port)

	printPrelude()

	for {
		recordEntry()
	}

}

func printPrelude() {

	fmt.Println("==================================================================")
	fmt.Println("Receipt Box!")
	fmt.Println("type in 'oops' to reset the current entry if you've made a mistake")
	fmt.Println("==================================================================")

}

func recordEntry() {
	defer fmt.Println("------------------------------------------------------------------")

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print("Restaurant       > ")
	restaurant, isOops := readString(inputReader)
	if isOops {
		return
	}

	fmt.Print("Date (mm-dd)     > ")
	date, isOops := readString(inputReader)
	if isOops {
		return
	}

	fmt.Print("Amount           > ")
	amount, isOops := readString(inputReader)
	if isOops {
		return
	}

	fmt.Print("Submit?          > ")
	_, isOops = readString(inputReader)
	if isOops {
		return
	}

	fmt.Print("...")

	message := map[string]string{
		"date":       date,
		"restaurant": restaurant,
		"amount":     amount,
	}

	messageAsBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	data := bytes.NewBuffer(messageAsBytes)

	response, err := http.Post(address, "application/json", data)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func readString(reader *bufio.Reader) (value string, isOops bool) {
	value, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	value = strings.TrimRight(value, "\r\n")
	isOops = value == "oops"
	return
}
