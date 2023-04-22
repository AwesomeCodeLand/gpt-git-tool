package config

// Cfg is the global config
type Cfg struct {
	Open OpenAI `yaml:"openai"`
}

type OpenAI struct {
	Token string `yaml:"token"`
}
