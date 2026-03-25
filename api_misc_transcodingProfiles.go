package broadpeakio

import (
	"encoding/json"
	"fmt"
)

type TranscodingProfileOutput struct {
	Name       string          `json:"name"`
	Content    json.RawMessage `json:"content"`
	Id         uint            `json:"id"`
	InternalId string          `json:"internalId"`
}

func (client BroadpeakClient) GetAllTranscodingProfiles(offset uint, limit uint) ([]TranscodingProfileOutput, error) {
	url := client.getBaseUrl() + "transcoding-profiles"

	url = addOffsetUrl(url, offset, limit)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return []TranscodingProfileOutput{}, err
	}

	var transcodingProfileApiResponse []TranscodingProfileOutput

	err = deserialiseApiResponse(resp, &transcodingProfileApiResponse)
	if err != nil {
		return []TranscodingProfileOutput{}, err
	}

	return transcodingProfileApiResponse, nil
}

func (client BroadpeakClient) GetTranscodingProfile(id uint) (TranscodingProfileOutput, error) {
	url := fmt.Sprintf(client.getBaseUrl()+"transcoding-profiles/%d", id)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return TranscodingProfileOutput{}, err
	}

	var transcodingProfileApiResponse TranscodingProfileOutput

	err = deserialiseApiResponse(resp, &transcodingProfileApiResponse)
	if err != nil {
		return TranscodingProfileOutput{}, err
	}

	return transcodingProfileApiResponse, nil
}
