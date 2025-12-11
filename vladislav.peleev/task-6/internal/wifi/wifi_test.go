package wifi

import (
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWiFiService_GetAddresses(t *testing.T) {
	addr1, _ := net.ParseMAC("00:11:22:33:44:55")
	addr2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")

	mockHandle := &mockWiFiHandle{
		interfaces: []*wifi.Interface{
			{HardwareAddr: addr1, Name: "wlan0"},
			{HardwareAddr: addr2, Name: "wlan1"},
		},
	}

	service := New(mockHandle)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)
	assert.Equal(t, []net.HardwareAddr{addr1, addr2}, addrs)
}

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	mockHandle := &mockWiFiHandle{
		err: assert.AnError,
	}

	service := New(mockHandle)

	_, err := service.GetAddresses()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "getting interfaces:")
}

func TestWiFiService_GetNames(t *testing.T) {
	addr1, _ := net.ParseMAC("00:11:22:33:44:55")

	mockHandle := &mockWiFiHandle{
		interfaces: []*wifi.Interface{
			{HardwareAddr: addr1, Name: "wlan0"},
			{HardwareAddr: nil, Name: "lo"}, // пример
		},
	}

	service := New(mockHandle)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{"wlan0", "lo"}, names)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	mockHandle := &mockWiFiHandle{
		err: assert.AnError,
	}

	service := New(mockHandle)

	_, err := service.GetNames()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "getting interfaces:")
}
