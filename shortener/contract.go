package shortener

type Shortener interface {
	GetLinkByShort() (string, error)
	GetAllLinks() ([]string, error)
	DeleteLink() (bool, error)
	ShortLink(source string, short string) (bool, error)
}
