package handlers

import (
	"context"

	"device-control-service/models"
	dv "device-control-service/proto/device"
	"device-control-service/repository"
)

type DeviceHandler struct {
	dv.UnimplementedDeviceServiceServer
	Repo repository.DeviceRepository
}

func (h *DeviceHandler) AddDevice(ctx context.Context, req *dv.AddDeviceRequest) (*dv.AddDeviceResponse, error) {
	device := models.Device{
		Name: req.Name,
		Type: req.Type,
	}
	id, err := h.Repo.AddDevice(device)
	if err != nil {
		return nil, err
	}
	return &dv.AddDeviceResponse{Id: id, Name: req.Name, Type: req.Type}, nil
}

func (h *DeviceHandler) GetDevices(ctx context.Context, req *dv.GetDevicesRequest) (*dv.GetDevicesResponse, error) {
	devices, err := h.Repo.GetDevices()
	if err != nil {
		return nil, err
	}
	var protoDevices []*dv.Device
	for _, device := range devices {
		protoDevices = append(protoDevices, &dv.Device{Id: device.ID, Name: device.Name, Type: device.Type})
	}
	return &dv.GetDevicesResponse{Devices: protoDevices}, nil
}

func (h *DeviceHandler) GetDevice(ctx context.Context, req *dv.GetDeviceRequest) (*dv.GetDeviceResponse, error) {
	device, err := h.Repo.GetDevice(req.Id)
	if err != nil {
		return nil, err
	}
	return &dv.GetDeviceResponse{Id: device.ID, Name: device.Name, Type: device.Type}, nil
}

func (h *DeviceHandler) UpdateDevice(ctx context.Context, req *dv.UpdateDeviceRequest) (*dv.UpdateDeviceResponse, error) {
	device := models.Device{
		ID:   req.Id,
		Name: req.Name,
		Type: req.Type,
	}
	err := h.Repo.UpdateDevice(device)
	if err != nil {
		return nil, err
	}
	return &dv.UpdateDeviceResponse{Id: device.ID, Name: device.Name, Type: device.Type}, nil
}

func (h *DeviceHandler) DeleteDevice(ctx context.Context, req *dv.DeleteDeviceRequest) (error) {
	err := h.Repo.DeleteDevice(req.Id)
	if err != nil {
		return err
	}
	return nil
}