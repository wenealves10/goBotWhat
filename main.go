package main

import (
	"fmt"
	"time"

	qr "github.com/Baozisoftware/qrcode-terminal-go"
	whatsapp "github.com/Rhymen/go-whatsapp"
)

func main() {

	wac, _ := whatsapp.NewConn(20 * time.Second)
	wac.SetClientVersion(2, 2142, 12)

	qrChan := make(chan string)
	go func() {
		obj := qr.New()
		obj.Get(<-qrChan).Print()
	}()
	wac.Login(qrChan)

	var numberWhat int

	fmt.Print("Número de whatsapp\nExemplo: 559988525464 \nDigite aqui: ")
	fmt.Scanln(&numberWhat)

	formatNumber := fmt.Sprintf("%d@s.whatsapp.net", numberWhat)

	text := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: formatNumber,
		},
		Text: "Golang Bot What",
	}

	wac.Send(text)

}
