package main

import (
	"fmt"

	"github.com/pyed/easyssh"
)

const (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
DILDJFLKJDLFJDLFKJLDKJFLKDJFLKDJFLDJFLKDJFLKDJFLKDJFLDLJKFLKDUfi
RL7W3I1DOYJ+OAScd1M/Hjez84UnD9xxVXlNiaG5zOhh5n+C3qDXXGHJbc3yiRS7
4Uw7lxlo9iv3WMznq6K6aSzAitNynER7xmGbPYkW0Quxv01xrWfUCm3I4KOULotT
h8aQJZx+IvpbJXRRRRRRRRRRRRRRRRRRRRRRRRRRRSHm6WWpUVwe8yu8FRGqp2Tw
3p4Guvg1m9ZbCapAyAoTzj4Mflbt7oetLqiv4Xxq/zrCyzxxTUT7zYk1s6USF+si
QjtzmLNEzaZcPLI4rcobpbhzEqXNdi1U2lT/XQIDAQABAoIBAQC5J1+l28Fi1Zxy
pwz52UiXvV2PGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOkyfbk99
YOWuPt+islkdGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOy70UJwG
LxsQEeBhBQCcGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOCusR90f
CR831+gGZLjuGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOdv1Hivv
EIAymyDfXY9dGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOnPgBlz0
NHI4EGaxAoGBGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOO0aOGgNX
tw/J1E/F0lYwGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOO2CYsyPy
MhuxxZjI1nEoGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOoGBAOJt
rMaWC5C2o1ogGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOCfAj3NH
0ebtJTg8GqJSGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOORkg1ekc
K9Y5QNLxckZVGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOHWMpvPJ
74GGa1lVbhIEGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOOmqk9GN
OXDc/EdgT9trGGGGGGGGGGGGGGGGGGGGGGGGOOOOOOOOOOOOOOOOOOOOOMB1aY0N
RstbLLLLLLLLLLLLLLLLLLLLLAAAAAAAAAAAAAAAANNNNNNNNNNNNNGGGGGGGbr6
9aZ5LLLLLLLLLLLLLLLLLLLLLAAAAAAAAAAAAAAAANNNNNNNNNNNNNGGGGGGGD1c
07ORLLLLLLLLLLLLLLLLLLLLLAAAAAAAAAAAAAAAANNNNNNNNNNNNNGGGGGGGYha
xnu7LLLLLLLLLLLLLLLLLLLLLAAAAAAAAAAAAAAAANNNNNNNNNNNNNGGGGGGG7VU
FI86LLLLLLLLLLLLLLLLLLLLLAAAAAAAAAAAAAAAANNNNNNNNNNNNNGGGGGGGjgh
kJgMzksd4y3LMsib5Pi/SVUTxoyZgYXksYTjdc/GqTP4+YJrvoZZ
-----END RSA PRIVATE KEY-----`
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "john",
		Server: "example.com",
		Key:    privateKey,
		Port:   "22",
	}

	// Call Run method with command you want to run on remote server.
	response, err := ssh.Run("ps ax")
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}
}
