package wifi_test

import (
	"errors"
	"net"
	"testing"

	mywifi "github.com/arinaklimova/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg=wifi_test --output=.

func createMockInterface(name string, mac string) *wifi.Interface {
	hwAddr, _ := net.ParseMAC(mac)
	return &wifi.Interface{
		Index:        1,
		Name:         name,
		HardwareAddr: hwAddr,
		PHY:          0,
		Device:       0,
		Type:         wifi.InterfaceTypeStation,
		Frequency:    2412,
	}
}

func parseMACs(addrs []string) []net.HardwareAddr {
	hwAddrs := make([]net.HardwareAddr, 0, len(addrs))
	for _, addr := range addrs {
		hwAddr, _ := net.ParseMAC(addr)
		hwAddrs = append(hwAddrs, hwAddr)
	}
	return hwAddrs
}

func TestNew(t *testing.T) {
	mockWiFi := &WiFiHandle{}
	service := mywifi.New(mockWiFi)

	assert.NotNil(t, service)
	assert.Equal(t, mockWiFi, service.WiFi)
}

func TestWiFiService_GetAddresses(t *testing.T) {
	t.Run("successful query", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		expectedInterfaces := []*wifi.Interface{
			createMockInterface("wlan0", "00:11:22:33:44:55"),
			createMockInterface("wlan1", "aa:bb:cc:dd:ee:ff"),
		}

		mockWiFi.On("Interfaces").Return(expectedInterfaces, nil)

		addresses, err := service.GetAddresses()

		assert.NoError(t, err)
		require.Len(t, addresses, 2)

		expectedAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
		expectedAddr2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")
		assert.Equal(t, expectedAddr1, addresses[0])
		assert.Equal(t, expectedAddr2, addresses[1])

		mockWiFi.AssertExpectations(t)
	})

	t.Run("interfaces error", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		mockWiFi.On("Interfaces").Return(nil, errors.New("wifi error"))

		addresses, err := service.GetAddresses()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "getting interfaces:")
		assert.Nil(t, addresses)
		mockWiFi.AssertExpectations(t)
	})

	t.Run("empty interfaces", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

		addresses, err := service.GetAddresses()

		assert.NoError(t, err)
		assert.Empty(t, addresses)
		mockWiFi.AssertExpectations(t)
	})

	t.Run("interface without MAC address", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		iface := &wifi.Interface{
			Index:        1,
			Name:         "wlan0",
			HardwareAddr: nil,
			PHY:          0,
			Device:       0,
			Type:         wifi.InterfaceTypeStation,
			Frequency:    2412,
		}

		mockWiFi.On("Interfaces").Return([]*wifi.Interface{iface}, nil)

		addresses, err := service.GetAddresses()

		assert.NoError(t, err)
		require.Len(t, addresses, 1)
		assert.Nil(t, addresses[0])
		mockWiFi.AssertExpectations(t)
	})
}

func TestWiFiService_GetNames(t *testing.T) {
	t.Run("successful query", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		expectedInterfaces := []*wifi.Interface{
			createMockInterface("wlan0", "00:11:22:33:44:55"),
			createMockInterface("wlan1", "aa:bb:cc:dd:ee:ff"),
			createMockInterface("eth0", "11:22:33:44:55:66"),
		}

		mockWiFi.On("Interfaces").Return(expectedInterfaces, nil)

		names, err := service.GetNames()

		assert.NoError(t, err)
		require.Len(t, names, 3)
		assert.Equal(t, "wlan0", names[0])
		assert.Equal(t, "wlan1", names[1])
		assert.Equal(t, "eth0", names[2])
		mockWiFi.AssertExpectations(t)
	})

	t.Run("interfaces error", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		mockWiFi.On("Interfaces").Return(nil, errors.New("wifi error"))

		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "getting interfaces:")
		assert.Nil(t, names)
		mockWiFi.AssertExpectations(t)
	})

	t.Run("empty interfaces", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

		names, err := service.GetNames()

		assert.NoError(t, err)
		assert.Empty(t, names)
		mockWiFi.AssertExpectations(t)
	})

	t.Run("duplicate interface names", func(t *testing.T) {
		mockWiFi := &WiFiHandle{}
		service := mywifi.New(mockWiFi)

		expectedInterfaces := []*wifi.Interface{
			createMockInterface("wlan0", "00:11:22:33:44:55"),
			createMockInterface("wlan0", "aa:bb:cc:dd:ee:ff"),
		}

		mockWiFi.On("Interfaces").Return(expectedInterfaces, nil)

		names, err := service.GetNames()

		assert.NoError(t, err)
		require.Len(t, names, 2)
		assert.Equal(t, "wlan0", names[0])
		assert.Equal(t, "wlan0", names[1])
		mockWiFi.AssertExpectations(t)
	})
}
