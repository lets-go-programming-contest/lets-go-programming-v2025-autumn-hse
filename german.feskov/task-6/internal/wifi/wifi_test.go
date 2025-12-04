package wifi_test

import (
	"errors"
	"net"
	"slices"
	"testing"

	wifimocks "github.com/6ermvH/german.feskov/task-6/internal/mocks/wifi"
	wifiModule "github.com/6ermvH/german.feskov/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
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
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	checkHardwareAddressesEqual(t, want, have)
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
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	checkHardwareAddressesEqual(t, want, have)
}

func TestGetAddressesErrorOnInterfaces(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return nil, errors.ErrUnsupported
	})

	_, err := service.GetAddresses()
	if err == nil {
		t.Fatalf("return no error")
	}
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
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	if !slices.Equal(want, have) {
		t.Fatalf("want: %v, have: %v", want, have)
	}
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
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	if !slices.Equal(want, have) {
		t.Fatalf("want: %v, have: %v", want, have)
	}
}

func TestGetNamesErrorOnInterfaces(t *testing.T) {
	t.Parallel()

	handle := wifimocks.NewWiFiHandle(t)
	service := wifiModule.New(handle)
	handle.On("Interfaces").Return(func() ([]*wifi.Interface, error) {
		return nil, errors.ErrUnsupported
	})

	_, err := service.GetNames()
	if err == nil {
		t.Fatalf("return no error")
	}
}

func checkHardwareAddressesEqual(t *testing.T, want, have []net.HardwareAddr) {
	t.Helper()

	if len(want) != len(have) {
		t.Fatalf("slices of addresses 'want' has %d size, 'have' has %d size", len(want), len(have))
	}

	for i := range want {
		if !slices.Equal(want[i], have[i]) {
			t.Fatalf("hardware addreses %d, want: %q, have: %q", i, want[i], have[i])
		}
	}
}
