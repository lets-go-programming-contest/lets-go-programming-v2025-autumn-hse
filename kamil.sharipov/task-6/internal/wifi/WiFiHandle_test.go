package wifi_test

import (
	"fmt"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
)

type MockWiFiHandle struct {
	mock.Mock
}

func (m *MockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()

	err := args.Error(1)
	if err != nil {
		return nil, fmt.Errorf("mock interfaces error: %w", err)
	}

	raw := args.Get(0)
	if raw == nil {
		return nil, nil
	}

	if interfaceSlice, ok := raw.([]*wifi.Interface); ok {
		return interfaceSlice, nil
	}

	return nil, fmt.Errorf("mock Interfaces() returned unexpected type %T, expected []*wifi.Interface", raw)
}
