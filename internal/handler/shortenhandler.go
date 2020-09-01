package handler

import (
	"net/http"

	"hello-shorturl-go/internal/logic"
	"hello-shorturl-go/internal/svc"
	"hello-shorturl-go/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func shortenHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShortenLogic(r.Context(), ctx)
		resp, err := l.Shorten(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
