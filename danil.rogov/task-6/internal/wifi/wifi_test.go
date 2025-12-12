package wifi_test

import (
	"net"
	"testing"

	myWifi "github.com/Tapochek2894/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .
const (
	gettingInterfacesError = "getting interfaces: "
	testWifi               = "wifi1"
	testMAC                = "00:11:22:33:44:55"
)

type MockWiFiHandle struct {
	interfaces []*wifi.Interface
	err        error
}

func (m *MockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	return m.interfaces, m.err
}

func createTestWiFiHandle(t *testing.T, interfaces []*wifi.Interface, err error) *MockWiFiHandle {
	t.Helper()

	return &MockWiFiHandle{
		interfaces: interfaces,
		err:        err,
	}
}

func TestCorrectGetNames(t *testing.T) {
	t.Parallel()

	mockHandle := createTestWiFiHandle(t, []*wifi.Interface{
		{Name: testWifi},
	}, nil)

	expected := []string{testWifi}
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
	assert.ErrorContains(t, err, gettingInterfacesError)
}

func TestCorrectGetAddresses(t *testing.T) {
	t.Parallel()

	mac, _ := net.ParseMAC(testMAC)

	mockHandle := createTestWiFiHandle(t, []*wifi.Interface{
		{HardwareAddr: mac},
	}, nil)

	expected := []net.HardwareAddr{mac}
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
	assert.ErrorContains(t, err, gettingInterfacesError)
}
