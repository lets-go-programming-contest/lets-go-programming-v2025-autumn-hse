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

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .
var ErrExpected = errors.New("expected error")

type rowTestSysInfo struct {
	addrs       []string
	names       []string
	errExpected error
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	var testTable = []rowTestSysInfo{
		{
			addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
			names: []string{"eth1", "eth2"},
		},
		{
			errExpected: ErrExpected,
		},
	}

	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs, row.names), row.errExpected)

		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			expectedErr := fmt.Errorf("getting interfaces: %w", row.errExpected)

			require.Error(t, err, "row: %d", i)
			require.EqualError(t, err, expectedErr.Error(), "row: %d", i)

			continue
		}

		require.NoError(t, err, "row: %d", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs, "row: %d", i)
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWifi := NewWiFiHandle(t)
	wifiService := myWifi.New(mockWifi)

	var testTable = []rowTestSysInfo{
		{
			addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
			names: []string{"eth1", "eth2"},
		},
		{
			errExpected: ErrExpected,
		},
	}

	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.addrs, row.names), row.errExpected)

		actualNames, err := wifiService.GetNames()

		if row.errExpected != nil {
			expectedErr := fmt.Errorf("getting interfaces: %w", row.errExpected)

			require.Error(t, err, "row: %d", i)
			require.EqualError(t, err, expectedErr.Error(), "row: %d", i)

			continue
		}

		require.NoError(t, err, "row: %d", i)
		require.Equal(t, row.names, actualNames, "row: %d", i)
	}
}

func mockIfaces(addrs, names []string) []*wifi.Interface {
	interfaces := make([]*wifi.Interface, 0, len(addrs))

	n := len(addrs)
	if len(names) > n {
		n = len(names)
	}

	for i := range n {
		var hwAddr net.HardwareAddr

		if i < len(addrs) {
			hwAddr = parseMAC(addrs[i])

			if hwAddr == nil {
				continue
			}
		} else {
			hwAddr = nil
		}

		name := fmt.Sprintf("eth%d", i+1)
		if i < len(names) {
			name = names[i]
		}

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         name,
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
