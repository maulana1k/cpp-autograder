#include <iostream>
#include <string>
using namespace std;
#define MAXQUEUE 3

typedef struct
{
    int count;
    int front;
    int rear;
    string orang[MAXQUEUE];
} QUEUE;

QUEUE antrian;

void inisialisasi(QUEUE *Q)
{
    Q->count = 0;
    Q->front = 0;
    Q->rear = 0;
}

int isFull(QUEUE *Q)
{
    int hasil = 0;
    if (Q->count == MAXQUEUE)
    {
        hasil = 1;
    }
    return hasil;
}

void enqueue(QUEUE *Q, string data)
{
    if (isFull(&antrian) == 1)
    {
        cout << "=== Maaf antrian penuh ===" << endl;
    }
    else
    {
        Q->orang[Q->rear] = data;
        Q->rear = (Q->rear + 1) % MAXQUEUE;
        cout << "=== " << data << " masuk antrian ===" << endl;
        ++(Q->count);
    }
}
void dequeue(QUEUE *q)
{
    string temp;
    if (q->count == 0)
    {
        cout << "Maaf antrian kosong" << endl;
    }
    else
    {
        temp = q->orang[q->front];
        q->orang[q->front] = "";
        q->front += 1 % MAXQUEUE;
        cout << temp << " telah keluar antrian" << endl;
        --(q->count);
    }
}
void cetak(QUEUE *Q)
{
    int depan = Q->front;
    for (int i = 0; i < Q->count; i++)
    {
        cout << i + 1 << ". " << Q->orang[depan] << endl;
        depan = (depan + 1) % MAXQUEUE;
        if (Q->orang[depan] == "")
        {
            cout << i + 2 << ". " << endl;
        }
    }
    cout << endl;
}

main()
{
    string nama;
    inisialisasi(&antrian);
    cout << "Jumlah Maksimum Antrian : " << MAXQUEUE << endl;
    for (int i = 0; i < MAXQUEUE + 1; i++)
    {
        cin >> nama;
        enqueue(&antrian, nama);
    }
    cetak(&antrian);
    dequeue(&antrian);
    cetak(&antrian);
    if (antrian.orang[MAXQUEUE] == "")
    {

        cout << "Menambahkan " << nama << " ke dalam antrian..." << endl;
        enqueue(&antrian, nama);
    }
    cetak(&antrian);
}