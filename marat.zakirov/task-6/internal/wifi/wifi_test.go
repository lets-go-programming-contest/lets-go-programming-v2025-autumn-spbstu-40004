package wifi_test

import (
	"net"
	"testing"

	"github.com/ZakirovMS/task-6/internal/wifi"
	mdlayherwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

type mockWiFiHandle struct {
	interfaces []*mdlayherwifi.Interface
	err        error
}

func (m *mockWiFiHandle) Interfaces() ([]*mdlayherwifi.Interface, error) {
	return m.interfaces, m.err
}

type wifiTestRow struct {
	name        string
	interfaces  []*mdlayherwifi.Interface
	errExpected error
}

type customError string

func (e customError) Error() string {
	return string(e)
}

var errFailedToGetInterfaces = customError("failed to get interfaces")

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mac1 := net.HardwareAddr{1, 2, 3, 4, 5, 6}
	mac2 := net.HardwareAddr{7, 8, 9, 10, 11, 12}

	testTableGetAddresses := []wifiTestRow{
		{
			name: "single interface",
			interfaces: []*mdlayherwifi.Interface{
				{HardwareAddr: mac1, Name: "nameSt"},
			},
		},
		{
			name: "multiple interfaces",
			interfaces: []*mdlayherwifi.Interface{
				{HardwareAddr: mac1, Name: "nameSt"},
				{HardwareAddr: mac2, Name: "NameNd"},
			},
		},
		{
			name:       "no interfaces",
			interfaces: []*mdlayherwifi.Interface{},
		},
		{
			name:        "error case",
			interfaces:  nil,
			errExpected: errFailedToGetInterfaces,
		},
	}

	for _, row := range testTableGetAddresses {
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			mockHandle := &mockWiFiHandle{
				interfaces: row.interfaces,
				err:        row.errExpected,
			}

			service := wifi.New(mockHandle)
			addresses, err := service.GetAddresses()

			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected)
				require.Nil(t, addresses)

				return
			}

			require.NoError(t, err)

			expectedAddrs := make([]net.HardwareAddr, 0, len(row.interfaces))
			for _, iface := range row.interfaces {
				expectedAddrs = append(expectedAddrs, iface.HardwareAddr)
			}

			require.Equal(t, expectedAddrs, addresses)
		})
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mac1 := net.HardwareAddr{1, 2, 3, 4, 5, 6}
	mac2 := net.HardwareAddr{7, 8, 9, 10, 11, 12}

	testTableGetNames := []wifiTestRow{
		{
			name: "single interface",
			interfaces: []*mdlayherwifi.Interface{
				{HardwareAddr: mac1, Name: "nameSt"},
			},
		},
		{
			name: "multiple interfaces",
			interfaces: []*mdlayherwifi.Interface{
				{HardwareAddr: mac1, Name: "nameSt"},
				{HardwareAddr: mac2, Name: "nameNd"},
			},
		},
		{
			name:       "no interfaces",
			interfaces: []*mdlayherwifi.Interface{},
		},
		{
			name:        "error case",
			interfaces:  nil,
			errExpected: errFailedToGetInterfaces,
		},
	}

	for _, row := range testTableGetNames {
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			mockHandle := &mockWiFiHandle{
				interfaces: row.interfaces,
				err:        row.errExpected,
			}

			service := wifi.New(mockHandle)
			names, err := service.GetNames()

			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected)
				require.Nil(t, names)

				return
			}

			require.NoError(t, err)

			expectedNames := make([]string, 0, len(row.interfaces))
			for _, iface := range row.interfaces {
				expectedNames = append(expectedNames, iface.Name)
			}

			require.Equal(t, expectedNames, names)
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockHandle := &mockWiFiHandle{}
	service := wifi.New(mockHandle)

	require.NotNil(t, service)
	require.NotNil(t, service.WiFi)
}
