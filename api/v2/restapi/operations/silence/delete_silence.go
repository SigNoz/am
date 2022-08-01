// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSilenceHandlerFunc turns a function with the right signature into a delete silence handler
type DeleteSilenceHandlerFunc func(DeleteSilenceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSilenceHandlerFunc) Handle(params DeleteSilenceParams) middleware.Responder {
	return fn(params)
}

// DeleteSilenceHandler interface for that can handle valid delete silence params
type DeleteSilenceHandler interface {
	Handle(DeleteSilenceParams) middleware.Responder
}

// NewDeleteSilence creates a new http.Handler for the delete silence operation
func NewDeleteSilence(ctx *middleware.Context, handler DeleteSilenceHandler) *DeleteSilence {
	return &DeleteSilence{Context: ctx, Handler: handler}
}

/* DeleteSilence swagger:route DELETE /silence/{silenceID} silence deleteSilence

Delete a silence by its ID
*/
type DeleteSilence struct {
	Context *middleware.Context
	Handler DeleteSilenceHandler
}

func (o *DeleteSilence) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteSilenceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
