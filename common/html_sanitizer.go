package common

import (
	"github.com/microcosm-cc/bluemonday"
)

// BlogHTMLPolicy tạo policy cho content blog
// Cho phép các thẻ an toàn cho blog như h1-h6, p, div, span, strong, em, ul, ol, li, a, img, br
func BlogHTMLPolicy() *bluemonday.Policy {
	p := bluemonday.NewPolicy()

	// Cho phép các thẻ heading
	p.AllowElements("h1", "h2", "h3", "h4", "h5", "h6")

	// Cho phép các thẻ text formatting
	p.AllowElements("p", "div", "span", "br")
	p.AllowElements("strong", "b", "em", "i", "u")

	// Cho phép danh sách
	p.AllowElements("ul", "ol", "li")

	// Cho phép blockquote
	p.AllowElements("blockquote")

	// Cho phép links với href (chỉ http, https)
	p.AllowAttrs("href").OnElements("a")
	p.AllowElements("a")
	p.RequireNoReferrerOnLinks(true)

	// Cho phép images với src, alt, width, height
	p.AllowAttrs("src", "alt", "width", "height").OnElements("img")
	p.AllowElements("img")

	// Cho phép tables
	p.AllowElements("table", "thead", "tbody", "tr", "th", "td")

	// Cho phép một số style attributes an toàn
	p.AllowAttrs("style").OnElements("p", "div", "span", "h1", "h2", "h3", "h4", "h5", "h6")

	// Cho phép class và id (để styling)
	p.AllowAttrs("class", "id").Globally()

	return p
}

// SanitizeBlogHTML sanitize HTML content cho blog
func SanitizeBlogHTML(html string) string {
	policy := BlogHTMLPolicy()
	return policy.Sanitize(html)
}

// StrictHTMLPolicy tạo policy nghiêm ngặt hơn (chỉ text formatting cơ bản)
func StrictHTMLPolicy() *bluemonday.Policy {
	p := bluemonday.NewPolicy()

	// Chỉ cho phép text formatting cơ bản
	p.AllowElements("p", "br", "strong", "b", "em", "i")

	return p
}

// SanitizeStrictHTML sanitize HTML với policy nghiêm ngặt
func SanitizeStrictHTML(html string) string {
	policy := StrictHTMLPolicy()
	return policy.Sanitize(html)
}
