package datastruct

func RollingHash(s string) int {
	h := 0
	A, B := 256, 3571
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}

	return h
}

type keyValue struct {
	key   string
	value string
}

type Map struct {
	keyArray [3571][]keyValue
}

func CreateMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	h := RollingHash(key)
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := RollingHash(key)

	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}
	return " "
}
