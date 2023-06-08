/* RESPONSI 2 STRUKTUR DATA
 * PAKET 3 QUEUE
 */

#include<iostream>
#define MAXQUEUE 99

using namespace std;

// Struktur data Queue
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

// Method inisialisasi
void inisialisasi(QUEUE *Q, int max) {
    Q->count = 0;
    Q->front = 0;
    Q->rear = 0;
    Q->max = max;
}

// Method cek apakah queue kosong
bool isEmpty(QUEUE *Q) {
    if (Q->count == 0) {
        return true;
    }
    return false;
}

// Method cek apakah queue penuh
bool isFull(QUEUE *Q) {
    if (Q->count == Q->max) {
        return true;
    }
    return false;
}

// Method memasukkan data ke queue
void enqueue(QUEUE *Q, int data) {
    Q->item[Q->rear] = data;
    Q->rear = (Q->rear + 1) % Q->max;
    ++(Q->count);
}

// Method mengambil data dari queue
int dequeue(QUEUE *Q) {
    int data = Q->item[Q->front];
    Q->front = (Q->front + 1) % Q->max;
    --(Q->count);
    
    return data;
}

// Method cek umur pengunjung
bool cekUmur(int data) {
    if (data > 13) {
        return true;
    }
    return false;
}

// Method masuk ke rumah hantu
void masukRumahHantu(int data) {
    /* == TODO ==
     * Lengkapi method ini.
     * 1. Method ini menerima satu argumen berupa usia pengunjung
     * 2. Masukkan pengunjung ke antrian dalam apabila belum penuh dan usia memenuhi syarat
     * 3. Apabila antrian dalam penuh, maka pengunjung dimasukkan ke antrian luar
     */
}

// Method keluar dari rumah hantu
void keluarRumahHantu(int jmlOrang) {
    while (jmlOrang > 0) {
        /* == TODO ==
         * Lengkapi blok kode while loop ini.
         * 1. Blok kode ini mengeluarkan pengunjung dari antrian dalam sebanyak jmlOrang
         * 2. Apabila antrian luar tidak kosong, masukkan pengunjung dari antrian luar ke antrian dalam
         * 3. Update variabel global total_orang untuk mendata jumlah orang yang sudah masuk dan keluar rumah hantu
         */
    }
}


int main() {
    /* == PERINGATAN ==
     * Kamu tidak perlu mengubah isi method main ini.
     * Kamu hanya perlu mengisi blok kode yang bertuliskan TODO saja.
     */

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
