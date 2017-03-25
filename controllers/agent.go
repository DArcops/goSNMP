package controllers

import (
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
	file, err := os.Open("./hosts.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	str := (hostname + " " + ip + " " + comunity)
	data := []byte(str)

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func Delete(hostname string, comunity string) error {
	return nil
}
