package weibo

type Weibo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Id   int    `json:"id"`
		Url  string `json:"url"`
		Name string `json:"name"`
	} `json:"data"`
}

type AlApi struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		HotWord    string `json:"hot_word"`
		HotWordNum string `json:"hot_word_num"`
	} `json:"data"`
	Author struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
	} `json:"author"`
}
