#include "kv739_client.h"
#include <grpcpp/grpcpp.h>
#include "kv739.grpc.pb.h"
#include <string>
#include <iostream>
#include <cctype>
#include <fstream>
#include <vector>
#include <arpa/inet.h> // For sockaddr_in and inet_pton
#include <regex>       // For std::regex and std::regex_match
#include <algorithm>   // For std::remove_if

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using kv739::CloseRequest;
using kv739::CloseResponse;
using kv739::GetRequest;
using kv739::GetResponse;
using kv739::KVStoreService;
using kv739::LeaveRequest;
using kv739::LeaveResponse;
using kv739::PutRequest;
using kv739::PutResponse;
using kv739::StartRequest;
using kv739::StartResponse;

bool is_valid_key(const std::string &key)
{
    if (key.empty() || key.length() > 128)
    {
        return false;
    }
    for (char c : key)
    {
        if (c == '[' || c == ']')
        {
            return false;
        }
        if (!std::isalnum(c) && c != '_')
        {
            return false;
        }
    }
    return true;
}

bool is_valid_value(const std::string &value)
{
    if (value.empty() || value.length() > 2048)
    {
        return false;
    }
    for (char c : value)
    {
        if (c == '[' || c == ']')
        {
            return false;
        }
        if (!std::isalnum(c) && c != '_')
        {
            return false;
        }
    }
    return true;
}

class KVStoreClient
{
public:
    // Constructor that initializes the gRPC stub for KVStoreService
    KVStoreClient(std::shared_ptr<Channel> channel)
        : stub_(KVStoreService::NewStub(channel)) {}

    // Get operation for retrieving a value by key
    int kv739_get(const std::string &key, std::string &value)
    {
        // Validate the key
        if (!is_valid_key(key))
        {
            std::cerr << "Invalid key. Keys must be 1-128 characters and contain only letters, digits, or underscore, and cannot include '[' or ']'." << std::endl;
            return -1; // Failure
        }

        GetRequest request;
        request.set_key(key);

        GetResponse response;
        ClientContext context;
        context.set_deadline(std::chrono::system_clock::now() + std::chrono::milliseconds(10000)); // Set a 3-second timeout

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
        // Validate the key and value
        if (!is_valid_key(key))
        {
            std::cerr << "Invalid key. Keys must be 1-128 characters and contain only letters, digits, or underscore, and cannot include '[' or ']'." << std::endl;
            return -1; // Failure
        }
        if (!is_valid_value(value))
        {
            std::cerr << "Invalid value. Values must be 1-2048 characters and contain only letters, digits, or underscore, and cannot include '[' or ']'." << std::endl;
            return -1; // Failure
        }

        PutRequest request;
        request.set_key(key);
        request.set_value(value);

        PutResponse response;
        ClientContext context;
        context.set_deadline(std::chrono::system_clock::now() + std::chrono::milliseconds(3000)); // Set a 3-second timeout

        Status status = stub_->Put(&context, request, &response);

        if (!status.ok())
        {
            std::cerr << "gRPC Put request failed: " << status.error_message() << std::endl;
            return -1; // Failure
        }

        old_value = response.old_value();
        return response.status(); // 0 if old value exists, 1 if no old value
    }

    int kv739_die(const std::string &server_name, int clean)
    {
        CloseRequest request;
        request.set_server_name(server_name);
        request.set_clean(clean);

        CloseResponse response;
        ClientContext context;
        context.set_deadline(std::chrono::system_clock::now() + std::chrono::milliseconds(3000)); // Set a 1-second timeout

        Status status = stub_->Close(&context, request, &response);
        if (!status.ok())
        {
            std::cerr << "gRPC Close request failed: " << status.error_message() << std::endl;
            return -1; // Failure
        }

        return response.status();
    }

    int kv739_start(const std::string &server_name, int is_new)
    {
        StartRequest request;
        request.set_server_name(server_name);
        request.set_new_(is_new);

        StartResponse response;
        ClientContext context;
        context.set_deadline(std::chrono::system_clock::now() + std::chrono::milliseconds(3000)); // Set a 1-second timeout

        Status status = stub_->Start(&context, request, &response);
        if (!status.ok())
        {
            std::cerr << "gRPC Start request failed: " << status.error_message() << std::endl;
            return -1; // Failure
        }

        return response.status();
    }

    int kv739_leave(const std::string &server_name, int clean)
    {
        LeaveRequest request;
        request.set_server_name(server_name);
        request.set_clean(clean);

        LeaveResponse response;
        ClientContext context;
        context.set_deadline(std::chrono::system_clock::now() + std::chrono::milliseconds(3000)); // Set a 1-second timeout

        Status status = stub_->Leave(&context, request, &response);
        if (!status.ok())
        {
            std::cerr << "gRPC Leave request failed: " << status.error_message() << std::endl;
            return -1; // Failure
        }

        return response.status();
    }

private:
    std::unique_ptr<KVStoreService::Stub> stub_; // gRPC stub to communicate with the server
};

// Helper function to validate if the host is a valid IP address
bool isValidIpAddress(const std::string &host)
{
    struct sockaddr_in sa;
    return inet_pton(AF_INET, host.c_str(), &(sa.sin_addr)) != 0;
}

// Helper function to validate the format of each instance (host:port)
bool validateInstanceFormat(const std::string &instance)
{
    std::regex pattern(R"((\w+|\d+\.\d+\.\d+\.\d+):(\d+))");
    std::smatch match;
    return std::regex_match(instance, match, pattern);
}

// Global pointer to the gRPC client object
std::vector<KVStoreClient *> clients;
const int Client_num = 3;
// Initialize the gRPC client with the given server address in "host:port" format.
int kv739_init(char *config_file)
{
    if (clients.size() == Client_num)
    {
        std::cerr << "Client is already initialized." << std::endl;
        return -1;
    }

    std::ifstream file(config_file);
    if (!file.is_open())
    {
        std::cerr << "Failed to open configuration file: " << config_file << std::endl;
        return -1;
    }

    std::string line;
    while (std::getline(file, line))
    {
        line.erase(remove_if(line.begin(), line.end(), isspace), line.end()); // remove whitespace

        if (line.empty() || line[0] == '#')
        {
            // Skip empty lines or comments
            continue;
        }

        // Validate the instance format
        if (!validateInstanceFormat(line))
        {
            std::cerr << "Invalid instance format in configuration file: " << line << std::endl;
            return -1;
        }

        // Extract host and port
        size_t colon_pos = line.find(':');
        std::string host = line.substr(0, colon_pos);
        std::string port_str = line.substr(colon_pos + 1);

        // Validate the port
        int port = std::stoi(port_str);
        if (port <= 0 || port > 65535)
        {
            std::cerr << "Invalid port number in configuration file: " << line << std::endl;
            return -1;
        }

        // If host is numeric, validate as IP; otherwise, assume DNS name
        if (isValidIpAddress(host) || !host.empty())
        {
            KVStoreClient *client = new KVStoreClient(grpc::CreateChannel(line, grpc::InsecureChannelCredentials()));
            // Check if client is created successfully
            if (client == nullptr)
            {
                std::cerr << "Failed to initialize client." << std::endl;
                return -1;
            }

            // Perform a health check to validate the server address
            std::string dummy_value;
            int result = client->kv739_get("dummy_key", dummy_value);

            if (result == -1)
            {
                std::cerr << "Failed to connect to server: " << line << std::endl;
                delete client;
                client = nullptr;
                return -1;
            }

            std::cout << "Successfully connected to server: " << line << std::endl;
            clients.push_back(client);
        }
        else
        {
            std::cerr << "Invalid host in configuration file: " << host << std::endl;
            return -1;
        }
    }

    if (clients.empty())
    {
        std::cerr << "No valid instances found in configuration file." << std::endl;
        return -1;
    }

    return 0;
}

// Shutdown the connection to the server and free up resources.
int kv739_shutdown(void)
{
    if (clients.size() == 0)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    // Delete the client and free resources
    clients.clear();
    clients.shrink_to_fit();

    return 0;
}

// Get the value corresponding to the given key.
int kv739_get(char *key, char *value)
{
    if (clients.size() == 0)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    std::string key_str(key);
    // Validate the key
    if (!is_valid_key(key_str))
    {
        std::cerr << "Invalid key. Keys must be 1-128 characters and contain only letters, digits, or underscore, and cannot include '[' or ']'." << std::endl;
        return -1; // Failure
    }
    std::string value_str;

    int result = -1;
    for (KVStoreClient *client : clients)
    {
        // Perform get operation
        result = client->kv739_get(key_str, value_str);

        if (result == 0)
        {
            // Copy the value to the provided buffer
            strcpy(value, value_str.c_str());
            break;
        }
    }

    return result;
}

// Perform a get operation on the current value into old_value and then store the specified value.
int kv739_put(char *key, char *value, char *old_value)
{
    if (clients.size() == 0)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    std::string key_str(key);
    std::string value_str(value);

    // Validate the key and value
    if (!is_valid_key(key_str))
    {
        std::cerr << "Invalid key. Keys must be 1-128 characters and contain only letters, digits, or underscore, and cannot include '[' or ']'." << std::endl;
        return -1; // Failure
    }
    if (!is_valid_value(value_str))
    {
        std::cerr << "Invalid value. Values must be 1-2048 characters and contain only letters, digits, or underscore, and cannot include '[' or ']'." << std::endl;
        return -1; // Failure
    }

    std::string old_value_str;
    int result = -1;

    for (KVStoreClient *client : clients)
    {
        result = client->kv739_put(key_str, value_str, old_value_str);
        if (result == 0 || result == 1)
        {
            strcpy(old_value, old_value_str.c_str());
            break;
        }
    }

    return result;
}

int kv739_die(char *server_name, int clean)
{
    if (clients.size() == 0)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    std::string server_address(server_name);

    int result = -1;

    for (KVStoreClient *client : clients)
    {
        result = client->kv739_die(server_address, clean);
        if (result == 0)
        {
            break;
        }
    }

    return result;
}

int kv739_start(char *server_name, int is_new)
{
    if (clients.size() == 0)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    std::string server_address(server_name);

    int result = -1;

    for (KVStoreClient *client : clients)
    {
        result = client->kv739_start(server_address, is_new);
        if (result == 0)
        {
            break;
        }
    }

    return result;
}

int kv739_leave(char *server_name, int clean)
{
    if (clients.size() == 0)
    {
        std::cerr << "Client not initialized." << std::endl;
        return -1;
    }

    std::string server_address(server_name);

    int result = -1;

    for (KVStoreClient *client : clients)
    {
        result = client->kv739_leave(server_address, clean);
        if (result == 0)
        {
            break;
        }
    }

    return result;
}