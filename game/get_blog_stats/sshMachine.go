package get_blog_stats

import (
	"os/user"
	"code.google.com/p/go.crypto/ssh"
	"io/ioutil"
)

func GetKeyFile() (key ssh.Signer, err error) {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.ssh/id_rsa"
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}

func GetConfigFromKeys(username string) (config *ssh.ClientConfig, err error) {
	key, err := GetKeyFile()

	if err != nil {
		panic(err)
	}

	config = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	return
}

func GetConfig(username string, password string) (config *ssh.ClientConfig) {
	config = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	return
}


