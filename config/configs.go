package giconfig

var (
	entries []Config
)

func Add(key string, value interface{}, description string) {
	entries = append(entries, Config{
		Key:         key,
		Value:       value,
		Description: description,
	})
}

func Entries() []Config {
	return entries
}
