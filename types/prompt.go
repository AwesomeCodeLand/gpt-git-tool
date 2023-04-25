package types

const DiffSystemPrompt = `Please answer the following questions as a git expert`
const DiffUserPrompt = `I hope you can understand the changes in the following files and generate a commit message. This is a content format example:
content: [
file: a.go
change:
+ a := 1
- b := 2
]
answer: files: <file name>
		commit: <commit message>

Here is a commit message example for your reference:

files: a.go
commit: add a variable a and delete b variable

Please understand the changes in the following file and generate a commit message for the changes.

[
file: %s
change: 
%s
]
`

const DiffMaxTokens = 310
