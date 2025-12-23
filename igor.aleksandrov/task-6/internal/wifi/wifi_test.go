package wifi

import (
	"errors"
	"net"
	"testing"

	wifilib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWiFiHandle struct {
	mock.Mock
}

func (m *MockWiFiHandle) Interfaces() ([]*wifilib.Interface, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*wifilib.Interface), args.Error(1)
}

func mustMAC(s string) net.HardwareAddr {
	mac, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}
	return mac
}

func TestWiFiService_GetAddresses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mockFn  func(m *MockWiFiHandle)
		want    []net.HardwareAddr
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{Name: "wlan0", HardwareAddr: mustMAC("00:11:22:33:44:55")},
					{Name: "wlan1", HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
				}, nil)
			},
			want: []net.HardwareAddr{
				mustMAC("00:11:22:33:44:55"),
				mustMAC("aa:bb:cc:dd:ee:ff"),
			},
			wantErr: false,
		},
		{
			name: "success_single",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{HardwareAddr: net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}},
				}, nil)
			},
			want:    []net.HardwareAddr{{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}},
			wantErr: false,
		},
		{
			name: "empty_list",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{}, nil)
			},
			want:    []net.HardwareAddr{},
			wantErr: false,
		},
		{
			name: "interface_with_nil_mac",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{Name: "wlan0", HardwareAddr: nil},
				}, nil)
			},
			want:    []net.HardwareAddr{nil},
			wantErr: false,
		},
		{
			name: "error_interfaces",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return(nil, errors.New("fail"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(MockWiFiHandle)
			tt.mockFn(m)

			service := New(m)
			got, err := service.GetAddresses()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			m.AssertExpectations(t)
		})
	}
}

func TestWiFiService_GetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mockFn  func(m *MockWiFiHandle)
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{Name: "wlan0"},
					{Name: "wlan1"},
				}, nil)
			},
			want:    []string{"wlan0", "wlan1"},
			wantErr: false,
		},
		{
			name: "empty_list",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{}, nil)
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "empty_name",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{Name: ""},
				}, nil)
			},
			want:    []string{""},
			wantErr: false,
		},
		{
			name: "duplicate_names",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{Name: "wlan0"},
					{Name: "wlan0"},
				}, nil)
			},
			want:    []string{"wlan0", "wlan0"},
			wantErr: false,
		},
		{
			name: "error_interfaces",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return(nil, errors.New("fail"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(MockWiFiHandle)
			tt.mockFn(m)

			service := New(m)
			got, err := service.GetNames()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			m.AssertExpectations(t)
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	m := new(MockWiFiHandle)
	service := New(m)

	assert.NotNil(t, service)
	assert.Equal(t, m, service.WiFi)
}
