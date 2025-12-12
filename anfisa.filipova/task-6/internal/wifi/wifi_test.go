//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg=wifi_test --output .
package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/Anfisa111/task-6/internal/wifi"
	wifipkg "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var (
	errPermission    = errors.New("permission denied")
	errGetInterfaces = errors.New("failed to get interfaces")
)

func TestGetAddresses(t *testing.T) {
	t.Parallel()
	// Тест 1: Успешное получение адресов
	t.Run("success with multiple addresses", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{HardwareAddr: parseMAC("00:11:22:33:44:55")},
			{HardwareAddr: parseMAC("aa:bb:cc:dd:ee:ff")},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		addrs, err := service.GetAddresses()

		require.NoError(t, err)
		require.Len(t, addrs, 2)
		require.Equal(t, "00:11:22:33:44:55", addrs[0].String())
		require.Equal(t, "aa:bb:cc:dd:ee:ff", addrs[1].String())
		mockWifi.AssertExpectations(t)
	})

	// Тест 2: Один адрес
	t.Run("success with single address", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{HardwareAddr: parseMAC("11:22:33:44:55:66")},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		addrs, err := service.GetAddresses()

		require.NoError(t, err)
		require.Len(t, addrs, 1)
		require.Equal(t, "11:22:33:44:55:66", addrs[0].String())
		mockWifi.AssertExpectations(t)
	})

	// Тест 3: Пустой результат
	t.Run("success with empty interfaces", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		mockWifi.On("Interfaces").Return([]*wifipkg.Interface{}, nil)

		addrs, err := service.GetAddresses()

		require.NoError(t, err)
		require.Empty(t, addrs)
		mockWifi.AssertExpectations(t)
	})

	// Тест 4: Интерфейс с nil адресом
	t.Run("interface with nil hardware address", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{HardwareAddr: nil},
			{HardwareAddr: parseMAC("aa:bb:cc:dd:ee:ff")},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		addrs, err := service.GetAddresses()

		require.NoError(t, err)
		require.Len(t, addrs, 2)
		require.Nil(t, addrs[0])
		require.NotNil(t, addrs[1])
		mockWifi.AssertExpectations(t)
	})

	// Тест 5: Ошибка при получении интерфейсов
	t.Run("error getting interfaces", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		mockWifi.On("Interfaces").Return(nil, errPermission)

		addrs, err := service.GetAddresses()

		require.Error(t, err)
		require.ErrorContains(t, err, "getting interfaces:")
		require.Nil(t, addrs)
		mockWifi.AssertExpectations(t)
	})

	// Тест 6: Невалидный MAC-адрес
	t.Run("invalid MAC address", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{HardwareAddr: net.HardwareAddr{}},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		addrs, err := service.GetAddresses()

		require.NoError(t, err)
		require.Len(t, addrs, 1)
		require.NotNil(t, addrs[0])
		mockWifi.AssertExpectations(t)
	})
}

func TestGetNames(t *testing.T) {
	t.Parallel()
	// Тест 1: Успешное получение имен
	t.Run("success with multiple names", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{Name: "wlan0"},
			{Name: "wlan1"},
			{Name: "eth0"},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		names, err := service.GetNames()

		require.NoError(t, err)
		require.Equal(t, []string{"wlan0", "wlan1", "eth0"}, names)
		mockWifi.AssertExpectations(t)
	})

	// Тест 2: Одно имя
	t.Run("success with single name", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{Name: "wlan0"},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		names, err := service.GetNames()

		require.NoError(t, err)
		require.Equal(t, []string{"wlan0"}, names)
		mockWifi.AssertExpectations(t)
	})

	// Тест 3: Пустой результат
	t.Run("success with empty interfaces", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		mockWifi.On("Interfaces").Return([]*wifipkg.Interface{}, nil)

		names, err := service.GetNames()

		require.NoError(t, err)
		require.Empty(t, names)
		mockWifi.AssertExpectations(t)
	})

	// Тест 4: Дубликаты имен
	t.Run("interfaces with duplicate names", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{Name: "wlan0"},
			{Name: "wlan0"},
			{Name: "wlan1"},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		names, err := service.GetNames()

		require.NoError(t, err)
		require.Equal(t, []string{"wlan0", "wlan0", "wlan1"}, names)
		mockWifi.AssertExpectations(t)
	})

	// Тест 5: Пустое имя
	t.Run("interface with empty name", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		interfaces := []*wifipkg.Interface{
			{Name: ""},
			{Name: "wlan0"},
		}
		mockWifi.On("Interfaces").Return(interfaces, nil)

		names, err := service.GetNames()

		require.NoError(t, err)
		require.Equal(t, []string{"", "wlan0"}, names)
		mockWifi.AssertExpectations(t)
	})

	// Тест 6: Ошибка при получении интерфейсов
	t.Run("error getting interfaces", func(t *testing.T) {
		t.Parallel()
		mockWifi := NewWiFiHandle(t)
		service := wifi.New(mockWifi)

		mockWifi.On("Interfaces").Return(nil, errGetInterfaces)

		names, err := service.GetNames()

		require.Error(t, err)
		require.ErrorContains(t, err, "getting interfaces:")
		require.Nil(t, names)
		mockWifi.AssertExpectations(t)
	})
}

func TestNew(t *testing.T) {
	t.Parallel()
	mockWifi := NewWiFiHandle(t)
	service := wifi.New(mockWifi)

	require.NotNil(t, service)
	require.Equal(t, mockWifi, service.WiFi)
}

func parseMAC(addr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(addr)
	if err != nil {
		return nil
	}

	return hwAddr
}
