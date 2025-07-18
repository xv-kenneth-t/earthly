package parser

// GetLexerModeNames returns the generated mode names.
func GetLexerModeNames() []string {
	EarthLexerInit()
	return EarthLexerLexerStaticData.ModeNames
}

// GetLexerSymbolicNames returns the generated token names.
func GetLexerSymbolicNames() []string {
	EarthLexerInit()
	return EarthLexerLexerStaticData.SymbolicNames
}

// GetLexerLiteralNames returns the generated literal names.
func GetLexerLiteralNames() []string {
	EarthLexerInit()
	return EarthLexerLexerStaticData.LiteralNames
}
