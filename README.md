<h1>mvola-go *** Work In Progress ***</h1>
<p>
  <a href="#" target="_blank">
    <img alt="License: BSD" src="https://img.shields.io/badge/License-BSD-yellow.svg" />
  </a>
  <a href="https://twitter.com/tsiry_sndr" target="_blank">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" />
  </a>
</p>

[MVola](https://www.mvola.mg/devportal) Go client library.

## Install

```sh
  go get -u github.com/mvola/mvola-go
```

## Usage

```go
import (
	"fmt"
	"log"
	"os"

	mvola "github.com/tsirysndr/mvola-go"
)

func main() {
	var (
		consumerKey    = os.Getenv("CONSUMER_KEY")
		consumerSecret = os.Getenv("CONSUMER_SECRET")
	)
	client := mvola.NewClient(mvola.SANDBOX_URL)
	res, err := client.Auth.GenerateToken(consumerKey, consumerSecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

```

## Author

👤 **Tsiry Sandratraina <tsiry.sndr@aol.com>**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
