package wifi_test

import (
	"errors"
	"fmt"
	"github.com/15446-rus75/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

type MockWiFi struct {
	mock.Mock
}

func (m *MockWiFi) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()
	return args.Get(0).([]*wifi.Interface), args.Error(1)
}

func parseMAC(addr string) net.HardwareAddr {
	hw, _ := net.ParseMAC(addr)
	return hw
}

func mockInterfaces(addrs []string) []*wifi.Interface {
	var ifaces []*wifi.Interface
	for i, addr := range addrs {
		ifaces = append(ifaces, &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("wlan%d", i+1),
			HardwareAddr: parseMAC(addr),
		})
	}
	return ifaces
}

func TestGetAddresses(t *testing.T) {
	mockWiFi := new(MockWiFi)
	service := wifi.New(mockWiFi)

	tests := []struct {
		name          string
		mockAddrs     []string
		mockError     error
		expectedAddrs []net.HardwareAddr
		expectedError error
	}{
		{
			name:          "success case",
			mockAddrs:     []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
			expectedAddrs: []net.HardwareAddr{parseMAC("00:11:22:33:44:55"), parseMAC("aa:bb:cc:dd:ee:ff")},
		},
		{
			name:          "no interfaces",
			mockAddrs:     []string{},
			expectedAddrs: []net.HardwareAddr{},
		},
		{
			name:          "error case",
			mockError:     errors.New("wifi error"),
			expectedError: errors.New("wifi error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockWiFi.On("Interfaces").Return(mockInterfaces(tc.mockAddrs), tc.mockError).Once()

			addrs, err := service.GetAddresses()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedAddrs, addrs)
			}

			mockWiFi.AssertExpectations(t)
		})
	}
}
