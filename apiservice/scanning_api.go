package apiservice

import (
	"strconv"

	"app-server/logic"
	"app-server/models"

	"github.com/gin-gonic/gin"
)

func GetScanListByTag1(index, tag string) *models.ScanningList {
	startIdx, err := strconv.ParseUint(index, 10, 32)
	if err != nil {
		return nil
	}

	return logic.GetScanListByTag1(uint32(startIdx), tag)
}

func CreateScanning(c *gin.Context) uint32 {
	var bind models.SendScannigBinding
	err := c.BindJSON(&bind)
	if err != nil {
		return 10356
	}

	return logic.CreateScanning(&bind)
}

func GetScanningRedpkt(redpktId, userId string) uint32 {
	return logic.GetScanningRedpkt(redpktId, userId)
}
