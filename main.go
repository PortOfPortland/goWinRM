package goWinRM

import (
	ps "github.com/portofportland/go-powershell"
	"github.com/portofportland/go-powershell/backend"
)


func RunWinRMCommand(username string, password string, server string, command string, usessl string, usessh string) (string, error) {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		//something bad happened - return an error
		return "", err
	}
	defer shell.Exit()

	// ... and interact with it
	var winRMPre string

    if (usessh == "1") {
    	winRMPre = "$s = New-PSSession -HostName " + server + " -Username " + username + " -SSHTransport"
    } else {
        winRMPre = "$SecurePassword = '" + password + "' | ConvertTo-SecureString -AsPlainText -Force; $cred = New-Object System.Management.Automation.PSCredential -ArgumentList '" + username + "', $SecurePassword; $s = New-PSSession -ComputerName " + server + " -Credential $cred"
	}

    var winRMPost string = "; Invoke-Command -Session $s -Scriptblock { " + command + " }; Remove-PSSession $s"

	// use SSL if requested
	var winRMCommand string
	if (usessl == "1") {
		winRMCommand = winRMPre + " -UseSSL" + winRMPost
	} else {
		winRMCommand = winRMPre + winRMPost
	}
	stdout, _, err := shell.Execute(winRMCommand)
	
	return stdout, err
}
