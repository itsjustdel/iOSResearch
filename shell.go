package main

import (
	"log"
	"os"
	"os/exec"
)

// Recommended way seems to be using openssl https://www.openssl.org
func main() {

	// scroll down to "From PC" for openssl commands. Openssl is available on linux
	//https://learn.pandasuite.com/article/646-step-2-generate-ios-distribution-certificate

	cmd := exec.Command("cmd.exe", "set", "RANDFILE=.rnd") // how to do this cross platform?
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	//openssl req -new -key mykey.key -out CertificateSigningRequest.certSigningRequest  -subj "/emailAddress=yourAddress@example.com, CN=John Doe, C=US"
	cmd = exec.Command("openssl", "genrsa", "-out", "private.key", "2048")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	cmd = exec.Command("openssl", "req", "-new", "-key", "private.key", "-out", "CertificateSigningRequest.certSigningRequest", "-subj", "/emailAddress=yourAddress@example.com, CN=John Doe, C=US")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// Now upload this "CertificateSigningRequest.certSigningRequest" file to apple and download the iOS Distribution file (CER).
	pathToiOSDistFile := "../"
	iOSDistFile, err := os.Open(pathToiOSDistFile)
	if err != nil {
		log.Println("Can't find ios file")
	}

	//p12 creation
	//openssl pkcs12 -export -out keyStore.p12 -inkey myKey.pem -in certs.pem
	//https://stackoverflow.com/questions/38847489/how-to-generate-a-pkcss12-file-given-private-key-and-certificate-in-golang
	cmd = exec.Command("openssl", "pkcs12", "-export", "-out", "keyStore.p12", "-inkey", "private.key", "-in", iOSDistFile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
