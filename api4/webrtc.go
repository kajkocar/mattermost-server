// Copyright (c) 2017 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/mattermost/mattermost-server/utils"
)

func (api *API) InitWebrtc() {
	l4g.Debug(utils.T("api.webrtc.init.debug"))

	api.BaseRoutes.Webrtc.Handle("/token", api.ApiSessionRequired(webrtcToken)).Methods("GET")
}

func webrtcToken(c *Context, w http.ResponseWriter, r *http.Request) {
	result, err := c.App.GetWebrtcInfoForSession(c.Session.Id)

	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(result.ToJson()))
}
