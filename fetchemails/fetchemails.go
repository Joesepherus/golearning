package fetchemails

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/joho/godotenv"
)

func fetchEmails(c *client.Client) {
	// Select the mailbox
	mailbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the latest messages
	seqset := new(imap.SeqSet)
	seqset.AddRange(mailbox.Messages, mailbox.Messages)

	messages := make(chan *imap.Message, 0)
	go func() {
		// Fetch the full message including all parts
		if err := c.Fetch(seqset, []imap.FetchItem{imap.FetchItem("BODY.PEEK[TEXT]"), imap.FetchEnvelope}, messages); err != nil {
			log.Fatal(err)
		}
		// Close the channel when done
		//              close(messages)
	}()

	for msg := range messages {
		fmt.Println("\n\n\n\n\n")
		fmt.Println("Message received:")
		fmt.Printf("From: %v\n", msg.Envelope.From)
		fmt.Printf("Subject: %v\n", msg.Envelope.Subject)
		fmt.Printf("Date: %v\n", msg.Envelope.Date)

		// Print message body
		if msg.Body != nil {

			fmt.Println("Body: \n")

			// Initialize a variable to collect body parts
			var bodyParts []string

			// Iterate over body parts
			for _, section := range msg.Body {
				if section != nil {
					// Read the section
					bodyBytes, err := ioutil.ReadAll(section)
					if err != nil {
						log.Fatal(err)
					}

					// Convert bytes to string
					bodyStr := string(bodyBytes)

					// Check if the body is base64-encoded
					// Decode the base64 string
					decodedBytes, err := base64.StdEncoding.DecodeString(bodyStr)
					if err != nil {
						log.Println(err)
						bodyParts = append(bodyParts, bodyStr)
						continue
					}
					bodyStr = string(decodedBytes)

					// Collect body parts
					bodyParts = append(bodyParts, bodyStr)
				}
			}

			// Print the combined body
			fmt.Println(strings.Join(bodyParts, "\n"))
		} else {
			fmt.Println("No body found.")
		}
	}
}

func Run() {
  err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Email server credentials
	server := "imap.hostinger.com:993"
	username := "joes@joesexperiences.com"
	password := os.Getenv("PASSWORD")

	// Connect to the server
	c, err := client.DialTLS(server, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	// Login
	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("kokot")
		fetchEmails(c)
		time.Sleep(1 * time.Second) // Adjust the sleep interval as needed
	}
}
