#include <iostream>
#include <string>
#include <cstdlib> // For getenv

#include "kv739_client.h"

// Helper function to initialize the client with the server address
void init_client(const std::string &server_address)
{
    char *server_name = const_cast<char *>(server_address.c_str());
    if (kv739_init(server_name) != 0)
    {
        std::cerr << "Failed to initialize client with server address: " << server_name << std::endl;
        exit(-1);
    }
    std::cout << "Client successfully initialized with server address: " << server_name << std::endl;
}

// Helper function to gracefully shut down the client
void shutdown_client()
{
    if (kv739_shutdown() != 0)
    {
        std::cerr << "Failed to shut down the client." << std::endl;
    }
    else
    {
        std::cout << "Client successfully shut down." << std::endl;
    }
}

int main(int argc, char *argv[])
{
    // Read the server address from the environment variable "SERVER_ADDRESS"
    const char *server_address_env = getenv("SERVER_ADDRESS");

    // If the environment variable is not set, use a default value
    std::string server_address = (server_address_env != nullptr) ? server_address_env : "localhost:8080";

    // Initialize client with the server address
    init_client(server_address);

    // Run the tests
    if (argc > 1 && std::string(argv[1]) == "--test-recovery")
    {
        // Run the recovery test only
    }
    else
    {
        // Run the regular tests
    }

    // Shutdown client after tests
    shutdown_client();

    return 0;
}