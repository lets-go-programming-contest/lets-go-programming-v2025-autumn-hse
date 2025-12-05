package wifi_test

import (
	"errors"
	"net"
	"testing"

	wifimocks "github.com/6ermvH/german.feskov/task-6/internal/mocks/wifi"
	wifiModule "github.com/6ermvH/german.feskov/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --dir=. --name=WiFiHandle --output=./../mocks/wifi --outpkg=wifimocks

func TestGetAddressesNoErrorOneAddr(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	want := []net.HardwareAddr{[]byte("addr")}

	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return []*wifi.Interface{
			{
				HardwareAddr: want[0],
			},
		}, nil
	})

	have, err := service.GetAddresses()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetAddressesNoErrorMoreAddrs(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	want := []net.HardwareAddr{[]byte("ade"), []byte("addr2"), []byte("addr3")}

	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return []*wifi.Interface{
			{
				HardwareAddr: want[0],
			},
			{
				HardwareAddr: want[1],
			},
			{
				HardwareAddr: want[2],
			},
		}, nil
	})

	have, err := service.GetAddresses()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetAddressesErrorOnInterfaces(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return nil, errors.ErrUnsupported
	})

	_, err := service.GetAddresses()
	require.Error(t, err)
}

func TestGetNamesNoErrorOneName(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	want := []string{"german"}

	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return []*wifi.Interface{
			{
				Name: want[0],
			},
		}, nil
	})

	have, err := service.GetNames()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetNamesNoErrorMoreNames(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	want := []string{"german", "anthon", "vitaly"}

	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return []*wifi.Interface{
			{
				Name: want[0],
			},
			{
				Name: want[1],
			},
			{
				Name: want[2],
			},
		}, nil
	})

	have, err := service.GetNames()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetNamesErrorOnInterfaces(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return nil, errors.ErrUnsupported
	})

	_, err := service.GetNames()
	require.Error(t, err)
}
