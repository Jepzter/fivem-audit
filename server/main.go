package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	portPtr := flag.String("port", "8090", "portnumber")
	flag.Parse()
	r := gin.Default()
	if _, err := os.Stat("./dist"); !os.IsNotExist(err) {
		r.Use(static.Serve("/", static.LocalFile("./dist", true)))
		r.NoRoute(func(c *gin.Context) {
			c.File("./dist/index.html")
		})
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // temp
	r.Use(cors.New(config))
	r.POST("/api/fivem-audit/auditlog-v1", storeAudit)
	r.GET("/api/fivem-audit/auditlog-v1", getAudit)

	r.Run(fmt.Sprintf(":%s", *portPtr))
}

func getAudit(c *gin.Context) {
	data, err := Load()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, data)
}

func storeAudit(c *gin.Context) {
	var log Audit
	err := c.BindJSON(&log)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	log.AuditID = uuid.New()
	log.Timestamp = time.Now()
	data, err := Load()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	data.Audits = append(data.Audits, log)
	handleAuditType(&data, log)
	err = Save(data)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(201, log)
}

func handleAuditType(data *Data, log Audit) error {
	switch log.AuditType {
	case PlayerJoining:
		data.Online++
	case PlayerDropping:
		if data.Online == 0 {
			return nil
		}
		data.Online--
	}
	return nil
}
