package main

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// base64でエンコードしたRSAキー
var privateKeyRaw = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBbytmVzlrdjJZNUZqa3k4VFFMSzNyWXRqcUNLSUVmMFlqbTFsWFRua3BXVkJqTzd4CkV6QjdIWUJrSG5JQTE5SEVkemplS0xsTGNKZXZPVTZHM3ArdDgvNXZWaGNpUG5oS3B3VTZaenJSMFAzUTJ0b0MKK0tndHJuUEhYcG04UHkvSFZ3QnhIV1ozNTJOTmE1SjNkSHY2YTdCMWs3SU1BYUJUMDUzUDk5bDFOUXJCQ1ZKZgprWGNrUUFVT25zWWsvUEtGRnVmTmhsdTNuQWQrZVVsMEl2MUlBV3NVYjVsSEtrSU9PV3hubHJlRk41Z2RDZ0w2Cll4THdYUkJWMDEwenUvWTl6Rjh6ZGhTWlhXb0l2RC9KcE92a0toOXltMVZFMFl6WWtTVyttMFhuVGE0Vk1kWmYKbVdjUU5LaFlaZE5nMFRxd1Jjbzg5M2ZWZW83NTZKek9IUTBLY1FJREFRQUJBb0lCQURueE80eldUY1BsSWM1bQpWZWdKUmVXVDRTY1BEZ3RON2VCcnkrbXBEYXRxb0VHeU5vY1NISFBSYjVOVHVmaVJyMkoyT0JNdGJmM2ZvWklnCnNJMEM1ZnZiZG9CNHJTSm9ZL3VuWFg0Z3pRdVVyc0N2VVYyV0ZwQVZYZVdUYzNqaTZ4VldrcVpTRXhFNWlhVDkKb2owTGx2dDY1aFhRVzI0M3YycVAzM1UrNnJCVkMvMk5vekd6bHE5dUh2TkpVUkYwL1p6dElvNDdLNnErZ2laQwpPTnZ5Y1hWR1Q4aDVQSXpmUXNubnJzYUxZQ05lK3lKMXg5elBUMzJ4U3RJOCtyRi9QRkJPMVQ4dW1uYTYra1kyCnBZNVhRZmlZSDdnZlVJcWZqQTQ4T3ZueGRteWs5aWI2RHpjZ0IrdmxZKzRjZ0NHWDI3c2JHOTlweEdLdThuR2cKaUpqQkdBRUNnWUVBMUM0TitSd2VGaTJ2ekNCTG5yU0NaMlU2WGpGSnlnOXVnTCtYMDFETTRyM09SVmxGVjVXWgoyZWJLUzBFZndoM0pnNmZ1dmYvMi82c0ZQekg0eDhHSUtiOFdsZHNoOW1BR1VRek9KK2xVUnFhRW5idys1eFNzCnk3N09STlVlcy9tZHI1WW1LcHlGc0taWE5ES2FzSWVRNy91c3A1QndjdU90azVGNys1NnUrODBDZ1lFQXhjR0cKZzNYTnBicVczbW1HZjU5YXUrZW9mWWJ6bGlXWEZLZjNTNkExTXAyRUFHMDJMYjdIMnNOUUtJMkdYTUFCTGVBZgp5bE5MaWdLblNCZ2MybjRJY2U5MWRka2pKaHo1emsxUDZScGp6NHlDaE1mQzBxcXFEYnJwd3JqdXpIUkt0V05KCnhoKzNxUGlaem1wZ2ZLODdkZE5vLzRwVFhZQmovcnRXb2NXdmpUVUNnWUJpZTducW5WMXRwK2tJRXhsbVlaeUIKaDEvUEpvdDhhU3MrUVMra1dzV3VuTERvU1daQkgrUVlXdUljaWU3R3QvSzMxRE9oSnZTcmVPTG5rVGRLNkk1ZAo2aDcra1lhekI2RUc3NjJrb3MyR09YWW1qS0NadTJQMDhleGwwSkgrc1dhNmdEUFk1V3U4TVlrYVpqNmNuMS9yCnMrSlFxRjg1UkdwbHEwcGo0U1NSSlFLQmdRQ3BGWDRRdXRIWnFQOUVMYy90SVBCd2gzTmQyTncrL2ViK3AxcmYKVTUwSXFQdHJiZldNQ1hwU0J2dHJvUTVJRVhjd3BWZ3BJeTBNVkpaNUl2elFxRUJLUXFjWTIwNmRVTnNhVktwRgpzZVd6V1AxajBIVTRzT2x6a2VRN05Pb2c4REhNZzVkWmlsYjdaNHdDbkp2aEgrYmtLSjIzR0t1TTM3RWY1VWYrClM4dDlIUUtCZ0QzYnc3WVpwbnBQMDVmTVBWY3BMTE9keE9tZDhoNmtFS3F6T3UzM1BIcmVNbDlHQ0s3SjhlT0YKb1dxQUsxdWI3dllra2U2dGVabFBqWXlycXV0VFZDR0JaWG4yY3R1RGMvbzR0MTdRazhoWGJRQm42VzArWGdnQgpzVjB0R3kxQ0wraGt5azNFdVVVRFBhcHdYbVNvQTVQNWVKcU4xNlhpcVNRQ3Y3Q2lzMHpyCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t"
var publickeyRaw = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFvK2ZXOWt2Mlk1RmpreThUUUxLMwpyWXRqcUNLSUVmMFlqbTFsWFRua3BXVkJqTzd4RXpCN0hZQmtIbklBMTlIRWR6amVLTGxMY0pldk9VNkczcCt0CjgvNXZWaGNpUG5oS3B3VTZaenJSMFAzUTJ0b0MrS2d0cm5QSFhwbThQeS9IVndCeEhXWjM1Mk5OYTVKM2RIdjYKYTdCMWs3SU1BYUJUMDUzUDk5bDFOUXJCQ1ZKZmtYY2tRQVVPbnNZay9QS0ZGdWZOaGx1M25BZCtlVWwwSXYxSQpBV3NVYjVsSEtrSU9PV3hubHJlRk41Z2RDZ0w2WXhMd1hSQlYwMTB6dS9ZOXpGOHpkaFNaWFdvSXZEL0pwT3ZrCktoOXltMVZFMFl6WWtTVyttMFhuVGE0Vk1kWmZtV2NRTktoWVpkTmcwVHF3UmNvODkzZlZlbzc1Nkp6T0hRMEsKY1FJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t"

func GenerateTokenByPrivateKey() (string, error) {
	privateKeyByte, err := base64.URLEncoding.DecodeString(privateKeyRaw)
	if err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := jwt.StandardClaims{
		Subject:   "{\"id\": 100}",
		IssuedAt:  now.Unix(),
		ExpiresAt: now.AddDate(0, 0, 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	return tokenString, nil
}

func ParseTokenByPublicKey(tokenString string) (*jwt.Token, error) {
	publickeyByte, err := base64.URLEncoding.DecodeString(publickeyRaw)
	if err != nil {
		return nil, err
	}
	publickey, err := jwt.ParseRSAPublicKeyFromPEM(publickeyByte)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publickey, nil
	})
	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}
