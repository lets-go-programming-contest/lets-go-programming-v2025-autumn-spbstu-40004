package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/faxryzen/task-6/internal/wifi"
	wifimock "github.com/faxryzen/task-6/internal/wifi/_mocks"

	wifilib "github.com/mdlayher/wifi"
)

var errFail = errors.New("fail")

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	interfaces := []*wifilib.Interface{
		{HardwareAddr: net.HardwareAddr{0xAA, 0xBB, 0xCC}},
		{HardwareAddr: net.HardwareAddr{0x11, 0x22, 0x33}},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	addrs, err := service.GetAddresses()
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}

	if len(addrs) != 2 {
		t.Fatalf("unexpected: %v", addrs)
	}
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errFail)

	_, err := service.GetAddresses()
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	interfaces := []*wifilib.Interface{
		{Name: "wlan0"},
		{Name: "wlan1"},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	names, err := service.GetNames()
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}

	if len(names) != 2 || names[0] != "wlan0" || names[1] != "wlan1" {
		t.Fatalf("unexpected: %v", names)
	}
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := new(wifimock.WiFiHandle)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return(nil, errFail)

	_, err := service.GetNames()
	if err == nil {
		t.Fatalf("expected error")
	}
}
