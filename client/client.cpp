#include "kv739_client.h"
#include <grpcpp/grpcpp.h>
#include "kv739.grpc.pb.h"
#include <string>
#include <iostream>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using kv739::GetRequest;
using kv739::GetResponse;
using kv739::KVStoreService;
using kv739::PutRequest;
using kv739::PutResponse;

class KVStoreClient
{
public:
    // Constructor that initializes the gRPC stub for KVStoreService
    KVStoreClient(std::shared_ptr<Channel> channel)
        : stub_(KVStoreService::NewStub(channel)) {}

    // Get operation for retrieving a value by key
    int kv739_get(const std::string &key, std::string &value)
    {
        GetRequest request;
        request.set_key(key);

        GetResponse response;
        ClientContext context;

        Status status = stub_->Get(&context, request, &response);

        if (!status.ok())
        {
            std::cerr << "gRPC Get request failed: " << status.error_message() << std::endl;
            return -1; // Failure
        }

        if (response.status() == 1)
        {
            return 1; // Key not found
        }

        value = response.value();
        return 0; // Success
    }

    // Put operation for storing a value and getting the old value, if any
    int kv739_put(const std::string &key, const std::string &value, std::string &old_value)
    {
        PutRequest request;
        request.set_key(key);
        request.set_value(value);

        PutResponse response;
        ClientContext context;

        Status status = stub_->Put(&context, request, &response);

        if (!status.ok())
        {
            std::cerr << "gRPC Put request failed: " << status.error_message() << std::endl;
            return -1; // Failure
        }

        old_value = response.old_value();
        return response.status(); // 0 if old value exists, 1 if no old value
    }

private:
    std::unique_ptr<KVStoreService::Stub> stub_; // gRPC stub to communicate with the server
};

// Global pointer to the gRPC client object
KVStoreClient *client = nullptr;

// Initialize the gRPC client with the given server address in "host:port" format.
int kv739_init(char *server_name)
{
    if (client != nullptr)
    {
        std::cerr << "Client is already initialized." << std::endl;
        return -1;
    }

    // Create a new KVStoreClient instance with the given server address
    std::string server_address(server_name);
    client = new KVStoreClient(grpc::CreateChannel(server_address, grpc::InsecureChannelCredentials()));

    // Check if client is created successfully
    if (client == nullptr)
    {
        std::cerr << "Failed to initialize client." << std::endl;
        return -1;
    }

    return 0;
}

// Shutdown the connection to the server and free up resources.
int kv739_shutdown(void)
{
    if (client == nullptr)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    // Delete the client and free resources
    delete client;
    client = nullptr;

    return 0;
}

// Get the value corresponding to the given key.
int kv739_get(char *key, char *value)
{
    if (client == nullptr)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    std::string key_str(key);
    std::string value_str;

    // Perform get operation
    int result = client->kv739_get(key_str, value_str);

    if (result == 0)
    {
        // Copy the value to the provided buffer
        strcpy(value, value_str.c_str());
    }

    return result;
}

// Perform a get operation on the current value into old_value and then store the specified value.
int kv739_put(char *key, char *value, char *old_value)
{
    if (client == nullptr)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    std::string key_str(key);
    std::string value_str(value);
    std::string old_value_str;

    // Perform put operation
    int result = client->kv739_put(key_str, value_str, old_value_str);

    if (result == 0 || result == 1)
    {
        // Copy the old value to the provided buffer
        strcpy(old_value, old_value_str.c_str());
    }

    return result;
}
