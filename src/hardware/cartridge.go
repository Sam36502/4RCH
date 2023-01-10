package hardware

import "io/ioutil"

const (
	SIZE_PROG = 256
	SIZE_DATA = 4096
	FILE_MODE = 0650
)

type Cart struct {
	Program [SIZE_PROG]nybble
	Data    [SIZE_DATA]nybble
}

func LoadCartFromFile(filename string) (*Cart, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cart := Cart{}
	for bi, b := range fileData {
		for hn := 0; hn < 2; hn++ {
			ni := (bi << 1) | hn // Calculate nybble index
			if ni <= SIZE_PROG {
				cart.Program[ni] = nybble(b)
			} else {
				ni -= SIZE_PROG
				cart.Data[ni] = nybble(b)
			}
		}
	}

	return &cart, nil
}

func SaveCartToFile(filename string, cart Cart) error {
	fileData := []byte{}

	for ni := 0; ni < SIZE_PROG+SIZE_DATA; ni++ {
		bi := ni >> 1
		hn := ni % 2

		b := byte(0)
		if ni <= SIZE_PROG {
			b = byte(cart.Program[ni])
		} else {
			ni -= SIZE_PROG
			b = byte(cart.Data[ni])
		}

		fileData[bi] |= b << (hn * 4)
	}

	return ioutil.WriteFile(filename, fileData, FILE_MODE)
}
