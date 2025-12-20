package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/MrMels625/task-6/internal/wifi"
	"github.com/MrMels625/task-6/internal/wifi/mocks"
	wifilib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
)

func TestWiFiService_GetAddresses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mockFn  func(m *mocks.WiFiHandle)
		want    []net.HardwareAddr
		wantErr bool
	}{
		{
			name: "success_two_interfaces",
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{
						Index:        1,
						Name:         "wlan0",
						HardwareAddr: net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
					},
					{
						Index:        2,
						Name:         "wlan1",
						HardwareAddr: net.HardwareAddr{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
					},
				}, nil)
			},
			want: []net.HardwareAddr{
				{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
				{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
			},
			wantErr: false,
		},
		{
			name: "driver_error",
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return(nil, errors.New("system failure"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWiFi := new(mocks.WiFiHandle)
			tt.mockFn(mockWiFi)

			service := wifi.New(mockWiFi)
			got, err := service.GetAddresses()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			mockWiFi.AssertExpectations(t)
		})
	}
}

func TestWiFiService_GetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mockFn  func(m *mocks.WiFiHandle)
		want    []string
		wantErr bool
	}{
		{
			name: "success_names",
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{Name: "wifi_home"},
					{Name: "wifi_office"},
				}, nil)
			},
			want:    []string{"wifi_home", "wifi_office"},
			wantErr: false,
		},
		{
			name: "error_on_names",
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return(nil, errors.New("fail"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWiFi := new(mocks.WiFiHandle)
			tt.mockFn(mockWiFi)

			service := wifi.New(mockWiFi)
			got, err := service.GetNames()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			mockWiFi.AssertExpectations(t)
		})
	}
}
