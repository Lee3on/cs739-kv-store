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

std::vector<std::string> hot_keys;
std::vector<std::string> cold_keys;
std::vector<std::string> uniform_keys;

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

    std::this_thread::sleep_for(std::chrono::milliseconds(1000)); // Add 1000 ms sleep

    // Overwrite value
    int result = kv739_put("overwrite_key", "new_value", old_value);

    // Validate the overwrite
    char new_value[1024] = {0};
    std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
    kv739_get("overwrite_key", new_value);

    bool passed = (result == 0) && (std::string(old_value) == "initial_value") && (std::string(new_value) == "new_value");

    if (!passed)
    {
        std::cout << "Expected result: 0, actual result: " << result << std::endl;
        std::cout << "Expected old_value: 'initial_value', actual old_value: '" << old_value << "'" << std::endl;
        std::cout << "Expected new_value: 'new_value', actual new_value: '" << new_value << "'" << std::endl;
    }
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

    std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
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
void gen_hot_cold_keys(int total_num, float percent)
{
    int num_hot_keys = total_num * percent;
    int num_cold_keys = total_num - num_hot_keys;
    for (int i = 0; i < num_hot_keys; ++i)
    {
        hot_keys.push_back("hot_key_" + std::to_string(i));
    }
    for (int i = 0; i < num_cold_keys; ++i)
    {
        std::string cold_key = "cold_key_" + std::to_string(i);
        cold_keys.push_back(cold_key);
    }
}

void gen_uniform_distribution_keys(int total_num)
{
    for (int i = 0; i < total_num; ++i)
    {
        uniform_keys.push_back("uniform_key_" + std::to_string(i));
    }
}

void test_hot_keys_performance(int num_keys, int num_requests)
{
    std::cout << "Running performance test: 10% Hot Keys" << std::endl;

    gen_hot_cold_keys(num_keys, 0.1);

    char old_value[1024] = {0};
    for (int i = 0; i < hot_keys.size(); ++i)
    {
        kv739_put(const_cast<char *>(hot_keys[i].c_str()), "initial_hot_value", old_value);
    }

    for (int i = 0; i < cold_keys.size(); ++i)
    {
        kv739_put(const_cast<char *>(cold_keys[i].c_str()), "initial_cold_value", old_value);
    }

    // Randomly generate 90% of requests for hot keys and 10% for cold keys
    std::mt19937 gen(std::random_device{}());
    std::uniform_int_distribution<> dist(0, num_keys - 1);

    std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
    measure_performance("10% Hot Keys Performance", num_requests, [&]()
                        {
        std::string key;
        if (dist(gen) < 90) // 90% of the time, pick a hot key
        {
            key = hot_keys[dist(gen) % hot_keys.size()];
        }
        else // 10% of the time, pick a random cold key
        {
            key = cold_keys[dist(gen) % cold_keys.size()];
        }
        char value[1024] = {0};
        kv739_get(const_cast<char *>(key.c_str()), value); });
}

// Test: Performance with uniformly distributed keys and requests
void test_uniform_distribution_performance(int num_keys, int num_requests)
{
    std::cout << "Running performance test: Uniform Distribution" << std::endl;

    gen_uniform_distribution_keys(num_keys);

    // Prepare a larger number of uniformly distributed keys
    char old_value[1024] = {0};
    for (int i = 0; i < uniform_keys.size(); ++i)
    {
        kv739_put(const_cast<char *>(uniform_keys[i].c_str()), "initial_uniform_value", old_value);
    }

    std::mt19937 gen(std::random_device{}());
    std::uniform_int_distribution<> dist(0, num_keys - 1);

    std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
    measure_performance("Uniform Distribution Performance", num_requests, [&]()
                        {
        std::string key = uniform_keys[dist(gen)];
        char value[1024] = {0};
        int result = kv739_get(const_cast<char *>(key.c_str()), value); });
}

// Special Recovery Test to check for hard-coded predefined keys
void test_recovery(int num_keys, int num_requests)
{
    std::cout << "Running test: Recovery Test" << std::endl;
    bool all_keys_found = true;

    // Hard-coded predefined keys used in the test cases
    std::vector<std::string> predefined_keys = {
        "test_key",      // Used in test_put()
        "overwrite_key", // Used in test_put_overwrite()
        "concurrent_key" // Used in test_concurrent_puts()
    };

    gen_hot_cold_keys(num_keys, 0.1);
    gen_uniform_distribution_keys(num_keys);
    std::copy(hot_keys.begin(), hot_keys.end(), std::back_inserter(predefined_keys));
    std::copy(uniform_keys.begin(), uniform_keys.end(), std::back_inserter(predefined_keys));

    // Check each predefined key to ensure it is still present in the key-value store
    for (const auto &key : predefined_keys)
    {
        char value[1024] = {0};
        int result = kv739_get(const_cast<char *>(key.c_str()), value);

        if (result != 0) // If any key is not found or the value is not correct
        {
            std::cout << "Key not found or value mismatch for key: " << key << std::endl;
            all_keys_found = false;
        }
    }

    print_test_result("Recovery Test", all_keys_found);
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
    int key_nums = 100;
    int num_requests = 10000;
    if (argc > 1 && std::string(argv[1]) == "--put")
    {
        test_put();
    }
    else if (argc > 1 && std::string(argv[1]) == "--get")
    {
        test_get();
    }
    else if (argc > 1 && std::string(argv[1]) == "--put-overwrite")
    {
        test_put_overwrite();
    }
    else if (argc > 1 && std::string(argv[1]) == "--get-non-existent-key")
    {
        test_get_non_existent_key();
    }
    else if (argc > 1 && std::string(argv[1]) == "--concurrent-puts")
    {
        test_concurrent_puts();
    }
    else if (argc > 1 && std::string(argv[1]) == "--test-recovery")
    {
        // Run the recovery test only
        test_recovery(key_nums, num_requests);
    }
    else if (argc == 3 && std::string(argv[1]) == "--close")
    {
        std::string server_name = argv[2];
        int result = kv739_die(const_cast<char *>(server_name.c_str()), 1);
        if (result == 0)
        {
            std::cout << "Server successfully closed." << std::endl;
        }
        else
        {
            std::cerr << "Failed to close the server." << std::endl;
        }
    }
    else if (argc == 3 && std::string(argv[1]) == "--start")
    {
        std::string server_name = argv[2];
        int result = kv739_start(const_cast<char *>(server_name.c_str()), 1);
        if (result == 0)
        {
            std::cout << "Server successfully started." << std::endl;
        }
        else
        {
            std::cerr << "Failed to start the server." << std::endl;
        }
    }
    else if (argc == 3 && std::string(argv[1]) == "--leave")
    {
        std::string server_name = argv[2];
        int result = kv739_leave(const_cast<char *>(server_name.c_str()), 1);
        if (result == 0)
        {
            std::cout << "Server successfully left." << std::endl;
        }
        else
        {
            std::cerr << "Failed to remove the server." << std::endl;
        }
    }
    else
    {
        test_put();
        std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
        test_get();
        std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
        test_put_overwrite();
        std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
        test_get_non_existent_key();
        std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
        test_concurrent_puts();
        std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
        // Performance tests
        test_hot_keys_performance(key_nums, num_requests);
        std::this_thread::sleep_for(std::chrono::milliseconds(100)); // Add 100 ms sleep
        test_uniform_distribution_performance(key_nums, num_requests);
    }

    // Shutdown client after tests
    shutdown_client();

    return 0;
}