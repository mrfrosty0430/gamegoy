//cpu of zilog-z80 microprocessor (8bits)



type Registers struct{
  a byte
  b byte
  c byte
  d byte
  e byte
  f byte
  h byte
  l byte
  sp uint16
  pc uint16
}

func (regs *Registers) getZero() {
  return (regs.f >> 8 ) & 0x1
}

func (regs *Registers) getAddSub(){
  return (regs.f >> 7) & 0x1
}

func (regs *Registers) getHalfCarry(){
  return (regs.f >> 6) & 0x1
}

func (regs *Registers) getCarry(){
  return (regs.f >> 5) & 0x1
}

func (regs *Registers) setZero(){
  regs.f =.f = (regs.f | 0x80)
}

func (regs *Registers) setAddSub(){
  regs.f =.f = (regs.f | 0x60)
}

func (regs *Registers) setHalfCarry(){
  regs.f = (regs.f | 0x40)
}

func (regs *Registers) setCarry(){
  regs.f = (regs.f | 0x20)
}

func (regs *Registers) resetZero(){
  regs.f = (regs.f & 0x7f)
}

func (regs *Registers) resetAddSub(){
  regs.f = (regs.f & 0xbf)
}

func (regs *Registers) resetHalfCarry(){
  regs.f = (regs.f & 0xdf)
}

func (regs *Registers) resetCarry(){
  regs.f = (regs.f & 0xef)
}
