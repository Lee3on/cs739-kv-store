#include <unistd.h> // For system()

#include <chrono>
#include <condition_variable>
#include <cstdio>
#include <cstdlib> // For getenv
#include <cstring>
#include <fstream>
#include <functional>
#include <iostream>
#include <mutex>
#include <random>
#include <string>
#include <thread>
#include <vector>
#include <sstream>
#include <unordered_map>

// Include the client header file
#include "kv739_client.h"

// Global server address variable
std::string g_server_address;

// Helper function to initialize the client with the config file
void init_client(const std::string &config_file)
{
    // g_server_address = server_address; // Store the server address for re-use
    char *config_file_name = const_cast<char *>(config_file.c_str());
    if (kv739_init(config_file_name) != 0)
    {
        std::cerr << "Failed to initialize client with config file: "
                  << config_file << std::endl;
        exit(-1);
    }
    std::cout << "Client successfully initialized with config file: "
              << config_file << std::endl;
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
    std::cout << "Test: " << test_name << " - "
              << (passed ? "PASSED" : "FAILED") << std::endl;
}

char *generateRandomString(int length)
{
    const char *chars =
        "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    char *randomStr = new char[length + 1]; // +1 for the null terminator

    for (int i = 0; i < length; ++i)
    {
        randomStr[i] = chars[rand() % std::strlen(chars)];
    }
    randomStr[length] = '\0'; // Null-terminate the string

    return randomStr;
}

std::vector<std::string> split(const std::string &str, char delimiter)
{
    std::vector<std::string> tokens;
    std::istringstream stream(str);
    std::string token;

    while (std::getline(stream, token, delimiter))
    {
        if (!token.empty())
        { // Ignore empty tokens if you want
            tokens.push_back(token);
        }
    }

    return tokens;
}

void read_config_file(std::unordered_map<std::string, std::string> &all_server)
{
    std::ifstream file("./config/kv_server_list");
    std::string line;
    while (std::getline(file, line))
    {
        char delimiter = ' ';
        std::vector<std::string> strs = split(line, delimiter);
        all_server[strs[0]] = strs[1];
    }
    file.close();
}

// Test functions

void test_reinitialization_with_shutdown()
{
    std::string test_name = "Reinitialization with Shutdown";
    printf("\nTest: %s\n", test_name.c_str());

    // The client is already initialized in main

    shutdown_client(); // Shutdown client

    int ret = kv739_init(
        const_cast<char *>(g_server_address.c_str())); // Re-initialize client
    if (ret == 0)
    {
        print_test_result(test_name, true);
    }
    else
    {
        print_test_result(test_name, false);
    }
}

void test_reinitialization_without_shutdown()
{
    std::string test_name = "Reinitialization without Shutdown";
    printf("\nTest: %s\n", test_name.c_str());

    // The client is already initialized in main

    int ret = kv739_init(const_cast<char *>(g_server_address.c_str()));
    if (ret == -1)
    {
        print_test_result(test_name, true);
    }
    else
    {
        print_test_result(test_name, false);
    }
}

void test_get_non_existent_key()
{
    std::string test_name = "Get Non-Existent Key";
    printf("\nTest: %s\n", test_name.c_str());

    char non_existent_key[] = "non_existent_key";
    char value[2049] = {0};

    // Attempt to get the non-existent key
    int ret = kv739_get(non_existent_key, value);
    if (ret == 1)
    {
        print_test_result(test_name, true);
    }
    else
    {
        print_test_result(test_name, false);
    }
}

void test_get_existent_key()
{
    std::string test_name = "Get Existent Key";
    printf("\nTest: %s\n", test_name.c_str());

    int ret;
    char key[] = "existing_key";
    char value[2049] = "existing_value";
    char old_value[2049] = {0};
    char get_value[2049] = {0};

    kv739_put(key, value, old_value);
    ret = kv739_get(key, get_value);
    if (ret != 0)
    {
        print_test_result(test_name, false);
    }
    else if (strcmp(get_value, value) != 0)
    {
        print_test_result(test_name, false);
    }
    else
    {
        print_test_result(test_name, true);
    }
}

void test_put_non_existent_key()
{
    std::string test_name = "Put Non-Existent Key";
    printf("\nTest: %s\n", test_name.c_str());

    int ret;
    char non_existent_key[] = "non_existent_key_put";
    char value[2049] = "non_existing_value";
    char old_value[2049] = {0};

    ret = kv739_put(non_existent_key, value, old_value);
    if (ret != 1)
    {
        print_test_result(test_name, false);
    }
    else if (strlen(old_value) != 0)
    {
        print_test_result(test_name, false);
    }
    else
    {
        print_test_result(test_name, true);
    }
}

void test_put_existent_key()
{
    std::string test_name = "Put Existent Key";
    printf("\nTest: %s\n", test_name.c_str());

    int ret;
    char existent_key[] = "existent_key_put";
    char value[2049] = "existing_value";
    char new_value[2049] = "new_value";
    char old_value[2049] = {0};

    kv739_put(existent_key, value, old_value);           // Initial put
    ret = kv739_put(existent_key, new_value, old_value); // Overwrite
    if (ret != 0)
    {
        print_test_result(test_name, false);
    }
    else if (strcmp(old_value, value) != 0)
    {
        print_test_result(test_name, false);
    }
    else
    {
        print_test_result(test_name, true);
    }
}

void test_put_get_order_1()
{
    std::string test_name = "Put/Get Order 1";
    printf("\nTest: %s\n", test_name.c_str());

    std::mutex mtx;
    std::condition_variable cv;
    bool is_first_put_done = false;
    bool passed = true;

    // Function for putting keys
    auto put_keys = [&]()
    {
        char key1[] = "key1_order1";
        char key2[] = "key2_order1";
        char key3[] = "key3_order1";

        char value1[] = "value1";
        char value2[] = "value2";
        char value3[] = "value3";

        char old_value[2049] = {0};

        kv739_put(key1, value1, old_value);

        // Notify that the first put is done
        {
            std::lock_guard<std::mutex> lock(mtx);
            is_first_put_done = true;
        }
        cv.notify_all();

        kv739_put(key2, value2, old_value);
        kv739_put(key3, value3, old_value);
    };

    // Function for getting keys
    auto get_keys = [&]()
    {
        {
            std::unique_lock<std::mutex> lock(mtx);
            cv.wait(lock, [&]
                    { return is_first_put_done; });
        }
        std::this_thread::sleep_for(std::chrono::milliseconds(200));
        char value[2049] = {0};
        for (int i = 1; i <= 3; ++i)
        {
            char key[20];
            snprintf(key, sizeof(key), "key%d_order1", i);
            int ret = kv739_get(key, value);
            if (ret == 0)
            {
                printf("Retrieved %s = %s\n", key, value);
            }
            else
            {
                printf("Could not retrieve %s\n", key);
                passed = false;
            }
        }
    };

    std::thread put_thread(put_keys);
    std::thread get_thread(get_keys);

    put_thread.join();
    get_thread.join();

    print_test_result(test_name, passed);
}

void test_put_get_order_2()
{
    std::string test_name = "Put/Get Order 2";
    printf("\nTest: %s\n", test_name.c_str());

    std::mutex mtx;
    std::condition_variable cv;
    bool is_second_put_done = false;
    bool is_get_attempt_done = false;
    bool is_all_puts_done = false;
    bool passed = true;

    // Use unique keys for this test
    auto put_keys = [&]()
    {
        char key1[] = "unq_key1";
        char key2[] = "unq_key2";
        char key3[] = "unq_key3";

        char value1[] = "value1";
        char value2[] = "value2";
        char value3[] = "value3";

        char old_value[2049] = {0};

        // Put key1 and key2
        kv739_put(key1, value1, old_value);
        kv739_put(key2, value2, old_value);

        // Notify that the second put (key2) is done
        {
            std::lock_guard<std::mutex> lock(mtx);
            is_second_put_done = true;
        }
        cv.notify_all();

        // Wait for the Get Thread to attempt getting key3
        {
            std::unique_lock<std::mutex> lock(mtx);
            cv.wait(lock, [&]
                    { return is_get_attempt_done; });
        }

        // Now put key3
        kv739_put(key3, value3, old_value);

        // Notify that all puts are done
        {
            std::lock_guard<std::mutex> lock(mtx);
            is_all_puts_done = true;
        }
        cv.notify_all();
    };

    // Function for getting keys in reverse order
    auto get_keys = [&]()
    {
        {
            std::unique_lock<std::mutex> lock(mtx);
            cv.wait(lock, [&]
                    { return is_second_put_done; });
        }

        char value[2049] = {0};

        // Attempt to get key3 before it's written
        int ret = kv739_get("unq_key3", value);
        if (ret == 1)
        {
            printf("unq_key3 not found before it is put\n");
        }
        else
        {
            printf("unq_key3 found before it is put\n");
            passed = false;
        }

        // Notify that the Get Thread has attempted to get key3
        {
            std::lock_guard<std::mutex> lock(mtx);
            is_get_attempt_done = true;
        }
        cv.notify_all();

        // Wait for all puts to be done
        {
            std::unique_lock<std::mutex> lock(mtx);
            cv.wait(lock, [&]
                    { return is_all_puts_done; });
        }

        // Now get the values in reverse order
        for (int i = 3; i >= 1; --i)
        {
            char key[20];
            snprintf(key, sizeof(key), "unq_key%d", i);
            ret = kv739_get(key, value);
            if (ret == 0)
            {
                printf("Retrieved %s = %s\n", key, value);
            }
            else
            {
                printf("Could not retrieve %s\n", key);
                passed = false;
            }
        }
    };

    std::thread put_thread(put_keys);
    std::thread get_thread(get_keys);

    put_thread.join();
    get_thread.join();

    print_test_result(test_name, passed);
}

void test_concurrent_writes()
{
    std::string test_name = "Concurrent Writes";
    printf("\nTest: %s\n", test_name.c_str());

    char key[] = "concurrent_key";
    char old_value[2049] = {0};
    bool passed = true;

    // Function for first thread's puts
    auto put_keys_first = [&]()
    {
        kv739_put(key, "1", old_value);
        kv739_put(key, "2", old_value);
        kv739_put(key, "3", old_value);
    };

    // Function for second thread's puts
    auto put_keys_second = [&]()
    {
        kv739_put(key, "4", old_value);
        kv739_put(key, "5", old_value);
        kv739_put(key, "6", old_value);
    };

    std::thread put_thread1(put_keys_first);
    std::thread put_thread2(put_keys_second);

    put_thread1.join();
    put_thread2.join();

    // Final get after concurrent writes
    char value[2049] = {0};
    kv739_get(key, value);
    printf("Retrieved %s = %s (Expected: '3' or '6')\n", key, value);
    // Assuming that either "3" or "6" is acceptable
    if (strcmp(value, "3") == 0 || strcmp(value, "6") == 0)
    {
        passed = true;
    }
    else
    {
        passed = false;
    }

    print_test_result(test_name, passed);
}

void test_server_recovery()
{
    std::string test_name = "Server Restart Recovery";
    printf("\nTest: %s\n", test_name.c_str());

    char key[] = "persistent_key";
    char value[] = "persistent_value";
    char old_value[2049] = {0};

    // Put a key-value pair
    kv739_put(key, value, old_value);

    // Simulate server crash. Note, this kills all the servers.
    system("pkill kvstore_server");

    // Wait for server to terminate
    sleep(5);

    // Shutdown client
    kv739_shutdown();

    // Restart the server process
    system("./startup_many.sh");

    // Wait for server to start
    sleep(4);

    // Re-initialize client
    kv739_init(const_cast<char *>(g_server_address.c_str()));

    // Attempt to get the key
    char get_value[2049] = {0};
    int ret = kv739_get(key, get_value);
    if (ret == 0 && strcmp(get_value, value) == 0)
    {
        print_test_result(test_name, true);
    }
    else
    {
        print_test_result(test_name, false);
    }
}

std::unordered_map<std::string, std::string> killed_server;
std::unordered_map<std::string, std::string> stored_kv;
void test_seq_kill_put(int opt)
{
    std::string test_name = "Sequential Kill Put";
    printf("\nTest: %s\n", test_name.c_str());

    std::unordered_map<std::string, std::string> all_server;
    read_config_file(all_server);

    char *key = generateRandomString(10);
    char value[] = "random_value";
    char old_value[2049] = {0};
    int ret;
    bool passed = true;

    std::unordered_map<std::string, std::string> notKilledIP = all_server;
    int i = 0;
    while (i < 3)
    {
        // std::string id = pair.first;
        // std::string server_name = pair.second;

        ret = kv739_put(key, value, old_value);
        if (ret == -1)
        {
            printf("Put failed after killing %d servers\n", i++);
            passed = false;
            return;
        }
        else
        {
            stored_kv[std::string(key)] = std::string(value);
            printf(
                "Put succeeded after killing %d servers with return code %d\n", i++, ret);
        }

        if (notKilledIP.size() > 1)
        {
            int killIndex = rand() % notKilledIP.size();
            auto it = notKilledIP.begin();
            std::advance(it, killIndex);

            std::string killed_id = it->first;
            std::string servername = notKilledIP[killed_id];
            printf("Killing server %s: %s\n", killed_id.c_str(), servername.c_str());
            int result;
            if (opt == 0)
            {
                result = kv739_die(const_cast<char *>(servername.c_str()), 1);
            }
            else
            {
                result = kv739_leave(const_cast<char *>(servername.c_str()), 1);
            }

            if (result == -1)
            {
                printf("Kill server %s failed\n", killed_id.c_str());
                return;
            }

            std::this_thread::sleep_for(std::chrono::milliseconds(1000));

            killed_server[killed_id] = servername;
            notKilledIP.erase(it);
        }

        std::this_thread::sleep_for(std::chrono::milliseconds(10000));
    }

    delete[] key;

    // print_test_result(test_name, passed);
}

void test_seq_recover(int opt)
{
    if (killed_server.empty())
    {
        return;
    }

    std::string test_name = "Sequential Recovery";

    bool pass = true;
    int count = 0;
    bool flag = false;
    for (const auto &pair : killed_server)
    {
        std::string id = pair.first;
        std::string server_name = pair.second;
        printf("id=%s, server_name=%s\n", id.c_str(), server_name.c_str());

        int resp = kv739_start(const_cast<char *>(server_name.c_str()), opt);
        if (resp == -1)
        {
            printf("Start server %s failed\n", server_name.c_str());
            continue;
        }
        count++;

        std::this_thread::sleep_for(std::chrono::milliseconds(5000));

        for (const auto &kvpair : stored_kv)
        {
            std::string k = kvpair.first;
            std::string v = kvpair.second;
            printf("key=%s, server_name=%s\n", k.c_str(), v.c_str());
            char value[2049] = {0};
            int ret = kv739_get(const_cast<char *>(k.c_str()), value);
            if (ret != -1 && !flag)
            {
                printf("Get succeeded after recovering %d nodes\n", count);
                flag = true;
            }
            if (ret != 0 || std::string(value) != v)
            {
                pass = false;
                printf("Get failed to get the expected value. key=%s, expected val=%s, actual val=%s\n", k, v, value);
            }
        }
    }

    print_test_result(test_name, pass);
}

int main(int argc, char *argv[])
{
    // Initialize client with the server address
    init_client("./config_file");

    // printf("Running Correctness Tests...\n");

    // test_reinitialization_with_shutdown();
    // test_reinitialization_without_shutdown();
    // test_get_non_existent_key();
    // test_get_existent_key();
    // test_put_non_existent_key();
    // test_put_existent_key();
    // test_put_get_order_1();
    // test_put_get_order_2();
    // test_concurrent_writes();
    // test_server_recovery();
    if (argc > 1 && std::string(argv[1]) == "--die")
    {
        test_seq_kill_put(0);
        test_seq_recover(0);
    }
    else if (argc > 1 && std::string(argv[1]) == "--leave")
    {
        test_seq_kill_put(1);
        test_seq_recover(1);
    }

    printf("\nAll tests completed\n");

    // Shutdown client after tests
    shutdown_client();

    return 0;
}