package handlers

import (
	"context"

	cn "device-control-service/proto/control"
)

type ControlHandler struct {
	cn.UnimplementedControllerServiceServer
}

func (h *ControlHandler) ControlDevice(ctx context.Context, req *cn.ControlDeviceRequest) (*cn.ControlDeviceResponse, error) {
	result := "Device " + req.DeviceId + " action " + req.Action + " performed"
	return &cn.ControlDeviceResponse{Result: result}, nil
}
