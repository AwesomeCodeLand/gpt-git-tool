package types

const DiffSystemPrompt = `Please answer the following questions as a git expert`
const DiffUserPrompt = `Please generate submission information that complies with best practice standards for the following changes.
file: %s
change: %s
`

const DiffMaxTokens = 310
