package db

import (
	"context"
	pb "gRPC/pb"
)

type DeviceService struct{}

func (s *DeviceService) GetAllDevices(ctx context.Context, empty *pb.Empty) (*pb.Devices, error) {
	//TODO implement me
	panic("implement me")
}

func (s *DeviceService) GetDeviceByID(ctx context.Context, id *pb.ID) (*pb.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (s *DeviceService) SwitchDevice(ctx context.Context, device *pb.UpdateDevice) (*pb.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (s *DeviceService) RegisterDevice(ctx context.Context, device *pb.Device) (*pb.Device, error) {
	//TODO implement me
	panic("implement me")
}
