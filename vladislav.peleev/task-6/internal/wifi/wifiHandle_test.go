package wifi_test

import wifilib "github.com/mdlayher/wifi"

type mockWiFiWiFiHandle struct {
	interfaces []*wifilib.Interface
	err        error
}

func (m *mockWiFiWiFiHandle) Interfaces() ([]*wifilib.Interface, error) {
	return m.interfaces, m.err
}
