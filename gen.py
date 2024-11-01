def generate_server_lists():
    kv_start_port = 6000
    raft_start_port = 5000
    num_servers = 20

    kv_file_name = "./config/kv_server_list"
    raft_file_name = "./config/raft_server_list"
    run_server_file_name = "run_servers.sh"

    with open(kv_file_name, 'w') as kv_file:
        for i in range(num_servers):
            kv_server_address = f"{i+1} localhost:{kv_start_port + i}\n"
            kv_file.write(kv_server_address)

    with open(raft_file_name, 'w') as raft_file:
        for i in range(num_servers):
            raft_server_address = f"{i+1} http://127.0.0.1:{raft_start_port + i}\n"
            raft_file.write(raft_server_address)

    with open(run_server_file_name, 'w') as run_server_file:
        for i in range(num_servers):
            run_server_file.write(f"./server --id {i+1} &\n")
        run_server_file.write("sleep 20\n")
        run_server_file.write("./load_balancer &")

    print(f"Server lists saved to {kv_file_name} and {raft_file_name}")


if __name__ == "__main__":
    generate_server_lists()