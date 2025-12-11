package wifi_test

import wifilib "github.com/mdlayher/wifi"

type mockWiFiHandle struct {
	interfaces []*wifilib.Interface
	err        error
}

func (m *mockWiFiHandle) Interfaces() ([]*wifilib.Interface, error) {
	return m.interfaces, m.err
}
