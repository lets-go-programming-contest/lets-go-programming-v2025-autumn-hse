package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	w "github.com/Ekaterina-101/task-6/internal/wifi"
)

//go:generate mockery --name=WiFi --dir=../../../../../../.. --output=. --outpkg=wifi_test --filename=mock_wifi.go --structname=MockWiFi

var (
	errInterfacesError = errors.New("interfaces error")
	errNamesError      = errors.New("names error")
)

func TestGetAddresses_Success(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFi)

	hwAddr, err := net.ParseMAC("38:d5:7a:eb:43:8f")
	require.NoError(t, err)

	interfaces := []*wifi.Interface{
		{HardwareAddr: hwAddr},
	}

	want := []net.HardwareAddr{hwAddr}

	mock.On("Interfaces").Return(interfaces, nil)

	service := w.New(mock)
	got, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Equal(t, want, got)

	mock.AssertExpectations(t)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFi)

	mock.On("Interfaces").Return(nil, errInterfacesError)

	service := w.New(mock)
	got, err := service.GetAddresses()

	require.Error(t, err)
	assert.Nil(t, got)
	assert.Contains(t, err.Error(), "getting interfaces: interfaces error")

	mock.AssertExpectations(t)
}

func TestGetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFi)

	mock.On("Interfaces").Return([]*wifi.Interface{}, nil)

	service := w.New(mock)
	got, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Equal(t, []net.HardwareAddr{}, got)

	mock.AssertExpectations(t)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFi)

	ifaceName := "wlp2s0"
	interfaces := []*wifi.Interface{
		{Name: ifaceName},
	}

	mock.On("Interfaces").Return(interfaces, nil)

	service := w.New(mock)
	got, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, []string{ifaceName}, got)

	mock.AssertExpectations(t)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFi)

	mock.On("Interfaces").Return(nil, errNamesError)

	service := w.New(mock)
	got, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, got)
	assert.Contains(t, err.Error(), "getting interfaces: names error")

	mock.AssertExpectations(t)
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()

	mock := new(MockWiFi)

	mock.On("Interfaces").Return([]*wifi.Interface{}, nil)

	service := w.New(mock)
	got, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, []string{}, got)

	mock.AssertExpectations(t)
}
