package pkg


switch (opcode){
  /*
  8 bit loads
  loading immediate value into register
  */

  case 0x06:
    n := regs.pc
    regs.pc ++
    regs.b = n

  case 0x0e:
    n := regs.pc
    regs.pc ++
    regs.c = n

  case 0x16:
    n := regs.pc
    regs.pc ++
    regs.d = n

  case 0x1e:
    n := regs.pc
    regs.pc ++
    regs.e = n

  case 0x26:
    n := regs.pc
    regs.pc ++
    regs.h = n

  case 0x2e:
    n := regs.pc
    regs.pc ++
    regs.l = n


  /*
  8 bit loads
  loading register values to Registers
  */

  case 0x7f:
    r2 :=
    regs.a = regs.a

  case 0x78:
    r2 :=
    regs.a = regs.b

  case 0x79:
    r2 :=
    regs.a = regs.c

  case 0x7a:
    r2 :=
    regs.a = regs.d

  case 0x7b:
    r2 :=
    regs.a = regs.e

  case 0x7c:
    r2 :=
    regs.a = regs.h

  case 0x7d:
    r2 :=
    regs.a = regs.l

  case 0x7e:
    r2 :=
    regs.a = regs.hl // memory

  case 0x40:
    r2 :=
    regs.b = regs.b

  case 0x41:
    r2 :=
    regs.b = regs.c

  case 0x42:
    r2 :=
    regs.b = regs.d

  case 0x43:
    r2 :=
    regs.b = regs.e

  case 0x44:
    r2 :=
    regs.b = regs.h

  case 0x45:
    r2 :=
    regs.b = regs.l

  case 0x46:
    r2 :=
    regs.b = regs.hl // memory

  case 0x48:
    r2 :=
    regs.c = regs.b

  case 0x49:
    r2 :=
    regs.c = regs.c

  case 0x4a:
    r2 :=
    regs.c = regs.d

  case 0x4b:
    r2 :=
    regs.c = regs.e

  case 0x4c:
    r2 :=
    regs.c = regs.h

  case 0x4d:
    r2 :=
    regs.c = regs.l

  case 0x4e:
    r2 :=
    regs.c = regs.hl //memory

  case 0x50:
    r2 :=
    regs.d = regs.b

  case 0x51:
    r2 :=
    regs.d = regs.c

  case 0x52:
    r2 :=
    regs.d = regs.d

  case 0x53:
    r2 :=
    regs.d = regs.e

  case 0x54:
    r2 :=
    regs.d = regs.h

  case 0x55:
    r2 :=
    regs.d = regs.l

  case 0x56:
    r2 :=
    regs.d = regs.hl //memory

  case 0x58:
    r2 :=
    regs.e = regs.b

  case 0x59:
    r2 :=
    regs.e = regs.c

  case 0x5a:
    r2 :=
    regs.e = regs.d

  case 0x5b:
    r2 :=
    regs.e = regs.e

  case 0x5c:
    r2 :=
    regs.e = regs.h

  case 0x5d:
    r2 :=
    regs.e = regs.l

  case 0x5e:
    r2 :=
    regs.e = regs.hl // memory

  case 0x60:
    r2 :=
    regs.h = regs.b

  case 0x61:
    r2 :=
    regs.h = regs.c

  case 0x62:
    r2 :=
    regs.h = regs.d

  case 0x63:
    r2 :=
    regs.h = regs.e

  case 0x64:
    r2 :=
    regs.h = regs.h

  case 0x65:
    r2 :=
    regs.h = regs.l

  case 0x66:
    r2 :=
    regs.h = regs.hl // memory

  case 0x68:
    r2 :=
    regs.l = regs.b

  case 0x69:
    r2 :=
    regs.l = regs.c

  case 0x6a:
    r2 :=
    regs.l = regs.d

  case 0x6b:
    r2 :=
    regs.l = regs.e

  case 0x6c:
    r2 :=
    regs.l = regs.h

  case 0x6d:
    r2 :=
    regs.l = regs.l

  case 0x6e:
    r2 :=
    regs.l = regs.hl // memory

  case 0x70:
    r2 :=
    regs.hl = regs.b

  case 0x71:
    r2 :=
    regs.hl = regs.c

  case 0x72:
    r2 :=
    regs.hl = regs.d

  case 0x73:
    r2 :=
    regs.hl = regs.e

  case 0x74:
    r2 :=
    regs.hl = regs.h

  case 0x75:
    r2 :=
    regs.hl = regs.l

  case 0x36:
    n :=
    regs.pc ++
    regs.hl = n // memory


  case 0x0a:
    temp := regs.b << 8 | regs.c
    regs.a = //memory

  case 0x1a:
    temp := regs.d << 8 | regs.e
    regs.a = //memoiry

  case 0x7e:
    temp := regs.h << 8 | regs.l
    regs.a = //memory

  case 0xfa:
    nn = //memory
    regs.pc ++
    regs.a = memory

}
