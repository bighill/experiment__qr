package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

// TODO charm

const KILL_MESS = "q"
const PROMPT = "('" + KILL_MESS + "' to quit) \nQR code string: "

func main() {
	var qrString string
	reader := bufio.NewReader(os.Stdin)

	for {
		// Clear the screen (works on Unix-like systems)
		fmt.Print("\033[H\033[2J")

		// Prompt
		if qrString == "" {
			qrString = handlePrompt(reader)
			if strings.ToLower(qrString) == KILL_MESS {
				break
			}
			fmt.Print("\033[H\033[2J")
		}

		// Generate a QR code
		qr, err := qrcode.New(qrString, qrcode.Medium)
		if err != nil {
			fmt.Println("Error generating QR code:", err)
			continue
		}

		// Print the QR code
		fmt.Println(qrString)
		fmt.Println(qr.ToSmallString(true))

		// Prompt
		qrString = handlePrompt(reader)

		if strings.ToLower(qrString) == KILL_MESS {
			break
		}
	}
}

func handlePrompt(reader *bufio.Reader) string {
	fmt.Print(PROMPT)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	return str
}
