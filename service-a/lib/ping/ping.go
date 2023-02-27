package ping

import (
	"context"
	"fmt"
	"net/http"

	libhttp "service-a/lib/http"
	"service-a/lib/tracing"

	"github.com/opentracing/opentracing-go"
)

func Ping(ctx context.Context, hostPort string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ping-send")
	defer span.Finish()

	url := fmt.Sprintf("http://%s/ping", hostPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	if err := tracing.Inject(span, req); err != nil {
		return "", err
	}
	return libhttp.Do(req)
}
