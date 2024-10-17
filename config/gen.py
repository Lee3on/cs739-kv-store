def generate_server_lists():
    # 定义起始端口和服务器数量
    kv_start_port = 6000
    raft_start_port = 5000
    num_servers = 100

    # 文件名
    kv_file_name = "kv_server_list"
    raft_file_name = "raft_server_list"

    # 打开并写入 kv_server_list 文件
    with open(kv_file_name, 'w') as kv_file:
        for i in range(num_servers):
            kv_server_address = f"localhost:{kv_start_port + i}\n"
            kv_file.write(kv_server_address)

    # 打开并写入 raft_server_list 文件
    with open(raft_file_name, 'w') as raft_file:
        for i in range(num_servers):
            raft_server_address = f"http://127.0.0.1:{raft_start_port + i}\n"
            raft_file.write(raft_server_address)

    print(f"Server lists saved to {kv_file_name} and {raft_file_name}")


if __name__ == "__main__":
    generate_server_lists()