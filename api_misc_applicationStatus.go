package broadpeakio

func (client BroadpeakClient) CheckApplicationStatus() (string, error) {
	url := client.getBaseUrl() + "status"

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return "", err
	}

	return resp, nil
}
