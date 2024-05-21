package api

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wallet/base"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func RouterRunOnPhone() {
	router := setupRouterOnPhone()
	err := router.Run("0.0.0.0:9090")
	if err != nil {
		return
	}
}

func setupRouterOnPhone() *gin.Engine {
	router := gin.Default()

	router.POST("/addInvoice", func(c *gin.Context) {

		id := uuid.New().String()
		amountStr := c.PostForm("amount")
		amountInt, err := strconv.Atoi(amountStr)
		result := true
		if err != nil || amountInt <= 0 {
			result = false
			fmt.Printf("%s amountInt less than or equal to zero || strconv.Atoi(amount) :%v\n", GetTimeNow(), err)
		}
		var invoiceStr string
		if result {
			invoiceStr = AddInvoice(int64(amountInt), "")
		}
		if invoiceStr == "" {
			result = false
		}

		err = InitPhoneDB()
		if err != nil {
			fmt.Printf("%s InitPhoneDB err :%v\n", GetTimeNow(), err)
		}

		db, err := bolt.Open(filepath.Join(base.QueryConfigByKey("dirpath"), "phone.db"), 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
		}
		defer func(db *bolt.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
			}
		}(db)
		s := &PhoneStore{DB: db}

		if result {
			invoiceStr = strings.ToUpper(invoiceStr)
			err = s.CreateOrUpdateInvoice("invoices", &Invoice{
				ID:         id,
				Amount:     amountInt,
				InvoiceStr: invoiceStr,
			})
			if err != nil {
				fmt.Printf("%s CreateOrUpdateInvoice err :%v\n", GetTimeNow(), err)
				result = false
			}
		} else {
			id = ""
		}

		c.JSON(http.StatusOK, gin.H{
			"time":    GetTimeNow(),
			"id":      id,
			"amount":  amountStr,
			"invoice": invoiceStr,
			"result":  result,
		})
	})

	//TODO: Need to test
	router.POST("/newAddr", func(c *gin.Context) {
		assetID := c.PostForm("asset_id")
		if assetID == "" {
			fmt.Printf("%s asset id null.\n", GetTimeNow())
			c.JSON(http.StatusOK, JsonResult{
				Success: false,
				Error:   "asset id null.",
				Data:    "",
			})
		}
		amountStr := c.PostForm("amount")
		amountInt, err := strconv.Atoi(amountStr)
		if err != nil || amountInt <= 0 {
			fmt.Printf("%s amountInt less than or equal to zero || strconv.Atoi(amount) :%v\n", GetTimeNow(), err)
			c.JSON(http.StatusOK, JsonResult{
				Success: false,
				Error:   "amountInt less than or equal to zero || strconv.Atoi(amount).",
				Data:    "",
			})
		}
		addr, err := newAddr(assetID, amountInt)
		if err != nil {
			LogError("new addr.", err)
			c.JSON(http.StatusOK, JsonResult{
				Success: false,
				Error:   "new addr. " + err.Error(),
				Data:    "",
			})
		}
		addrStr := addr.Encoded
		c.JSON(http.StatusOK, JsonResult{
			Success: true,
			Error:   "",
			Data:    addrStr,
		})
	})
	return router
}
