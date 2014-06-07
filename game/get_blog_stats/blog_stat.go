package get_blog_stats

import (
	"fmt"
	curlCmd "github.com/golang-basic/go-curl"
	"code.google.com/p/go.crypto/ssh"
	"bytes"
)

var username = "root"
var password = "ribbon7apple"

func ExecuteBlog() {
	/**
	config is same for all germany hosts
	 */
	config := GetConfig(username, password)
	/**
	For each host
	1. dial in
	2. run curl
	 */
	for _, host := range Hosts {

		client, err := ssh.Dial("tcp", host+":22", config)
		if err != nil {
			panic("Failed to dial: " + err.Error())
		}
		runCmdCountry(client)
	}
}

var value = "co.uk"
var times = 100

func runCmdCountry(client *ssh.Client) {
	for i := 0; i <= 100; i++ {
		for _, value := range country {
			runCmd(client, value)
		}
	}
}


func runCmd(client *ssh.Client, country string) {
	for _, urll := range Gobasic_urls {
		// Each ClientConn can support multiple interactive sessions,
		// represented by a Session.
		session, err := client.NewSession()
		if err != nil {
			panic("Failed to create session: " + err.Error())
		}
		defer session.Close()

		// Once a Session is created, you can execute a single command on
		// the remote side using the Run method.
		var b bytes.Buffer

		urll = blog+country+urll
		session.Stdout = &b
		if err := session.Run("cd /tmp && curl -k -v " + urll); err != nil {
			panic("Failed to run: " + err.Error())
		}
//		fmt.Println(b.String())
	}
}

func runcurl() {

	easy := curlCmd.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curlCmd.OPT_URL, "http://golang-basic.blogspot.com/2014/05/why-go-dont-use-space-indentation-as.html?view=sidebar")

	// make a callback function
	fooTest := func(buf []byte, userdata interface{}) bool {
		//		println("DEBUG: size=>", len(buf))
		//		println("DEBUG: content=>", string(buf))
		return true
	}

	//	easy.Setopt(curlCmd.OPT_WRITEFUNCTION, fooTest)
	easy.Setopt(curlCmd.OPT_HEADERFUNCTION, fooTest)

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}

func exampleFromGoLang() {
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
		User:
		"root",
		Auth: []ssh.AuthMethod{
			ssh.Password("ribbon7apple"),
		},
	}

	client, err := ssh.Dial("tcp", "letter.boxes.cf-app.com:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer

	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())


	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session1, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session1.Close()
	var b1 bytes.Buffer

	session1.Stdout = &b1
	if err := session1.Run("pwd"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b1.String())
}
