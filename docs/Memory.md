
# Arch40 Memory Mapping
Arch40 uses 8-bit addressing to be able to access up to
256 nybles of data. Additionally, Program Memory and RAM
are on separate busses, meaning one can address up to 256 instructions
(See [Instruction Set] for info on instruction size) and 256 RAM nybles.

## RAM Map

                +-------------+
    0x00 - 0x0F |  Zero-Page  |
                |-------------|
    0x10 - 0x1F |    Stack    |
                |-------------|
           0x20 |             |
                |             |
                |             |
                |             |
                |             |
                |             |
                |             |
                |             |
                |             |
                |             |
                |             |
           0xDF |             |
                |-------------|
    0xE0 - 0xEF | Peripherals |
                |-------------|
    0xF0 - 0xFF | Expansions  |
                +-------------+

### Zero-Page
*might not be in the final version...*

### Stack
The stack is a page of 16 nybles which is mainly used for storing return addresses
when subroutines are called. The stack-pointer (SP) is handled by the CPU and
**can** under-/overflow, so watch out for that.

When the "JSR" (Jump to SubRoutine) instruction is called, the program address
of the instruction **after it** is pushed onto the stack, the address is 2 nybles
and is pushed on little-endian (high-nyble first).

When the "RTS" (ReTurn from Subroutine) instruction is called, the program
adress nybles are popped of the stack and the instruction pointer is moved
to the return address.

### Peripherals
This page is designated for reading and writing to and from basic peripherals.
It can be used for reading the state of controller buttons or a keyboard for example,
and could also be used to set basic outputs like LEDs.

*TODO: Figure out format*

### Expansions
This block of nybles is for more advanced peripherals like graphics cards or
sound boards. It is split into 4 groups of 4 nybles which correspond to the
four expansion slots in the computer. When the "TEX" (Trigger EXpansion Card)
instruction is called, the upper 2 bits of the argument will indicate which
expansion card is called.

| Start | End  | Expansion Slot |
|-------|------|----------------|
| 0xF0  | 0xF3 | Slot 0         |
| 0xF4  | 0xF7 | Slot 1         |
| 0xF8  | 0xFB | Slot 2         |
| 0xFC  | 0xFF | Slot 3         |


