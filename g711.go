package audio

const AmiMask16 int16 = 0x55
const AmiMask8 byte = 0x55

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
		mask = AmiMask16 | 0x80
	} else {
		mask = AmiMask16
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

func Alaw2linear(alaw byte) (linear int16) {
	var i int16

	alawBuf := alaw ^ AmiMask8
	i = (int16(alawBuf & 0x0F) << 4);
	seg := (int16(alawBuf) & 0x70) >> 4;
	if seg != 0 {
		i = (i + 0x100) << (seg - 1)
	}
		
	if (alawBuf & 0x80) != 0 {
		return i
	}
	return -i
}