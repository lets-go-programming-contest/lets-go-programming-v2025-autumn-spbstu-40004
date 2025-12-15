package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	mywifi "github.com/Ksenia-rgb/task-6/internal/wifi"
	wifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output

var ErrExpected = errors.New("error expected")

var testTable = [][]string{
	{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
	{"00:01:02:03:04:05", "aa:ab:ac:ad:ae:af"},
}

func TestGetAddressesSuccess(t *testing.T) {
	mockWifi := NewWiFiHandle(t)
	wifiServece := mywifi.New(mockWifi)

	for _, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row), nil)
		actualAddrs, err := wifiServece.GetAddresses()

		require.Equal(t, parseMACs(row), actualAddrs)
		require.NoError(t, err)
	}
}

func TestGetAddressesError(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	wifiServece := mywifi.New(mockWifi)

	mockWifi.On("Interfaces").Unset()
	mockWifi.On("Interfaces").Return(nil, ErrExpected)
	actualAddrs, err := wifiServece.GetAddresses()

	require.Nil(t, actualAddrs)
	require.ErrorIs(t, err, ErrExpected)
	require.ErrorContains(t, err, "getting interfaces")
}

func TestGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	wifiServece := mywifi.New(mockWifi)

	for _, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row), nil)
		actualNames, err := wifiServece.GetNames()

		require.NoError(t, err)
		require.Equal(t, parseName(row), actualNames)
	}
}

func TestGetNamesError(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	wifiServece := mywifi.New(mockWifi)

	mockWifi.On("Interfaces").Unset()
	mockWifi.On("Interfaces").Return(nil, ErrExpected)
	actualNames, err := wifiServece.GetNames()

	require.Nil(t, actualNames)
	require.ErrorIs(t, err, ErrExpected)
	require.ErrorContains(t, err, "getting interfaces")
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var interfaces []*wifi.Interface

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

func parseMACs(maxStr []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr

	for _, addr := range maxStr {
		addrs = append(addrs, parseMAC(addr))
	}

	return addrs
}

func parseMAC(maxStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(maxStr)
	if err != nil {
		return nil
	}

	return hwAddr
}

func parseName(addrs []string) []string {
	netNames := []string{}

	for i := range addrs {
		netNames = append(netNames, fmt.Sprintf("eth%d", i+1))
	}

	return netNames
}
