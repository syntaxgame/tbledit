package editor

import (
	"log"
	"strconv"
	"tbl-editor/utils"
)

type Buffer struct {
	data   []byte
	offset uint64
}

func NewBuffer(data []byte) *Buffer {
	return &Buffer{data: data, offset: 0}
}

func (b *Buffer) Read(t ColType) interface{} {

	data := b.ReadN(typeSizes[t])
	switch t {
	case BYTE:
		return data[0]

	case INT16:
		return int16(utils.BytesToInt(data, true))

	case UINT16:
		return uint16(utils.BytesToInt(data, true))

	case INT32:
		return int32(utils.BytesToInt(data, true))

	case UINT32:
		return uint32(utils.BytesToInt(data, true))

	case FLOAT:
		return utils.BytesToFloat(data, true)

	default:
		return nil
	}
}

func (b *Buffer) ReadN(n uint64) []byte {

	if b.offset == uint64(len(b.data)) {
		return []byte{}
	}

	if uint64(len(b.data)) >= b.offset+n {
		b.offset += n
		return b.data[b.offset-n : b.offset]
	}

	data := b.data[b.offset:]
	b.offset = uint64(len(b.data))
	return data
}

func (b *Buffer) Write(value string, t ColType) {

	if t == STRING {
		data := value
		length := uint64(len(data))
		b.data = append(b.data, utils.IntToBytes(length, 4, true)...)
		b.offset += 4

		b.data = append(b.data, []byte(data)...)
		b.offset += length
		return
	}

	val, err := strconv.ParseFloat(value, 10)
	if err != nil {
		log.Fatal(err)
	}

	switch t {
	case BYTE:
		data := byte(val)
		b.data = append(b.data, data)
		b.offset++

	case INT16:
		data := utils.IntToBytes(uint64(int16(val)), 2, true)
		b.data = append(b.data, data...)
		b.offset += 2

	case UINT16:
		data := utils.IntToBytes(uint64(uint16(val)), 2, true)
		b.data = append(b.data, data...)
		b.offset += 2

	case INT32:
		data := utils.IntToBytes(uint64(int32(val)), 4, true)
		b.data = append(b.data, data...)
		b.offset += 4

	case UINT32:
		data := utils.IntToBytes(uint64(uint32(val)), 4, true)
		b.data = append(b.data, data...)
		b.offset += 4

	case FLOAT:
		data := utils.FloatToBytes(float32(val), 4, true)
		b.data = append(b.data, data...)
		b.offset += 4
	}
}

func (b *Buffer) Overwrite(data []byte, i int) {
	_data := make([]byte, len(data))
	copy(_data, data)
	(*b).data = append(((*b).data)[:i], append(_data, ((*b).data)[i+len(data):]...)...)
}

func (b *Buffer) GetBytes() []byte {
	return b.data
}

func (b *Buffer) GetOffset() uint64 {
	return b.offset
}
