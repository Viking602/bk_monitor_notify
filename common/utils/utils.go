package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func DingTalkSign(secret string, timestamp int64) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(stringToSign))
	signData := hash.Sum(nil)
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
	return sign
}

// TimeToTimestampAdd8 转换时间为时间戳并加八小时
func TimeToTimestampAdd8(timeStr string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		return 0
	}
	sr := theTime.Add(8 * time.Hour).Unix()
	return sr
}
