package main

import (
	"bytes"
	"net"
	"os"
	"os/exec"
	"strings"
)

// Obtener UUID del sistema usando PowerShell
func GetSystemUUID() (string, error) {
	cmd := exec.Command("powershell", "Get-WmiObject -Class Win32_ComputerSystemProduct | Select-Object -ExpandProperty UUID")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// Obtener la primera MAC address válida
func GetMACAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.HardwareAddr != nil && len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String(), nil
		}
	}
	return "", nil
}
func GetHostname() (string, error) {
	return os.Hostname()
}
