// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// unnamed API: Application Controllers
//
// Command:
// $ goagen
// --design=github.com/aya-eiya/goa-test/design/app
// --out=$(GOPATH)/src/github.com/aya-eiya/goa-test/apps
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// MyResorceController is the controller interface for the MyResorce actions.
type MyResorceController interface {
	goa.Muxer
	MyAction(*MyActionMyResorceContext) error
}

// MountMyResorceController "mounts" a MyResorce resource controller on the given service.
func MountMyResorceController(service *goa.Service, ctrl MyResorceController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewMyActionMyResorceContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*MyType)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.MyAction(rctx)
	}
	service.Mux.Handle("POST", "/actions/my_action", ctrl.MuxHandler("my_action", h, unmarshalMyActionMyResorcePayload))
	service.LogInfo("mount", "ctrl", "MyResorce", "action", "MyAction", "route", "POST /actions/my_action")
}

// unmarshalMyActionMyResorcePayload unmarshals the request body into the context request data Payload field.
func unmarshalMyActionMyResorcePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	var err error
	var payload myType
	rawArrs := req.FormValue("arrs")
	rawIntVal := req.FormValue("intVal")
	if intVal, err2 := strconv.Atoi(rawIntVal); err2 == nil {
		tmp2 := intVal
		tmp1 := &tmp2
		payload.IntVal = tmp1
	} else {
		err = goa.MergeErrors(err, goa.InvalidParamTypeError("intVal", rawIntVal, "integer"))
	}
	if err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
