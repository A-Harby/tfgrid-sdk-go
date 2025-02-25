// Package integration for integration tests
package integration

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/deployer"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
	"golang.org/x/crypto/ssh"
)

var (
	trueVal   = true
	falseVal  = false
	statusUp  = "up"
	value1    = uint64(1)
	minRootfs = *convertGBToBytes(2)
)

var nodeFilter = types.NodeFilter{
	Status:  &statusUp,
	FreeSRU: convertGBToBytes(10),
	FreeHRU: convertGBToBytes(2),
	FreeMRU: convertGBToBytes(2),
	FarmIDs: []uint64{1},
	Rented:  &falseVal,
}

func convertGBToBytes(gb uint64) *uint64 {
	bytes := gb * 1024 * 1024 * 1024
	return &bytes
}

func setup() (deployer.TFPluginClient, error) {
	mnemonics := os.Getenv("MNEMONICS")
	log.Printf("mnemonics: %s", mnemonics)

	network := os.Getenv("NETWORK")
	log.Printf("network: %s", network)

	return deployer.NewTFPluginClient(mnemonics, "sr25519", network, "", "", "", 0, false)
}

// TestConnection used to test connection
func TestConnection(addr string, port string) bool {
	for t := time.Now(); time.Since(t) < 3*time.Second; {
		con, err := net.DialTimeout("tcp", net.JoinHostPort(addr, port), time.Second*12)
		if err == nil {
			con.Close()
			return true
		}
	}
	return false
}

// RemoteRun used for running cmd remotely using ssh
func RemoteRun(user string, addr string, cmd string, privateKey string) (string, error) {
	key, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", errors.Wrapf(err, "could not parse ssh private key %v", key)
	}
	// Authentication
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}

	// Connect
	port := "22"
	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, port), config)
	if err != nil {
		return "", errors.Wrapf(err, "could not start ssh connection")
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrapf(err, "could not create new session with message error")
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", errors.Wrapf(err, "could not execute command on remote with output %s", output)
	}
	return string(output), nil
}

// GenerateSSHKeyPair creates the public and private key for the machine
func GenerateSSHKeyPair() (string, string, error) {

	rsaKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", errors.Wrapf(err, "could not generate rsa key")
	}

	pemKey := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rsaKey)}
	privateKey := pem.EncodeToMemory(pemKey)

	pub, err := ssh.NewPublicKey(&rsaKey.PublicKey)
	if err != nil {
		return "", "", errors.Wrapf(err, "could not extract public key")
	}
	authorizedKey := ssh.MarshalAuthorizedKey(pub)
	return string(authorizedKey), string(privateKey), nil
}
