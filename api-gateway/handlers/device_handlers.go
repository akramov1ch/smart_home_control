package handlers

import (
	"api-gateway/config"
	dv "api-gateway/proto/device"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var deviceService dv.DeviceServiceClient

func init() {
	conn, err := grpc.NewClient(config.GetConfig().DeviceServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to device service: %v", err)
	}
	deviceService = dv.NewDeviceServiceClient(conn)
}

func AddDevice(c *gin.Context) {
	var req dv.AddDeviceRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := deviceService.AddDevice(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetDevices(c *gin.Context) {
	resp, err := deviceService.GetDevices(context.Background(), &dv.GetDevicesRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetDevice(c *gin.Context) {
	id := c.Param("id")

	resp, err := deviceService.GetDevice(context.Background(), &dv.GetDeviceRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateDevice(c *gin.Context) {
	id := c.Param("id")
	var req dv.UpdateDeviceRequest
	req.Id = id

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := deviceService.UpdateDevice(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteDevice(c *gin.Context) {
	id := c.Param("id")

	_, err := deviceService.DeleteDevice(context.Background(), &dv.DeleteDeviceRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device deleted"})
}

func ControlDevice(c *gin.Context) {
	id := c.Param("id")
	var req dv.ControlDeviceRequest
	req.Id = id

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := deviceService.ControlDevice(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
