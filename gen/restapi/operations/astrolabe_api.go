// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAstrolabeAPI creates a new Astrolabe instance
func NewAstrolabeAPI(spec *loads.Document) *AstrolabeAPI {
	return &AstrolabeAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		GetAstrolabeTasksNexusTaskNexusIDHandler: GetAstrolabeTasksNexusTaskNexusIDHandlerFunc(func(params GetAstrolabeTasksNexusTaskNexusIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetAstrolabeTasksNexusTaskNexusID has not yet been implemented")
		}),
		PostAstrolabeTasksNexusHandler: PostAstrolabeTasksNexusHandlerFunc(func(params PostAstrolabeTasksNexusParams) middleware.Responder {
			return middleware.NotImplemented("operation PostAstrolabeTasksNexus has not yet been implemented")
		}),
		CopyProtectedEntityHandler: CopyProtectedEntityHandlerFunc(func(params CopyProtectedEntityParams) middleware.Responder {
			return middleware.NotImplemented("operation CopyProtectedEntity has not yet been implemented")
		}),
		CreateSnapshotHandler: CreateSnapshotHandlerFunc(func(params CreateSnapshotParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateSnapshot has not yet been implemented")
		}),
		DeleteProtectedEntityHandler: DeleteProtectedEntityHandlerFunc(func(params DeleteProtectedEntityParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteProtectedEntity has not yet been implemented")
		}),
		GetProtectedEntityInfoHandler: GetProtectedEntityInfoHandlerFunc(func(params GetProtectedEntityInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation GetProtectedEntityInfo has not yet been implemented")
		}),
		GetTaskInfoHandler: GetTaskInfoHandlerFunc(func(params GetTaskInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation GetTaskInfo has not yet been implemented")
		}),
		ListProtectedEntitiesHandler: ListProtectedEntitiesHandlerFunc(func(params ListProtectedEntitiesParams) middleware.Responder {
			return middleware.NotImplemented("operation ListProtectedEntities has not yet been implemented")
		}),
		ListServicesHandler: ListServicesHandlerFunc(func(params ListServicesParams) middleware.Responder {
			return middleware.NotImplemented("operation ListServices has not yet been implemented")
		}),
		ListSnapshotsHandler: ListSnapshotsHandlerFunc(func(params ListSnapshotsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListSnapshots has not yet been implemented")
		}),
		ListTaskNexusHandler: ListTaskNexusHandlerFunc(func(params ListTaskNexusParams) middleware.Responder {
			return middleware.NotImplemented("operation ListTaskNexus has not yet been implemented")
		}),
		ListTasksHandler: ListTasksHandlerFunc(func(params ListTasksParams) middleware.Responder {
			return middleware.NotImplemented("operation ListTasks has not yet been implemented")
		}),
	}
}

/*AstrolabeAPI Astrolabe data protection framework API */
type AstrolabeAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// GetAstrolabeTasksNexusTaskNexusIDHandler sets the operation handler for the get astrolabe tasks nexus task nexus ID operation
	GetAstrolabeTasksNexusTaskNexusIDHandler GetAstrolabeTasksNexusTaskNexusIDHandler
	// PostAstrolabeTasksNexusHandler sets the operation handler for the post astrolabe tasks nexus operation
	PostAstrolabeTasksNexusHandler PostAstrolabeTasksNexusHandler
	// CopyProtectedEntityHandler sets the operation handler for the copy protected entity operation
	CopyProtectedEntityHandler CopyProtectedEntityHandler
	// CreateSnapshotHandler sets the operation handler for the create snapshot operation
	CreateSnapshotHandler CreateSnapshotHandler
	// DeleteProtectedEntityHandler sets the operation handler for the delete protected entity operation
	DeleteProtectedEntityHandler DeleteProtectedEntityHandler
	// GetProtectedEntityInfoHandler sets the operation handler for the get protected entity info operation
	GetProtectedEntityInfoHandler GetProtectedEntityInfoHandler
	// GetTaskInfoHandler sets the operation handler for the get task info operation
	GetTaskInfoHandler GetTaskInfoHandler
	// ListProtectedEntitiesHandler sets the operation handler for the list protected entities operation
	ListProtectedEntitiesHandler ListProtectedEntitiesHandler
	// ListServicesHandler sets the operation handler for the list services operation
	ListServicesHandler ListServicesHandler
	// ListSnapshotsHandler sets the operation handler for the list snapshots operation
	ListSnapshotsHandler ListSnapshotsHandler
	// ListTaskNexusHandler sets the operation handler for the list task nexus operation
	ListTaskNexusHandler ListTaskNexusHandler
	// ListTasksHandler sets the operation handler for the list tasks operation
	ListTasksHandler ListTasksHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AstrolabeAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AstrolabeAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AstrolabeAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AstrolabeAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AstrolabeAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AstrolabeAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AstrolabeAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AstrolabeAPI
func (o *AstrolabeAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetAstrolabeTasksNexusTaskNexusIDHandler == nil {
		unregistered = append(unregistered, "GetAstrolabeTasksNexusTaskNexusIDHandler")
	}
	if o.PostAstrolabeTasksNexusHandler == nil {
		unregistered = append(unregistered, "PostAstrolabeTasksNexusHandler")
	}
	if o.CopyProtectedEntityHandler == nil {
		unregistered = append(unregistered, "CopyProtectedEntityHandler")
	}
	if o.CreateSnapshotHandler == nil {
		unregistered = append(unregistered, "CreateSnapshotHandler")
	}
	if o.DeleteProtectedEntityHandler == nil {
		unregistered = append(unregistered, "DeleteProtectedEntityHandler")
	}
	if o.GetProtectedEntityInfoHandler == nil {
		unregistered = append(unregistered, "GetProtectedEntityInfoHandler")
	}
	if o.GetTaskInfoHandler == nil {
		unregistered = append(unregistered, "GetTaskInfoHandler")
	}
	if o.ListProtectedEntitiesHandler == nil {
		unregistered = append(unregistered, "ListProtectedEntitiesHandler")
	}
	if o.ListServicesHandler == nil {
		unregistered = append(unregistered, "ListServicesHandler")
	}
	if o.ListSnapshotsHandler == nil {
		unregistered = append(unregistered, "ListSnapshotsHandler")
	}
	if o.ListTaskNexusHandler == nil {
		unregistered = append(unregistered, "ListTaskNexusHandler")
	}
	if o.ListTasksHandler == nil {
		unregistered = append(unregistered, "ListTasksHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AstrolabeAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AstrolabeAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *AstrolabeAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *AstrolabeAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *AstrolabeAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AstrolabeAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the astrolabe API
func (o *AstrolabeAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AstrolabeAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe/tasks/nexus/{taskNexusID}"] = NewGetAstrolabeTasksNexusTaskNexusID(o.context, o.GetAstrolabeTasksNexusTaskNexusIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/astrolabe/tasks/nexus"] = NewPostAstrolabeTasksNexus(o.context, o.PostAstrolabeTasksNexusHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/astrolabe/{service}"] = NewCopyProtectedEntity(o.context, o.CopyProtectedEntityHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/astrolabe/{service}/{protectedEntityID}/snapshots"] = NewCreateSnapshot(o.context, o.CreateSnapshotHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/astrolabe/{service}/{protectedEntityID}"] = NewDeleteProtectedEntity(o.context, o.DeleteProtectedEntityHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe/{service}/{protectedEntityID}"] = NewGetProtectedEntityInfo(o.context, o.GetProtectedEntityInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe/tasks/{taskID}"] = NewGetTaskInfo(o.context, o.GetTaskInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe/{service}"] = NewListProtectedEntities(o.context, o.ListProtectedEntitiesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe"] = NewListServices(o.context, o.ListServicesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe/{service}/{protectedEntityID}/snapshots"] = NewListSnapshots(o.context, o.ListSnapshotsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe/tasks/nexus"] = NewListTaskNexus(o.context, o.ListTaskNexusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/astrolabe/tasks"] = NewListTasks(o.context, o.ListTasksHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AstrolabeAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AstrolabeAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AstrolabeAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AstrolabeAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *AstrolabeAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
