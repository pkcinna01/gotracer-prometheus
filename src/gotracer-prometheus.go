package main

# uses serial drivers from epsolar project:
#pkcinna@chronos:~/proj/gotracer-prometheus$ sudo insmod ./xr_usb_serial_common-1a/xr_usb_serial_common.ko
#pkcinna@chronos:~/proj/gotracer-prometheus$ sudo modprobe usbserial

import (
	"fmt"
	"log"
	"os"

	"github.com/spagettikod/gotracer"
)

type MpptInfo struct {
	pvGroup string
	title   string
}

var mpptByPort map[string]MpptInfo = map[string]MpptInfo{"/dev/ttyXRUSB0": MpptInfo{"HQST", "EPever 40A"}, "/dev/ttyXRUSB1": MpptInfo{"Renogy", "EPever 30A"}, "*": MpptInfo{"All", "All"}}

func prometheusExportField(port string, key string, value string) string {
	var deviceInfo string = ""
	if mpptInfo, found := mpptByPort[port]; found {
		deviceInfo = fmt.Sprintf(",title=\"%s\",pvGroup=\"%s\"", mpptInfo.title, mpptInfo.pvGroup)
	}
	recField := fmt.Sprintf("epsolar_%s{port=\"%s\"%s} %s", key, port, deviceInfo, value)
	return recField
}

func prometheusExport(usbPortPath string, tracerStatus *gotracer.TracerStatus) {
	fmt.Println(prometheusExportField(usbPortPath, "ArrayVoltage", fmt.Sprintf("%.2f", tracerStatus.ArrayVoltage)))
	fmt.Println(prometheusExportField(usbPortPath, "ArrayCurrent", fmt.Sprintf("%.2f", tracerStatus.ArrayCurrent)))
	fmt.Println(prometheusExportField(usbPortPath, "ArrayPower", fmt.Sprintf("%.2f", tracerStatus.ArrayPower)))
	fmt.Println(prometheusExportField(usbPortPath, "BatteryVoltage", fmt.Sprintf("%.2f", tracerStatus.BatteryVoltage)))
	fmt.Println(prometheusExportField(usbPortPath, "BatteryCurrent", fmt.Sprintf("%.2f", tracerStatus.BatteryCurrent)))
	fmt.Println(prometheusExportField(usbPortPath, "BatteryPower", fmt.Sprintf("%.2f", tracerStatus.BatteryVoltage*tracerStatus.BatteryCurrent)))
	fmt.Println(prometheusExportField(usbPortPath, "BatterySOC", fmt.Sprintf("%v", tracerStatus.BatterySOC)))
	fmt.Println(prometheusExportField(usbPortPath, "BatteryTemp", fmt.Sprintf("%.2f", tracerStatus.BatteryTemp)))
	fmt.Println(prometheusExportField(usbPortPath, "BatteryMaxVoltage", fmt.Sprintf("%.2f", tracerStatus.BatteryMaxVoltage)))
	fmt.Println(prometheusExportField(usbPortPath, "BatteryMinVoltage", fmt.Sprintf("%.2f", tracerStatus.BatteryMinVoltage)))
	fmt.Println(prometheusExportField(usbPortPath, "DeviceTemp", fmt.Sprintf("%.2f", tracerStatus.DeviceTemp)))

	fmt.Println(prometheusExportField(usbPortPath, "LoadVoltage", fmt.Sprintf("%.2f", tracerStatus.LoadVoltage)))
	fmt.Println(prometheusExportField(usbPortPath, "LoadCurrent", fmt.Sprintf("%.2f", tracerStatus.LoadCurrent)))
	fmt.Println(prometheusExportField(usbPortPath, "LoadPower", fmt.Sprintf("%.2f", tracerStatus.LoadPower)))

	fmt.Println(prometheusExportField(usbPortPath, "EnergyConsumedDaily", fmt.Sprintf("%.2f", tracerStatus.EnergyConsumedDaily)))
	fmt.Println(prometheusExportField(usbPortPath, "EnergyConsumedMonthly", fmt.Sprintf("%.2f", tracerStatus.EnergyConsumedMonthly)))
	fmt.Println(prometheusExportField(usbPortPath, "EnergyConsumedAnnual", fmt.Sprintf("%.2f", tracerStatus.EnergyConsumedAnnual)))
	fmt.Println(prometheusExportField(usbPortPath, "EnergyConsumedTotal", fmt.Sprintf("%.2f", tracerStatus.EnergyConsumedTotal)))
	fmt.Println(prometheusExportField(usbPortPath, "EnergyGeneratedDaily", fmt.Sprintf("%.2f", tracerStatus.EnergyGeneratedDaily)))
	fmt.Println(prometheusExportField(usbPortPath, "EnergyGeneratedMonthly", fmt.Sprintf("%.2f", tracerStatus.EnergyGeneratedMonthly)))
	fmt.Println(prometheusExportField(usbPortPath, "EnergyGeneratedAnnual", fmt.Sprintf("%.2f", tracerStatus.EnergyGeneratedAnnual)))
	fmt.Println(prometheusExportField(usbPortPath, "EnergyGeneratedTotal", fmt.Sprintf("%.2f", tracerStatus.EnergyGeneratedTotal)))
	//t.BatterySOC, t.BatteryTemp, t.BatteryMaxVoltage, t.BatteryMinVoltage, t.DeviceTemp, t.LoadVoltage, t.LoadCurrent, t.LoadPower, t.Load, t.EnergyConsumedDaily, t.EnergyConsumedMonthly, t.EnergyConsumedAnnual, t.EnergyConsumedTotal, t.EnergyGeneratedDaily, t.EnergyGeneratedMonthly, t.EnergyGeneratedAnnual, t.EnergyGeneratedTotal)
}

func main() {

	usbPortPaths := []string{"/dev/ttyXRUSB0", "/dev/ttyXRUSB1" }
	errCnt := 0

	for _, usbPortPath := range usbPortPaths {
		//fmt.Printf("%s Status:\n", usbPortPath)
		status, err := gotracer.Status(usbPortPath)
		if err != nil {
			errCnt++
			log.Printf("ERROR %s: %+v", usbPortPath, err)
		} else {
			prometheusExport(usbPortPath, &status)
		}
		serialReadOk := "1" 
		if err != nil {
			serialReadOk = "0"
		}
		fmt.Println(prometheusExportField(usbPortPath, "SerialReadOk", serialReadOk))
	}

	fmt.Println(prometheusExportField("*", "SerialReadErrorCnt", fmt.Sprintf("%d",errCnt)))
	os.Exit(errCnt)
}
