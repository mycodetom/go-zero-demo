// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ExpandReq struct {
	Shorten string `form:"shorten"`
}

type ExpandResp struct {
	Url string `json:"url"`
}

type ShortenReq struct {
	Url string `form:"url"`
}

type ShortenResp struct {
	Shorten string `json:"shorten"`
}
