#include<iostream>
#define MAXQUEUE 99

using namespace std;

typedef struct {
    int count;
    int front;
    int rear;
    int max;
    int item[MAXQUEUE];
} QUEUE;

QUEUE antrian_dalam;
QUEUE antrian_luar;
int total_orang = 0;

void inisialisasi(QUEUE *Q, int max) {
    Q->count = 0;
    Q->front = 0;
    Q->rear = 0;
    Q->max = max;
}

bool isEmpty(QUEUE *Q) {
    if (Q->count == 0) {
        return true;
    }
    return false;
}

bool isFull(QUEUE *Q) {
    if (Q->count == Q->max) {
        return true;
    }
    return false;
}

void enqueue(QUEUE *Q, int data) {
    Q->item[Q->rear] = data;
    Q->rear = (Q->rear + 1) % Q->max;
    ++(Q->count);
}

int dequeue(QUEUE *Q) {
    int data = Q->item[Q->front];
    Q->front = (Q->front + 1) % Q->max;
    --(Q->count);
    
    return data;
}

bool cekUmur(int data) {
    if (data > 13) {
        return true;
    }
    return false;
}

void masukRumahHantu(int data) {
    if (isFull(&antrian_dalam)) {
        enqueue(&antrian_luar, data);
    } else {
        if (cekUmur(data)) {
            enqueue(&antrian_dalam, data);
        }
    }
}

void keluarRumahHantu(int jmlOrang) {
    while (jmlOrang > 0) {
        if (isEmpty(&antrian_dalam)) break;
        dequeue(&antrian_dalam);
        total_orang++;
        if (!isEmpty(&antrian_luar)) {
            masukRumahHantu(dequeue(&antrian_luar));
        }
        jmlOrang--;
    }
}


int main() {
    inisialisasi(&antrian_dalam, 3);
    inisialisasi(&antrian_luar, 99);
    
    string input;
    cin >> input;
    while (input != "x") {
        if (input == "Masuk") {
            int usia;
            cin >> usia;
            masukRumahHantu(usia);
        } else if (input == "Keluar") {
            int jmlOrang;
            cin >> jmlOrang;
            keluarRumahHantu(jmlOrang);
        }
        cin >> input;
    }

    cout << total_orang << endl;
}
