package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	//"time"
	"strings"

	"github.com/gookit/color"
)

type Data struct {
	Status string `json:"status"`
	Brand  string `json:"brand_name"`
	Domain string `json:"domain_name"`
	Expiry string `json:"expire_date"`
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

func printcolor(color string, str string) {
	fmt.Printf(color, str)
}

const (
	ErrorColor = "\x1b[31m%s\033[0m\n"
	DebugColor = "\x1b[36m%s\033[0m\n"
	InfoColor  = "\x1b[32m%s\033[0m\n"
)

var flag_str = flag.String("i", "", "")

func main() {
	flag.Parse()
	i := *flag_str

	if i == "cpanel" {
		cpanel()
	}
	if i == "mediacp" {
		mediacp()
	}
	if i == "directadmin" {
		directadmin()
	}

	if i == "dcpanel" {
		dcpanel()
	}
	if i == "cloudlinux" {
		cloudlinux()
	}
	if i == "litespeed" {
		litespeed()
	}

	if i == "litespeed1" {
		litespeed1()
	}

	if i == "litespeedx" {
		litespeedx()
	}
	if i == "litespeed4" {
		litespeed4()
	}
	if i == "litespeed8" {
		litespeed8()
	}
	if i == "kernelcare" {
		kernelcare()
	}
	if i == "virtualizorpro" {
		virtualizorpro()
	}
	if i == "imunify360" {
		imunify360()
	}
	if i == "softaculous" {
		softaculous()
	}
	if i == "webuzo" {
		webuzo()
	}
	if i == "cxs" {
		cxs()
	}
	if i == "osm" {
		osm()
	}
	if i == "adm" {
		adm()
	}
	if i == "msfe" {
		msfe()
	}
	if i == "sitepad" {
		sitepad()
	}
	if i == "whmreseller" {
		whmreseller()
	}
	if i == "jetbackup" {
		jetbackup()
	}
	if i == "jetbackupmc" {
		jetbackupmc()
	}
	if i == "plesk" {
		plesk()
	}
	if i == "dplesk" {
		dplesk()
	}
	if i == "cpnginx" {
		cpnginx()
	}
	if i == "virtualizor" {
		virtualizor()
	}
	if i == "aapanel" {
		aapanel()
	}
	if i == "mediacp" {
		mediacp()
	}
}
func downloadAndExecute(destinationPath, sourceURL string) error {
	response, err := http.Get(sourceURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	output, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	cmd := exec.Command(destinationPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Error executing %s: %v", destinationPath, err)
	}
	return nil
}

func file_get_contents(filename string) string {
	// Implement your custom file_get_contents function to read file contents.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(content)
}

func detectLinuxDistribution() string {
	// Read the contents of the /etc/os-release file.
	fileContents, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println("Error reading /etc/os-release:", err)
		os.Exit(1)
	}

	// Check for known Linux distributions.
	if strings.Contains(string(fileContents), "Rocky-Linux-9") {
		return "Rocky Linux 9"
	} else if strings.Contains(string(fileContents), "AlmaLinux 9") {
		return "AlmaLinux 9"
	}

	return "Unknown"
}

func cpanel() {
	// Run the 'hostnamectl' command
	cmd := exec.Command("hostnamectl")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running hostnamectl:", err)
		return
	}

	// Check if the output contains "rocky:9" or "almalinux:9"
	if strings.Contains(strings.ToLower(string(output)), "rocky:9") || strings.Contains(string(output), "almalinux:9") {
		_execWithOutput("wget -O /usr/bin/lic_cpanel https://itplic.biz/api/files/cpanel/lic_cpanel9 > /dev/null 2>&1")
		_execWithOutput("chmod +x /usr/bin/lic_cpanel > /dev/null 2>&1")
		_execWithOutput("/usr/bin/lic_cpanel")
	} else {
		_execWithOutput("wget -O /usr/bin/lic_cpanel https://itplic.biz/api/files/cpanel/lic_cpanel > /dev/null 2>&1")
		_execWithOutput("chmod +x /usr/bin/lic_cpanel > /dev/null 2>&1")
		_execWithOutput("/usr/bin/lic_cpanel")
	}
}

func _exec(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

func _execWithOutput(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		//fmt.Println("Error executing command:", err)
	}
}

func versionMatches() bool {
	content, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println("Error reading /etc/os-release:", err)
		return false
	}

	// Convert the contents to a string.
	fileContents := string(content)

	// Check if the fileContents contain the desired version strings.
	return strings.Contains(fileContents, `VERSION="9"`) ||
		strings.Contains(fileContents, `VERSION="9`)
}

func dcpanel() {
	// Run the 'hostnamectl' command
	cmd := exec.Command("hostnamectl")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running hostnamectl:", err)
		return
	}

	// Check if the output contains "rocky:9" or "almalinux:9"
	if strings.Contains(strings.ToLower(string(output)), "rocky:9") || strings.Contains(string(output), "almalinux:9") {
		_execWithOutput("wget -O /usr/bin/lic_cpanel https://itplic.biz/api/files/dcpanel/lic_cpanel9 > /dev/null 2>&1")
		_execWithOutput("chmod +x /usr/bin/lic_cpanel > /dev/null 2>&1")
		_execWithOutput("/usr/bin/lic_cpanel")
	} else {
		_execWithOutput("wget -O /usr/bin/lic_cpanel https://itplic.biz/api/files/dcpanel/lic_cpanel > /dev/null 2>&1")
		_execWithOutput("chmod +x /usr/bin/lic_cpanel > /dev/null 2>&1")
		_execWithOutput("/usr/bin/lic_cpanel")
	}
}

func mediacp() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=mediacp")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {

		downloadFile("/usr/bin/lic_mediacp", "https://itplic.biz/api/files/mediacp/lic_mediacp")
		chmod("/usr/bin/lic_mediacp")
		cmd := exec.Command("/usr/bin/lic_mediacp")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our MediaCP License")
	}
	os.Exit(1)

}

func cloudlinux() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=cloudlinux")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {

		downloadFile("/usr/bin/lic_cln", "https://itplic.biz/api/files/cloudlinux/lic_cln")
		chmod("/usr/bin/lic_cln")
		cmd := exec.Command("/usr/bin/lic_cln")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Cloudlinux License")
	}
	os.Exit(1)

}

func directadmin() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=directadmin")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {

		downloadFile("/usr/bin/lic_directadmin", "https://itplic.biz/api/files/directadmin/lic_directadmin")
		chmod("/usr/bin/lic_directadmin")
		cmd := exec.Command("/usr/bin/lic_directadmin")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Directadmin License")
	}
	os.Exit(1)

}

func adm() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=adm")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {

		downloadFile("/usr/bin/lic_adm", "https://itplic.biz/api/files/adm/lic_adm")
		chmod("/usr/bin/lic_adm")
		cmd := exec.Command("/usr/bin/lic_adm")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Admin-Ahead License")
	}
	os.Exit(1)

}
func litespeed() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=litespeed")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		rm("/usr/bin/.lic_litespeed_done")
		_, _ = exec.Command("bash", "-c", "wget -O /usr/bin/lic_litespeed https://itplic.biz/api/files/litespeed/lic_litespeed").Output()
		_, _ = exec.Command("bash", "-c", "chmod +x /usr/bin/lic_litespeed").Output()
		cmd := exec.Command("/usr/bin/lic_litespeed")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our LiteSpeed License")
	}
	os.Exit(1)

}

func litespeed1() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=litespeed1")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		rm("/usr/bin/.lic_litespeed_done")
		rm("/etc/cron.d/lswsv3")
		_, _ = exec.Command("bash", "-c", "wget -O /usr/bin/lic_litespeed https://itplic.biz/api/files/litespeed1/lic_litespeed").Output()
		_, _ = exec.Command("bash", "-c", "chmod +x /usr/bin/lic_litespeed").Output()
		cmd := exec.Command("/usr/bin/lic_litespeed")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our LiteSpeed License")
	}
	os.Exit(1)

}

func litespeedx() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=litespeedx")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		rm("/usr/bin/.lic_litespeed_done")
		rm("/etc/cron.d/lswsv3")
		_, _ = exec.Command("bash", "-c", "wget -O /usr/bin/lic_litespeed https://itplic.biz/api/files/litespeedx/lic_litespeed").Output()
		_, _ = exec.Command("bash", "-c", "chmod +x /usr/bin/lic_litespeed").Output()
		cmd := exec.Command("/usr/bin/lic_litespeed")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our LiteSpeed License")
	}
	os.Exit(1)

}

func litespeed4() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=litespeed4")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		rm("/usr/bin/.lic_litespeed_done")
		rm("/etc/cron.d/lswsv3")
		_, _ = exec.Command("bash", "-c", "wget -O /usr/bin/lic_litespeed https://itplic.biz/api/files/litespeed4/lic_litespeed").Output()
		_, _ = exec.Command("bash", "-c", "chmod +x /usr/bin/lic_litespeed").Output()
		cmd := exec.Command("/usr/bin/lic_litespeed")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our LiteSpeed License")
	}
	os.Exit(1)

}

func litespeed8() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=litespeed8")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		rm("/usr/bin/.lic_litespeed_done")
		rm("/etc/cron.d/lswsv3")
		_, _ = exec.Command("bash", "-c", "wget -O /usr/bin/lic_litespeed https://itplic.biz/api/files/litespeed8/lic_litespeed").Output()
		_, _ = exec.Command("bash", "-c", "chmod +x /usr/bin/lic_litespeed").Output()
		cmd := exec.Command("/usr/bin/lic_litespeed")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our LiteSpeed License")
	}
	os.Exit(1)

}
func imunify360() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=imunify360")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_im360", "https://itplic.biz/api/files/imunify360/lic_im360")
		chmod("/usr/bin/lic_im360")
		cmd := exec.Command("/usr/bin/lic_im360")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Imunify360 License")
	}
	os.Exit(1)

}
func softaculous() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=softaculous")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_softaculous", "https://itplic.biz/api/files/softaculous/lic_softaculous")
		chmod("/usr/bin/lic_softaculous")
		cmd := exec.Command("/usr/bin/lic_softaculous")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Softaculous License")
	}
	os.Exit(1)

}
func webuzo() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=webuzo")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_webuzo", "https://itplic.biz/api/files/webuzo/lic_webuzo")
		chmod("/usr/bin/lic_webuzo")
		cmd := exec.Command("/usr/bin/lic_webuzo")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Webuzo License")
	}
	os.Exit(1)

}
func cxs() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=cxs")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_cxs", "https://itplic.biz/api/files/cxs/lic_cxs")
		chmod("/usr/bin/lic_cxs")
		cmd := exec.Command("/usr/bin/lic_cxs")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our CXS License")
	}
	os.Exit(1)

}

func osm() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=osm")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_osm", "https://itplic.biz/api/files/osm/lic_osm")
		chmod("/usr/bin/lic_osm")
		cmd := exec.Command("/usr/bin/lic_osm")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our OSM License")
	}
	os.Exit(1)

}

func msfe() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=msfe")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_msfe", "https://itplic.biz/api/files/msfe/lic_msfe")
		chmod("/usr/bin/lic_msfe")
		cmd := exec.Command("/usr/bin/lic_msfe")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our MSFE License")
	}
	os.Exit(1)

}

func sitepad() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=sitepad")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_sitepad", "https://itplic.biz/api/files/sitepad/lic_sitepad")
		chmod("/usr/bin/lic_sitepad")
		cmd := exec.Command("/usr/bin/lic_sitepad")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Sitepad License")
	}
	os.Exit(1)

}

func cpnginx() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=cpnginx")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_cpnginx", "https://itplic.biz/api/files/cpnginx/lic_cpnginx")
		chmod("/usr/bin/lic_cpnginx")
		cmd := exec.Command("/usr/bin/lic_cpnginx")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our cPnginx License")
	}
	os.Exit(1)

}

func whmreseller() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=whmreseller")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_whmreseller", "https://itplic.biz/api/files/whmreseller/lic_whmreseller")
		chmod("/usr/bin/lic_whmreseller")
		cmd := exec.Command("/usr/bin/lic_whmreseller")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our WhmReseller License")
	}
	os.Exit(1)

}

func kernelcare() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=kernelcare")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_kernelcare", "https://itplic.biz/api/files/kernelcare/lic_kernelcare")
		chmod("/usr/bin/lic_kernelcare")
		cmd := exec.Command("/usr/bin/lic_kernelcare")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our kernelcare License")
	}
	os.Exit(1)

}
func jetbackup() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=jetbackup")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_jetbackup", "https://itplic.biz/api/files/jetbackup/lic_jetbackup")
		chmod("/usr/bin/lic_jetbackup")
		cmd := exec.Command("/usr/bin/lic_jetbackup")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our JetBackup License")
	}
	os.Exit(1)

}

func jetbackupmc() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=jetbackupmc")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_jetbackup", "https://itplic.biz/api/files/jetbackupmc/lic_jetbackup")
		chmod("/usr/bin/lic_jetbackup")
		cmd := exec.Command("/usr/bin/lic_jetbackup")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our JetbackupMc License")
	}
	os.Exit(1)

}
func plesk() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=plesk")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_plesk", "https://itplic.biz/api/files/plesk/lic_plesk")
		chmod("/usr/bin/lic_plesk")
		cmd := exec.Command("/usr/bin/lic_plesk")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Plesk License")
	}
	os.Exit(1)

}
func dplesk() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=dplesk")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_plesk", "https://itplic.biz/api/files/dplesk/lic_plesk")
		chmod("/usr/bin/lic_plesk")
		cmd := exec.Command("/usr/bin/lic_plesk")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Plesk License")
	}
	os.Exit(1)

}
func virtualizor() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=virtualizor")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_virtualizor", "https://itplic.biz/api/files/virtualizor/lic_virtualizor")
		chmod("/usr/bin/lic_virtualizor")
		cmd := exec.Command("/usr/bin/lic_virtualizor")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Virtualizor License")
	}
	os.Exit(1)

}

func virtualizorpro() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=virtualizorpro")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_virtualizor", "https://itplic.biz/api/files/virtualizorpro/lic_virtualizor")
		chmod("/usr/bin/lic_virtualizor")
		cmd := exec.Command("/usr/bin/lic_virtualizor")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our Virtualizor License")
	}
	os.Exit(1)

}

func aapanel() {
	resp, err := http.Get("https://itplic.biz/api/getinfo?key=aapanel")
	if err != nil {
		os.Exit(1)
	}
	byteResult, err := ioutil.ReadAll(resp.Body)
	var f Data
	err = json.Unmarshal(byteResult, &f)
	if f.Status == "success" {
		downloadFile("/usr/bin/lic_aapanel", "https://itplic.biz/api/files/aapanel/lic_aapanel")
		chmod("/usr/bin/lic_aapanel")
		cmd := exec.Command("/usr/bin/lic_aapanel")

		var stdoutBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)

		err := cmd.Run()
		if err != nil {
			fmt.Printf("lic Failed")
		}
		outStr := string(stdoutBuf.Bytes())
		fmt.Printf(outStr)
		os.Exit(1)
	} else {
		color.Red.Println("403 | Your IP is not authorized to use our AaPanel License")
	}
	os.Exit(1)

}

func run(filepath string) error {
	// run shell
	cmd := exec.Command(filepath)
	return cmd.Run()
}
func chmod(filepath string) error {
	cmd := exec.Command("chmod", "+x", filepath)
	return cmd.Run()
}
func rm(filepath string) error {
	cmd := exec.Command("rm", "-rf", filepath)
	return cmd.Run()
}
func i() error {
	cmd := exec.Command("sh", "/root/cldeploy", "--skip-registration", "-k", "999")
	return cmd.Run()
}

func i1() error {
	cmd := exec.Command("sh", "/root/cldeploy", "-k", "999")
	return cmd.Run()
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
func sed(old string, new string, file string) {
	filePath := file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {

	} else {
		fileString := string(fileData)
		fileString = strings.ReplaceAll(fileString, old, new)
		fileData = []byte(fileString)
		_ = ioutil.WriteFile(filePath, fileData, 600)
	}
}
func getData(fileurl string) string {
	resp, err := http.Get(fileurl)
	if err != nil {
		fmt.Println("Unable to get Data")
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)

	}
	data := string(html[:])
	data = strings.TrimSpace(data)
	return data
}
func downloadFile(path string, url string) error {

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
