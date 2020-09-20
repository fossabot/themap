package themap

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestApplePayBlock(t *testing.T) {

	payload := []byte(strings.TrimSpace(`{"version":"EC_v1","data":"TGuomSxhVblEfMiJdbVe9JeN9K52znlr1Q3n2OsMwz4luSJAZbumg/ZFgBYyty2Dhst4GQinLffz11vO3VH/kdRShRlM74esNW3cz0hRC0Lfg0to2ddqfxzwun4euELpo0Q3CZeBbfC7Hq5WZtnApTUEvpOrshCaJIgj5n7LHcgu0BLpyeENgPuEHFamejzAyRW4vQ2gw4dg1MsHN0WxC5OQ7vLvRyIfXLifIwQb6zSOHGRh3izuAV8V5Gk0CbYq6x+DnKa9MRSm9efF8Uv1yZFyxI13BjD6yfFJtRa7LLOXnn6I+8OekKxDcNxxPmhRK6JW/5KXDrLa4/XgFGr+cxb/jG9F7uyEgXLNveQsQBoIBlcLEqQYSl7UD783XbUYmUwWhwLBh6196hllOJqVPM3J8Q060+5V4nLWnmtcWuGO3A==","signature":"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","header":{"ephemeralPublicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXZRD8nW8mbVLw+fAvjCWUjjUQHHqqi6Fd2VuDD9LZpsJX7aErN1B3HEm7GHPmr4q3oJ6f+yQOVRvtU/hxXMrcw==","publicKeyHash":"eRBhXMP8PITmbMLkhyVMtz4RMb6UmKEC8SpEr54AZdE=","transactionId":"b027c4fc440641a58e1496f23c6f9abfbdaac5292988a55861f77ad6216c34ad"}}`))

	var base64 string = `eyJ2ZXJzaW9uIjoiRUNfdjEiLCJkYXRhIjoiVEd1b21TeGhWYmxFZk1pSmRiVmU5SmVOOUs1MnpubHIxUTNuMk9zTXd6NGx1U0pBWmJ1bWcvWkZnQll5dHkyRGhzdDRHUWluTGZmejExdk8zVkgva2RSU2hSbE03NGVzTlczY3owaFJDMExmZzB0bzJkZHFmeHp3dW40ZXVFTHBvMFEzQ1plQmJmQzdIcTVXWnRuQXBUVUV2cE9yc2hDYUpJZ2o1bjdMSGNndTBCTHB5ZUVOZ1B1RUhGYW1lanpBeVJXNHZRMmd3NGRnMU1zSE4wV3hDNU9RN3ZMdlJ5SWZYTGlmSXdRYjZ6U09IR1JoM2l6dUFWOFY1R2swQ2JZcTZ4K0RuS2E5TVJTbTllZkY4VXYxeVpGeXhJMTNCakQ2eWZGSnRSYTdMTE9Ybm42SSs4T2VrS3hEY054eFBtaFJLNkpXLzVLWERyTGE0L1hnRkdyK2N4Yi9qRzlGN3V5RWdYTE52ZVFzUUJvSUJsY0xFcVFZU2w3VUQ3ODNYYlVZbVV3V2h3TEJoNjE5NmhsbE9KcVZQTTNKOFEwNjArNVY0bkxXbm10Y1d1R08zQT09Iiwic2lnbmF0dXJlIjoiTUlBR0NTcUdTSWIzRFFFSEFxQ0FNSUFDQVFFeER6QU5CZ2xnaGtnQlpRTUVBZ0VGQURDQUJna3Foa2lHOXcwQkJ3RUFBS0NBTUlJRDR6Q0NBNGlnQXdJQkFnSUlUREJCU1ZHZFZEWXdDZ1lJS29aSXpqMEVBd0l3ZWpFdU1Dd0dBMVVFQXd3bFFYQndiR1VnUVhCd2JHbGpZWFJwYjI0Z1NXNTBaV2R5WVhScGIyNGdRMEVnTFNCSE16RW1NQ1FHQTFVRUN3d2RRWEJ3YkdVZ1EyVnlkR2xtYVdOaGRHbHZiaUJCZFhSb2IzSnBkSGt4RXpBUkJnTlZCQW9NQ2tGd2NHeGxJRWx1WXk0eEN6QUpCZ05WQkFZVEFsVlRNQjRYRFRFNU1EVXhPREF4TXpJMU4xb1hEVEkwTURVeE5qQXhNekkxTjFvd1h6RWxNQ01HQTFVRUF3d2NaV05qTFhOdGNDMWljbTlyWlhJdGMybG5ibDlWUXpRdFVGSlBSREVVTUJJR0ExVUVDd3dMYVU5VElGTjVjM1JsYlhNeEV6QVJCZ05WQkFvTUNrRndjR3hsSUVsdVl5NHhDekFKQmdOVkJBWVRBbFZUTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFd2hWMzdldld4N0loajJqZGNKQ2hJWTNIc0wxdkxDZzloR0NWMlVyMHBVRWJnMElPMkJIelFINkRNeDhjVk1QMzZ6SWcxcnJWMU8vMGtvbUpQbndQRTZPQ0FoRXdnZ0lOTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVSS9KSnhFK1Q1TzhuNXNUMktHdy9vcnY5TGtzd1JRWUlLd1lCQlFVSEFRRUVPVEEzTURVR0NDc0dBUVVGQnpBQmhpbG9kSFJ3T2k4dmIyTnpjQzVoY0hCc1pTNWpiMjB2YjJOemNEQTBMV0Z3Y0d4bFlXbGpZVE13TWpDQ0FSMEdBMVVkSUFTQ0FSUXdnZ0VRTUlJQkRBWUpLb1pJaHZkalpBVUJNSUgrTUlIREJnZ3JCZ0VGQlFjQ0FqQ0J0Z3lCczFKbGJHbGhibU5sSUc5dUlIUm9hWE1nWTJWeWRHbG1hV05oZEdVZ1lua2dZVzU1SUhCaGNuUjVJR0Z6YzNWdFpYTWdZV05qWlhCMFlXNWpaU0J2WmlCMGFHVWdkR2hsYmlCaGNIQnNhV05oWW14bElITjBZVzVrWVhKa0lIUmxjbTF6SUdGdVpDQmpiMjVrYVhScGIyNXpJRzltSUhWelpTd2dZMlZ5ZEdsbWFXTmhkR1VnY0c5c2FXTjVJR0Z1WkNCalpYSjBhV1pwWTJGMGFXOXVJSEJ5WVdOMGFXTmxJSE4wWVhSbGJXVnVkSE11TURZR0NDc0dBUVVGQndJQkZpcG9kSFJ3T2k4dmQzZDNMbUZ3Y0d4bExtTnZiUzlqWlhKMGFXWnBZMkYwWldGMWRHaHZjbWwwZVM4d05BWURWUjBmQkMwd0t6QXBvQ2VnSllZamFIUjBjRG92TDJOeWJDNWhjSEJzWlM1amIyMHZZWEJ3YkdWaGFXTmhNeTVqY213d0hRWURWUjBPQkJZRUZKUlgyMi9WZElHR2lZbDJMMzVYaFFmbm0xZ2tNQTRHQTFVZER3RUIvd1FFQXdJSGdEQVBCZ2txaGtpRzkyTmtCaDBFQWdVQU1Bb0dDQ3FHU000OUJBTUNBMGtBTUVZQ0lRQytDVmNmNXg0ZWMxdFY1YStzdE1jdjYwUmZNQmhTSXNjbEVBSzJIcjF2VlFJaEFOR0xOUXBkMXQxdXNYUmdOYkVlc3M2SHo2UG1yMnk5ZzRDSkRjZ3MzYXBqTUlJQzdqQ0NBbldnQXdJQkFnSUlTVzB2dnpxWTJwY3dDZ1lJS29aSXpqMEVBd0l3WnpFYk1Ca0dBMVVFQXd3U1FYQndiR1VnVW05dmRDQkRRU0F0SUVjek1TWXdKQVlEVlFRTERCMUJjSEJzWlNCRFpYSjBhV1pwWTJGMGFXOXVJRUYxZEdodmNtbDBlVEVUTUJFR0ExVUVDZ3dLUVhCd2JHVWdTVzVqTGpFTE1Ba0dBMVVFQmhNQ1ZWTXdIaGNOTVRRd05UQTJNak0wTmpNd1doY05Namt3TlRBMk1qTTBOak13V2pCNk1TNHdMQVlEVlFRRERDVkJjSEJzWlNCQmNIQnNhV05oZEdsdmJpQkpiblJsWjNKaGRHbHZiaUJEUVNBdElFY3pNU1l3SkFZRFZRUUxEQjFCY0hCc1pTQkRaWEowYVdacFkyRjBhVzl1SUVGMWRHaHZjbWwwZVRFVE1CRUdBMVVFQ2d3S1FYQndiR1VnU1c1akxqRUxNQWtHQTFVRUJoTUNWVk13V1RBVEJnY3Foa2pPUFFJQkJnZ3Foa2pPUFFNQkJ3TkNBQVR3RnhHRUdkZGtoZFVhWGlXQkIzYm9nS0x2M251dVRlQ04vRXVUNFROVzFXWmJOYTRpMEpkMkRTSk9lN29JL1hZWHpvakxkcnRtY0w3STZDbUUvMVJGbzRIM01JSDBNRVlHQ0NzR0FRVUZCd0VCQkRvd09EQTJCZ2dyQmdFRkJRY3dBWVlxYUhSMGNEb3ZMMjlqYzNBdVlYQndiR1V1WTI5dEwyOWpjM0F3TkMxaGNIQnNaWEp2YjNSallXY3pNQjBHQTFVZERnUVdCQlFqOGtuRVQ1UGs3eWZteFBZb2JEK2l1LzB1U3pBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUI4R0ExVWRJd1FZTUJhQUZMdXczcUZZTTRpYXBJcVozcjY5NjYvYXl5U3JNRGNHQTFVZEh3UXdNQzR3TEtBcW9DaUdKbWgwZEhBNkx5OWpjbXd1WVhCd2JHVXVZMjl0TDJGd2NHeGxjbTl2ZEdOaFp6TXVZM0pzTUE0R0ExVWREd0VCL3dRRUF3SUJCakFRQmdvcWhraUc5Mk5rQmdJT0JBSUZBREFLQmdncWhrak9QUVFEQWdObkFEQmtBakE2ejNLRFVSYVpzWWI3TmNOV3ltSy85QmZ0MlE5MVRhS092dkdjZ1Y1Q3Q0bjRtUGViV1orWTFVRU5qNTNwd3Y0Q01ESXQxVVFoc0tNRmQyeGQ4emc3a0dmOUYzd3NJVzJXVDhaeWFZSVNiMVQ0ZW4wYm1jdWJDWWtoWVFhWkR3bVNIUUFBTVlJQmpEQ0NBWWdDQVFFd2dZWXdlakV1TUN3R0ExVUVBd3dsUVhCd2JHVWdRWEJ3YkdsallYUnBiMjRnU1c1MFpXZHlZWFJwYjI0Z1EwRWdMU0JITXpFbU1DUUdBMVVFQ3d3ZFFYQndiR1VnUTJWeWRHbG1hV05oZEdsdmJpQkJkWFJvYjNKcGRIa3hFekFSQmdOVkJBb01Da0Z3Y0d4bElFbHVZeTR4Q3pBSkJnTlZCQVlUQWxWVEFnaE1NRUZKVVoxVU5qQU5CZ2xnaGtnQlpRTUVBZ0VGQUtDQmxUQVlCZ2txaGtpRzl3MEJDUU14Q3dZSktvWklodmNOQVFjQk1Cd0dDU3FHU0liM0RRRUpCVEVQRncweU1EQTRNRFF4TlRRME5UWmFNQ29HQ1NxR1NJYjNEUUVKTkRFZE1Cc3dEUVlKWUlaSUFXVURCQUlCQlFDaENnWUlLb1pJemowRUF3SXdMd1lKS29aSWh2Y05BUWtFTVNJRUlLdjk4WVFzenNoSyswallBWDhyNk1jc2hyMC83cUM2NUx6blZya2Zkc0VPTUFvR0NDcUdTTTQ5QkFNQ0JFY3dSUUlnU29kV01rV09LRlkxdHYweUx1MjV0UUZ3SXVnZ3J6VlZjNHlnZzBsTE1OOENJUURrOWtKbUQ4QmdNSm1nSEhwbmgwMUpXSGpwZ3ZJVnlMTHVCR0dVTmoxanhBQUFBQUFBQUE9PSIsImhlYWRlciI6eyJlcGhlbWVyYWxQdWJsaWNLZXkiOiJNRmt3RXdZSEtvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVYWlJEOG5XOG1iVkx3K2ZBdmpDV1VqalVRSEhxcWk2RmQyVnVERDlMWnBzSlg3YUVyTjFCM0hFbTdHSFBtcjRxM29KNmYreVFPVlJ2dFUvaHhYTXJjdz09IiwicHVibGljS2V5SGFzaCI6ImVSQmhYTVA4UElUbWJNTGtoeVZNdHo0Uk1iNlVtS0VDOFNwRXI1NEFaZEU9IiwidHJhbnNhY3Rpb25JZCI6ImIwMjdjNGZjNDQwNjQxYTU4ZTE0OTZmMjNjNmY5YWJmYmRhYWM1MjkyOTg4YTU1ODYxZjc3YWQ2MjE2YzM0YWQifX0=`

	reply := `{
    "Success": true,
    "OrderId": "ApplePayTestOrder-001",
    "Amount": 300,
    "ErrCode": ""
}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", APILink+"/Block",
		httpmock.NewStringResponder(200, reply))

	trans := New("SendtipsTestTerminal", "ApplePayTestOrder-001")
	ctx := context.Background()

	err := trans.ApplePayBlock(ctx, 300, payload)
	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.ApplePayToken != base64 {
		t.Error("Base64 encoding fail", trans.ApplePayToken)
	}

}

func TestApplePay(t *testing.T) {

	payload := []byte(strings.TrimSpace(` {"version":"EC_v1","data":"TGuomSxhVblEfMiJdbVe9JeN9K52znlr1Q3n2OsMwz4luSJAZbumg/ZFgBYyty2Dhst4GQinLffz11vO3VH/kdRShRlM74esNW3cz0hRC0Lfg0to2ddqfxzwun4euELpo0Q3CZeBbfC7Hq5WZtnApTUEvpOrshCaJIgj5n7LHcgu0BLpyeENgPuEHFamejzAyRW4vQ2gw4dg1MsHN0WxC5OQ7vLvRyIfXLifIwQb6zSOHGRh3izuAV8V5Gk0CbYq6x+DnKa9MRSm9efF8Uv1yZFyxI13BjD6yfFJtRa7LLOXnn6I+8OekKxDcNxxPmhRK6JW/5KXDrLa4/XgFGr+cxb/jG9F7uyEgXLNveQsQBoIBlcLEqQYSl7UD783XbUYmUwWhwLBh6196hllOJqVPM3J8Q060+5V4nLWnmtcWuGO3A==","signature":"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","header":{"ephemeralPublicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXZRD8nW8mbVLw+fAvjCWUjjUQHHqqi6Fd2VuDD9LZpsJX7aErN1B3HEm7GHPmr4q3oJ6f+yQOVRvtU/hxXMrcw==","publicKeyHash":"eRBhXMP8PITmbMLkhyVMtz4RMb6UmKEC8SpEr54AZdE=","transactionId":"b027c4fc440641a58e1496f23c6f9abfbdaac5292988a55861f77ad6216c34ad"}}`))

	var base64 string = `eyJ2ZXJzaW9uIjoiRUNfdjEiLCJkYXRhIjoiVEd1b21TeGhWYmxFZk1pSmRiVmU5SmVOOUs1MnpubHIxUTNuMk9zTXd6NGx1U0pBWmJ1bWcvWkZnQll5dHkyRGhzdDRHUWluTGZmejExdk8zVkgva2RSU2hSbE03NGVzTlczY3owaFJDMExmZzB0bzJkZHFmeHp3dW40ZXVFTHBvMFEzQ1plQmJmQzdIcTVXWnRuQXBUVUV2cE9yc2hDYUpJZ2o1bjdMSGNndTBCTHB5ZUVOZ1B1RUhGYW1lanpBeVJXNHZRMmd3NGRnMU1zSE4wV3hDNU9RN3ZMdlJ5SWZYTGlmSXdRYjZ6U09IR1JoM2l6dUFWOFY1R2swQ2JZcTZ4K0RuS2E5TVJTbTllZkY4VXYxeVpGeXhJMTNCakQ2eWZGSnRSYTdMTE9Ybm42SSs4T2VrS3hEY054eFBtaFJLNkpXLzVLWERyTGE0L1hnRkdyK2N4Yi9qRzlGN3V5RWdYTE52ZVFzUUJvSUJsY0xFcVFZU2w3VUQ3ODNYYlVZbVV3V2h3TEJoNjE5NmhsbE9KcVZQTTNKOFEwNjArNVY0bkxXbm10Y1d1R08zQT09Iiwic2lnbmF0dXJlIjoiTUlBR0NTcUdTSWIzRFFFSEFxQ0FNSUFDQVFFeER6QU5CZ2xnaGtnQlpRTUVBZ0VGQURDQUJna3Foa2lHOXcwQkJ3RUFBS0NBTUlJRDR6Q0NBNGlnQXdJQkFnSUlUREJCU1ZHZFZEWXdDZ1lJS29aSXpqMEVBd0l3ZWpFdU1Dd0dBMVVFQXd3bFFYQndiR1VnUVhCd2JHbGpZWFJwYjI0Z1NXNTBaV2R5WVhScGIyNGdRMEVnTFNCSE16RW1NQ1FHQTFVRUN3d2RRWEJ3YkdVZ1EyVnlkR2xtYVdOaGRHbHZiaUJCZFhSb2IzSnBkSGt4RXpBUkJnTlZCQW9NQ2tGd2NHeGxJRWx1WXk0eEN6QUpCZ05WQkFZVEFsVlRNQjRYRFRFNU1EVXhPREF4TXpJMU4xb1hEVEkwTURVeE5qQXhNekkxTjFvd1h6RWxNQ01HQTFVRUF3d2NaV05qTFhOdGNDMWljbTlyWlhJdGMybG5ibDlWUXpRdFVGSlBSREVVTUJJR0ExVUVDd3dMYVU5VElGTjVjM1JsYlhNeEV6QVJCZ05WQkFvTUNrRndjR3hsSUVsdVl5NHhDekFKQmdOVkJBWVRBbFZUTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFd2hWMzdldld4N0loajJqZGNKQ2hJWTNIc0wxdkxDZzloR0NWMlVyMHBVRWJnMElPMkJIelFINkRNeDhjVk1QMzZ6SWcxcnJWMU8vMGtvbUpQbndQRTZPQ0FoRXdnZ0lOTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVSS9KSnhFK1Q1TzhuNXNUMktHdy9vcnY5TGtzd1JRWUlLd1lCQlFVSEFRRUVPVEEzTURVR0NDc0dBUVVGQnpBQmhpbG9kSFJ3T2k4dmIyTnpjQzVoY0hCc1pTNWpiMjB2YjJOemNEQTBMV0Z3Y0d4bFlXbGpZVE13TWpDQ0FSMEdBMVVkSUFTQ0FSUXdnZ0VRTUlJQkRBWUpLb1pJaHZkalpBVUJNSUgrTUlIREJnZ3JCZ0VGQlFjQ0FqQ0J0Z3lCczFKbGJHbGhibU5sSUc5dUlIUm9hWE1nWTJWeWRHbG1hV05oZEdVZ1lua2dZVzU1SUhCaGNuUjVJR0Z6YzNWdFpYTWdZV05qWlhCMFlXNWpaU0J2WmlCMGFHVWdkR2hsYmlCaGNIQnNhV05oWW14bElITjBZVzVrWVhKa0lIUmxjbTF6SUdGdVpDQmpiMjVrYVhScGIyNXpJRzltSUhWelpTd2dZMlZ5ZEdsbWFXTmhkR1VnY0c5c2FXTjVJR0Z1WkNCalpYSjBhV1pwWTJGMGFXOXVJSEJ5WVdOMGFXTmxJSE4wWVhSbGJXVnVkSE11TURZR0NDc0dBUVVGQndJQkZpcG9kSFJ3T2k4dmQzZDNMbUZ3Y0d4bExtTnZiUzlqWlhKMGFXWnBZMkYwWldGMWRHaHZjbWwwZVM4d05BWURWUjBmQkMwd0t6QXBvQ2VnSllZamFIUjBjRG92TDJOeWJDNWhjSEJzWlM1amIyMHZZWEJ3YkdWaGFXTmhNeTVqY213d0hRWURWUjBPQkJZRUZKUlgyMi9WZElHR2lZbDJMMzVYaFFmbm0xZ2tNQTRHQTFVZER3RUIvd1FFQXdJSGdEQVBCZ2txaGtpRzkyTmtCaDBFQWdVQU1Bb0dDQ3FHU000OUJBTUNBMGtBTUVZQ0lRQytDVmNmNXg0ZWMxdFY1YStzdE1jdjYwUmZNQmhTSXNjbEVBSzJIcjF2VlFJaEFOR0xOUXBkMXQxdXNYUmdOYkVlc3M2SHo2UG1yMnk5ZzRDSkRjZ3MzYXBqTUlJQzdqQ0NBbldnQXdJQkFnSUlTVzB2dnpxWTJwY3dDZ1lJS29aSXpqMEVBd0l3WnpFYk1Ca0dBMVVFQXd3U1FYQndiR1VnVW05dmRDQkRRU0F0SUVjek1TWXdKQVlEVlFRTERCMUJjSEJzWlNCRFpYSjBhV1pwWTJGMGFXOXVJRUYxZEdodmNtbDBlVEVUTUJFR0ExVUVDZ3dLUVhCd2JHVWdTVzVqTGpFTE1Ba0dBMVVFQmhNQ1ZWTXdIaGNOTVRRd05UQTJNak0wTmpNd1doY05Namt3TlRBMk1qTTBOak13V2pCNk1TNHdMQVlEVlFRRERDVkJjSEJzWlNCQmNIQnNhV05oZEdsdmJpQkpiblJsWjNKaGRHbHZiaUJEUVNBdElFY3pNU1l3SkFZRFZRUUxEQjFCY0hCc1pTQkRaWEowYVdacFkyRjBhVzl1SUVGMWRHaHZjbWwwZVRFVE1CRUdBMVVFQ2d3S1FYQndiR1VnU1c1akxqRUxNQWtHQTFVRUJoTUNWVk13V1RBVEJnY3Foa2pPUFFJQkJnZ3Foa2pPUFFNQkJ3TkNBQVR3RnhHRUdkZGtoZFVhWGlXQkIzYm9nS0x2M251dVRlQ04vRXVUNFROVzFXWmJOYTRpMEpkMkRTSk9lN29JL1hZWHpvakxkcnRtY0w3STZDbUUvMVJGbzRIM01JSDBNRVlHQ0NzR0FRVUZCd0VCQkRvd09EQTJCZ2dyQmdFRkJRY3dBWVlxYUhSMGNEb3ZMMjlqYzNBdVlYQndiR1V1WTI5dEwyOWpjM0F3TkMxaGNIQnNaWEp2YjNSallXY3pNQjBHQTFVZERnUVdCQlFqOGtuRVQ1UGs3eWZteFBZb2JEK2l1LzB1U3pBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUI4R0ExVWRJd1FZTUJhQUZMdXczcUZZTTRpYXBJcVozcjY5NjYvYXl5U3JNRGNHQTFVZEh3UXdNQzR3TEtBcW9DaUdKbWgwZEhBNkx5OWpjbXd1WVhCd2JHVXVZMjl0TDJGd2NHeGxjbTl2ZEdOaFp6TXVZM0pzTUE0R0ExVWREd0VCL3dRRUF3SUJCakFRQmdvcWhraUc5Mk5rQmdJT0JBSUZBREFLQmdncWhrak9QUVFEQWdObkFEQmtBakE2ejNLRFVSYVpzWWI3TmNOV3ltSy85QmZ0MlE5MVRhS092dkdjZ1Y1Q3Q0bjRtUGViV1orWTFVRU5qNTNwd3Y0Q01ESXQxVVFoc0tNRmQyeGQ4emc3a0dmOUYzd3NJVzJXVDhaeWFZSVNiMVQ0ZW4wYm1jdWJDWWtoWVFhWkR3bVNIUUFBTVlJQmpEQ0NBWWdDQVFFd2dZWXdlakV1TUN3R0ExVUVBd3dsUVhCd2JHVWdRWEJ3YkdsallYUnBiMjRnU1c1MFpXZHlZWFJwYjI0Z1EwRWdMU0JITXpFbU1DUUdBMVVFQ3d3ZFFYQndiR1VnUTJWeWRHbG1hV05oZEdsdmJpQkJkWFJvYjNKcGRIa3hFekFSQmdOVkJBb01Da0Z3Y0d4bElFbHVZeTR4Q3pBSkJnTlZCQVlUQWxWVEFnaE1NRUZKVVoxVU5qQU5CZ2xnaGtnQlpRTUVBZ0VGQUtDQmxUQVlCZ2txaGtpRzl3MEJDUU14Q3dZSktvWklodmNOQVFjQk1Cd0dDU3FHU0liM0RRRUpCVEVQRncweU1EQTRNRFF4TlRRME5UWmFNQ29HQ1NxR1NJYjNEUUVKTkRFZE1Cc3dEUVlKWUlaSUFXVURCQUlCQlFDaENnWUlLb1pJemowRUF3SXdMd1lKS29aSWh2Y05BUWtFTVNJRUlLdjk4WVFzenNoSyswallBWDhyNk1jc2hyMC83cUM2NUx6blZya2Zkc0VPTUFvR0NDcUdTTTQ5QkFNQ0JFY3dSUUlnU29kV01rV09LRlkxdHYweUx1MjV0UUZ3SXVnZ3J6VlZjNHlnZzBsTE1OOENJUURrOWtKbUQ4QmdNSm1nSEhwbmgwMUpXSGpwZ3ZJVnlMTHVCR0dVTmoxanhBQUFBQUFBQUE9PSIsImhlYWRlciI6eyJlcGhlbWVyYWxQdWJsaWNLZXkiOiJNRmt3RXdZSEtvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVYWlJEOG5XOG1iVkx3K2ZBdmpDV1VqalVRSEhxcWk2RmQyVnVERDlMWnBzSlg3YUVyTjFCM0hFbTdHSFBtcjRxM29KNmYreVFPVlJ2dFUvaHhYTXJjdz09IiwicHVibGljS2V5SGFzaCI6ImVSQmhYTVA4UElUbWJNTGtoeVZNdHo0Uk1iNlVtS0VDOFNwRXI1NEFaZEU9IiwidHJhbnNhY3Rpb25JZCI6ImIwMjdjNGZjNDQwNjQxYTU4ZTE0OTZmMjNjNmY5YWJmYmRhYWM1MjkyOTg4YTU1ODYxZjc3YWQ2MjE2YzM0YWQifX0=`

	reply := `{
    "Success": true,
    "OrderId": "ApplePayTestOrder-001",
    "Amount": 300,
    "ErrCode": ""
}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", APILink+"/Pay",
		httpmock.NewStringResponder(200, reply))

	trans := New("SendtipsTestTerminal", "ApplePayTestOrder-001")
	ctx := context.Background()

	err := trans.ApplePay(ctx, 300, payload)
	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.ApplePayToken != base64 {
		t.Error("Base64 encoding fail", trans.ApplePayToken)
	}

}

func ExampleApplePay() {

	payload := []byte(strings.TrimSpace(` {"version":"EC_v1","data":"TGuomSxhVblEfMiJdbVe9JeN9K52znlr1Q3n2OsMwz4luSJAZbumg/ZFgBYyty2Dhst4GQinLffz11vO3VH/kdRShRlM74esNW3cz0hRC0Lfg0to2ddqfxzwun4euELpo0Q3CZeBbfC7Hq5WZtnApTUEvpOrshCaJIgj5n7LHcgu0BLpyeENgPuEHFamejzAyRW4vQ2gw4dg1MsHN0WxC5OQ7vLvRyIfXLifIwQb6zSOHGRh3izuAV8V5Gk0CbYq6x+DnKa9MRSm9efF8Uv1yZFyxI13BjD6yfFJtRa7LLOXnn6I+8OekKxDcNxxPmhRK6JW/5KXDrLa4/XgFGr+cxb/jG9F7uyEgXLNveQsQBoIBlcLEqQYSl7UD783XbUYmUwWhwLBh6196hllOJqVPM3J8Q060+5V4nLWnmtcWuGO3A==","signature":"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","header":{"ephemeralPublicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXZRD8nW8mbVLw+fAvjCWUjjUQHHqqi6Fd2VuDD9LZpsJX7aErN1B3HEm7GHPmr4q3oJ6f+yQOVRvtU/hxXMrcw==","publicKeyHash":"eRBhXMP8PITmbMLkhyVMtz4RMb6UmKEC8SpEr54AZdE=","transactionId":"b027c4fc440641a58e1496f23c6f9abfbdaac5292988a55861f77ad6216c34ad"}}`))

	// check themap hostname env is set, otherwise use default host
	apihost, ok := os.LookupEnv("THEMAPAPIHOST")
	if ok {
		APILink = apihost
	}

	pay := New(os.Getenv("THEMAPTERMID"), "ApplePayTestOrder-001")

	err := pay.ApplePay(context.TODO(), 300, payload)
	if err != nil {
		fmt.Errorf("Error occurred: %v", err)
	}

	fmt.Printf("%v", pay.Success)
	// Output: false
}
