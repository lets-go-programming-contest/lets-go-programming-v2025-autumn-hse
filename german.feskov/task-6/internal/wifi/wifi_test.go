package wifi

import (
	"errors"
	"net"
	"slices"
	"testing"

	wifimocks "github.com/6ermvH/german.feskov/task-6/internal/mocks/wifi"
	"github.com/mdlayher/wifi"
)

//go:generate mockery --dir=. --name=WiFiHandle --output=./../mocks/wifi --outpkg=wifimocks

func TestGetAddresses(t *testing.T) {

	t.Run("no error, one addr", func(t *testing.T) {
		handle := wifimocks.NewWiFiHandle(t)
		service := New(handle)
		want := []net.HardwareAddr{[]byte("addr")}

		handle.
			On("Interfaces").
			Return(func() ([]*wifi.Interface, error) {
				return []*wifi.Interface{
					&wifi.Interface{
						HardwareAddr: want[0],
					},
				}, nil
			})

		have, err := service.GetAddresses()

		if err != nil {
			t.Fatalf("return error: %s", err.Error())
		}

		checkHardwareAddressesEqual(t, want, have)

	})

	t.Run("no error, more addrs", func(t *testing.T) {
		handle := wifimocks.NewWiFiHandle(t)
		service := New(handle)
		want := []net.HardwareAddr{[]byte("ade"), []byte("addr2"), []byte("addr3")}

		handle.
			On("Interfaces").
			Return(func() ([]*wifi.Interface, error) {
				return []*wifi.Interface{
					&wifi.Interface{
						HardwareAddr: want[0],
					},
					&wifi.Interface{
						HardwareAddr: want[1],
					},
					&wifi.Interface{
						HardwareAddr: want[2],
					},
				}, nil
			})

		have, err := service.GetAddresses()

		if err != nil {
			t.Fatalf("return error: %s", err.Error())
		}

		checkHardwareAddressesEqual(t, want, have)

	})

	t.Run("error on Interfaces", func(t *testing.T) {
		handle := wifimocks.NewWiFiHandle(t)
		service := New(handle)
		handle.
			On("Interfaces").
			Return(func() ([]*wifi.Interface, error) {
				return nil, errors.ErrUnsupported
			})
		_, err := service.GetAddresses()

		if err == nil {
			t.Fatalf("return no error")
		}

	})

}

func TestGetNames(t *testing.T) {

	t.Run("no error, one name", func(t *testing.T) {
		handle := wifimocks.NewWiFiHandle(t)
		service := New(handle)
		want := []string{"german"}

		handle.
			On("Interfaces").
			Return(func() ([]*wifi.Interface, error) {
				return []*wifi.Interface{
					&wifi.Interface{
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

	})

	t.Run("no error, more names", func(t *testing.T) {
		handle := wifimocks.NewWiFiHandle(t)
		service := New(handle)
		want := []string{"german", "anthon", "vitaly"}

		handle.
			On("Interfaces").
			Return(func() ([]*wifi.Interface, error) {
				return []*wifi.Interface{
					&wifi.Interface{
						Name: want[0],
					},
					&wifi.Interface{
						Name: want[1],
					},
					&wifi.Interface{
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

	})

	t.Run("error on Interfaces", func(t *testing.T) {
		handle := wifimocks.NewWiFiHandle(t)
		service := New(handle)
		handle.
			On("Interfaces").
			Return(func() ([]*wifi.Interface, error) {
				return nil, errors.ErrUnsupported
			})
		_, err := service.GetNames()

		if err == nil {
			t.Fatalf("return no error")
		}

	})

}

func checkHardwareAddressesEqual(t *testing.T, want, have []net.HardwareAddr) {
	t.Helper()

	if len(want) != len(have) {
		t.Fatalf("slices of addresses 'want' has %d size, 'have' has %d size", len(want), len(have))
	}

	for i := 0; i < len(want); i++ {
		if !slices.Equal(want[i], have[i]) {
			t.Fatalf("hardware addreses %d, want: %q, have: %q", i, want[i], have[i])
		}
	}

}
