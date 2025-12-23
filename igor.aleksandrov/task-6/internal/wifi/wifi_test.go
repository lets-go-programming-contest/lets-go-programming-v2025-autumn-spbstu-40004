package wifi

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWiFiHandle struct {
	mock.Mock
}

func (m *MockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*wifi.Interface), args.Error(1)
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
				m.On("Interfaces").Return([]*wifi.Interface{
					{HardwareAddr: mustMAC("00:11:22:33:44:55")},
				}, nil)
			},
			want:    []net.HardwareAddr{mustMAC("00:11:22:33:44:55")},
			wantErr: false,
		},
		{
			name: "success_multiple",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					{HardwareAddr: mustMAC("00:11:22:33:44:55")},
					{HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
				}, nil)
			},
			want: []net.HardwareAddr{
				mustMAC("00:11:22:33:44:55"),
				mustMAC("aa:bb:cc:dd:ee:ff"),
			},
			wantErr: false,
		},
		{
			name: "empty_list",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{}, nil)
			},
			want:    []net.HardwareAddr{},
			wantErr: false,
		},
		{
			name: "interface_with_nil_mac",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					{HardwareAddr: nil},
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
				assert.Contains(t, err.Error(), "getting interfaces")
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
				m.On("Interfaces").Return([]*wifi.Interface{
					{Name: "wlan0"},
				}, nil)
			},
			want:    []string{"wlan0"},
			wantErr: false,
		},
		{
			name: "success_multiple",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					{Name: "wlan0"},
					{Name: "wlan1"},
					{Name: "eth0"},
				}, nil)
			},
			want:    []string{"wlan0", "wlan1", "eth0"},
			wantErr: false,
		},
		{
			name: "empty_list",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{}, nil)
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "empty_name",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
					{Name: ""},
				}, nil)
			},
			want:    []string{""},
			wantErr: false,
		},
		{
			name: "duplicate_names",
			mockFn: func(m *MockWiFiHandle) {
				m.On("Interfaces").Return([]*wifi.Interface{
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
				assert.Contains(t, err.Error(), "getting interfaces")
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

func TestNew_NilHandle(t *testing.T) {
	t.Parallel()

	service := New(nil)
	assert.NotNil(t, service)
	assert.Nil(t, service.WiFi)
}
