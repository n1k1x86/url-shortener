package models

import "url-shortener/shortener/models"

type GetAllLinksResponse struct {
	Data []models.LinkRecord `json:"data"`
}

type ShortLinkBody struct {
	Short  string `json:"short"`
	Source string `json:"source"`
}
