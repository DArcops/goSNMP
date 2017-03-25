package controllers

import (
	io "io/ioutil"
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
	return nil
}

func Delete(hostname string, comunity string) error {
	return nil
}
