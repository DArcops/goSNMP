package main

import (
	"github.com/DArcops/goSNMP/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "/assets")
	apiRoutes := router.Group("api")
	{
		agent := apiRoutes.Group("agent")
		{
			agent.POST("", api.AddAgent)
			agent.GET("", api.GetAgents)
			agent.GET(":ip", api.GetAgent)
			agent.DELETE(":ip", api.DeleteAgent)
		}
	}
	router.Run(":8000")
}

/*
func main() {
	//check how many agents are active
	agents, err := c.ActiveAgents()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("from main function", agents)

	//When you recive request for add a new agent , you need hostname, comunity and ip
	//maybe in the API recive that params by get request
	agentIP := "8.25.100.12"
	agentComunity := "SNMPCom"
	agentHostname := "windows"

	err = c.Add(agentHostname, agentIP, agentComunity)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Agent added")

	c.Delete(agentHostname, "8.25.100.11", agentComunity)

	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = "192.168.1.10"
	g.Default.Version = 0x1
	err = g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
}*/
