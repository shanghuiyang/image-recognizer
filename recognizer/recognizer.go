package recognizer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/shanghuiyang/image-recognizer/oauth"
)

const (
	baiduURL = "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general"
)

// Recognizer ...
type Recognizer struct {
	auth *oauth.Oauth
}

type response struct {
	ResultNum int32    `json:"result_num"`
	Results   []result `json:"result"`
	ErrorCode int      `json:"error_code"`
	ErrorMsg  string   `json:"error_msg"`
}

type result struct {
	Score   float32 `json:"score"`
	Root    string  `json:"root"`
	Keyword string  `json:"keyword"`
}

// New ...
func New(auth *oauth.Oauth) *Recognizer {
	return &Recognizer{
		auth: auth,
	}
}

// Recognize ...
func (r *Recognizer) Recognize(imageFile string) (string, error) {
	token, err := r.auth.GetToken()
	if err != nil {
		return "", err
	}

	b64img, err := r.b64Image(imageFile)
	if err != nil {
		return "", err
	}

	formData := url.Values{
		"access_token": {token},
		"image":        {b64img},
	}
	resp, err := http.PostForm(baiduURL, formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	if res.ErrorCode > 0 {
		return "", fmt.Errorf("error_code: %v, error_msg: %v\n", res.ErrorCode, res.ErrorMsg)
	}

	if res.ResultNum > 0 {
		return res.Results[0].Keyword, nil
	}
	return "", fmt.Errorf("failed to recognize")
}

func (r *Recognizer) b64Image(imageFile string) (string, error) {
	file, err := os.Open(imageFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	image, err := ioutil.ReadAll(file)
	if err != nil {
		return "nil", err
	}
	b64img := base64.StdEncoding.EncodeToString(image)
	return b64img, nil
}
