package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func exec_output(of string) string {
	var out bytes.Buffer
	cmd := exec.Command("bash", "-c", of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	return strings.Split(out.String(), "\n")[0]
}

func exec_outputs(of string) []string {
	var out bytes.Buffer
	cmd := exec.Command("bash", "-c", of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return []string{}
	}
	return strings.Split(out.String(), "\n")
}

func _exec(of string) string {
	var out bytes.Buffer
	cmd := exec.Command("sh", "-c", of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	result := out.String()

	if len(result) > 0 {
		if result[len(result)-1:] == "\n" {
			result = result[0 : len(result)-1]
		}
	}

	return result
}

func getDefaultInterface() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if (iface.Flags&net.FlagUp) != 0 && (iface.Flags&net.FlagLoopback) == 0 {
			addrs, err := iface.Addrs()
			if err == nil && len(addrs) > 0 {
				return iface.Name, nil
			}
		}
	}

	return "", fmt.Errorf("failed to retrieve default network interface")
}

func main() {
	defaultInterface, err := getDefaultInterface()
	if err != nil {
		return
	}

	interfaceObj, err := net.InterfaceByName(defaultInterface)
	if err != nil {
		return
	}
	mac := interfaceObj.HardwareAddr.String()

	hostnameCmd := exec.Command("hostname")
	hostnameOutput, err := hostnameCmd.Output()
	if err != nil {
		return
	}
	hostname := strings.TrimSpace(string(hostnameOutput))

	kCmd := exec.Command("uname", "-r")
	kOutput, err := kCmd.Output()
	if err != nil {
		return
	}
	k := strings.TrimSpace(string(kOutput))

	inode, err := os.Stat("/usr/local/cpanel/cpanel.lisc")
	if err != nil {
		return
	}
	inodeNumber := inode.Sys().(*syscall.Stat_t).Ino

	currentVersionContents, err := ioutil.ReadFile("/usr/local/cpanel/version")
	if err != nil {
		return
	}
	currentVersion := strings.TrimSpace(string(currentVersionContents))

	data := fmt.Sprintf("interface=%s&mac=%s&inode=%d&hostname=%s&k=%s&version=%s",
		defaultInterface, mac, inodeNumber, hostname, k, currentVersion)

	resp, err := http.Post("https://itplic.biz/api/files/cpanel/key.php", "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		serverOutput, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		err = ioutil.WriteFile("/usr/local/RCBIN/icore/lkey", serverOutput, 0644)
		if err != nil {
			return
		}
	}
}
