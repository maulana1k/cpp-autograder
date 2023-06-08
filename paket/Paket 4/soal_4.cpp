/* RESPONSI 2 STRUKTUR DATA
 * PAKET 4 BUBBLE SORT
 */

#include <iostream>

using namespace std;

// Method bubble sort
void bubble_sort(char a[], int size) {
    /* == TODO ==
     * Lengkapi method bubble sort ini.
     * 1. Method ini menerima dua buah argumen, yaitu sebuah array berisi char dan size dari array tsb
     * 2. Method ini melakukan sorting pada array a[] sehingga urut secara ascending menggunakan teknik bubble sort
     */
}


int main(void) {
    /* == PERINGATAN ==
     * Kamu tidak perlu mengubah isi method main ini.
     * Kamu hanya perlu mengisi blok kode yang bertuliskan TODO saja.
     */

    int n;
    cin >> n;
    char deret[n];
    for (int i = 0; i < n; i++) {
        cin >> deret[i];
    }

    bubble_sort(deret, n);

    for (int i = 0; i < n; i++) {
      cout << deret[i];
      if (i != n-1) cout << " ";
    }
    cout << endl;
}
