package wifi

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAddressesSuccess(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFiHandle)

	hwAddr, err := net.ParseMAC("38:d5:7a:eb:43:8f")
	require.NoError(t, err)

	interfaces := []*wifi.Interface{
		{HardwareAddr: hwAddr},
	}

	want := []net.HardwareAddr{hwAddr}
	mock.On("Interfaces").Return(interfaces, nil)

	service := New(mock)
	have, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Equal(t, want, have)

	mock.AssertExpectations(t)
}

func TestGetAddressesError(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFiHandle)

	mock.On("Interfaces").Return(nil, errors.New("interfaces error"))

	service := New(mock)
	have, err := service.GetAddresses()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "getting interfaces: interfaces error")

	mock.AssertExpectations(t)
}

func TestGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFiHandle)

	wifiInterfaceName := "wlp2s0"
	interfaces := []*wifi.Interface{
		{Name: wifiInterfaceName},
	}

	mock.On("Interfaces").Return(interfaces, nil)

	want := []string{wifiInterfaceName}
	service := New(mock)
	have, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, want, have)

	mock.AssertExpectations(t)
}

func TestGetNamesError(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFiHandle)

	mock.On("Interfaces").Return(nil, errors.New("names error"))

	service := New(mock)
	have, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "getting interfaces: names error")

	mock.AssertExpectations(t)
}
