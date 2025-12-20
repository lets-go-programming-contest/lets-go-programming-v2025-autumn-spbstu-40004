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
					{Name: "wlan0", HardwareAddr: mustMAC("00:11:22:33:44:55")},
					{Name: "wlan1", HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
				}, nil)
			},
			wantLen: 2,
			check: func(t *testing.T, addrs []net.HardwareAddr) {
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
				assert.Empty(t, addrs)
			},
		},
		{
			name: "interface withnil MAC",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					{Name: "wlan0", HardwareAddr: nil},
				}, nil)
			},
			wantLen: 1,
			check: func(t *testing.T, addrs []net.HardwareAddr) {
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
				assert.NoError(t, err)
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
					{Name: "wlan0", HardwareAddr: mustMAC("00:11:22:33:44:55")},
					{Name: "wlan1", HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
					{Name: "eth0", HardwareAddr: mustMAC("11:22:33:44:55:66")},
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
					{Name: "wlan0", HardwareAddr: mustMAC("00:11:22:33:44:55")},
					{Name: "wlan0", HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
				}, nil)
			},
			want: []string{"wlan0", "wlan0"},
		},
		{
			name: "interface with nil MAC",
			mockFunc: func(m *WiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					{Name: "wlan0", HardwareAddr: nil},
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
				assert.NoError(t, err)
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
		assert.NoError(t, err)
		assert.Equal(t, []string{""}, names)

		addrs, err := service.GetAddresses()
		assert.NoError(t, err)
		assert.Len(t, addrs, 1)

		mock.AssertExpectations(t)
	})

	t.Run("multiple calls", func(t *testing.T) {
		t.Parallel()

		mock := &WiFiHandle{}
		ifaces := []*wifi.Interface{
			{Name: "wlan0", HardwareAddr: mustMAC("00:11:22:33:44:55")},
		}
		mock.On("Interfaces").Return(ifaces, nil).Twice()

		service := mywifi.New(mock)

		addrs1, err1 := service.GetAddresses()
		assert.NoError(t, err1)

		names1, err2 := service.GetNames()
		assert.NoError(t, err2)

		assert.Len(t, addrs1, 1)
		assert.Len(t, names1, 1)
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
