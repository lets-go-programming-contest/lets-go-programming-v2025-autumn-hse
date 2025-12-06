package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/OlesiaOl/task-6/internal/mocks"
	"github.com/OlesiaOl/task-6/internal/wifi"

	wifipkg "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

func TestWiFiGetNamesOneNoError(t *testing.T) {
	wifiHandler := mocks.NewWiFiHandle(t)
	service := wifi.New(wifiHandler)
	expectedResult := []string{"nevergivup"}

	wifiHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			Name: "nevergivup",
		},
	}, nil)

	realNames, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, expectedResult, realNames)

	wifiHandler.AssertExpectations(t)
}

func TestWiFiGetNamesMultipleNoError(t *testing.T) {
	wifiHandler := mocks.NewWiFiHandle(t)
	service := wifi.New(wifiHandler)
	expectedResult := []string{"Naf-Naf", "Nuf-Nuf", "Nif-Nif"}

	wifiHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			Name: "Naf-Naf",
		},
		{
			Name: "Nuf-Nuf",
		},
		{
			Name: "Nif-Nif",
		},
	}, nil)

	realNames, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, expectedResult, realNames)

	wifiHandler.AssertExpectations(t)
}

func TestWiFIGetNamesWithError(t *testing.T) {
	wifiHandler := mocks.NewWiFiHandle(t)
	service := wifi.New(wifiHandler)
	expectedErr := errors.New("not exist")

	wifiHandler.On("Interfaces").Return(nil, expectedErr)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "getting interfaces")
	require.ErrorIs(t, err, expectedErr)

	wifiHandler.AssertExpectations(t)
}

func TestWiFiGetAddressesOneNoError(t *testing.T) {
	mockHandler := mocks.NewWiFiHandle(t)
	service := wifi.New(mockHandler)

	mac, _ := net.ParseMAC("00:11:22:33:44:55")
	expectedResult := []net.HardwareAddr{mac}

	mockHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			HardwareAddr: mac,
		},
	}, nil)

	actualAddresses, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, expectedResult, actualAddresses)
	require.Len(t, actualAddresses, 1)

	mockHandler.AssertExpectations(t)
}

func TestWiFiGetAddressesMultipleNoError(t *testing.T) {
	mockHandler := mocks.NewWiFiHandle(t)
	service := wifi.New(mockHandler)

	mac1, _ := net.ParseMAC("00:11:22:33:44:55")
	mac2, _ := net.ParseMAC("AA:BB:CC:DD:EE:FF")
	mac3, _ := net.ParseMAC("11:22:33:44:55:66")

	expectedResult := []net.HardwareAddr{mac1, mac2, mac3}

	mockHandler.On("Interfaces").Return([]*wifipkg.Interface{
		{
			Name:         "Naf-Naf",
			HardwareAddr: mac1,
		},
		{
			Name:         "Nuf-Nuf",
			HardwareAddr: mac2,
		},
		{
			Name:         "Nif-Nif",
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
	mockHandler := mocks.NewWiFiHandle(t)
	service := wifi.New(mockHandler)

	expectedErr := errors.New("WiFi disabled")
	mockHandler.On("Interfaces").Return(nil, expectedErr)

	actualAddresses, err := service.GetAddresses()

	require.Error(t, err)
	require.Nil(t, actualAddresses)
	require.Contains(t, err.Error(), "getting interfaces")
	require.ErrorIs(t, err, expectedErr)

	mockHandler.AssertExpectations(t)
}
