package controllers

import (
	"fmt"
	io "io/ioutil"
	"os"
	"strings"
)

//Agent is an device registered in SNMP.
type Agent struct {
	IP       string `json:"ip" binding:"required"`
	Hostname string `json:"hostname" binding:"required"`
	Comunity string `json:"comunity" binding:"required"`
}

//ActiveAgents returns all agents registered in hosts file.
func ActiveAgents() ([]Agent, error) {
	var agents []Agent
	content, err := io.ReadFile("./hosts.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		data := strings.Split(line, " ")
		if len(data) > 1 {
			a := Agent{
				IP:       data[1],
				Hostname: data[0],
				Comunity: data[2],
			}
			agents = append(agents, a)
		}
	}

	return agents, nil
}

//Add registers a new Agent in hosts file.
func (a Agent) Add() error {
	file, err := os.OpenFile("./hosts.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := a.Get(); err == nil {
		return ErrDuplicate
	}
	str := (a.Hostname + " " + a.IP + " " + a.Comunity + "\n")
	_, err = file.WriteString(str)
	if err != nil {
		return err
	}
	return nil
}

//Get returns an agent using its IP.
func (a *Agent) Get() error {
	content, err := io.ReadFile("./hosts.txt")
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.Contains(line, a.IP) {
			data := strings.Split(line, " ")
			a.Hostname = data[0]
			a.Comunity = data[2]
			return nil
		}
	}
	return ErrNotFound
}

//Delete removes an specific agent using its IP.
func (a *Agent) Delete() error {
	content, err := io.ReadFile("./hosts.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	idx := -1
	for i, line := range lines {
		if strings.Contains(line, a.IP) {
			idx = i
			break
		}
	}
	if idx < 0 {
		return ErrNotFound
	}
	lines = append(lines[:idx], lines[idx+1:]...)
	output := strings.Join(lines, "\n")
	err = io.WriteFile("./hosts.txt", []byte(output), 0644)
	if err != nil {
		return err
	}
	fmt.Println("deleted")
	return nil
}
