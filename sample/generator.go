package sample

import (
	"github.com/tatrasoft/grpc-basics/pb"
	"time"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	noOfCores := randomInt(2, 8)
	noOfThreads := randomInt(noOfCores, 12)
	minGhz := randomFloat32(2.0, 3.5)
	maxGhz := randomFloat32(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(noOfCores),
		NumberThreads: uint32(noOfThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	minGhz := randomFloat32(2.0, 3.5)
	maxGhz := randomFloat32(minGhz, 5.0)
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   randomGPUName(brand),
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: NewMemory(),
	}

	return gpu
}

func NewMemory() *pb.Memory {

	memory := &pb.Memory{
		Value: uint64(randomInt(1, 16)),
		Unit:  randomMemoryUnit(),
	}

	return memory
}

func NewScreen() *pb.Screen {
	height := uint32(randomInt(1080, 4320))
	width := height * 16 / 9

	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: &pb.Screen_Resolution{
			Width:  width,
			Height: height,
		},
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128,1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1,6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()

	laptop := &pb.Laptop{
		Id:          randomID(),
		Brand:       brand,
		Name:        randomLaptopName(brand),
		Cpu:         NewCPU(),
		Ram:         NewMemory(),
		Gpus:        []*pb.GPU{NewGPU()},
		Storages:    []*pb.Storage{NewHDD(), NewSSD()},
		Screen:      NewScreen(),
		Keyboard:    NewKeyboard(),
		Weight:      &pb.Laptop_WightKg{WightKg:float64(randomFloat32(1, 3))},
		PriceUsd:    float64(randomFloat32(1500, 3000)),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   time.Now().Unix(),
	}

	return laptop
}

func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}

