// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: kv739.proto

#include "kv739.pb.h"
#include "kv739.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/message_allocator.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace kv739 {

static const char* KVStoreService_method_names[] = {
  "/kv739.KVStoreService/Get",
  "/kv739.KVStoreService/Put",
};

std::unique_ptr< KVStoreService::Stub> KVStoreService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< KVStoreService::Stub> stub(new KVStoreService::Stub(channel, options));
  return stub;
}

KVStoreService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Get_(KVStoreService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Put_(KVStoreService_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status KVStoreService::Stub::Get(::grpc::ClientContext* context, const ::kv739::GetRequest& request, ::kv739::GetResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::kv739::GetRequest, ::kv739::GetResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Get_, context, request, response);
}

void KVStoreService::Stub::async::Get(::grpc::ClientContext* context, const ::kv739::GetRequest* request, ::kv739::GetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::kv739::GetRequest, ::kv739::GetResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Get_, context, request, response, std::move(f));
}

void KVStoreService::Stub::async::Get(::grpc::ClientContext* context, const ::kv739::GetRequest* request, ::kv739::GetResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Get_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::kv739::GetResponse>* KVStoreService::Stub::PrepareAsyncGetRaw(::grpc::ClientContext* context, const ::kv739::GetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::kv739::GetResponse, ::kv739::GetRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Get_, context, request);
}

::grpc::ClientAsyncResponseReader< ::kv739::GetResponse>* KVStoreService::Stub::AsyncGetRaw(::grpc::ClientContext* context, const ::kv739::GetRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status KVStoreService::Stub::Put(::grpc::ClientContext* context, const ::kv739::PutRequest& request, ::kv739::PutResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::kv739::PutRequest, ::kv739::PutResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Put_, context, request, response);
}

void KVStoreService::Stub::async::Put(::grpc::ClientContext* context, const ::kv739::PutRequest* request, ::kv739::PutResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::kv739::PutRequest, ::kv739::PutResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Put_, context, request, response, std::move(f));
}

void KVStoreService::Stub::async::Put(::grpc::ClientContext* context, const ::kv739::PutRequest* request, ::kv739::PutResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Put_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::kv739::PutResponse>* KVStoreService::Stub::PrepareAsyncPutRaw(::grpc::ClientContext* context, const ::kv739::PutRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::kv739::PutResponse, ::kv739::PutRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Put_, context, request);
}

::grpc::ClientAsyncResponseReader< ::kv739::PutResponse>* KVStoreService::Stub::AsyncPutRaw(::grpc::ClientContext* context, const ::kv739::PutRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncPutRaw(context, request, cq);
  result->StartCall();
  return result;
}

KVStoreService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      KVStoreService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< KVStoreService::Service, ::kv739::GetRequest, ::kv739::GetResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](KVStoreService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::kv739::GetRequest* req,
             ::kv739::GetResponse* resp) {
               return service->Get(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      KVStoreService_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< KVStoreService::Service, ::kv739::PutRequest, ::kv739::PutResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](KVStoreService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::kv739::PutRequest* req,
             ::kv739::PutResponse* resp) {
               return service->Put(ctx, req, resp);
             }, this)));
}

KVStoreService::Service::~Service() {
}

::grpc::Status KVStoreService::Service::Get(::grpc::ServerContext* context, const ::kv739::GetRequest* request, ::kv739::GetResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status KVStoreService::Service::Put(::grpc::ServerContext* context, const ::kv739::PutRequest* request, ::kv739::PutResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace kv739

