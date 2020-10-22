// You can edit this code!
// Click here and start typing.
package resserver_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/liyustar/nuts/resserver"
)

var context = resserver.Context {
	Conf: resserver.Config {
		ClientId: "89f849642e8967122a2c",
		ClientSecret: "7f649d478d5b815fa4849c6a9da5c1e72943647b",
		RedirectUrl: "http://localhost:9090/oauth/redirect",
	},
}

func TestResServer(t *testing.T) {
	fmt.Println("start done.")
	http.HandleFunc("/", context.Hello)
	http.HandleFunc("/oauth/redirect", context.Oauth)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("监听失败，错误信息为：", err)
		return
	}
}
