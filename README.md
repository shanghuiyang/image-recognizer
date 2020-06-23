# image-recognizer
image-recognizer uses the Baidu API to implement image recognition

## Usage
Suppose that you have had the `AppKey` and `SecretKey` from the [Baidu AI Platform](https://ai.baidu.com). Replace `your_app_key` and `your_secret_key` with yours in main.go.
```go
// ...
const (
	// replace your_app_key with your app key
	appKey    = "your_app_key"
	// replace your_secret_key with your secret key
	secretKey = "your_secret_key"
)
// ...
```

build and test.
```shell
# cd to this project directory
$ go build
$ ./image-recognizer images/orenge.jpg
```
