package audio

const AmiMask int16 = 0x55

var (
	segEnd = [8]int16{
		0xFF, 0x1FF, 0x3FF, 0x7FF, 0xFFF, 0x1FFF, 0x3FFF, 0x7FFF,
	}
)

func Linear2alaw(linear int16) (alaw byte) {
	var seg int16
	var mask int16

	pcmVal := linear
	if pcmVal >= 0 {
		mask = AmiMask | 0x80
	} else {
		mask = AmiMask
		pcmVal = -pcmVal
	}
	for seg = 0; seg < 8; seg++ {
		if pcmVal <= segEnd[seg] {
			break
		}
	}

	if seg != 0 {
		return byte(((seg << 4) | ((pcmVal >> uint(seg+3)) & 0x0F)) ^ mask)
	}
	return byte(((seg << 4) | ((pcmVal >> 4) & 0x0F)) ^ mask)
}