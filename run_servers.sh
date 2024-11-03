./server --id 1 &
./server --id 2 &
./server --id 3 &
sleep 20
./load_balancer &