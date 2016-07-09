package api

type Request struct {

}

type Response struct {
	Title string `json:"title"`
	Locale string `json:"locale"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
