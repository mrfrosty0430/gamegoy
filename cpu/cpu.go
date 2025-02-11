package cpu

type CPU struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8
	F uint8
	H uint8
	L uint8

	SP uint16
	PC uint16
}

const (
	FlagZ = 7 // Zero flag
	FlagN = 6 // Subtract flag
	FlagH = 5 // Half carry flag
	FlagC = 4 // Carry flag
)

func (c *CPU) Execute(opcode uint8) {
	block := c.GetBlock(opcode)
	switch block {
	case 0:
		c.ExecuteBlock0(opcode)
	case 1:
		c.ExecuteBlock1(opcode)
	case 2:
		c.ExecuteBlock2(opcode)
	case 3:
		c.ExecuteBlock3(opcode)
	default:
		break

	}

}

func (c *CPU) SetZFlag(value uint8) {
	zerod = c.F && (1 << 7)
}

func (c *CPU) ExecuteBlock0(opcode uint8) {
	switch opcode {
	case 0x00:
		break
	case 0x07:
	}

}

func (c *CPU) ExecuteBlock1(opcode uint8) {

}
func (c *CPU) ExecuteBlock2(opcode uint8) {

}
func (c *CPU) ExecuteBlock3(opcode uint8) {

}

func (c *CPU) GetBlock(opcode uint8) uint8 {
	return opcode >> 6
}
