package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	myWifi "spbstu.ru/nadia.voronina/task-6/internal/wifi"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var (
	errExpected  = errors.New("ExpectedError")
	errSomeError = errors.New("some error")
)

type rowTestSysInfo struct {
	addrs       []string
	errExpected error
}

func getTestTable() []rowTestSysInfo {
	return []rowTestSysInfo{
		{
			addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		},
		{
			errExpected: errExpected,
		},
	}
}

func TestGetAddresses_EmptyInterfaces(t *testing.T) {
	t.Parallel()
	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	mockWifi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	addrs, err := wifiService.GetAddresses()
	require.NoError(t, err)
	require.Empty(t, addrs)
}

func TestGetAddresses_InvalidMAC(t *testing.T) {
	t.Parallel()
	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	ifaces := []*wifi.Interface{
		{
			Index:        1,
			Name:         "eth1",
			HardwareAddr: nil,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		},
	}

	mockWifi.On("Interfaces").Return(ifaces, nil)

	addrs, err := wifiService.GetAddresses()
	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{nil}, addrs)
}

func TestGetAddresses_MultipleInterfaces(t *testing.T) {
	t.Parallel()
	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	macs := []string{"01:23:45:67:89:ab", "de:ad:be:ef:00:01"}
	mockWifi.On("Interfaces").Return(mockIfaces(macs), nil)

	addrs, err := wifiService.GetAddresses()
	require.NoError(t, err)
	require.Equal(t, parseMACs(macs), addrs)
}

func TestGetName(t *testing.T) {
	t.Parallel()
	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	for i, row := range getTestTable() {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)

		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)

			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs,
			"row: %d, expected addrs: %s, actual addrs: %s", i,
			parseMACs(row.addrs), actualAddrs)
	}
}

func TestGetNames_EmptyInterfaces(t *testing.T) {
	t.Parallel()

	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	mockWifi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	names, err := wifiService.GetNames()
	require.NoError(t, err)
	require.Empty(t, names)
}

func TestGetNames_MultipleInterfaces(t *testing.T) {
	t.Parallel()

	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)
	ifaces := []*wifi.Interface{
		{Index: 1, Name: "wlan0"},
		{Index: 2, Name: "eth1"},
		{Index: 3, Name: "wifi42"},
	}
	mockWifi.On("Interfaces").Return(ifaces, nil)

	names, err := wifiService.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "eth1", "wifi42"}, names)
}

func TestGetNames_ErrorFromInterfaces(t *testing.T) {
	t.Parallel()
	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	mockWifi.On("Interfaces").Return(nil, errSomeError)

	names, err := wifiService.GetNames()
	require.ErrorIs(t, err, errSomeError)
	require.Nil(t, names)
}

func TestGetNames_NilName(t *testing.T) {
	t.Parallel()
	mockWifi := myWifi.NewMockWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	ifaces := []*wifi.Interface{
		{Index: 1, Name: ""},
		{Index: 2, Name: "eth2"},
	}
	mockWifi.On("Interfaces").Return(ifaces, nil)

	names, err := wifiService.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"", "eth2"}, names)
}

func mockIfaces(addrs []string) []*wifi.Interface {
	interfaces := make([]*wifi.Interface, 0, len(addrs))

	for i, addrStr := range addrs {
		hwAddr := parseMAC(addrStr)
		if hwAddr == nil {
			continue
		}

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i+1),
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces = append(interfaces, iface)
	}

	return interfaces
}

func parseMACs(macStr []string) []net.HardwareAddr {
	addrs := make([]net.HardwareAddr, 0, len(macStr))
	for _, addr := range macStr {
		addrs = append(addrs, parseMAC(addr))
	}

	return addrs
}

func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}

	return hwAddr
}
