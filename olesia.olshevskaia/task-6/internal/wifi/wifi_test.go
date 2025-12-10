//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .
package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/OlesiaOl/task-6/internal/wifi"

	wifipkg "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var (
	errTest = errors.New("test error")
)

const (
	interfaceName1 = "nevergivup"
	interfaceName2 = "Naf-Naf"
	interfaceName3 = "Nuf-Nuf"
	interfaceName4 = "Nif-Nif"
)

var (
	mac1, _ = net.ParseMAC("00:11:22:33:44:55")
	mac2, _ = net.ParseMAC("AA:BB:CC:DD:EE:FF")
	mac3, _ = net.ParseMAC("11:22:33:44:55:66")
)

func TestWiFiGetNamesOneNoError(t *testing.T) {
	t.Parallel()

	wifiHandler := NewWiFiHandle(t)
	service := wifi.New(wifiHandler)
	expectedResult := []string{interfaceName1}

	wifiHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			Name: interfaceName1,
		},
	}, nil)

	realNames, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, expectedResult, realNames)

	wifiHandler.AssertExpectations(t)
}

func TestWiFiGetNamesMultipleNoError(t *testing.T) {
	t.Parallel()

	wifiHandler := NewWiFiHandle(t)
	service := wifi.New(wifiHandler)
	expectedResult := []string{interfaceName2, interfaceName3, interfaceName4}

	wifiHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			Name: interfaceName2,
		},
		{
			Name: interfaceName3,
		},
		{
			Name: interfaceName4,
		},
	}, nil)

	realNames, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, expectedResult, realNames)

	wifiHandler.AssertExpectations(t)
}

func TestWiFIGetNamesWithError(t *testing.T) {
	t.Parallel()

	wifiHandler := NewWiFiHandle(t)
	service := wifi.New(wifiHandler)

	wifiHandler.On("Interfaces").Return(nil, errTest)

	_, err := service.GetNames()

	require.ErrorContains(t, err, "getting interfaces")
	require.ErrorIs(t, err, errTest)

	wifiHandler.AssertExpectations(t)
}

func TestWiFiGetAddressesOneNoError(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := wifi.New(mockHandler)

	expectedResult := []net.HardwareAddr{mac1}

	mockHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			HardwareAddr: mac1,
		},
	}, nil)

	actualAddresses, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, expectedResult, actualAddresses)
	require.Len(t, actualAddresses, 1)

	mockHandler.AssertExpectations(t)
}

func TestWiFiGetAddressesMultipleNoError(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := wifi.New(mockHandler)

	expectedResult := []net.HardwareAddr{mac1, mac2, mac3}

	mockHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			Name:         interfaceName2,
			HardwareAddr: mac1,
		},
		{
			Name:         interfaceName3,
			HardwareAddr: mac2,
		},
		{
			Name:         interfaceName4,
			HardwareAddr: mac3,
		},
	}, nil)

	actualAddresses, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, expectedResult, actualAddresses)
	require.Len(t, actualAddresses, 3)

	mockHandler.AssertExpectations(t)
}

func TestWiFiGetAddressesWithError(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := wifi.New(mockHandler)

	mockHandler.On("Interfaces").Return(nil, errTest)

	_, err := service.GetAddresses()

	require.ErrorContains(t, err, "getting interfaces")
	require.ErrorIs(t, err, errTest)

	mockHandler.AssertExpectations(t)
}
