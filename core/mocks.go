package main

type mockI2CBus struct{}

func (bus mockI2CBus) ReadBytes(addr byte, num int) (value []byte, err error) {
	return []byte{1}, nil
}

func (bus mockI2CBus) ReadByte(addr byte) (value byte, err error) {
	return 1, nil
}

func (bus mockI2CBus) WriteByte(addr, value byte) error {
	return nil
}

func (bus mockI2CBus) Close() error {
	return nil
}
