#include <bits/stdc++.h>
using namespace std;

typedef struct {
    int count;
    int front;
    int rear;
    int list[150];
} Queue;

Queue antrian;

void init(Queue *Q) {
    Q->count = 0;
    Q->front = 0;
    Q->rear = 0;
}

void add(Queue *Q, int num, int count) {
    // fungsi untuk menambahkan data
}

void del(Queue *Q, int count) {
    // fungsi untuk menghapus data
}

void print(Queue *Q) {
    cout << Q->count << "\n";
    for(int i = Q->front; i < Q->rear; i++) {
        cout << Q->list[i] << " "; 
    }
    cout << "\n";
}

main() {
    init(&antrian);

    int n;
    cin >> n;

    for (int i = 0; i < n; i++) {
        int num, count;
        string type;
        cin >> type;
        if(type == "add") {
            cin >> num >> count;
            add(&antrian, num, count);
        } else {
            cin >> count;
            del(&antrian, count);
        }
    }

    // output
    print(&antrian);
    return 0;
}