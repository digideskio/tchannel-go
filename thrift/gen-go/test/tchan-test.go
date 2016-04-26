// @generated Code generated by thrift-gen. Do not modify.

// Package test is generated code used to make or handle TChannel calls using Thrift.
package test

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanSecondService is the interface that defines the server handler and client interface.
type TChanSecondService interface {
	Echo(ctx thrift.Context, arg string) (string, error)
}

// TChanSimpleService is the interface that defines the server handler and client interface.
type TChanSimpleService interface {
	Call(ctx thrift.Context, arg *Data) (*Data, error)
	Simple(ctx thrift.Context) error
}

// Implementation of a client and service handler.

type tchanSecondServiceClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanSecondServiceInheritedClient(thriftService string, client thrift.TChanClient) *tchanSecondServiceClient {
	return &tchanSecondServiceClient{
		thriftService,
		client,
	}
}

// NewTChanSecondServiceClient creates a client that can be used to make remote calls.
func NewTChanSecondServiceClient(client thrift.TChanClient) TChanSecondService {
	return NewTChanSecondServiceInheritedClient("SecondService", client)
}

func (c *tchanSecondServiceClient) Echo(ctx thrift.Context, arg string) (string, error) {
	var resp SecondServiceEchoResult
	args := SecondServiceEchoArgs{
		Arg: arg,
	}
	success, err := c.client.Call(ctx, c.thriftService, "Echo", &args, &resp)
	if err == nil && !success {
	}

	return resp.GetSuccess(), err
}

type tchanSecondServiceServer struct {
	handler TChanSecondService

	interceptors []thrift.Interceptor
}

// NewTChanSecondServiceServer wraps a handler for TChanSecondService so it can be
// registered with a thrift.Server.
func NewTChanSecondServiceServer(handler TChanSecondService) thrift.TChanServer {
	return &tchanSecondServiceServer{
		handler: handler,
	}
}

func (s *tchanSecondServiceServer) Service() string {
	return "SecondService"
}

func (s *tchanSecondServiceServer) Methods() []string {
	return []string{
		"Echo",
	}
}

// RegisterInterceptors registers the provided interceptors with the server.
func (s *tchanSecondServiceServer) RegisterInterceptors(interceptors ...thrift.Interceptor) {
	if s.interceptors == nil {
		interceptorsLength := len(interceptors)
		s.interceptors = make([]thrift.Interceptor, interceptorsLength, interceptorsLength)
	}

	s.interceptors = append(s.interceptors, interceptors...)
}

func (s *tchanSecondServiceServer) callInterceptorsPre(ctx thrift.Context, method string, args athrift.TStruct) error {
	if s.interceptors == nil {
		return nil
	}
	var firstErr error
	for _, interceptor := range s.interceptors {
		err := interceptor.Pre(ctx, method, args)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (s *tchanSecondServiceServer) callInterceptorsPost(ctx thrift.Context, method string, args, response athrift.TStruct, err error) error {
	if s.interceptors == nil {
		return err
	}
	transformedErr := err
	for _, interceptor := range s.interceptors {
		transformedErr = interceptor.Post(ctx, method, args, response, transformedErr)
	}
	return transformedErr
}

func (s *tchanSecondServiceServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "Echo":
		return s.handleEcho(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanSecondServiceServer) handleEcho(ctx thrift.Context, protocol athrift.TProtocol) (handled bool, resp athrift.TStruct, err error) {
	var req SecondServiceEchoArgs
	var res SecondServiceEchoResult
	serviceMethod := "Echo.Echo"

	defer func() {
		if uncaught := recover(); uncaught != nil {
			err = thrift.PanicErr{Value: uncaught}
		}
		err = s.callInterceptorsPost(ctx, serviceMethod, &req, resp, err)
		if err != nil {
			resp = nil
		}
	}()

	if readErr := req.Read(protocol); readErr != nil {
		return false, nil, readErr
	}

	err = s.callInterceptorsPre(ctx, serviceMethod, &req)
	if err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Echo(ctx, req.Arg)

	if err == nil {
		res.Success = &r
	}

	return err == nil, &res, err
}

type tchanSimpleServiceClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanSimpleServiceInheritedClient(thriftService string, client thrift.TChanClient) *tchanSimpleServiceClient {
	return &tchanSimpleServiceClient{
		thriftService,
		client,
	}
}

// NewTChanSimpleServiceClient creates a client that can be used to make remote calls.
func NewTChanSimpleServiceClient(client thrift.TChanClient) TChanSimpleService {
	return NewTChanSimpleServiceInheritedClient("SimpleService", client)
}

func (c *tchanSimpleServiceClient) Call(ctx thrift.Context, arg *Data) (*Data, error) {
	var resp SimpleServiceCallResult
	args := SimpleServiceCallArgs{
		Arg: arg,
	}
	success, err := c.client.Call(ctx, c.thriftService, "Call", &args, &resp)
	if err == nil && !success {
	}

	return resp.GetSuccess(), err
}

func (c *tchanSimpleServiceClient) Simple(ctx thrift.Context) error {
	var resp SimpleServiceSimpleResult
	args := SimpleServiceSimpleArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "Simple", &args, &resp)
	if err == nil && !success {
		if e := resp.SimpleErr; e != nil {
			err = e
		}
	}

	return err
}

type tchanSimpleServiceServer struct {
	handler TChanSimpleService

	interceptors []thrift.Interceptor
}

// NewTChanSimpleServiceServer wraps a handler for TChanSimpleService so it can be
// registered with a thrift.Server.
func NewTChanSimpleServiceServer(handler TChanSimpleService) thrift.TChanServer {
	return &tchanSimpleServiceServer{
		handler: handler,
	}
}

func (s *tchanSimpleServiceServer) Service() string {
	return "SimpleService"
}

func (s *tchanSimpleServiceServer) Methods() []string {
	return []string{
		"Call",
		"Simple",
	}
}

// RegisterInterceptors registers the provided interceptors with the server.
func (s *tchanSimpleServiceServer) RegisterInterceptors(interceptors ...thrift.Interceptor) {
	if s.interceptors == nil {
		interceptorsLength := len(interceptors)
		s.interceptors = make([]thrift.Interceptor, interceptorsLength, interceptorsLength)
	}

	s.interceptors = append(s.interceptors, interceptors...)
}

func (s *tchanSimpleServiceServer) callInterceptorsPre(ctx thrift.Context, method string, args athrift.TStruct) error {
	if s.interceptors == nil {
		return nil
	}
	var firstErr error
	for _, interceptor := range s.interceptors {
		err := interceptor.Pre(ctx, method, args)
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (s *tchanSimpleServiceServer) callInterceptorsPost(ctx thrift.Context, method string, args, response athrift.TStruct, err error) error {
	if s.interceptors == nil {
		return err
	}
	transformedErr := err
	for _, interceptor := range s.interceptors {
		transformedErr = interceptor.Post(ctx, method, args, response, transformedErr)
	}
	return transformedErr
}

func (s *tchanSimpleServiceServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "Call":
		return s.handleCall(ctx, protocol)
	case "Simple":
		return s.handleSimple(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanSimpleServiceServer) handleCall(ctx thrift.Context, protocol athrift.TProtocol) (handled bool, resp athrift.TStruct, err error) {
	var req SimpleServiceCallArgs
	var res SimpleServiceCallResult
	serviceMethod := "Call.Call"

	defer func() {
		if uncaught := recover(); uncaught != nil {
			err = thrift.PanicErr{Value: uncaught}
		}
		err = s.callInterceptorsPost(ctx, serviceMethod, &req, resp, err)
		if err != nil {
			resp = nil
		}
	}()

	if readErr := req.Read(protocol); readErr != nil {
		return false, nil, readErr
	}

	err = s.callInterceptorsPre(ctx, serviceMethod, &req)
	if err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Call(ctx, req.Arg)

	if err == nil {
		res.Success = r
	}

	return err == nil, &res, err
}

func (s *tchanSimpleServiceServer) handleSimple(ctx thrift.Context, protocol athrift.TProtocol) (handled bool, resp athrift.TStruct, err error) {
	var req SimpleServiceSimpleArgs
	var res SimpleServiceSimpleResult
	serviceMethod := "Simple.Simple"

	defer func() {
		if uncaught := recover(); uncaught != nil {
			err = thrift.PanicErr{Value: uncaught}
		}
		err = s.callInterceptorsPost(ctx, serviceMethod, &req, resp, err)
		if err != nil {
			switch v := err.(type) {
			case *SimpleErr:
				if v == nil {
					resp = nil
					err = fmt.Errorf("Handler for simpleErr returned non-nil error type *SimpleErr but nil value")
				}
				res.SimpleErr = v
				err = nil
			default:
				resp = nil
			}
		}
	}()

	if readErr := req.Read(protocol); readErr != nil {
		return false, nil, readErr
	}

	err = s.callInterceptorsPre(ctx, serviceMethod, &req)
	if err != nil {
		return false, nil, err
	}

	err =
		s.handler.Simple(ctx)

	return err == nil, &res, err
}
