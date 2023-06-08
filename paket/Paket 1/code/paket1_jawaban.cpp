#include <iostream>
#define MAXQUEUE 5
using namespace std;

typedef struct {
    string nama;
    int berat;
} Person;

typedef struct {
    int count;
    int front;
    int rear;
    Person listPerson[5];
} Queue;

Queue antrian;

void init(Queue *Q) {
    Q->count = 0;
    Q->front = 0;
    Q->rear = 0;
}

bool isFull(Queue *Q) {
    if (Q->count == MAXQUEUE) {
        return true;
    }
    return false;
}

void insert(Queue *Q, Person data) {
    if (isFull(&antrian)) {
        cout << data.nama << " tidak mendapatkan antrian" << endl;
    } else {
        if (data.berat < 55) {
            Q->listPerson[Q->rear] = data;
            Q->rear = (Q->rear + 1) % MAXQUEUE;
            ++(Q->count);
        } else {
            cout << "Berat badan " << data.nama << " tidak memenuhi syarat"
                 << endl;
        }
    }
}

main() {
    Person orang;
    init(&antrian);

    int n;
    cin >> n;

    for (int i = 0; i < n; i++) {
        cin >> orang.nama;
        cin >> orang.berat;
        insert(&antrian, orang);
    }

    for (int i = 0; i < 5; i++) {
        cout << antrian.listPerson[i].nama << " " << antrian.listPerson[i].berat
             << endl;
    }
}