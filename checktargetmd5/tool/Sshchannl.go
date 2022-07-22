package tool

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

type Sshinfo struct {
	Ipaddr    string `json:"ipaddr"`
	User      string `json:"user"`
	Passwd    string `json:"passwd"`
	Cmd       string `json:"cmd"`
	Result    string `json:"result"`
	sshClient *ssh.Client
}

func NewSSHClient(ipaddr string, user string, pwd string, cmd string) Sshinfo {
	return Sshinfo{
		Ipaddr: ipaddr,
		User:   user,
		Passwd: pwd,
		Cmd:    cmd,
	}
}

func (si *Sshinfo) GetConfig() *ssh.ClientConfig {
	conf := &ssh.ClientConfig{
		User: si.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(si.Passwd),
		},
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return conf
}

func (si *Sshinfo) Connect() error {
	conf := si.GetConfig()

	client, err := ssh.Dial("tcp", si.Ipaddr+":22", conf)
	if err != nil {

		return err
	}
	si.sshClient = client
	return nil
}

func (si *Sshinfo) Run() error {
	if si.sshClient == nil {
		if err := si.Connect(); err != nil {
			fmt.Println("=========", err.Error())
			return err
		}
	}

	session, err := si.sshClient.NewSession()
	if err != nil {
		return err
	}

	defer session.Close()

	buf, err := session.CombinedOutput(si.Cmd)
	if err != nil {
		return err
	}
	si.Result = string(buf)
	return nil

}
