package examples

import (
	"github.com/yedf/dtm/dtmcli"
	"github.com/yedf/dtm/dtmcli/dtmimp"
)

func init() {
	addSample("msg", func() string {
		dtmimp.Logf("a busi transaction begin")
		req := &TransReq{Amount: 30}
		msg := dtmcli.NewMsg(DtmServer, dtmcli.MustGenGid(DtmServer)).
			Add(Busi+"/TransOut", req).
			Add(Busi+"/TransIn", req)
		err := msg.Prepare(Busi + "/query")
		dtmimp.FatalIfError(err)
		dtmimp.Logf("busi trans submit")
		err = msg.Submit()
		dtmimp.FatalIfError(err)
		return msg.Gid
	})
}
