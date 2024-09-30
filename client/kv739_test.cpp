#include <iostream>
#include <thread>
#include <chrono>
#include <string>
#include <cstdlib> // For getenv
#include <vector>
#include <random>
#include <functional>

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

// Helper function to measure and print latency and throughput
void measure_performance(const std::string &test_name, int num_requests, const std::function<void()> &test_function)
{
    auto start = std::chrono::high_resolution_clock::now();

    for (int i = 0; i < num_requests; ++i)
    {
        test_function();
    }

    auto end = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double> duration = end - start;

    double total_time = duration.count();
    double latency = total_time / num_requests;
    double throughput = num_requests / total_time;

    std::cout << "Performance Test: " << test_name << std::endl;
    std::cout << "Total Time: " << total_time << " seconds" << std::endl;
    std::cout << "Latency: " << latency * 1000 << " ms/request" << std::endl;
    std::cout << "Throughput: " << throughput << " requests/second" << std::endl;
}

// Test: Performance with 10% hot keys and 90% of requests to those hot keys
void test_hot_keys_performance(int num_requests)
{
    std::cout << "Running performance test: 10% Hot Keys" << std::endl;

    // Prepare hot keys (10% of total)
    int num_hot_keys = 10;
    std::vector<std::string> hot_keys;
    char old_value[1024] = {0};
    for (int i = 0; i < num_hot_keys; ++i)
    {
        hot_keys.push_back("hot_key_" + std::to_string(i));
        kv739_put(const_cast<char *>(hot_keys[i].c_str()), "initial_hot_value", old_value);
    }

    // Prepare cold keys (remaining 90%)
    int num_cold_keys = 90;
    std::vector<std::string> cold_keys;
    for (int i = 0; i < num_cold_keys; ++i)
    {
        std::string cold_key = "cold_key_" + std::to_string(i);
        cold_keys.push_back(cold_key);
        kv739_put(const_cast<char *>(cold_key.c_str()), "initial_cold_value", old_value);
    }

    // Randomly generate 90% of requests for hot keys and 10% for cold keys
    std::mt19937 gen(std::random_device{}());
    std::uniform_int_distribution<> dist(0, 99); // Generate numbers between 0 and 99

    measure_performance("10% Hot Keys Performance", num_requests, [&]()
    {
        std::string key;
        if (dist(gen) < 90) // 90% of the time, pick a hot key
        {
            key = hot_keys[dist(gen) % num_hot_keys];
        }
        else // 10% of the time, pick a random cold key
        {
            key = cold_keys[dist(gen) % num_cold_keys];
        }
        char value[1024] = {0};
        kv739_get(const_cast<char *>(key.c_str()), value);
    });
}

// Test: Performance with uniformly distributed keys and requests
void test_uniform_distribution_performance(int num_requests)
{
    std::cout << "Running performance test: Uniform Distribution" << std::endl;

    // Prepare a larger number of uniformly distributed keys
    int num_keys = 100;
    std::vector<std::string> keys;
    char old_value[1024] = {0};
    for (int i = 0; i < num_keys; ++i)
    {
        keys.push_back("uniform_key_" + std::to_string(i));
        kv739_put(const_cast<char *>(keys[i].c_str()), "initial_uniform_value", old_value);
    }

    // Randomly pick keys for each request
    std::mt19937 gen(std::random_device{}());
    std::uniform_int_distribution<> dist(0, num_keys - 1);

    measure_performance("Uniform Distribution Performance", num_requests, [&]()
    {
        std::string key = keys[dist(gen)];
        char value[1024] = {0};
        kv739_get(const_cast<char *>(key.c_str()), value);
    });
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

    // Performance tests
    int num_requests = 10000;
    test_hot_keys_performance(num_requests);
    test_uniform_distribution_performance(num_requests);

    // Shutdown client after tests
    shutdown_client();

    return 0;
}