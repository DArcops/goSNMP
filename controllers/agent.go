package controllers

import (
	"fmt"
	io "io/ioutil"
	"os"
	"strings"
)

func ActiveAgents() (int, error) {
	content, err := io.ReadFile("./hosts.txt")
	if err != nil {
		return 0, err
	}
	strContent := string(content)
	lines := strings.Count(strContent, "\n")
	return lines, nil
}

func Add(hostname string, ip string, comunity string) error {
	file, err := os.OpenFile("./hosts.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	str := (hostname + " " + ip + " " + comunity + "\n")

	_, err = file.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}

func Delete(hostname string, ip string, comunity string) error {
	content, err := io.ReadFile("./hosts.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if strings.Contains(line, ip) {
			fmt.Println("!!!match", ip)
			lines[i] = ""
		}
	}

	output := strings.Join(lines, "\n")
	err = io.WriteFile("./hosts.txt", []byte(output), 0644)
	if err != nil {
		return err
	}
	fmt.Println("deleted")
	return nil
}
