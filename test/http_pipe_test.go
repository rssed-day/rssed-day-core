package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rssed-day/rssed-day-core/context"
	"github.com/rssed-day/rssed-day-core/http/dtos"
	"github.com/rssed-day/rssed-day-core/plugins/inputs"
	"github.com/rssed-day/rssed-day-core/plugins/outputs"
	"net/http"
	"testing"
)

func TestHttpPipe(t *testing.T) {
	url := "http://localhost:51401/api/v1/pipelines/action"
	body := &dtos.PipelineActionModel{
		Action: "pipe",
		Config: context.Config{
			InputConfigs: []inputs.InputConfig{
				{
					Name: "file",
					Args: map[string]interface{}{
						"path": "assets/helloworld.json",
					},
				},
			},
			OutputConfigs: []outputs.OutputConfig{
				{
					Name: "cmd",
				},
			},
		},
	}
	bts, err := json.Marshal(body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(bts))
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 || res.StatusCode < 200 {
		t.Errorf(fmt.Sprintf("%s", res.Status))
		return
	}
	return
}
