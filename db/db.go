package db

import (
	"context"
	"encoding/json"
	"fmt"
	pb "gRPC/pb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
)

type DeviceService struct{}

// Devices stores the device definition
var Devices pb.Devices

//var Device pb.Device
var tmpDevices devicesLoad
var tmpDevice deviceLoad

type devicesLoad struct {
	Items []deviceLoad
}

type deviceLoad struct {
	ID       int    `json:"id"`
	Hardware string `json:"hardware"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	State    int    `json:"state"`
}

// InitDB initialize the "db" with data.json file
func InitDB() {

	dataFile, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println("error reading data file: ", err)
	}
	if err := json.Unmarshal(dataFile, &tmpDevices); err != nil {
		log.Println("error parsing data file ", err)
	}
	fmt.Println("db init ok: ", tmpDevices)
	for _, val := range tmpDevices.Items {
		Device := new(pb.Device)
		Device.Id = int32(val.ID)
		Device.Hardware = val.Hardware
		Device.Location = val.Location
		Device.Name = val.Name
		Device.State = int32(val.State)
		Device.Type = getDeviceType(val.Type)

		Devices.Device = append(Devices.Device, Device)
	}
}

func getDeviceType(s string) pb.Device_DeviceType {
	var ret pb.Device_DeviceType
	for v, k := range pb.Device_DeviceType_value {
		if v == s {
			ret = pb.Device_DeviceType(k)
			return ret
		}
	}
	return ret
}

func (s *DeviceService) GetAllDevices(ctx context.Context, empty *pb.Empty) (*pb.Devices, error) {
	return nil, nil
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

func Run() {
	InitDB()
	grpcPort := "localhost:8082"
	// start listening for grpc
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}
	// Create new grpc server
	server := grpc.NewServer()
	// Register service
	pb.RegisterDeviceServiceServer(server, new(DeviceService))
	// Start serving requests
	server.Serve(listen)
}
