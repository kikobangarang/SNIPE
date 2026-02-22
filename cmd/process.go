package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func processResult(result string) (*SocketInfo, error) {
	fields := strings.Fields(result)
	if len(fields) < 3 {
		return nil, fmt.Errorf("wrong format from ss command: %s", result)
	}

	info := &SocketInfo{
		State: fields[0],
	}

	// Parse local address (*:3000)
	localParts := strings.Split(fields[1], ":")
	if len(localParts) == 2 {
		info.LocalPort = localParts[1]
	}

	// Regex to extract process info
	re := regexp.MustCompile(`"([^"]+)",pid=(\d+),fd=(\d+)`)
	matches := re.FindStringSubmatch(result)
	if len(matches) == 4 {
		info.ProcessName = matches[1]
		pid, _ := strconv.Atoi(matches[2])
		info.PID = pid
	}

	return info, nil
}

type SocketInfo struct {
	State       string
	LocalPort   string
	ProcessName string
	PID         int
	sudo        bool
}
