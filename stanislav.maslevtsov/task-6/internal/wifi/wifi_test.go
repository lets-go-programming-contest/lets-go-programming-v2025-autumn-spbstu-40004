package wifi_test

import (
	"errors"
	"net"
	"testing"

	myWiFi "github.com/jambii1/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

var errExpected = errors.New("expected error")

type testTable struct {
	addrs, names []string
}

func parseMAC(t *testing.T, macStr string) net.HardwareAddr {
	t.Helper()

	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}

	return hwAddr
}

func parseMACs(t *testing.T, macStr []string) []net.HardwareAddr {
	t.Helper()

	addrs := make([]net.HardwareAddr, 0, len(macStr))

	for _, addr := range macStr {
		addrs = append(addrs, parseMAC(t, addr))
	}

	return addrs
}

func mockIfaces(t *testing.T, table testTable) []*wifi.Interface {
	t.Helper()

	require.Equal(t, len(table.addrs), len(table.names))

	interfaces := make([]*wifi.Interface, 0, len(table.addrs))

	for addrIdx, addrStr := range table.addrs {
		hwAddr := parseMAC(t, addrStr)
		if hwAddr == nil {
			continue
		}

		iface := &wifi.Interface{
			Index:        addrIdx + 1,
			Name:         table.names[addrIdx],
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

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	testData := testTable{
		addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		names: []string{"eth1", "eth2"},
	}

	mockWifi := NewWiFiHandle(t)
	wifiService := myWiFi.New(mockWifi)

	mockWifi.On("Interfaces").Unset()
	mockWifi.On("Interfaces").Return(mockIfaces(t, testData), nil)

	addrs, err := wifiService.GetAddresses()

	require.NoError(t, err, "error must be nil")
	require.Equal(t, parseMACs(t, testData.addrs), addrs,
		"expected addrs: %s, actual addrs: %s", parseMACs(t, testData.addrs), addrs)

	mockWifi.On("Interfaces").Unset()
	mockWifi.On("Interfaces").Return(nil, errExpected)

	addrs, err = wifiService.GetAddresses()

	require.ErrorIs(t, err, errExpected, "expected error: %w, actual error: %w", errExpected, err)
	require.Nil(t, addrs, "addrs must be nil")
	require.ErrorContains(t, err, "getting interfaces")
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	testData := testTable{
		addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		names: []string{"eth1", "eth2"},
	}

	mockWifi := NewWiFiHandle(t)
	wifiService := myWiFi.New(mockWifi)

	mockWifi.On("Interfaces").Unset()
	mockWifi.On("Interfaces").Return(mockIfaces(t, testData), nil)

	names, err := wifiService.GetNames()

	require.NoError(t, err, "error must be nil")
	require.Equal(t, testData.names, names,
		"expected names: %s, actual names: %s", testData.names, names)

	mockWifi.On("Interfaces").Unset()
	mockWifi.On("Interfaces").Return(nil, errExpected)

	names, err = wifiService.GetNames()

	require.ErrorIs(t, err, errExpected, "expected error: %w, actual error: %w", errExpected, err)
	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "getting interfaces")
}
