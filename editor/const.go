package editor

type ColType byte

const (
	BYTE   ColType = 2
	INT16  ColType = 3
	UINT16 ColType = 4
	INT32  ColType = 5
	UINT32 ColType = 6
	STRING ColType = 7
	FLOAT  ColType = 8
)

var (
	ansiChars   = []rune{'€', '\x81', '‚', 'ƒ', '„', '…', '†', '‡', 'ˆ', '‰', 'Š', '‹', 'Œ', '\x8d', 'Ž', '\x8f', '\x90', '‘', '’', '“', '”', '•', '–', '—', '˜', '™', 'š', '›', 'œ', '\x9d', 'ž', 'Ÿ', '\xa0', '¡', '¢', '£', '¤', '¥', '¦', '§', '¨', '©', 'ª', '«', '¬', '\xad', '®', '¯', '°', '±', '²', '³', '´', 'µ', '¶', '·', '¸', '¹', 'º', '»', '¼', '½', '¾', '¿', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Æ', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï', 'Ð', 'Ñ', 'Ò', 'Ó', 'Ô', 'Õ', 'Ö', '×', 'Ø', 'Ù', 'Ú', 'Û', 'Ü', 'Ý', 'Þ', 'ß', 'à', 'á', 'â', 'ã', 'ä', 'å', 'æ', 'ç', 'è', 'é', 'ê', 'ë', 'ì', 'í', 'î', 'ï', 'ð', 'ñ', 'ò', 'ó', 'ô', 'õ', 'ö', '÷', 'ø', 'ù', 'ú', 'û', 'ü', 'ý', 'þ', 'ÿ'}
	typeSizes   = map[ColType]uint64{BYTE: 1, INT16: 2, UINT16: 2, INT32: 4, UINT32: 4, FLOAT: 4}
	typeTitles  = map[ColType]string{BYTE: "Byte", INT16: "Int16", UINT16: "UInt16", INT32: "Int32", UINT32: "UInt32", FLOAT: "Float", STRING: "String"}
	rTypeTitles = make(map[string]ColType)
)

func init() {
	for typ, title := range typeTitles {
		rTypeTitles[title] = typ
	}
}
