package main

import (
	"encoding/json"
	"fmt"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"io"
	"log"
	"math/big"
	"os"
)

var configPath = "./sign.json"

func main() {
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("config file open failed ,err: %s", err)
		return
	}
	defer file.Close()
	signMsg := new(SignMsg)
	err = loadMsg(file, signMsg)
	if err != nil {
		log.Printf("load msg failed ,err: %s", err)
		return
	}
	ok := PerpSign(signMsg)
	log.Printf("verify %v", ok)
}

func loadMsg(reader io.Reader, signMsg *SignMsg) error {
	dec := json.NewDecoder(reader)
	return dec.Decode(signMsg)
}

func PerpSign(signMsg *SignMsg) bool {
	msg := new(big.Int)
	msg, ok := msg.SetString(signMsg.SignedMsg, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}
	X := new(big.Int)
	X, ok = X.SetString(signMsg.PubX, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}
	Y := new(big.Int)
	Y, ok = Y.SetString(signMsg.PubY, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}
	pubKey := &babyjub.Point{X: X, Y: Y}
	pk := babyjub.PublicKey(*pubKey)
	S := new(big.Int)
	S, ok = S.SetString(signMsg.SignatureS, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}
	R8X := new(big.Int)
	R8X, ok = R8X.SetString(signMsg.SignatureR8X, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}
	R8Y := new(big.Int)
	R8Y, ok = R8Y.SetString(signMsg.SignatureR8Y, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}

	sig := &babyjub.Signature{R8: &babyjub.Point{X: R8X, Y: R8Y}, S: S}
	ok = pk.VerifyPoseidon(msg, sig)
	return ok
}
