// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(app{}),
		Listeners: []string{"bin"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦c9ebda1b:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→harmony-engine/Prompter⟧\n⟦cba7df56:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→bin⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:  "harmony-engine/Prompter",
		Iface: reflect.TypeOf((*Prompter)(nil)).Elem(),
		Impl:  reflect.TypeOf(prompter{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return prompter_local_stub{impl: impl.(Prompter), tracer: tracer, binMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Bin", Remote: false}), fetchMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Fetch", Remote: false}), listMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "List", Remote: false}), registerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Register", Remote: false}), retrieveMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Retrieve", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return prompter_client_stub{stub: stub, binMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Bin", Remote: true}), fetchMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Fetch", Remote: true}), listMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "List", Remote: true}), registerMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Register", Remote: true}), retrieveMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "harmony-engine/Prompter", Method: "Retrieve", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return prompter_server_stub{impl: impl.(Prompter), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return prompter_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[weaver.Main] = (*app)(nil)
var _ weaver.InstanceOf[Prompter] = (*prompter)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*app)(nil)
var _ weaver.Unrouted = (*prompter)(nil)

// Local stub implementations.

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

type prompter_local_stub struct {
	impl            Prompter
	tracer          trace.Tracer
	binMetrics      *codegen.MethodMetrics
	fetchMetrics    *codegen.MethodMetrics
	listMetrics     *codegen.MethodMetrics
	registerMetrics *codegen.MethodMetrics
	retrieveMetrics *codegen.MethodMetrics
}

// Check that prompter_local_stub implements the Prompter interface.
var _ Prompter = (*prompter_local_stub)(nil)

func (s prompter_local_stub) Bin(ctx context.Context, a0 *Prompt) (r0 *IPromptResult, err error) {
	// Update metrics.
	begin := s.binMetrics.Begin()
	defer func() { s.binMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Prompter.Bin", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Bin(ctx, a0)
}

func (s prompter_local_stub) Fetch(ctx context.Context, a0 *User) (r0 *User, err error) {
	// Update metrics.
	begin := s.fetchMetrics.Begin()
	defer func() { s.fetchMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Prompter.Fetch", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Fetch(ctx, a0)
}

func (s prompter_local_stub) List(ctx context.Context) (r0 *IPromptListResult, err error) {
	// Update metrics.
	begin := s.listMetrics.Begin()
	defer func() { s.listMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Prompter.List", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.List(ctx)
}

func (s prompter_local_stub) Register(ctx context.Context, a0 *User) (r0 *IUserResult, err error) {
	// Update metrics.
	begin := s.registerMetrics.Begin()
	defer func() { s.registerMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Prompter.Register", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Register(ctx, a0)
}

func (s prompter_local_stub) Retrieve(ctx context.Context, a0 *Prompt) (r0 *Prompt, err error) {
	// Update metrics.
	begin := s.retrieveMetrics.Begin()
	defer func() { s.retrieveMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Prompter.Retrieve", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Retrieve(ctx, a0)
}

// Client stub implementations.

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

type prompter_client_stub struct {
	stub            codegen.Stub
	binMetrics      *codegen.MethodMetrics
	fetchMetrics    *codegen.MethodMetrics
	listMetrics     *codegen.MethodMetrics
	registerMetrics *codegen.MethodMetrics
	retrieveMetrics *codegen.MethodMetrics
}

// Check that prompter_client_stub implements the Prompter interface.
var _ Prompter = (*prompter_client_stub)(nil)

func (s prompter_client_stub) Bin(ctx context.Context, a0 *Prompt) (r0 *IPromptResult, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.binMetrics.Begin()
	defer func() { s.binMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Prompter.Bin", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_ptr_Prompt_085838cb(a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_ptr_Prompt_085838cb(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_IPromptResult_c276f7f7(dec)
	err = dec.Error()
	return
}

func (s prompter_client_stub) Fetch(ctx context.Context, a0 *User) (r0 *User, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.fetchMetrics.Begin()
	defer func() { s.fetchMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Prompter.Fetch", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_ptr_User_29f1f4c9(a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_ptr_User_29f1f4c9(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_User_29f1f4c9(dec)
	err = dec.Error()
	return
}

func (s prompter_client_stub) List(ctx context.Context) (r0 *IPromptListResult, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.listMetrics.Begin()
	defer func() { s.listMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Prompter.List", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 2, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_IPromptListResult_c31a6ab9(dec)
	err = dec.Error()
	return
}

func (s prompter_client_stub) Register(ctx context.Context, a0 *User) (r0 *IUserResult, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.registerMetrics.Begin()
	defer func() { s.registerMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Prompter.Register", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_ptr_User_29f1f4c9(a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_ptr_User_29f1f4c9(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_IUserResult_a051ac57(dec)
	err = dec.Error()
	return
}

func (s prompter_client_stub) Retrieve(ctx context.Context, a0 *Prompt) (r0 *Prompt, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.retrieveMetrics.Begin()
	defer func() { s.retrieveMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Prompter.Retrieve", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_ptr_Prompt_085838cb(a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	serviceweaver_enc_ptr_Prompt_085838cb(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 4, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_Prompt_085838cb(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.22.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

type prompter_server_stub struct {
	impl    Prompter
	addLoad func(key uint64, load float64)
}

// Check that prompter_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*prompter_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s prompter_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Bin":
		return s.bin
	case "Fetch":
		return s.fetch
	case "List":
		return s.list
	case "Register":
		return s.register
	case "Retrieve":
		return s.retrieve
	default:
		return nil
	}
}

func (s prompter_server_stub) bin(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *Prompt
	a0 = serviceweaver_dec_ptr_Prompt_085838cb(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Bin(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_IPromptResult_c276f7f7(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s prompter_server_stub) fetch(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *User
	a0 = serviceweaver_dec_ptr_User_29f1f4c9(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Fetch(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_User_29f1f4c9(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s prompter_server_stub) list(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.List(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_IPromptListResult_c31a6ab9(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s prompter_server_stub) register(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *User
	a0 = serviceweaver_dec_ptr_User_29f1f4c9(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Register(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_IUserResult_a051ac57(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s prompter_server_stub) retrieve(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *Prompt
	a0 = serviceweaver_dec_ptr_Prompt_085838cb(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Retrieve(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_Prompt_085838cb(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)

type prompter_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that prompter_reflect_stub implements the Prompter interface.
var _ Prompter = (*prompter_reflect_stub)(nil)

func (s prompter_reflect_stub) Bin(ctx context.Context, a0 *Prompt) (r0 *IPromptResult, err error) {
	err = s.caller("Bin", ctx, []any{a0}, []any{&r0})
	return
}

func (s prompter_reflect_stub) Fetch(ctx context.Context, a0 *User) (r0 *User, err error) {
	err = s.caller("Fetch", ctx, []any{a0}, []any{&r0})
	return
}

func (s prompter_reflect_stub) List(ctx context.Context) (r0 *IPromptListResult, err error) {
	err = s.caller("List", ctx, []any{}, []any{&r0})
	return
}

func (s prompter_reflect_stub) Register(ctx context.Context, a0 *User) (r0 *IUserResult, err error) {
	err = s.caller("Register", ctx, []any{a0}, []any{&r0})
	return
}

func (s prompter_reflect_stub) Retrieve(ctx context.Context, a0 *Prompt) (r0 *Prompt, err error) {
	err = s.caller("Retrieve", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*IPromptListResult)(nil)

type __is_IPromptListResult[T ~struct {
	weaver.AutoMarshal
	Prompts []Prompt "json:\"prompts\""
}] struct{}

var _ __is_IPromptListResult[IPromptListResult]

func (x *IPromptListResult) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("IPromptListResult.WeaverMarshal: nil receiver"))
	}
	serviceweaver_enc_slice_Prompt_e51fb0c6(enc, x.Prompts)
}

func (x *IPromptListResult) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("IPromptListResult.WeaverUnmarshal: nil receiver"))
	}
	x.Prompts = serviceweaver_dec_slice_Prompt_e51fb0c6(dec)
}

func serviceweaver_enc_slice_Prompt_e51fb0c6(enc *codegen.Encoder, arg []Prompt) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Prompt_e51fb0c6(dec *codegen.Decoder) []Prompt {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Prompt, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

var _ codegen.AutoMarshal = (*IPromptResult)(nil)

type __is_IPromptResult[T ~struct {
	weaver.AutoMarshal
	Id int64 "json:\"id\""
}] struct{}

var _ __is_IPromptResult[IPromptResult]

func (x *IPromptResult) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("IPromptResult.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.Id)
}

func (x *IPromptResult) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("IPromptResult.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.Int64()
}

var _ codegen.AutoMarshal = (*IUserResult)(nil)

type __is_IUserResult[T ~struct {
	weaver.AutoMarshal
	Id int64 "json:\"id\""
}] struct{}

var _ __is_IUserResult[IUserResult]

func (x *IUserResult) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("IUserResult.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.Id)
}

func (x *IUserResult) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("IUserResult.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.Int64()
}

var _ codegen.AutoMarshal = (*Prompt)(nil)

type __is_Prompt[T ~struct {
	weaver.AutoMarshal
	Id        int    "json:\"id\""
	Text      string "json:\"text\""
	Model     string "json:\"model\""
	Tags      string "json:\"tags\""
	CreatedAt int64  "json:\"created_at\""
	UpdatedAt int64  "json:\"updated_at\""
}] struct{}

var _ __is_Prompt[Prompt]

func (x *Prompt) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Prompt.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.Id)
	enc.String(x.Text)
	enc.String(x.Model)
	enc.String(x.Tags)
	enc.Int64(x.CreatedAt)
	enc.Int64(x.UpdatedAt)
}

func (x *Prompt) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Prompt.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.Int()
	x.Text = dec.String()
	x.Model = dec.String()
	x.Tags = dec.String()
	x.CreatedAt = dec.Int64()
	x.UpdatedAt = dec.Int64()
}

var _ codegen.AutoMarshal = (*User)(nil)

type __is_User[T ~struct {
	weaver.AutoMarshal
	Id       int    "json:\"id\""
	Username string "json:\"username\""
	Password string "json:\"password\""
}] struct{}

var _ __is_User[User]

func (x *User) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.Id)
	enc.String(x.Username)
	enc.String(x.Password)
}

func (x *User) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.Int()
	x.Username = dec.String()
	x.Password = dec.String()
}

// Encoding/decoding implementations.

func serviceweaver_enc_ptr_Prompt_085838cb(enc *codegen.Encoder, arg *Prompt) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_Prompt_085838cb(dec *codegen.Decoder) *Prompt {
	if !dec.Bool() {
		return nil
	}
	var res Prompt
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_ptr_IPromptResult_c276f7f7(enc *codegen.Encoder, arg *IPromptResult) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_IPromptResult_c276f7f7(dec *codegen.Decoder) *IPromptResult {
	if !dec.Bool() {
		return nil
	}
	var res IPromptResult
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_ptr_User_29f1f4c9(enc *codegen.Encoder, arg *User) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_User_29f1f4c9(dec *codegen.Decoder) *User {
	if !dec.Bool() {
		return nil
	}
	var res User
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_ptr_IPromptListResult_c31a6ab9(enc *codegen.Encoder, arg *IPromptListResult) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_IPromptListResult_c31a6ab9(dec *codegen.Decoder) *IPromptListResult {
	if !dec.Bool() {
		return nil
	}
	var res IPromptListResult
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_ptr_IUserResult_a051ac57(enc *codegen.Encoder, arg *IUserResult) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_IUserResult_a051ac57(dec *codegen.Decoder) *IUserResult {
	if !dec.Bool() {
		return nil
	}
	var res IUserResult
	(&res).WeaverUnmarshal(dec)
	return &res
}

// Size implementations.

// serviceweaver_size_ptr_Prompt_085838cb returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_ptr_Prompt_085838cb(x *Prompt) int {
	if x == nil {
		return 1
	} else {
		return 1 + serviceweaver_size_Prompt_575b9be2(&*x)
	}
}

// serviceweaver_size_ptr_User_29f1f4c9 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_ptr_User_29f1f4c9(x *User) int {
	if x == nil {
		return 1
	} else {
		return 1 + serviceweaver_size_User_f98d3709(&*x)
	}
}

// serviceweaver_size_Prompt_575b9be2 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_Prompt_575b9be2(x *Prompt) int {
	size := 0
	size += 0
	size += 8
	size += (4 + len(x.Text))
	size += (4 + len(x.Model))
	size += (4 + len(x.Tags))
	size += 8
	size += 8
	return size
}

// serviceweaver_size_User_f98d3709 returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_User_f98d3709(x *User) int {
	size := 0
	size += 0
	size += 8
	size += (4 + len(x.Username))
	size += (4 + len(x.Password))
	return size
}
