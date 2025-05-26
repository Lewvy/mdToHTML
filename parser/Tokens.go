package parser

const (
	LBRACE          string = "["
	RBRACE                 = "]"
	H1                     = "# "
	H2                     = "## "
	H3                     = "### "
	BACKSLASH              = `\`
	STAR                   = "*"
	CHECKBOX               = "- [ ] "
	CheckTokenStart        = `<label><input type="checkbox" name="option%d" value=%d> `
	CheckTokenEnd          = "</label><br>"
)
