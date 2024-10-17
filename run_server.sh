./server --id=1 &
./server --id=2 &
./server --id=3 &
./server --id=4 &
./server --id=5 &

sleep 5

./load_balancer &