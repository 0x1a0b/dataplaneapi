// Code generated by go-swagger; DO NOT EDIT.

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLogTargetsHandlerFunc turns a function with the right signature into a get log targets handler
type GetLogTargetsHandlerFunc func(GetLogTargetsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLogTargetsHandlerFunc) Handle(params GetLogTargetsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetLogTargetsHandler interface for that can handle valid get log targets params
type GetLogTargetsHandler interface {
	Handle(GetLogTargetsParams, interface{}) middleware.Responder
}

// NewGetLogTargets creates a new http.Handler for the get log targets operation
func NewGetLogTargets(ctx *middleware.Context, handler GetLogTargetsHandler) *GetLogTargets {
	return &GetLogTargets{Context: ctx, Handler: handler}
}

/*GetLogTargets swagger:route GET /services/haproxy/configuration/log_targets LogTarget getLogTargets

Return an array of all Log Targets

Returns all Log Targets that are configured in specified parent.

*/
type GetLogTargets struct {
	Context *middleware.Context
	Handler GetLogTargetsHandler
}

func (o *GetLogTargets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLogTargetsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}