package config

var (
	entries  []config
)


func Add(key string, value interface{}, description string) {
	entries = append(entries, config{
		key:         key,
		value:       value,
		description: description,
	})
}



