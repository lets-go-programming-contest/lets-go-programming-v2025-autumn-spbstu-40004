package wifi_test

import (
    "errors"
    "net"
    "testing"
    
    "github.com/15446-rus75/task-6/internal/wifi"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

type MockWiFi struct {
    interfacesFunc func() ([]wifi.Interface, error)
}

func (m *MockWiFi) Interfaces() ([]wifi.Interface, error) {
    if m.interfacesFunc != nil {
        return m.interfacesFunc()
    }
    return nil, nil
}

func TestNew(t *testing.T) {
    mockWiFi := &MockWiFi{}
    service := wifi.New(mockWiFi)
    assert.NotNil(t, service)
    assert.Equal(t, mockWiFi, service.WiFi)
}

func TestWiFiService_GetAddresses(t *testing.T) {
    t.Run("successful get addresses", func(t *testing.T) {
        mac1, _ := net.ParseMAC("00:11:22:33:44:55")
        mac2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")
        
        mockWiFi := &MockWiFi{
            interfacesFunc: func() ([]wifi.Interface, error) {
                return []wifi.Interface{
                    {HardwareAddr: mac1},
                    {HardwareAddr: mac2},
                }, nil
            },
        }
        
        service := wifi.New(mockWiFi)
        
        addresses, err := service.GetAddresses()
        
        require.NoError(t, err)
        require.Len(t, addresses, 2)
        assert.Equal(t, mac1, addresses[0])
        assert.Equal(t, mac2, addresses[1])
    })
    
    t.Run("empty interfaces", func(t *testing.T) {
        mockWiFi := &MockWiFi{
            interfacesFunc: func() ([]wifi.Interface, error) {
                return []wifi.Interface{}, nil
            },
        }
        
        service := wifi.New(mockWiFi)
        
        addresses, err := service.GetAddresses()
        
        require.NoError(t, err)
        assert.Empty(t, addresses)
    })
    
    t.Run("interfaces error", func(t *testing.T) {
        expectedErr := errors.New("wifi not available")
        
        mockWiFi := &MockWiFi{
            interfacesFunc: func() ([]wifi.Interface, error) {
                return nil, expectedErr
            },
        }
        
        service := wifi.New(mockWiFi)
        
        addresses, err := service.GetAddresses()
        
        assert.Error(t, err)
        assert.Nil(t, addresses)
        assert.Equal(t, expectedErr, err)
    })
    
    t.Run("interfaces with nil addresses", func(t *testing.T) {
        mockWiFi := &MockWiFi{
            interfacesFunc: func() ([]wifi.Interface, error) {
                return []wifi.Interface{
                    {HardwareAddr: nil},
                    {HardwareAddr: nil},
                }, nil
            },
        }
        
        service := wifi.New(mockWiFi)
        
        addresses, err := service.GetAddresses()
        
        require.NoError(t, err)
        require.Len(t, addresses, 2)
        assert.Nil(t, addresses[0])
        assert.Nil(t, addresses[1])
    })
    
    t.Run("single interface", func(t *testing.T) {
        mac, _ := net.ParseMAC("11:22:33:44:55:66")
        
        mockWiFi := &MockWiFi{
            interfacesFunc: func() ([]wifi.Interface, error) {
                return []wifi.Interface{
                    {HardwareAddr: mac},
                }, nil
            },
        }
        
        service := wifi.New(mockWiFi)
        
        addresses, err := service.GetAddresses()
        
        require.NoError(t, err)
        require.Len(t, addresses, 1)
        assert.Equal(t, mac, addresses[0])
    })
    
    t.Run("mixed addresses", func(t *testing.T) {
        mac, _ := net.ParseMAC("11:22:33:44:55:66")
        
        mockWiFi := &MockWiFi{
            interfacesFunc: func() ([]wifi.Interface, error) {
                return []wifi.Interface{
                    {HardwareAddr: nil},
                    {HardwareAddr: mac},
                    {HardwareAddr: nil},
                }, nil
            },
        }
        
        service := wifi.New(mockWiFi)
        
        addresses, err := service.GetAddresses()
        
        require.NoError(t, err)
        require.Len(t, addresses, 3)
        assert.Nil(t, addresses[0])
        assert.Equal(t, mac, addresses[1])
        assert.Nil(t, addresses[2])
    })
}
