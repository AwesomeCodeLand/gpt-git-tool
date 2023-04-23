package helper

import (
	"encoding/json"
	"fmt"
	"ggt/models"
	"ggt/types"
	"io/ioutil"
	"net/http"
	"strings"
)

// OpenAI struct
type OpenAI struct {
	url  string
	key  string
	msgs []models.OpenAIMessages
}

func NewOpenAI(key string) *OpenAI {
	return &OpenAI{
		url:  "https://bot.devexp.cn/api/chat-stream",
		key:  key,
		msgs: []models.OpenAIMessages{},
	}
}

func (ai *OpenAI) Diff(diff map[string]string) (answer string, err error) {
	ai.WithSystemPrompt(types.DiffSystemPrompt)
	for file, content := range diff {
		ai.WithUserPrompt(fmt.Sprintf(types.DiffUserPrompt, file, content))
		_answer, err := ai.do(types.DiffMaxTokens)
		if err != nil {
			return "", &types.GptError{
				Err:  err,
				Code: types.GptFailedError,
			}
		}

		answer = fmt.Sprintf("%s\n\t%s\n%s", answer, file, _answer)
		fmt.Printf("[%s] 处理完成\n", file)
		ai.CleanUserPrompt()
	}

	return
}

func (ai *OpenAI) CleanUserPrompt() {
	ai.msgs = ai.msgs[:1]
}

func (ai *OpenAI) WithSystemPrompt(prompt string) *OpenAI {
	ai.msgs = append(ai.msgs, models.OpenAIMessages{
		Role:    "system",
		Content: prompt,
	})
	return ai
}

func (ai *OpenAI) WithUserPrompt(prompt string) *OpenAI {
	ai.msgs = append(ai.msgs, models.OpenAIMessages{
		Role:    "user",
		Content: prompt,
	})
	return ai
}

// Path: helper/gpt.go
// do Invoke openAI GPT-3.5 API
func (ai *OpenAI) do(maxTokens int) (answer string, err error) {
	url := ai.url
	method := "POST"

	gptRequest := models.OpenAI{
		Messages:        ai.msgs,
		Stream:          false,
		Model:           "gpt-3.5-turbo",
		Temperature:     0.7,
		PresencePenalty: 2,
		MaxTokens:       maxTokens,
	}

	data, err := json.Marshal(gptRequest)
	if err != nil {
		return "", err
	}

	// fmt.Println(string(data))
	payload := strings.NewReader(string(data))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Host", "bot.devexp.cn")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Origin", "https://bot.devexp.cn")
	req.Header.Add("Referer", "https://bot.devexp.cn/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Safari/605.1.15")
	// req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("path", "v1/chat/completions")
	req.Header.Add("access-code", ai.key)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	str := string(body)
	// fmt.Println(str)

	str = strings.TrimPrefix(str, "```json")
	str = strings.TrimSuffix(str, "```")

	// fmt.Println(str)
	var gptResponse models.OpenAIResponse
	err = json.Unmarshal([]byte(str), &gptResponse)
	if err != nil {
		return "", err
	}

	if gptResponse.Error.Message != "" {
		return "", fmt.Errorf("%v", gptResponse.Error)
	}

	return gptResponse.Choices[0].Message.Content, nil
}
