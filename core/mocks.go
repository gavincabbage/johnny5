package main

type mockI2CBus struct {}

func (bus mockI2CBus) ReadByte(addr byte) (value byte, err error) {
    return 1, nil
}

func (bus mockI2CBus) WriteByte(addr, value byte) error {
    return nil
}

func (bus mockI2CBus) WriteBytes(addr byte, value []byte) error {
    return nil
}

func (bus mockI2CBus) ReadFromReg(addr, reg byte, value []byte) error {
    return nil
}

func (bus mockI2CBus) ReadByteFromReg(addr, reg byte) (value byte, err error) {
    return 1, nil
}

func (bus mockI2CBus) ReadWordFromReg(addr, reg byte) (value uint16, err error) {
    return 1, nil
}

func (bus mockI2CBus) WriteToReg(addr, reg byte, value []byte) error {
    return nil
}

func (bus mockI2CBus) WriteByteToReg(addr, reg, value byte) error {
    return nil
}

func (bus mockI2CBus) WriteWordToReg(addr, reg byte, value uint16) error {
    return nil
}

func (bus mockI2CBus) Close() error {
    return nil
}
