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
			name: "success",
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{HardwareAddr: net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}},
				}, nil)
			},
			want:    []net.HardwareAddr{{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}},
			wantErr: false,
		},
		{
			name: "error_interfaces",
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return(nil, errors.New("fail"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(mocks.WiFiHandle)
			tt.mockFn(m)

			service := wifi.New(m)
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
		mockFn  func(m *mocks.WiFiHandle)
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return([]*wifilib.Interface{
					{Name: "wlan0"},
					{Name: "wlan1"},
				}, nil)
			},
			want:    []string{"wlan0", "wlan1"},
			wantErr: false,
		},
		{
			name: "error_interfaces", // Важно: этот тест покрывает if err != nil внутри GetNames()
			mockFn: func(m *mocks.WiFiHandle) {
				m.On("Interfaces").Return(nil, errors.New("fail"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := new(mocks.WiFiHandle)
			tt.mockFn(m)

			service := wifi.New(m)
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
