package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/caarlos0/env"
)

type config struct {
	BasicAuthUser string `env:"BASIC_AUTH_USER" envDefault:""`
	BasicAuthPass string `env:"BASIC_AUTH_PASS" envDefault:""`
	Port          string `env:"PORT" envDefault:"80"`
}

var (
	cfg config
	svc *translate.Translate
)

func main() {
	cfg = config{}
	env.Parse(&cfg)

	awscfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("Unable to load AWS SDK config, " + err.Error())
	}
	awscfg.Region = endpoints.UsEast2RegionID
	svc = translate.New(awscfg)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.Handle("/", wrap(cfg, logic))

	log.Printf("[service] listening on %s", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil))
}

type custom struct {
	io.Writer
	http.ResponseWriter
	status int
}

func (r *custom) Write(b []byte) (int, error) {
	if r.Header().Get("Content-Type") == "" {
		r.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return r.Writer.Write(b)
}

func (r *custom) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.status = status
}

func auth(c config, r *http.Request) bool {
	if username, password, ok := r.BasicAuth(); ok {
		return username == c.BasicAuthUser && password == c.BasicAuthPass
	}
	return false
}

func wrap(c config, f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (len(c.BasicAuthUser) > 0) && (len(c.BasicAuthPass) > 0) && !auth(c, r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="REALM"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		proc := time.Now()
		addr := r.RemoteAddr
		if candidate := r.Header["X-Forwarded-For"]; len(candidate) > 0 {
			addr = candidate[0]
		}
		ioWriter := w.(io.Writer)
		writer := &custom{Writer: ioWriter, ResponseWriter: w, status: http.StatusOK}
		f(writer, r)

		log.Printf("[%s] %.3f %d %s %s",
			addr, time.Now().Sub(proc).Seconds(),
			writer.status, r.Method, r.URL)
	})
}

func logic(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{}

	if strings.EqualFold(r.Method, http.MethodPost) {
		from := r.FormValue("f")

		res, err := svc.TextRequest(&translate.TextInput{
			SourceLanguageCode: aws.String("ja"),
			TargetLanguageCode: aws.String("en"),
			Text:               aws.String(from),
		}).Send()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data["from"] = from
		data["to"] = aws.StringValue(res.TranslatedText)
	}
	t := template.Must(template.New("html").Parse(`<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8">
  <title>AWS SDK sample</title>
  <link rel="shortcut icon" href="data:image/x-icon;," type="image/x-icon"> 
  <style>
  body {margin: 50px;}
  textarea {width: 400px; height: 12rem; font-size: 1.1rem;}
  </style>
</head>
<body>
  <form action="/" method="post">
    <textarea name="f" placeholder="日本語">{{ .from }}</textarea>
    <textarea name="t" placeholder="English" readonly="readonly">{{ .to }}</textarea>
    <input type="submit" value="Translate!">
  </form>
</body>
`))
	t.Execute(w, data)
}
