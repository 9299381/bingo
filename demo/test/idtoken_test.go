package test

import (
	"fmt"
	"github.com/9299381/bingo/package/id"
	"github.com/9299381/bingo/package/token"
	"testing"
)

func TestID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := id.New()
		fmt.Println(id)
	}
}

func TestToken(t *testing.T) {
	claims := &token.Claims{
		Id:   id.New(),
		Name: "无所谓",
		Role: "user",
	}
	auth, _ := token.GetToken(claims)
	fmt.Println(auth)

}

func TestVerify(t *testing.T) {
	auth := "eyJpZCI6IjEzNzM2ODkzNDUzNzE3OTk1NTIiLCJuYW1lIjoi5peg5omA6LCTIiwicm9sZSI6InVzZXIiLCJpYXQiOjE1ODk3ODgyNjEsImV4cCI6MTU5MjM4MDI2MX0=.c16db89dc8cf119ab139b4fa4fcc8ad7"
	claim, err := token.CheckToken(auth)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(claim)

}
