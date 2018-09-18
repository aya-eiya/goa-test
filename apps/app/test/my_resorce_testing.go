// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// unnamed API: my_resorce TestHelpers
//
// Command:
// $ goagen
// --design=github.com/aya-eiya/goa-test/design/app
// --out=$(GOPATH)/src/github.com/aya-eiya/goa-test/apps
// --version=v1.3.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aya-eiya/goa-test/apps/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// MyActionMyResorceOK runs the method MyAction of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func MyActionMyResorceOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.MyResorceController, payload *app.MyType) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/actions/my_action"),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "MyResorceTest"), rw, req, prms)
	myActionCtx, _err := app.NewMyActionMyResorceContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}
	myActionCtx.Payload = payload

	// Perform action
	_err = ctrl.MyAction(myActionCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
