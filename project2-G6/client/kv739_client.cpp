#include <iostream>
#include <chrono>
#include <thread>
#include <string>
#include <cstdlib> // For getenv
#include <atomic>

#include "kv739_client.h"

// Global variable to track server health status
std::atomic<bool> server_healthy(true);

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

// Helper function to check server health
bool check_server_health()
{
    // Use a specific key to test the server's health. This key should be constant and not interfere with other keys.
    const char *test_key = "health_check_key";
    char value[1024] = {0};

    int result = kv739_get(const_cast<char *>(test_key), value);

    // If `kv739_get` returns 0, the key was found (which means the server is healthy).
    if (result == 0 || result == 1) // Status 0: key found, Status 1: key not found (server is still responsive)
    {
        return true;
    }

    return false;
}

// Health check loop to continuously monitor the server
void health_check_loop(int interval_seconds)
{
    while (true)
    {
        bool healthy = check_server_health();

        if (healthy)
        {
            if (!server_healthy.load())
            {
                std::cout << "Server has recovered and is now healthy." << std::endl;
            }
            server_healthy.store(true);
            std::cout << "Server health check passed." << std::endl;
        }
        else
        {
            if (server_healthy.load())
            {
                std::cout << "Server health check failed. Server is down or unreachable." << std::endl;
            }
            server_healthy.store(false);
            break; // Exit the health check loop if the server becomes unhealthy
        }

        // Sleep for the specified interval before the next health check
        std::this_thread::sleep_for(std::chrono::seconds(interval_seconds));
    }
}

int main()
{
    // Read the server address from the environment variable "SERVER_ADDRESS"
    const char *server_address_env = getenv("SERVER_ADDRESS");

    // If the environment variable is not set, use a default value
    std::string server_address = (server_address_env != nullptr) ? server_address_env : "localhost:8080";

    // Initialize client with the server address
    init_client(server_address);

    // Health check interval (seconds)
    int health_check_interval = 5;

    // Start the health check loop in the main thread
    health_check_loop(health_check_interval);

    // If health check loop exits, it means the server is down. Shutdown the client.
    std::cout << "Server is not healthy. Shutting down the client..." << std::endl;
    shutdown_client();

    return 0;
}