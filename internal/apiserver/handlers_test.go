package apiserver

import (
	"bytes"
	"github.com/Srgkharkov/test-game/internal/game"
	"github.com/Srgkharkov/test-game/internal/metrics"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type config struct {
	path       string
	KeyValues  [][2]string
	configType string
	configName string
}

func TestAPIHandlers(t *testing.T) {
	game := game.NewGame()
	metrics := metrics.NewMetrics()
	metrics.Run()
	APIServer := NewAPIServer(game, metrics)
	go func() {
		err := APIServer.Run(":8081")
		assert.Nil(t, err, err)
	}()
	time.Sleep(time.Second)

	urladdconfig := "http://localhost:8081/addconfig"
	urlgetresult := "http://localhost:8081/getresult"
	configs := []config{
		{
			KeyValues: [][2]string{
				{"configType", "reels"},
				{"configName", "confreels_1"},
			},
			path: "../../test/confreels.json",
		},
		{
			KeyValues: [][2]string{
				{"configType", "lines"},
				{"configName", "conflines_1"},
			},
			path: "../../test/conflines.json",
		},
		{
			KeyValues: [][2]string{
				{"configType", "payouts"},
				{"configName", "confpayouts_1"},
			},
			path: "../../test/confpayouts.json",
		},
	}
	expectedresult := "{\"lines\":[{\"line\":1,\"payout\":50},{\"line\":2,\"payout\":0},{\"line\":3,\"payout\":0}],\"total\":50}"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	for _, config := range configs {
		// New multipart writer.
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for _, keyValue := range config.KeyValues {
			err := writer.WriteField(keyValue[0], keyValue[1])
			assert.Nil(t, err, err)
		}
		file, err := os.Open(config.path)
		assert.Nil(t, err, err)
		defer file.Close()

		part, err := writer.CreateFormFile("config", filepath.Base(config.path))
		assert.Nil(t, err, err)
		_, err = io.Copy(part, file)
		assert.Nil(t, err, err)

		err = writer.Close()
		assert.Nil(t, err, err)

		request, err := http.NewRequest("POST", urladdconfig, body)
		assert.Nil(t, err, err)

		request.Header.Set("Content-Type", writer.FormDataContentType())

		//client := &http.Client{}
		resp, err := client.Do(request)
		assert.Nil(t, err, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code", resp.StatusCode, config.path)

	}

	var jsonStr = []byte(`{"config_reels_name":"confreels_1", "config_lines_name":"conflines_1", "config_payouts_name":"confpayouts_1"}`)
	request, err := http.NewRequest("POST", urlgetresult, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")

	//client := &http.Client{}
	resp, err := client.Do(request)
	assert.Nil(t, err, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err, err)

	assert.Equal(t, expectedresult, string(body))

	//file, err := os.Open("../test/confreels.json")
	//assert.Nil(t, err, err)
	//defer file.Close()
	//
	//body := &bytes.Buffer{}
	//writer := multipart.NewWriter(body)
	//part, err := writer.CreateFormFile("config", filepath.Base("../test/confreels.json"))
	//assert.Nil(t, err, err)
	//_, err = io.Copy(part, file)
	//assert.Nil(t, err, err)
	//
	//err = writer.WriteField("key", "val")
	//assert.Nil(t, err, err)
	////httptest.NewRequest(http.MethodPost)
	//w := multipart.NewWriter()
}
