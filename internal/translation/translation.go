package translation

type Translation struct {
	Source      string `example:"auto"`
	Destination string `example:"en"`
	Original    string `example:"текст для перевода"`
	Translation string `example:"text for translation"`
}
