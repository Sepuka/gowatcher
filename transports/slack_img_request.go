package transports

import (
	"bytes"
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

const (
	fileType = "image/png"
	fileName = "file.png"
	fileExt  = "png"
	title    = "Load average"
	channels = "CBYQ32MN3" //TODO move to config
)

type uploadedFile struct {
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

type filesUploadAPIResponse struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	File  uploadedFile `json:"file"`
}

func (obj slack) sendImg(msg command.Result) (err error) {
	request, err := buildImgRequest(msg, obj.cfg.FileUploadUrl, obj.cfg.Token)

	if err != nil {
		obj.
			logger.
			With(zap.Error(err)).
			Error("Build slack request failed")

		return err
	}

	sender := &loggedRequestSender{
		obj.httpClient,
		obj.logger,
		new(filesUploadAPIResponse),
	}

	err = sender.sendRequest(request)
	if err != nil {
		return err
	}

	answer := sender.answer.(*filesUploadAPIResponse)
	if answer.Ok {
		obj.
			logger.
			With(zap.Any("file", answer.File)).
			Info("Img was uploaded to slack")
	} else {
		obj.
			logger.
			With(zap.Any("err", answer.Error)).
			Error("Img not uploaded to slack with error")
	}

	return err
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
	writer.WriteField("Token", token)

	req, _ := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}

func writeFile(w *multipart.Writer, content []byte) error {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="file"; fileName="%v"`, fileName))
	header.Set("Content-Type", fileType)
	part, err := w.CreatePart(header)
	part.Write(content)

	return err
}
