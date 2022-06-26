package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAmOIbW29is4NJJk7zIcHJaSN4TCd13nxfSml+Op3QMdNDl6ai
L5rSv3UfqUBV51UmJOw0nBZsI1BUJK8NG7N/rIYs8ap8ERJwSeaKEwPA1/yD1VlZ
NC/b8BKC9+SqWnLfix8ngLfNcluaOoTdDdnclz3sGRMsjtTNmGtuxvyNtbU7nvUx
bK2O5ul7tIFwJ0N7FK6F6ugLyL+/AGY1T9FE/6a2lFEbDUwTzhhaCh3k4HgVTTl1
itIKxRZfGPoFhYJSFeD5N/VFtm45ikN9X2rDOe9AI5Ky1fRPFr/wNWCXWdJKlY1o
V1r6O/K5DZYypSoFaJ91QqNvOWcIBF/F5/Da5QIDAQABAoIBAGb9+5wPX5i7Wu4Z
xaT6Hatcn8/9zWQCuVcx1j26zuSiBCkHsr8LN+qPBrNIihZ7wGSjU5XLbTIlDWph
Gx2MQPiCs4hiZ45As7C1RFMm9iULEe0KchO8IhiK//vr6e9g78OTS1Nlf1wI5c+C
bkDEgdzJdGI4Do5yQcbqqQgYX4cG8QknG2+mkwpTKfJlLLljWyitgaiBDZbPh0+F
Z/Wk1sNSZ4XHqfg1pujwug4gL+YYnQH6WQ2XfbSKylLlNnl12pbSJGF4/Zo2Hg4p
2Dk7stpyyRkrozxu3WhAVea0Gyxik3egVns8KkqQo+5A4thW5Fib7R2PCpMbFYUc
qu5etsECgYEA5bbsoUu4tVWr/uL4y44O6iQJX8lDEVQYZNh/vP7DUfl51iqxyZfL
T416DC5U09WKLUQ5t6eH5heEjVgFtGsgJJ2uivvwV7K3/+dQnwuNpDWwBmQ9gR2s
SkNyi9BiEkkWpwpRQA8Ftopjs7vmZqkhv7PCMtRRlyKwn2h20MUII60CgYEAqmCL
/4AS1dP0n4TNeanqeIpDsue/u0/N2lpLwnvdh7q1PSPHg6abLX1+A0Y+LsxrYK9k
3yjVKJYAgY1IFZiIKyVwEodp9IBy7VzPySsxskK7l7SFMO5MfAvB1cVLU4hPL6fV
enjE7tvJ9u8QAmPEeScmQ1yWVY3QXswGLmsuuxkCgYAzLLU6mavkedl/RwE2F8eq
0axk8mlGiv2EOdb7O6Y3tOQ3mftRdceNSW9PF2M+bewCeZiCGYhk8ghNlLZwowze
G2KvA8FfSClFkTqcs+4yDuPQCLTK9tlTEgOKsjmm6TFqtRm0s6QKLnpXqByD8lna
Yyl4OWSTzt6aJKOTjtFpYQKBgQClHjBPi7W1WVcoCbKBCHVeINF/Xy3nwG3GAeCO
OTD2y7G98SD4q8yUB5zKW5cED8S4zQK7a305ejY/V8bWdx7wgbXdnzLGbH31IA+X
7K6bDiVz3tV+GFQzm8lc/XoFGIN1sfgoW0awHn3bPNCNIFdW+uQQQHjJrUiVtrD+
541AcQKBgAan598sudDfmw64uIuaxfeIuMRmpwmlPzJq/Q1XjGdxeM/FpogYEHM1
Rc2UMkTb8aY4sjHEzmMWdANdT1KWvhseDgbv99vrx631arGFp7ymzxmqShVTkNGs
2XAKNXc6z4O6cbdK7/OfoK5gy5ZyUR6pPZ2KXYE0Dx2QkDZbXquU
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parase private key: %v", err)
	}
	g := NewJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516246222, 0)
	}
	tkn, err := g.GenerateToken("123123123", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token %v", err)
	}
	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNTM0MjIsImlhdCI6MTY1NTkwNzE2MywiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMTIzMTIzMTIzIn0.fA6Su5ah4My0AD-Ir88QzKXTrx5bvONmYghjnr9DqdDffdFp5Nb3E2ISfh76XiximTjWhWebrngFc-GXXBw0tD2WJGV3Js1-8uSmM9Q5O0bBK_S7Oqp-6hTe4lg_ri5mTiSFxE2P3B3UhumzFlA5IVNJKzQXr9Tmam6cxLiPBYCHkuSs4tOfVBXK_SyUwQo4d_ZJJWoBqxJ-65_Y4BP-xh0DeW4qD8Q-puzw4I6M6gVrwcPlgJOrDwnHSQ57gbSeBwQGSycWrA64NL1acjT_ecqocAmDDFLDAFEHy6wFvmfCrv5TGYwUFr3-UajxfxdV6RpD5c72m512Kix37tFdQw"
	if tkn != want {
		t.Errorf("wrong token gennerated. want %q; got: %q", want, tkn)
	}
}
