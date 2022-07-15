package main

type SignMsg struct {
	SignedMsg    string `json:"signed_msg"`
	PubX         string `json:"pub_x"`
	PubY         string `json:"pub_y"`
	SignatureR8X string `json:"signature_r_8_x"`
	SignatureR8Y string `json:"signature_r_8_y"`
	SignatureS   string `json:"signature_s"`
}
