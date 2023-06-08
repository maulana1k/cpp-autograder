/*

        !!! Perhatian !!!
* Sebelum mengerjaiakn baca doa
* Kerjakan hanya pada method saja
* Dilarang merubah main
* Good Luck all!!!

*/

#include <iostream>
#include <string>
using namespace std;
#define MAXQUEUE 3

typedef struct
{
    /*
    TODO : Lengkapi struct queue dengan ketentuan sebagai berikut
    * - rear (int)
    * - front (int)
    * - count (int)
    * - orang (string)
    */
} QUEUE;

QUEUE antrian;

void inisialisasi(QUEUE *Q)
{
    // TODO : Implementasi method inisialisasi
}

int isFull(QUEUE *Q)
{
    // TODO : Implementasi method isFull
}

void enqueue(QUEUE *Q, string data)
{
    // TODO : Implementasi method enqueue dengan output yang sesuai dengan soal
}
void dequeue(QUEUE *q)
{
    // TODO : Implementasi method dequeue dengan output yang sesuai dengan soal
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
        cout << "Masukkan Nama : ";
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