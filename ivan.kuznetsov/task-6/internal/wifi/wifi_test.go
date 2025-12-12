package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	myWifi "github.com/kuzid-17/task-6/internal/wifi"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

var errExpected = errors.New("ExpectedError")

type rowTestSysInfo struct {
	addrs       []string
	names       []string
	errExpected error
}

func getTestTable() []rowTestSysInfo {
	return []rowTestSysInfo{
		{
			addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
			names: []string{"wlan0", "wlan1"},
		},
		{
			errExpected: errExpected,
		},
	}
}

func TestNew(t *testing.T) {
	t.Parallel()
	mockWifi := NewWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	require.NotNil(t, wifiService, "service should not be nil")
	require.Equal(t, mockWifi, wifiService.WiFi, "WiFi field should be set correctly")
}

func TestGetAddress(t *testing.T) {
	t.Parallel()
	mockWifi := NewWiFiHandle(t)
	wifiService := myWifi.WiFiService{WiFi: mockWifi}

	for i, row := range getTestTable() {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs), row.errExpected)

		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)

			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs,
			"row: %d, expected addrs: %s, actual addrs: %s",
			i, parseMACs(row.addrs), actualAddrs)
	}
}

func TestGetName(t *testing.T) {
	t.Parallel()
	mockWifi := NewWiFiHandle(t)
	wifiService := myWifi.WiFiService{WiFi: mockWifi}

	for i, row := range getTestTable() {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfacesForNames(row.names), row.errExpected)

		actualNames, err := wifiService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)

			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, actualNames, "row: %d, expected names: %s, actual names: %s", i, row.names, actualNames)
	}
}

func mockIfacesForNames(names []string) []*wifi.Interface {
	interfaces := make([]*wifi.Interface, 0, len(names))

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

func mockIfaces(addrs []string) []*wifi.Interface {
	interfaces := make([]*wifi.Interface, 0, len(addrs))

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
	addrs := make([]net.HardwareAddr, 0, len(macStr))

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
