package utils

import "net/url"

func BuildNewSessionUrl(appUrl, localUrl, fileId, authToken string) (string, error) {
	baseUrl, err := url.Parse(appUrl)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("baseUrl", localUrl)
	params.Add("fileId", fileId)
	params.Add("token", authToken)
	params.Add("version", "v1")

	return baseUrl.String() + "/new#" + params.Encode(), nil
}
