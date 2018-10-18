package transports

import (
	"github.com/sepuka/gowatcher/command"
	"bytes"
	"mime/multipart"
	"net/textproto"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/sepuka/gowatcher/config"
)

const (
	fileType = "image/png"
	fileName = "file.png"
	fileExt  = "png"
	title    = "Load average"
	channels = "CBYQ32MN3"//TODO move to config
)

type UploadedFile struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	Name               string `json:"name"`
	MimeType           string `json:"mimetype"`
	FileType           string `json:"filetype"`
	User               string `json:"user"`
	PrivateUrl         string `json:"url_private"`
	PrivateDownloadUrl string `json:"url_private_download"`
	Permalink          string `json:"permalink"`
	PublicPermalink    string `json:"permalink_public"`
}

type FilesUploadAPIResponse struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	File  UploadedFile `json:"file"`
}

func sendImgRequest(httpClient *http.Client, msg command.Result, cfg config.TransportSlack) {
	request, err := buildImgRequest(msg, cfg.FileUploadUrl, cfg.Token)
	if err != nil {
		log.Println("Build slack request failed: ", err)
	}

	resp, _ := httpClient.Do(request)
	body, _ := ioutil.ReadAll(resp.Body)
	res := new(FilesUploadAPIResponse)
	_ = json.Unmarshal(body, res)

	defer resp.Body.Close()
}

func buildImgRequest(msg command.Result, url string, token string) (*http.Request, error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	err := writeFile(writer, []byte(msg.GetContent()))
	if err != nil {
		return nil, err
	}

	writer.WriteField("title", title)
	writer.WriteField("fileName", fileName)
	writer.WriteField("filetype", fileExt)
	writer.WriteField("channels", channels)
	writer.WriteField("token", token)

	req, _ := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func writeFile(w *multipart.Writer, content []byte) error {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="file"; fileName="%s"`, fileName))
	header.Set("Content-Type", fileType)
	part, err := w.CreatePart(header)
	part.Write(content)

	return err
}