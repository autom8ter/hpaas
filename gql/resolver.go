package gql

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/graphikDB/generic"
	"github.com/graphikDB/kdeploy/app"
	"github.com/graphikDB/kdeploy/gen/gql/go/generated"
	"github.com/graphikDB/kdeploy/helpers"
	"github.com/graphikDB/kdeploy/logger"
	"github.com/graphikDB/kubego"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/metadata"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client      *app.Manager
	cors        *cors.Cors
	store       *generic.Cache
	config      *oauth2.Config
	tokenCookie string
	stateCookie string
	logger      *logger.Logger
	jwtCache    *generic.Cache
	userInfo    string
}

func NewResolver(client *kubego.Client, cors *cors.Cors, config *oauth2.Config, logger *logger.Logger, userInfoEndpoint string) *Resolver {
	return &Resolver{
		client:      app.New(client),
		cors:        cors,
		config:      config,
		tokenCookie: "graphik-playground-token",
		stateCookie: "graphik-playground-state",
		store:       generic.NewCache(5 * time.Minute),
		logger:      logger,
		jwtCache:    generic.NewCache(1 * time.Minute),
		userInfo:    userInfoEndpoint,
	}
}

func (r *Resolver) QueryHandler() http.Handler {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers:  r,
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			auth := initPayload.Authorization()
			ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", auth)
			return ctx, nil
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(&apollotracing.Tracer{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	return r.cors.Handler(r.authMiddleware(srv))
}

func (r *Resolver) authMiddleware(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var token *oauth2.Token
		if r.store != nil && r.config != nil && r.config.ClientID != "" {
			token, _ = r.getToken(req)
			if token != nil && req.Header.Get("Authorization") == "" {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
			}
		}
		var authHeader = req.Header.Get("Authorization")
		tokenHash := helpers.Hash([]byte(authHeader))
		if val, ok := r.jwtCache.Get(tokenHash); ok {
			payload := val.(map[string]interface{})
			ctx, err := r.checkRequest(req, payload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}
		userinfoReq, err := http.NewRequest(http.MethodGet, r.userInfo, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get userinfo: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		userinfoReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := http.DefaultClient.Do(userinfoReq)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get userinfo: %s", err.Error()), http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			http.Error(w, fmt.Sprintf("failed to get userinfo: %v", resp.StatusCode), http.StatusUnauthorized)
			return
		}
		bits, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get userinfo: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		payload := map[string]interface{}{}
		if err := json.Unmarshal(bits, &payload); err != nil {
			http.Error(w, fmt.Sprintf("failed to get userinfo: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		r.jwtCache.Set(tokenHash, payload, 1*time.Hour)
		ctx, err = r.checkRequest(req, payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		handler.ServeHTTP(w, req.WithContext(ctx))
	}
}

func (r *Resolver) Playground() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if r.config == nil || r.config.ClientID == "" {
			http.Error(w, "playground disabled", http.StatusNotFound)
			return
		}
		authToken, err := r.getToken(req)
		if err != nil {
			r.logger.Error("playground: failed to get session - redirecting", zap.Error(err))
			r.redirectLogin(w, req)
			return
		}
		if authToken == nil {
			r.redirectLogin(w, req)
			return
		}
		if !authToken.Valid() {
			r.redirectLogin(w, req)
			return
		}
		w.Header().Add("Content-Type", "text/html")
		var playground = template.Must(template.New("playground").Parse(`<!DOCTYPE html>
<html>

<head>
  <meta charset=utf-8/>
  <meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
  <title>Graphik Playground</title>
  <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/css/index.css" />
  <link rel="shortcut icon" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/favicon.png" />
  <script src="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/js/middleware.js"></script>
</head>

<body>
  <div id="root">
    <style>
      body {
        background-color: rgb(23, 42, 58);
        font-family: Open Sans, sans-serif;
        height: 90vh;
      }

      #root {
        height: 100%;
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
      }

      .loading {
        font-size: 32px;
        font-weight: 200;
        color: rgba(255, 255, 255, .6);
        margin-left: 20px;
      }

      img {
        width: 78px;
        height: 78px;
      }

      .title {
        font-weight: 400;
      }
    </style>
    <img src='//cdn.jsdelivr.net/npm/graphql-playground-react/build/logo.png' alt=''>
    <div class="loading"> Loading
      <span class="title">Graphik Playground</span>
    </div>
  </div>
  <script>window.addEventListener('load', function (event) {
 		const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
      GraphQLPlayground.init(document.getElementById('root'), {
		endpoint: location.protocol + '//' + location.host,
		subscriptionsEndpoint: wsProto + '//' + location.host,
		shareEnabled: true,
		settings: {
			'request.credentials': 'same-origin',
			'prettier.useTabs': true
		}
      })
    })</script>
</body>

</html>
`))

		playground.Execute(w, map[string]string{})
	}
}

func (r *Resolver) redirectLogin(w http.ResponseWriter, req *http.Request) {
	state := helpers.Hash([]byte(fmt.Sprint(rand.Int())))
	r.setState(w, state)
	redirect := r.config.AuthCodeURL(state)
	http.Redirect(w, req, redirect, http.StatusTemporaryRedirect)
}

func (r *Resolver) PlaygroundCallback(playgroundRedirect string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if r.config == nil || r.config.ClientID == "" {
			http.Error(w, "playground disabled", http.StatusNotFound)
			return
		}
		code := req.URL.Query().Get("code")
		state := req.URL.Query().Get("state")
		if code == "" {
			r.logger.Error("playground: empty authorization code - redirecting")
			r.redirectLogin(w, req)
			return
		}
		if state == "" {
			r.logger.Error("playground: empty authorization state - redirecting")
			r.redirectLogin(w, req)
			return
		}

		stateVal, err := r.getState(req)
		if err != nil {
			r.logger.Error("playground: failed to get session state - redirecting", zap.Error(err))
			r.redirectLogin(w, req)
			return
		}
		if stateVal != state {
			r.logger.Error("playground: session state mismatch - redirecting")
			r.redirectLogin(w, req)
			return
		}
		token, err := r.config.Exchange(req.Context(), code)
		if err != nil {
			r.logger.Error("playground: failed to exchange authorization code - redirecting", zap.Error(err))
			r.redirectLogin(w, req)
			return
		}
		r.setToken(w, req, token)
		http.Redirect(w, req, playgroundRedirect, http.StatusTemporaryRedirect)
	}
}

func (r *Resolver) refreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	return r.config.TokenSource(oauth2.NoContext, token).Token()
}

func (r *Resolver) getToken(req *http.Request) (*oauth2.Token, error) {
	cookie, err := req.Cookie(r.tokenCookie)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get cookie: %s", r.tokenCookie)
	}
	val, ok := r.store.Get(cookie.Value)
	if !ok || val == nil {
		return nil, ErrNoTokenSession
	}
	return r.refreshToken(val.(*oauth2.Token))
}

func (r *Resolver) setToken(w http.ResponseWriter, req *http.Request, token *oauth2.Token) {
	id := helpers.Hash([]byte(fmt.Sprint(rand.Int())))
	r.store.Set(id, token, 1*time.Hour)
	cookie := &http.Cookie{
		Name:    r.tokenCookie,
		Value:   id,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, cookie)
}

func (r *Resolver) getState(req *http.Request) (string, error) {
	cookie, err := req.Cookie(r.stateCookie)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get cookie: %s", r.stateCookie)
	}
	val, ok := r.store.Get(cookie.Value)
	if !ok || val == nil {
		return "", ErrNoStateSession
	}
	return val.(string), nil
}

func (r *Resolver) setState(w http.ResponseWriter, state string) {
	id := helpers.Hash([]byte(fmt.Sprint(rand.Int())))
	r.store.Set(id, state, 5*time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    r.stateCookie,
		Value:   id,
		Expires: time.Now().Add(5 * time.Minute),
		Path:    "/",
	})
}

func (r *Resolver) checkRequest(req *http.Request, userData map[string]interface{}) (context.Context, error) {
	ctx := req.Context()
	ctx = context.WithValue(req.Context(), userInfo, userData)
	return ctx, nil
}

func (r *Resolver) getUserInfo(ctx context.Context) map[string]interface{} {
	if ctx.Value(userInfo) == nil {
		return map[string]interface{}{}
	}
	return ctx.Value(userInfo).(map[string]interface{})
}