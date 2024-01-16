package main

/*            API:
* Import Canvas Users to UnlockEd
* canvasservice import -u <url> -k <api_key>
*
* Match Canvas Users to UnlockEd Users
* canvasservice match -u <url> -k <api_key>
*
* OR: ENV VARS
*
* CANVAS_URL
* UNLOCKED_URL
*
 */

type Args struct {
	BaseUrl string
	Key     string
	Import  bool
	Match   bool
}

func main() {
}
