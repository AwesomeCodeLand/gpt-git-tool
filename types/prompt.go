package types

const DiffSystemPrompt = `Please answer the following questions as a git expert`
const DiffUserPrompt = `Please understand the changes in the following file and generate a commit message for the changes.
file: %s
change: %s
`

const DiffMaxTokens = 310
