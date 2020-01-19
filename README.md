## HOTP, TOTP (short-lived OTP values, which are desirable for enhanced security)

#### Usage

```go
package main

import (
	"fmt"
	"github.com/kbukum/hotp"
)

func main() {
	totp := hotp.TOTP{
		OTP: hotp.OTP{
			SharedSecret: "Secret Key",
			Digits:       8,
			Crypto:hotp.SHA1,
		},
		StartTime: 0,
		TimeStep:  30,
	}
	fmt.Println(totp.Password())
}

```
   
#### REFERENCES
* RFC4226 [https://tools.ietf.org/html/rfc4226](https://tools.ietf.org/html/rfc4226)
* RFC6238 [https://tools.ietf.org/html/rfc6238](https://tools.ietf.org/html/rfc6238)


