package messagefactory

import "fmt"

func Text(text string) string {
	return fmt.Sprintf("%s \n", text)
}

func TextInline(text string) string {
	return fmt.Sprintf("%s ", text)
}

func Bold(text string) string {
	return fmt.Sprintf("<b>%s</b> \n", text)
}

func Italic(text string) string {
	return fmt.Sprintf("<i>%s</i> \n", text)
}

func ItalicInline(text string) string {
	return fmt.Sprintf("<i>%s</i> ", text)
}

func Strike(text string) string {
	return fmt.Sprintf("<s>%s</s> \n", text)
}

func Underline(text string) string {
	return fmt.Sprintf("<u>%s</u> \n", text)
}

func Spoiler(text string) string {
	return fmt.Sprintf("<tg-spoiler>%s</tg-spoiler> \n", text)
}

func link(url string, text string) string {
	return fmt.Sprintf("<a href='%s'>%s</a> \n", url, text)
}

func Code(text string) string {
	return fmt.Sprintf("<pre>%s</pre> \n", text)
}

func Copy(text string) string {
	return fmt.Sprintf("<code>%s</code> \n", text)
}

func CopyInline(text string) string {
	return fmt.Sprintf("<code>%s</code> ", text)
}

func Seperator() string {
	return fmt.Sprintf("<code>-------------------------------</code> \n")
}
