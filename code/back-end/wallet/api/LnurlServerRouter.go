package api

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/wallet/base"
	"net/http"
	"strconv"
	"time"
)

func RouterRunOnServer() {
	router := setupRouterOnServer()
	err := router.Run("0.0.0.0:9080")
	if err != nil {
		return
	}
}

func setupRouterOnServer() *gin.Engine {
	router := gin.Default()

	router.POST("/upload/user", func(c *gin.Context) {
		id := c.PostForm("id")
		name := c.PostForm("name")
		ip := c.ClientIP()
		localPort := c.PostForm("local_port")
		socket := ip + ":" + localPort
		remotePort := c.PostForm("remote_port")
		result := true
		if name == "" || localPort == "" {
			result = false
		}
		user := &User{
			ID:         id,
			Name:       name,
			Socket:     socket,
			RemotePort: remotePort,
		}
		err := InitServerDB()
		if err != nil {
			fmt.Printf("%s InitServerDB err :%v\n", GetTimeNow(), err)
		}
		db, err := bolt.Open("server.db", 0600, &bolt.Options{
			Timeout: 1 * time.Second,
		})
		if err != nil {
			fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
		}
		defer func(db *bolt.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
			}
		}(db)
		s := &ServerStore{DB: db}
		if result {
			err = s.CreateOrUpdateUser("users", user)
			if err != nil {
				fmt.Printf("%s CreateOrUpdateUser err :%v\n", GetTimeNow(), err)
				result = false
			}
		}
		var lnurlStr string
		serverDomainOrSocket := base.QueryConfigByKey("LnurlServerHost")
		if result {
			lnurlStr = Encode("http://" + serverDomainOrSocket + "/pay?id=" + id)
		} else {
			id = ""
		}
		c.JSON(http.StatusOK, gin.H{
			"time":        GetTimeNow(),
			"id":          id,
			"name":        name,
			"socket":      socket,
			"remote_port": remotePort,
			"result":      result,
			"lnurl":       lnurlStr,
		})
	})

	router.POST("/pay", func(c *gin.Context) {
		id := c.Query("id")
		amount := c.PostForm("amount")
		result := true
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			result = false
			fmt.Printf("%s strconv.Atoi(amount) :%v\n", GetTimeNow(), err)
		}
		if id == "" || amount == "" || amountInt <= 0 {
			result = false
		}
		err = InitServerDB()
		if err != nil {
			fmt.Printf("%s InitServerDB err :%v\n", GetTimeNow(), err)
		}
		db, err := bolt.Open("server.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
		}
		defer func(db *bolt.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
			}
		}(db)
		s := &ServerStore{DB: db}
		user, err := s.ReadUser("users", id)
		if err != nil {
			fmt.Printf("%s ReadUser err :%v\n", GetTimeNow(), err)
		}

		var invoice string
		if result {
			//@dev: RemotePort
			invoice = PostPhoneToAddInvoice(user.RemotePort, amount)
		}
		if invoice == "" {
			result = false
		}

		c.JSON(http.StatusOK, gin.H{
			"time":    GetTimeNow(),
			"id":      id,
			"amount":  amount,
			"invoice": invoice,
			"result":  result,
		})
	})

	router.GET("/availablePort", func(c *gin.Context) {
		remotePort := QueryAvailablePort()
		result := true
		if remotePort == 0 {
			result = false
		}
		c.JSON(http.StatusOK, gin.H{
			"time":        GetTimeNow(),
			"remote_port": remotePort,
			"result":      result,
		})
	})

	return router
}
