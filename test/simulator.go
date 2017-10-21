package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	createuserEndpoint = "http://localhost:8080/api/create/"
	userJSON           = "{\"Name\":\"#ID - Nagu Kothapalli\",\"Gender\":\"male\",\"Age\":25,\"Id\":\"#ID\"}"
)

func main() {
	fmt.Println("Create user Test  - Strated at", time.Now())

	for i := 1; i <= 100; i++ {
		createuserTest(i)
	}

	fmt.Println("Create user Test  - Ended at ", time.Now())

}

func createuserTest(id int) {

	req, err := http.NewRequest("POST", createuserEndpoint, bytes.NewBuffer([]byte(strings.Replace(userJSON, "#ID", strconv.Itoa(id), -1))))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response Status: %s  Response: %s \n", resp.Status, string(body))
}
