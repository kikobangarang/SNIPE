package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
)

func SNIPE(port string, force bool) {
	cmdStruct := exec.Command("ss", "-QHlntp", fmt.Sprintf("sport = :%s", port))
	out, err := cmdStruct.Output()
	if err != nil {
		fmt.Println(err)
	}

	processInfo, err := processResult(string(out))
	if err != nil {
		fmt.Println(err)
	}

	processInfo = handleSudo(processInfo)

	fmt.Printf("Port %s, with state '%s', is being used by process '%s' (PID: %d)\n", port, processInfo.State, processInfo.ProcessName, processInfo.PID)
	if force {
		killProcess(processInfo.PID, processInfo.sudo)
		return
	}
	fmt.Printf("Kill process %d ? (y/n): ", processInfo.PID)

	var res string
	_, err = fmt.Scanf("%s", &res)
	if err != nil {
		fmt.Println(err)
	}

	switch res {
	case "y":
		killProcess(processInfo.PID, processInfo.sudo)
	case "n":
		fmt.Println("Aborting...")
	default:
		fmt.Println("Invalid input. Aborting...")
	}
}

func killProcess(pid int, sudo bool) {
	fmt.Println("Killing process...")
	var cmdStruct *exec.Cmd
	if sudo {
		cmdStruct = exec.Command("sudo", "kill", strconv.Itoa(pid))
	} else {
		cmdStruct = exec.Command("kill", strconv.Itoa(pid))
	}
	_, err := cmdStruct.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Process %d killed successfully\n", pid)
	}
}

func handleSudo(info *SocketInfo) *SocketInfo {
	if info.PID == 0 {
		fmt.Printf("Process has root privileges. Trying with sudo...\n")
		cmdStruct := exec.Command("sudo", "ss", "-QHlntp", fmt.Sprintf("sport = :%s", info.LocalPort))
		out, err := cmdStruct.Output()
		if err != nil {
			fmt.Println(err)
		}

		processInfo, err := processResult(string(out))
		if err != nil {
			fmt.Println(err)
		}
		processInfo.sudo = true
		return processInfo
	}
	return info
}
