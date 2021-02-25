package resample

import (
	"encoding/binary"
	"math"
)

func Resample(data []byte, srcSampleRate uint, dstSampleRate uint) []byte {
	factor := float64(dstSampleRate) / float64(srcSampleRate)
	length := int(math.Ceil(float64(len(data)) * factor))
	resampleData := make([]byte, length)
	step := int(factor)
	bitDepth := int(16 / 8)
	postSampleSize := step * bitDepth

	for i, j := 0, 0; i+1 < len(data) && j < length; i, j = i+bitDepth, j+postSampleSize {
		upSampleByte := make([]byte, postSampleSize)
		sample := binary.LittleEndian.Uint16(data[i : i+step])
		upSample := uint32(sample) << uint(bitDepth*(step-1)*8)

		binary.LittleEndian.PutUint32(upSampleByte, upSample)
		for k, s := range upSampleByte {
			resampleData[j+k] = s
		}
	}

	return resampleData
}
