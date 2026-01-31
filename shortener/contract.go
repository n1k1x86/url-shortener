package shortener

type Shortener interface {
	GetLinkByShort(short string) (string, error)
	GetAllLinks() ([]string, error)
	DeleteLink(short string) (bool, error)
	ShortLink(source string, short string) (bool, error)
}
