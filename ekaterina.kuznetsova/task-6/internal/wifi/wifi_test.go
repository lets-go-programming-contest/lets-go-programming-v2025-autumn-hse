package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	myWifi "github.com/Ekaterina-101/task-6/internal/wifi"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

func TestNew(t *testing.T) {
	t.Parallel()

	mockWifi := &MockWiFi{}
	wifiService := myWifi.New(mockWifi)

	require.NotNil(t, wifiService, "WiFiService should not be nil")
	require.Equal(t, mockWifi, wifiService.WiFi, "WiFi field should be set correctly")
}

type rowTestSysInfo struct {
	addrs       []string
	errExpected error
}

var testTable = []rowTestSysInfo{
	{
		addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
	},
	{
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWifi := &MockWiFi{}
	wifiService := myWifi.WiFiService{WiFi: mockWifi}

	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)

		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected,
				"row: %d, expected error: %v, actual error: %v",
				i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs,
			"row: %d, expected addrs: %v, actual addrs: %v",
			i, parseMACs(row.addrs), actualAddrs)
	}
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var interfaces []*wifi.Interface

	for i, addrStr := range addrs {
		hwAddr := parseMAC(addrStr)
		if hwAddr == nil {
			continue
		}

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i+1),
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}

		interfaces = append(interfaces, iface)
	}

	return interfaces
}

func parseMACs(macStr []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, addr := range macStr {
		addrs = append(addrs, parseMAC(addr))
	}
	return addrs
}

func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}
	return hwAddr
}

type rowTestGetNames struct {
	names       []string
	errExpected error
}

var testTableGetNames = []rowTestGetNames{
	{
		names: []string{"1111", "WifiIvan"},
	},
	{
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWifi := &MockWiFi{}
	wifiService := myWifi.WiFiService{WiFi: mockWifi}

	for i, row := range testTableGetNames {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockInterfacesByNames(row.names), row.errExpected)

		actualNames, err := wifiService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected,
				"row %d: expected error %v, got %v",
				i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row %d: unexpected error", i)
		require.Equal(t, row.names, actualNames,
			"row %d: expected names %v, got %v",
			i, row.names, actualNames)
	}
}

func mockInterfacesByNames(names []string) []*wifi.Interface {
	var interfaces []*wifi.Interface

	for i, name := range names {
		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         name,
			HardwareAddr: parseMAC("00:00:00:00:00:00"),
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces = append(interfaces, iface)
	}

	return interfaces
}
