package api

import (
	"net/http"

	"github.com/DArcops/goSNMP/controllers"
	"github.com/gin-gonic/gin"
)

//AddAgent registers an agent recieved in JSON request.
func AddAgent(c *gin.Context) {
	var agent controllers.Agent
	if err := c.Bind(&agent); err != nil {
		JSONError(http.StatusBadRequest, err, c)
		return
	}
	err := agent.Add()
	if err != nil {
		status := http.StatusInternalServerError
		if err == controllers.ErrDuplicate {
			status = http.StatusConflict
		}
		JSONError(status, err, c)
		return
	}
	c.JSON(http.StatusOK, agent)
}

//DeleteAgent deletes an agent using param ip from request.
func DeleteAgent(c *gin.Context) {
	agent := controllers.Agent{
		IP: c.Param("ip"),
	}
	err := agent.Delete()
	if err != nil {
		status := http.StatusInternalServerError
		if err == controllers.ErrNotFound {
			status = http.StatusNotFound
		}
		JSONError(status, err, c)
		return
	}
	c.JSON(http.StatusOK, agent)
}

//GetAgent returns information about agent using param ip from request.
func GetAgent(c *gin.Context) {
	agent := controllers.Agent{
		IP: c.Param("ip"),
	}
	err := agent.Get()
	if err != nil {
		status := http.StatusInternalServerError
		if err == controllers.ErrNotFound {
			status = http.StatusNotFound
		}
		JSONError(status, err, c)
		return
	}
	//c.JSON(http.StatusOK, agent)
	c.HTML(http.StatusOK, "agent.tmpl", agent)
}

//GetAgents returns all agents registered.
func GetAgents(c *gin.Context) {
	var agents []controllers.Agent
	agents, err := controllers.ActiveAgents()
	if err != nil {
		JSONError(http.StatusInternalServerError, err, c)
		return
	}
	//c.JSON(http.StatusOK, agents)
	c.HTML(http.StatusOK, "index.tmpl", agents)

}
