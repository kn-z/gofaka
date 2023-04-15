package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/go-acme/lego/v4/registration"
	"gofaka/utils"
	"log"
	"os"
)

// MyUser You'll need a user or account type that implements acme.User
type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *MyUser) GetEmail() string {
	return u.Email
}
func (u MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func SaveToFile(domain string, path string, content []byte) {
	file, err := os.Create(fmt.Sprintf(path, domain))
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(content)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	utils.Init()
	// Create a user. New accounts need an email and private key to start.
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	myUser := MyUser{
		Email: utils.AuthEmail,
		key:   privateKey,
	}

	config := lego.NewConfig(&myUser)

	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	// CADirURL have 2 type -- Staging and Production.
	// Staging mode can apply for certificates without limit but cert is invalid.
	config.CADirURL = lego.LEDirectoryProduction
	config.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	cfConfig := cloudflare.NewDefaultConfig()
	cfConfig.AuthEmail = utils.AuthEmail
	cfConfig.AuthKey = utils.AuthKey

	// Create Cloudflare DNS Provider client
	cfClient, err := cloudflare.NewDNSProviderConfig(cfConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = client.Challenge.SetDNS01Provider(cfClient)

	// New users will need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Fatal(err)
	}
	myUser.Registration = reg

	request := certificate.ObtainRequest{
		Domains: []string{utils.Domain},
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Fatal(err)
	}

	// Each certificate comes back with the ct bytes, the bytes of the client's
	// private key, and a certificate URL. SAVE THESE TO DISK.
	//fmt.Printf("%#v\n", certificates)
	// ... all done.
	// Save certificate and private key to files
	SaveToFile(utils.Domain, "../cert/certificates/%s.crt", certificates.Certificate)
	SaveToFile(utils.Domain, "../cert/certificates/%s.issuer.crt", certificates.IssuerCertificate)
	SaveToFile(utils.Domain, "../cert/certificates/%s.key", certificates.PrivateKey)
}
