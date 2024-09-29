// kv739_test.cpp
#include "kv739_client.h"
#include <iostream>
#include <cstring> // For strcpy
#include <cstdlib> // For getenv

#define MAX_VALUE_SIZE 1024 // Define a maximum value size

// Function to test kv739_init
int test_kv739_init(const char *server_name)
{
    std::cout << "Testing kv739_init with server: " << server_name << std::endl;
    if (kv739_init(const_cast<char *>(server_name)) == 0)
    {
        std::cout << "kv739_init: SUCCESS" << std::endl;
        return 0;
    }
    else
    {
        std::cerr << "kv739_init: FAILURE" << std::endl;
        return -1;
    }
}

// Function to test kv739_get
int test_kv739_get(const char *key)
{
    char value[MAX_VALUE_SIZE + 1] = {0}; // Buffer for value to be retrieved

    std::cout << "Testing kv739_get with key: " << key << std::endl;
    int get_result = kv739_get(const_cast<char *>(key), value);
    if (get_result == 0)
    {
        std::cout << "kv739_get: SUCCESS, value is: " << value << std::endl;
        return 0;
    }
    else if (get_result == 1)
    {
        std::cout << "kv739_get: SUCCESS, key not found." << std::endl;
        return 1;
    }
    else
    {
        std::cerr << "kv739_get: FAILURE" << std::endl;
        return -1;
    }
}

// Function to test kv739_put
int test_kv739_put(const char *key, const char *value)
{
    char old_value[MAX_VALUE_SIZE + 1] = {0}; // Buffer for old_value

    std::cout << "Testing kv739_put with key: " << key << ", value: " << value << std::endl;
    int put_result = kv739_put(const_cast<char *>(key), const_cast<char *>(value), old_value);
    if (put_result == 0)
    {
        std::cout << "kv739_put: SUCCESS, old value was: " << old_value << std::endl;
        return 0;
    }
    else if (put_result == 1)
    {
        std::cout << "kv739_put: SUCCESS, no previous value existed." << std::endl;
        return 1;
    }
    else
    {
        std::cerr << "kv739_put: FAILURE" << std::endl;
        return -1;
    }
}

// Function to test kv739_shutdown
int test_kv739_shutdown()
{
    std::cout << "Testing kv739_shutdown" << std::endl;
    if (kv739_shutdown() == 0)
    {
        std::cout << "kv739_shutdown: SUCCESS" << std::endl;
        return 0;
    }
    else
    {
        std::cerr << "kv739_shutdown: FAILURE" << std::endl;
        return -1;
    }
}

int main()
{
    // Define server address in "host:port" format (update this as needed for your setup)
    // Read the server address from the environment variable "SERVER_ADDRESS"
    const char *server_name_env = getenv("SERVER_ADDRESS");

    // If the environment variable is not set, use a default value
    const char *server_name = (server_name_env != nullptr) ? server_name_env : "localhost:6666";

    // Run kv739_init test
    if (test_kv739_init(server_name) != 0)
    {
        std::cerr << "Test kv739_init failed. Exiting tests." << std::endl;
        return -1;
    }

    // Run kv739_put test
    const char test_key[] = "test_key";
    const char test_value[] = "test_value";
    if (test_kv739_put(test_key, test_value) == -1)
    {
        std::cerr << "Test kv739_put failed. Exiting tests." << std::endl;
        kv739_shutdown();
        return -1;
    }

    // Run kv739_get test
    if (test_kv739_get(test_key) == -1)
    {
        std::cerr << "Test kv739_get failed. Exiting tests." << std::endl;
        kv739_shutdown();
        return -1;
    }

    // Run kv739_put test again with a different value
    const char new_value[] = "new_test_value";
    if (test_kv739_put(test_key, new_value) == -1)
    {
        std::cerr << "Test kv739_put with new value failed. Exiting tests." << std::endl;
        kv739_shutdown();
        return -1;
    }

    // Run kv739_get test again to verify new value
    if (test_kv739_get(test_key) == -1)
    {
        std::cerr << "Test kv739_get after put failed. Exiting tests." << std::endl;
        kv739_shutdown();
        return -1;
    }

    // Run kv739_shutdown test
    if (test_kv739_shutdown() != 0)
    {
        std::cerr << "Test kv739_shutdown failed." << std::endl;
        return -1;
    }

    std::cout << "All tests passed successfully!" << std::endl;
    return 0;
}