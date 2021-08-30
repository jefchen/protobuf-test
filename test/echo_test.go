package tests

import (
	"fmt"
	"log"
	"testing"

	"google.golang.org/protobuf/proto"
	"jeffrey.com/proto/echo"
)

func TestEcho(t *testing.T) {
	req := &echo.EchoRequest{Name: "Sushil"}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("Error while marshalling the object : %v", err)
	}

	res := &echo.EchoRequest{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatalf("Error while un-marshalling the object : %v", err)
	}
	fmt.Printf("Value from un-marshalled data is %v", res.GetName())
}
