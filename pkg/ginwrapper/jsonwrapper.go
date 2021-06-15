package ginwrapper

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Req map[string]interface{}
type JsonHandler func(reqBody Req) (resp interface{}, code int)

func JsonWrapper(h JsonHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqB := make(Req)
		data, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Bad input!"})
			return
		}
		if e := json.Unmarshal(data, &reqB); e != nil {
			ctx.JSON(400, gin.H{"error": e.Error()})
			return
		}
		for k, v := range ctx.Request.URL.Query() {
			if len(v) == 1 {
				s := v[0]
				reqB[k] = s
			} else {
				reqB[k] = v
			}
		}
		resp, code := h(reqB)
		err, iserr := resp.(error)
		if iserr {
			ctx.JSON(code, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(code, resp)
	}
}
