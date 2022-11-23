# The PP4 Graphics Card
This is a basic "graphics chip" that is meant
to somewhat emulate what the NES PPU could do.

## General Description
The Card stores up to 16 'Tiles' and 16 'Palettes' and can draw them on the screen
wherever you like within a 16x16 grid using whatever colours are in the palette.
Furthermore, basic alterations including flipping, rotating, and scaling tiles are
supported to make the most of the little tile memory available.

(Note for possible future iteration: have 2n tile/palette ids and draw them seperately somehow)

## Instructions
| Nr. | Bits | Description |
|-----|------|-------------|
|  0  |  00  | Takes information about the contents of a Tile and stores it with the given ID |
|  1  |  01  | Takes information about the colours of a palette and stores it with the given ID |
|  2  |  10  | Takes a tile and palette ID and draws it at the grid coordinates provided |
|  3  |  11  | Takes a bitmask of possible alterations and applies them to the pixture |

### Set Tile Info
This instruction takes the arguments as follows:

    0x0 0000 - First nyble is the Tile-ID/Tile address
    0x1 0000 - Second nyble is the Tile section ID
    0x2 0000 - Third nyble is the high-bits of the section
    0x3 0000 - Fourth nyble is the low-bits of the section

#### Sections
The sections of the tile are laid out as follows:

    0------- 1-------
    2------- 3-------
    4------- 5-------
    6------- 7-------
    8------- 9-------
    10------ 11------
    12------ 13------
    14------ 15------

Each pixel in the 8x8 Tile is a 2-bit number with the colour
information to be retrieved from the palette.

#### Example
To clarify, here's an example:
The arguments `0001 0101 1100 1010` would break down like so:

    0001 0101 ---> Section 5 of Tile 1 is being edited
    1100 1010 ---> The pixels in Section 5 should be set like this:

    1 1 0 0   --\
     1 0 1 0  --/  11 10 01 00

Note: The pixels are in reverse order from this in the section.
i.e. The pixels in the section of this example would be: `00 01 10 11`

The specific colour values of these 2-bit pixels are set by the [Palette](###SetPaletteInfo)

### Set Palette Info
This instruction takes the arguments as follows:

    0x0 0000 - First nyble is the Palette-ID/Palette address
    0x1 0000 - Second nyble is the red component of the 4 3-bit colours
    0x2 0000 - Third nyble is the green component of the 4 3-bit colours
    0x3 0000 - Fourth nyble is the blue component of the 4 3-bit colours

The last three nybles interleave like so:

    1100 0101 1100:

    1  1  0  0
     0  1  0  1   ---> 100 110 001 011
      0  0  1  1

These colours are then referred to from the tiles starting with the 
least-significant bits:

    00 ---> 011
    01 ---> 001
    10 ---> 110
    11 ---> 100

It's also important to note that the 0th palette element will be considered the
transparency colour for the palette and any pixels in the tile with that colour
will be ignored when drawing

### Draw Tile
This instruction takes the arguments as follows:

    0x0 0000 - First nyble is the Tile-ID/Tile address
    0x0 0000 - Second nyble is the Palette-ID/Palette address
    0x2 0000 - Third nyble is the X-Coordinate in a 16x16 grid to draw the tile
    0x3 0000 - Fourth nyble is the Y-Coordinate in a 16x16 grid to draw the tile


### Set Palette Info
This instruction takes the arguments as follows:

    0x0 0000 - First nyble is the Palette-ID/Palette address
    0x1 0000 - Second nyble is the red component of the 4 3-bit colours
    0x2 0000 - Third nyble is the green component of the 4 3-bit colours
    0x3 0000 - Fourth nyble is the blue component of the 4 3-bit colours