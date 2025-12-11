package wifi_test

import (
	"github.com/mdlayher/wifi"
)

type mockWiFiHandle struct {
	interfaces []*wifi.Interface
	err        error
}

func (m *mockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	return m.interfaces, m.err
}
