package main

import (
	"46HW/pb"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req := pb.Request{
		Message: "Hello, World!",
	}
	reqBytes, _ := json.Marshal(req)
	resp, err := client.Post("http://localhost:8080/hash", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	var res pb.Response
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println("Hashed Message:", res.Message)
}
