package parser

// Content holds the parsing content of scripts
type Content []byte

// Token holds the type of token that is read
type Token uint32

// Block is the structure for what to store
type Block struct {
	content Content
	token   Token
	lineNo  uint64
	charNo  uint64
}

// Parser holds a list of ParserBlock
type Parser []Block
