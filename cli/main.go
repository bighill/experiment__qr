package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

const KILL_MESS = "q"
const PROMPT = "('" + KILL_MESS + "' to quit) \nQR code string: "

func main() {
	var qrString string
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()

		if qrString != "" {

			// Generate a QR code
			qr, err := qrcode.New(qrString, qrcode.Medium)
			if err != nil {
				fmt.Println("Error generating QR code:", err)
				continue
			}

			// Print the QR code
			fmt.Println(qrString)
			fmt.Println(qr.ToSmallString(true))

		}

		// Prompt
		fmt.Print(PROMPT)
		qrString, _ = reader.ReadString('\n')
		qrString = strings.TrimSpace(qrString)

		// Exit
		if strings.ToLower(qrString) == KILL_MESS {
			break
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
