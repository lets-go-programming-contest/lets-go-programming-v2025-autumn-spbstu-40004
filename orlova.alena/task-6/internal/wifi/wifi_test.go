package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mywifi "github.com/widgeiw/task-6/internal/wifi"
)

var errWiFi = errors.New("wifi error")

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

func createMockInterfaceWithNilMAC(name string) *wifi.Interface {
	return &wifi.Interface{
		Index:        1,
		Name:         name,
		HardwareAddr: nil,
		PHY:          0,
		Device:       0,
		Type:         wifi.InterfaceTypeStation,
		Frequency:    2412,
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	mock := &WiFiHandle{}
	service := mywifi.New(mock)

	assert.NotNil(t, service)
	assert.Equal(t, mock, service.WiFi)
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mockFunc func(*WiFiHandle)
		wantErr  bool
		errMsg   string
		wantLen  int
		check    func(*testing.T, []net.HardwareAddr)
	}{
		{
			name: "success of getting addresses",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					createMockInterface("wlan0", "00:11:22:33:44:55"),
					createMockInterface("wlan1", "aa:bb:cc:dd:ee:ff"),
				}, nil)
			},
			wantLen: 2,
			check: func(t *testing.T, addrs []net.HardwareAddr) {
				t.Helper()

				assert.Equal(t, mustMAC("00:11:22:33:44:55"), addrs[0])
				assert.Equal(t, mustMAC("aa:bb:cc:dd:ee:ff"), addrs[1])
			},
		},
		{
			name: "empty list of interfaces",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{}, nil)
			},
			wantLen: 0,
			check: func(t *testing.T, addrs []net.HardwareAddr) {
				t.Helper()

				assert.Empty(t, addrs)
			},
		},
		{
			name: "interface withnil MAC",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					createMockInterfaceWithNilMAC("wlan0"),
				}, nil)
			},
			wantLen: 1,
			check: func(t *testing.T, addrs []net.HardwareAddr) {
				t.Helper()

				assert.Nil(t, addrs[0])
			},
		},
		{
			name: "error of getting interfaces",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return(nil, errWiFi)
			},
			wantErr: true,
			errMsg:  "getting interfaces",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mock := &WiFiHandle{}
			tc.mockFunc(mock)

			service := mywifi.New(mock)
			got, err := service.GetAddresses()

			if tc.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Len(t, got, tc.wantLen)

				if tc.check != nil {
					tc.check(t, got)
				}
			}

			mock.AssertExpectations(t)
		})
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mockFunc func(*WiFiHandle)
		wantErr  bool
		errMsg   string
		want     []string
	}{
		{
			name: "success of getting names",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					createMockInterface("wlan0", "00:11:22:33:44:55"),
					createMockInterface("wlan1", "aa:bb:cc:dd:ee:ff"),
					createMockInterface("eth0", "11:22:33:44:55:66"),
				}, nil)
			},
			want: []string{"wlan0", "wlan1", "eth0"},
		},
		{
			name: "empty list of interfaces",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{}, nil)
			},
			want: []string{},
		},
		{
			name: "duplicated names",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					createMockInterface("wlan0", "00:11:22:33:44:55"),
					createMockInterface("wlan0", "aa:bb:cc:dd:ee:ff"),
				}, nil)
			},
			want: []string{"wlan0", "wlan0"},
		},
		{
			name: "interface with nil MAC",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					createMockInterfaceWithNilMAC("wlan0"),
				}, nil)
			},
			want: []string{"wlan0"},
		},
		{
			name: "error of getting interfaces",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return(nil, errWiFi)
			},
			wantErr: true,
			errMsg:  "getting interfaces",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mock := &WiFiHandle{}
			tc.mockFunc(mock)

			service := mywifi.New(mock)
			got, err := service.GetNames()

			if tc.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}

			mock.AssertExpectations(t)
		})
	}
}

func TestEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("nil WiFiHandle in constructor", func(t *testing.T) {
		t.Parallel()

		service := mywifi.New(nil)
		assert.NotNil(t, service)
		assert.Nil(t, service.WiFi)
	})

	t.Run("empty name of interface", func(t *testing.T) {
		t.Parallel()

		mock := &WiFiHandle{}
		mock.On("Interfaces").Return([]*wifi.Interface{
			{Name: "", HardwareAddr: mustMAC("00:11:22:33:44:55")},
		}, nil)

		service := mywifi.New(mock)

		names, err := service.GetNames()
		require.NoError(t, err)
		assert.Equal(t, []string{""}, names)

		addrs, err := service.GetAddresses()
		require.NoError(t, err)
		assert.Len(t, addrs, 1)

		mock.AssertExpectations(t)
	})

	t.Run("multiple calls", func(t *testing.T) {
		t.Parallel()

		mock := &WiFiHandle{}
		ifaces := []*wifi.Interface{
			createMockInterface("wlan0", "00:11:22:33:44:55"),
		}
		mock.On("Interfaces").Return(ifaces, nil).Twice()

		service := mywifi.New(mock)

		addrs, err := service.GetAddresses()
		require.NoError(t, err)

		names, err := service.GetNames()
		require.NoError(t, err)

		assert.Len(t, addrs, 1)
		assert.Len(t, names, 1)
		mock.AssertExpectations(t)
	})
}

func mustMAC(s string) net.HardwareAddr {
	mac, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}

	return mac
}
