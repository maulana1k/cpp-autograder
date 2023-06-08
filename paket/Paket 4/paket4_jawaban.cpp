#include <iostream>

using namespace std;

void bubble_sort(char a[], int size) {
    int switched = 1;
    char hold;
    int i = 0;
    int j = 0;
    size -= 1;
    for(i = 0; i < size && switched; i++) {
        switched = 0;
        for(j = 0; j < size - i; j++) {
            if(a[j] > a[j+1]) {
                switched = 1;
                hold = a[j];
                a[j] = a[j + 1];
                a[j + 1] = hold;
            }
        }
    }
}

int main(void) {
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
