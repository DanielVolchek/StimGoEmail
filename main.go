package main

import (
	"fmt"
	"os"

	"github.com/danielvolchek/stim-email/pkg/email"
)

type Arguments struct {
	generate_auth bool
}

func main() {

	// args.CmdArg

	// args := Arguments{}
	// args.checkArgs()
	//
	// // If run with --generate-auth
	// // Hit db route and generate new auth token
	// // post on /auth
	// // Delete any other auth token
	//
	// if args.generate_auth {
	// 	fmt.Println("Generate auth")
	// 	os.Exit(0)
	// }

	// else
	// connect to SES
	// Create a POST route (https://vercel.com/docs/concepts/functions/serverless-functions/runtimes/go)
	// listen for a raw html string + sender address + auth token
	// check for correct auth by hitting db /auth
	// and then send an email through SES
	// On receiving HTML string

	fmt.Println("Hello world")
	email.CreateTemplateTest("Jacob", "Daniel")
}

func (args *Arguments) checkArgs() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {

			if arg == "--generate-auth" {
				args.generate_auth = true
			}
		}

	}
}
