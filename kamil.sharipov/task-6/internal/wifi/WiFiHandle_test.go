package wifi_test

import (
	"errors"
	"fmt"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
)

type MockWiFiHandle struct {
	mock.Mock
}

var errMockTypeMismatch = errors.New("mock Interfaces() returned unexpected type, expected []*wifi.Interface")

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

	return nil, fmt.Errorf("%w: %T", errMockTypeMismatch, raw)
}
