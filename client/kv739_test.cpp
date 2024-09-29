#include <iostream>
#include <thread>
#include <chrono>
#include <string>
#include <cstdlib> // For getenv

// Include your client header file
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

// Helper function to print test results
void print_test_result(const std::string &test_name, bool passed)
{
    std::cout << "Test: " << test_name << " - " << (passed ? "PASSED" : "FAILED") << std::endl;
}

// Test: Put Operation
void test_put()
{
    std::cout << "Running test: Put Operation" << std::endl;
    char old_value[1024] = {0};
    int result = kv739_put("test_key", "test_value", old_value);

    bool passed = (result == 1); // Status 1 indicates the key was newly inserted
    print_test_result("Put Operation", passed);
}

// Test: Get Operation
void test_get()
{
    std::cout << "Running test: Get Operation" << std::endl;
    char value[1024] = {0};
    int result = kv739_get("test_key", value);

    bool passed = (result == 0) && (std::string(value) == "test_value"); // Status 0 indicates the key was found
    print_test_result("Get Operation", passed);
}

// Test: Put Overwrite Operation
void test_put_overwrite()
{
    std::cout << "Running test: Put Overwrite Operation" << std::endl;
    char old_value[1024] = {0};

    // Put initial value
    kv739_put("overwrite_key", "initial_value", old_value);

    // Overwrite value
    int result = kv739_put("overwrite_key", "new_value", old_value);

    // Validate the overwrite
    char new_value[1024] = {0};
    kv739_get("overwrite_key", new_value);

    bool passed = (result == 0) && (std::string(old_value) == "initial_value") && (std::string(new_value) == "new_value");
    print_test_result("Put Overwrite Operation", passed);
}

// Test: Get Non-Existent Key
void test_get_non_existent_key()
{
    std::cout << "Running test: Get Non-Existent Key" << std::endl;
    char value[1024] = {0};
    int result = kv739_get("non_existent_key", value);

    bool passed = (result == 1); // Status 1 indicates the key was not found
    print_test_result("Get Non-Existent Key", passed);
}

// Test: Concurrent Puts
void test_concurrent_puts()
{
    std::cout << "Running test: Concurrent Puts Operation" << std::endl;

    // Use multiple threads to simulate concurrent puts
    std::thread threads[10];
    for (int i = 0; i < 10; i++)
    {
        threads[i] = std::thread([i]()
                                 {
            char old_value[1024] = {0};
            std::string key = "concurrent_key";
            std::string value = "value_" + std::to_string(i);
            kv739_put(const_cast<char *>(key.c_str()), const_cast<char *>(value.c_str()), old_value); });
    }

    // Join threads to wait for completion
    for (int i = 0; i < 10; i++)
    {
        threads[i].join();
    }

    // Get the final value after concurrent puts
    char final_value[1024] = {0};
    int result = kv739_get("concurrent_key", final_value);

    // Check that a value was successfully retrieved
    bool passed = (result == 0) && (std::string(final_value).find("value_") == 0);
    print_test_result("Concurrent Puts Operation", passed);
}

int main()
{
    // Read the server address from the environment variable "SERVER_ADDRESS"
    const char *server_address_env = getenv("SERVER_ADDRESS");

    // If the environment variable is not set, use a default value
    std::string server_address = (server_address_env != nullptr) ? server_address_env : "localhost:6666";

    // Initialize client with the server address
    init_client(server_address);

    // Run the tests
    test_put();
    test_get();
    test_put_overwrite();
    test_get_non_existent_key();
    test_concurrent_puts();

    // Shutdown client after tests
    shutdown_client();

    return 0;
}