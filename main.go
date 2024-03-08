package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// GPU Temperature
	tempFilePath := "/sys/class/drm/card0/device/hwmon/hwmon3/temp1_input"
	tempStr, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		fmt.Println("Error reading temperature:", err)
		return
	}

	// The temperature is usually reported in millidegrees Celsius; convert to degrees
	tempMilliDegrees := strings.TrimSpace(string(tempStr))
	tempDegrees, err := strconv.Atoi(tempMilliDegrees)
	if err != nil {
		fmt.Println("Error converting temperature:", err)
		return
	}

	// GPU Average Power

	tempFilePath := "/sys/class/drm/card0/device/hwmon/hwmon3/temp1_input"
	tempStr, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		fmt.Println("Error reading temperature:", err)
		return
	}

	// The temperature is usually reported in millidegrees Celsius; convert to degrees
	tempMilliDegrees := strings.TrimSpace(string(tempStr))
	tempDegrees, err := strconv.Atoi(tempMilliDegrees)
	if err != nil {
		fmt.Println("Error converting temperature:", err)
		return
	}

	fmt.Printf("GPU Temperature: %.2f°C\n", float64(tempDegrees)/1000.0)
}

// package main

// //#define CL_TARGET_OPENCL_VERSION 300
// // #cgo LDFLAGS: -lOpenCL
// // #include <CL/cl.h>
// import "C"
// import (
// 	"fmt"
// 	"time"
// 	"unsafe"
// )

// func main() {
// 	for {
// 		displayGPUs()
// 		fmt.Println("Updating in 5 seconds...")
// 		time.Sleep(5 * time.Second) // Wait for 5 seconds before updating
// 	}
// }
// func displayGPUs() {
// 	// Get the number of OpenCL platforms
// 	var numPlatforms C.cl_uint
// 	if C.clGetPlatformIDs(0, nil, &numPlatforms) != C.CL_SUCCESS {
// 		fmt.Println("Unable to get platform count.")
// 		return
// 	}

// 	// Get the list of platform IDs
// 	platforms := make([]C.cl_platform_id, numPlatforms)
// 	if C.clGetPlatformIDs(numPlatforms, &platforms[0], nil) != C.CL_SUCCESS {
// 		fmt.Println("Unable to get platform IDs.")
// 		return
// 	}

// 	fmt.Printf("Found %d OpenCL platforms.\n", numPlatforms)

// 	// Iterate over platforms to get GPU devices
// 	for i, platform := range platforms {
// 		fmt.Printf("Platform %d:\n", i)

// 		// Get the number of GPU devices available
// 		var numDevices C.cl_uint
// 		status := C.clGetDeviceIDs(platform, C.CL_DEVICE_TYPE_GPU, 0, nil, &numDevices)
// 		if status != C.CL_SUCCESS {
// 			fmt.Printf("Unable to get device count. Error code: %d\n", status)
// 			continue
// 		}

// 		// Get the list of GPU device IDs
// 		devices := make([]C.cl_device_id, numDevices)
// 		if C.clGetDeviceIDs(platform, C.CL_DEVICE_TYPE_GPU, numDevices, &devices[0], nil) != C.CL_SUCCESS {
// 			fmt.Println("Unable to get device IDs.")
// 			continue
// 		}

// 		fmt.Printf("Found %d GPU devices.\n", numDevices)

// 		// Iterate over each device to print its information
// 		for j, device := range devices {
// 			var deviceName [1024]byte
// 			if C.clGetDeviceInfo(device, C.CL_DEVICE_NAME, C.size_t(unsafe.Sizeof(deviceName)), unsafe.Pointer(&deviceName), nil) != C.CL_SUCCESS {
// 				fmt.Println("Unable to get device name.")
// 				continue
// 			}
// 			fmt.Printf("Device %d: %s\n", j, C.GoString((*C.char)(unsafe.Pointer(&deviceName))))
// 		}
// 	}
// }
