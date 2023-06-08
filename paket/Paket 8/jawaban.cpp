#include <bits/stdc++.h>
using namespace std;

long long memory[50];

int fibonacci(int n) {
    if(memory[n] != -1) {
        return memory[n];
    }
    if (n <= 1) {
        return n;
    } else {
        memory[n] = fibonacci(n - 1) + fibonacci(n - 2);
        return memory[n];
    }
}

main() {
    memset(memory, -1, sizeof(memory));
    int n;
    cin >> n;
    while(n--) {
        int x;
        cin >> x;
        cout << fibonacci(x) << "\n";
    }
    return 0;
}