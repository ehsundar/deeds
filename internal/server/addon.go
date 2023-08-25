package server

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var apiKey string

func init() {
	var ok bool
	apiKey, ok = os.LookupEnv("API_KEY")
	if !ok {
		panic("please provide API_KEY environ variable")
	}
}

func (s *Server) attachAddon(ctx context.Context, token string) error {
	url := fmt.Sprintf("https://api.divar.ir/v1/open-platform/add-ons/post/%s", token)
	addon := fmt.Sprintf(addonFormat, token, token)
	payload := []byte(addon) // Request body

	// Create a new HTTP request with the POST method and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	return nil
}

var addonFormat = `
{"widgets":{"widgetList":[{"data":{"@type":"type.googleapis.com/widgets.LegendTitleRowData","title":"سند اظهار شده ملک توسط آگهی‌گذار","imageUrl":"logo","hasDivider":true},"widgetType":"LEGEND_TITLE_ROW"},{"data":{"text":"با مشاهده سند ملک، از صحت و ثقم اطلاعات آگهی مطمئن شوید و وقت خود را صرف یافتن خانه دلخواه خود کنید.","@type":"type.googleapis.com/widgets.DescriptionRowData"},"widgetType":"DESCRIPTION_ROW"},{"data":{"@type":"type.googleapis.com/widgets.WideButtonBarWidgetData","style":"SECONDARY","button":{"title":"مشاهده رایگان سند ملک","action":{"type":"LOAD_WEB_VIEW_PAGE","payload":{"url":"https://ehsandar.com/view?token=%s","@type":"type.googleapis.com/widgets.LoadWebViewPagePayload"},"fallbackLink":"https://ehsandar.com/view?token=%s"}}},"widgetType":"WIDE_BUTTON_BAR"}]}}
`
