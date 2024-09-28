#include <grpcpp/grpcpp.h>
#include "kvstore.grpc.pb.h"
#include <string>
#include <iostream>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using kv739::KVStoreService;
using kv739::GetRequest;
using kv739::GetResponse;
using kv739::PutRequest;
using kv739::PutResponse;

class KVStoreClient {
public:
    // Constructor that initializes the gRPC stub for KVStoreService
    KVStoreClient(std::shared_ptr<Channel> channel)
        : stub_(KVStoreService::NewStub(channel)) {}

    // Get operation for retrieving a value by key
    int kv739_get(const std::string& key, std::string& value) {
        GetRequest request;
        request.set_key(key);

        GetResponse response;
        ClientContext context;

        Status status = stub_->Get(&context, request, &response);

        if (!status.ok()) {
            std::cerr << "gRPC Get request failed: " << status.error_message() << std::endl;
            return -1;  // Failure
        }

        if (response.status() == 1) {
            return 1;  // Key not found
        }

        value = response.value();
        return 0;  // Success
    }

    // Put operation for storing a value and getting the old value, if any
    int kv739_put(const std::string& key, const std::string& value, std::string& old_value) {
        PutRequest request;
        request.set_key(key);
        request.set_value(value);

        PutResponse response;
        ClientContext context;

        Status status = stub_->Put(&context, request, &response);

        if (!status.ok()) {
            std::cerr << "gRPC Put request failed: " << status.error_message() << std::endl;
            return -1;  // Failure
        }

        old_value = response.old_value();
        return response.status();  // 0 if old value exists, 1 if no old value
    }

private:
    std::unique_ptr<KVStoreService::Stub> stub_;  // gRPC stub to communicate with the server
};

int main() {
    // Specify the server address
    std::string server_address = "localhost:50051";

    // Create a KVStoreClient instance
    KVStoreClient client(grpc::CreateChannel(server_address, grpc::InsecureChannelCredentials()));

    // Example usage of the client

    // Perform a Put operation
    std::string old_value;
    int put_status = client.kv739_put("exampleKey", "exampleValue", old_value);
    if (put_status == 0 || put_status == 1) {
        std::cout << "Put operation successful. Old value: " << old_value << std::endl;
    } else {
        std::cout << "Put operation failed." << std::endl;
    }

    // Perform a Get operation
    std::string value;
    int get_status = client.kv739_get("exampleKey", value);
    if (get_status == 0) {
        std::cout << "Get operation successful. Value: " << value << std::endl;
    } else if (get_status == 1) {
        std::cout << "Key not found." << std::endl;
    } else {
        std::cout << "Get operation failed." << std::endl;
    }

    return 0;
}
