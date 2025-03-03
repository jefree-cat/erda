// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: projectpipeline.proto

package pb

import (
	context "context"
	http1 "net/http"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// ProjectPipelineServiceHandler is the server API for ProjectPipelineService service.
type ProjectPipelineServiceHandler interface {
	// POST /api/project-pipeline
	Create(context.Context, *CreateProjectPipelineRequest) (*CreateProjectPipelineResponse, error)
	// GET /api/project-pipeline/actions/get-my-apps
	ListApp(context.Context, *ListAppRequest) (*ListAppResponse, error)
	// GET /api/project-pipeline/actions/get-pipeline-yml-list
	ListPipelineYml(context.Context, *ListAppPipelineYmlRequest) (*ListAppPipelineYmlResponse, error)
	// GET /api/project-pipeline/actions/name-pre-check
	CreateNamePreCheck(context.Context, *CreateProjectPipelineNamePreCheckRequest) (*CreateProjectPipelineNamePreCheckResponse, error)
	// GET /api/project-pipeline/actions/source-pre-check
	CreateSourcePreCheck(context.Context, *CreateProjectPipelineSourcePreCheckRequest) (*CreateProjectPipelineSourcePreCheckResponse, error)
}

// RegisterProjectPipelineServiceHandler register ProjectPipelineServiceHandler to http.Router.
func RegisterProjectPipelineServiceHandler(r http.Router, srv ProjectPipelineServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_Create := func(method, path string, fn func(context.Context, *CreateProjectPipelineRequest) (*CreateProjectPipelineResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateProjectPipelineRequest))
		}
		var Create_info transport.ServiceInfo
		if h.Interceptor != nil {
			Create_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Create", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, Create_info)
				}
				r = r.WithContext(ctx)
				var in CreateProjectPipelineRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListApp := func(method, path string, fn func(context.Context, *ListAppRequest) (*ListAppResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListAppRequest))
		}
		var ListApp_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListApp_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListApp", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListApp_info)
				}
				r = r.WithContext(ctx)
				var in ListAppRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListPipelineYml := func(method, path string, fn func(context.Context, *ListAppPipelineYmlRequest) (*ListAppPipelineYmlResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListAppPipelineYmlRequest))
		}
		var ListPipelineYml_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListPipelineYml_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListPipelineYml", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListPipelineYml_info)
				}
				r = r.WithContext(ctx)
				var in ListAppPipelineYmlRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateNamePreCheck := func(method, path string, fn func(context.Context, *CreateProjectPipelineNamePreCheckRequest) (*CreateProjectPipelineNamePreCheckResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateProjectPipelineNamePreCheckRequest))
		}
		var CreateNamePreCheck_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateNamePreCheck_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "CreateNamePreCheck", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateNamePreCheck_info)
				}
				r = r.WithContext(ctx)
				var in CreateProjectPipelineNamePreCheckRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateSourcePreCheck := func(method, path string, fn func(context.Context, *CreateProjectPipelineSourcePreCheckRequest) (*CreateProjectPipelineSourcePreCheckResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateProjectPipelineSourcePreCheckRequest))
		}
		var CreateSourcePreCheck_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateSourcePreCheck_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "CreateSourcePreCheck", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateSourcePreCheck_info)
				}
				r = r.WithContext(ctx)
				var in CreateProjectPipelineSourcePreCheckRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_Create("POST", "/api/project-pipeline", srv.Create)
	add_ListApp("GET", "/api/project-pipeline/actions/get-my-apps", srv.ListApp)
	add_ListPipelineYml("GET", "/api/project-pipeline/actions/get-pipeline-yml-list", srv.ListPipelineYml)
	add_CreateNamePreCheck("GET", "/api/project-pipeline/actions/name-pre-check", srv.CreateNamePreCheck)
	add_CreateSourcePreCheck("GET", "/api/project-pipeline/actions/source-pre-check", srv.CreateSourcePreCheck)
}
