package connection

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"test-server/proofofwork"
	"test-server/quotes"
	"time"
)

type Session struct {
	Conn     net.Conn
	hashCash string
}

func OpenServer(host string, port string) (listener net.Listener, err error) {
	listener, err = net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal("Server error: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started (" + host + ":" + port + ")")
	accepter(listener)
	return
}

func accepter(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accepter error: ", err)
		}
		session := &Session{
			Conn: conn,
		}
		go handleConnection(session)
	}
}

func handleConnection(session *Session) {
	defer session.Conn.Close()
	for {
		session.Conn.SetReadDeadline(time.Now().Add(time.Second * 20))
		receive := []string{}

		decoder := json.NewDecoder(session.Conn)
		err := decoder.Decode(&receive)
		if err == io.EOF {
			continue
		}
		if err != nil {
			fmt.Println("Error json: ", err)
			return
		}
		if len(receive) == 0 {
			fmt.Println("Message not found")
			continue
		}
		fmt.Println("<< ", receive)

		switch receive[0] {
		case "CHOOSE":
			choose(session)
			continue
		case "VERIFY":
			verify(session, receive)
			continue
		default:
			fmt.Println("type message not found")
			continue
		}
	}
}

func choose(session *Session) {
	session.hashCash = proofofwork.PrepareHashCash()
	message := [...]string{
		"SOLVE",
		os.Getenv("NAME"),
		session.hashCash,
	}
	send(session.Conn, message)
}

func verify(session *Session, receive []string) {
	nonce, err := strconv.Atoi(receive[2])
	if err != nil {
		fmt.Println("nonce not found")
		return
	}

	if !hashCheck(session, nonce) {
		fmt.Println("Recieved erorr hash")
		return
	}

	message := [...]string{
		"GRANT",
		os.Getenv("NAME"),
		quotes.GetRandom(),
	}
	send(session.Conn, message)
}

func send(conn net.Conn, message interface{}) {
	fmt.Println(">> ", message)
	encoder := json.NewEncoder(conn)
	err := encoder.Encode(&message)
	if err != nil {
		fmt.Println("Encode error = ", err)
		return
	}
}

func hashCheck(session *Session, nonce int) bool {
	pow := &proofofwork.ProofOfWork{
		HashCash: session.hashCash,
	}
	return pow.Validate(nonce)
}
