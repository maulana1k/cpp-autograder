#include <bits/stdc++.h>
using namespace std;

long long memory[50];

int fibonacci(int n) {
    // isi fungsi fibonacci 
}

main() {
    memset(memory, -1, sizeof(memory));
    int n;
    cin >> n;
    while(n--) {
        int x;
        cin >> x;
        cout << fibonacci(x) << '\n';
    }
    return 0;
}