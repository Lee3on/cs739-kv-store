#ifndef KV739CLIENT_H
#define KV739CLIENT_H

#include <iostream>
#include <memory>
#include <string>

// Initialize the gRPC client with the given server address in "host:port" format.
// Returns 0 on success and -1 on failure.
int kv739_init(char *server_name);

// Initialize the gRPC client with the given configuration file, which contains a list of server addresses.
// Returns 0 on success and -1 on failure.
int kv739_init(char *config_file);

// Shutdown the connection to the server and free up resources.
// Returns 0 on success and -1 on failure.
int kv739_shutdown(void);

// Get the value corresponding to the given key.
// Returns 0 on success and key is present (value will be stored in provided buffer),
//         1 if key is not present,
//         -1 on failure.
int kv739_get(char *key, char *value);

// Perform a get operation on the current value into old_value and then store the specified value.
// Returns 0 on success if there was an old value,
//         1 on success if there was no old value,
//         -1 on failure.
int kv739_put(char *key, char *value, char *old_value);

// Tell the server to terminate itself.
// If clean == 1, then the server can flush state and notify other machines it is failing.
// If clean == 0, then the server should terminate immediately.
// Returns 0 on success and -1 on failure.
int kv739_die(char * server_name, int clean);

#endif