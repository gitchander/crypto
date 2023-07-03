package base26

func EncodeToString(src []byte) string {
	dst := make([]byte, EncodedLenMax(len(src)))
	n := Encode(dst, src)
	return string(dst[:n])
}

func DecodeString(s string) ([]byte, error) {
	src := []byte(s)
	n, err := Decode(src, src)
	return src[:n], err
}
