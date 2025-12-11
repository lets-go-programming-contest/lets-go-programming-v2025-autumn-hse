package wifi_test

import (
	"net"
	"testing"

	myWifi "github.com/Tapochek2894/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	gettingInterfacesError = "getting interfaces: "
)

type MockWiFiHandle struct {
	interfaces []*wifi.Interface
	err        error
}

func (m *MockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	return m.interfaces, m.err
}

func createTestWiFiHandle(
	t *testing.T,
	interfaces []*wifi.Interface,
	err error,
) myWifi.WiFiHandle {
	t.Helper()
	return &MockWiFiHandle{
		interfaces: interfaces,
		err:        err,
	}
}

func TestCorrectGetNames(t *testing.T) {
	t.Parallel()

	mockHandle := createTestWiFiHandle(t, []*wifi.Interface{
		{Name: "wlan0"},
		{Name: "wifi0"},
	}, nil)

	expected := []string{"wlan0", "wifi0"}
	service := myWifi.New(mockHandle)
	got, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestIncorrectGetNames(t *testing.T) {
	t.Parallel()

	mockHandle := createTestWiFiHandle(t, nil, assert.AnError)

	service := myWifi.New(mockHandle)
	got, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, got)
	assert.Contains(t, err.Error(), gettingInterfacesError)
}

func TestCorrectGetAddresses(t *testing.T) {
	t.Parallel()

	mac1, _ := net.ParseMAC("00:11:22:33:44:55")
	mac2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")

	mockHandle := createTestWiFiHandle(t, []*wifi.Interface{
		{HardwareAddr: mac1},
		{HardwareAddr: mac2},
	}, nil)

	expected := []net.HardwareAddr{mac1, mac2}
	service := myWifi.New(mockHandle)
	got, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestIncorrectGetAddresses(t *testing.T) {
	t.Parallel()

	mockHandle := createTestWiFiHandle(t, nil, assert.AnError)

	service := myWifi.New(mockHandle)
	got, err := service.GetAddresses()

	require.Error(t, err)
	assert.Nil(t, got)
	assert.Contains(t, err.Error(), gettingInterfacesError)
}
