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
			SharedSecret: "KAMILBUKUM@GMAIL.COMHENNGECHALLENGE003",
			Digits:       10,
			Crypto:hotp.SHA512,
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


