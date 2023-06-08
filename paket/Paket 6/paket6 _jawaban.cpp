#include <iostream>
using namespace std;

int faktorial(int angka)
{
    if (angka < 0)
    {
        return 0;
    }
    if ((angka == 0) || (angka == 1))
    {
        return 1;
    }
    else
    {
        return (angka * faktorial(angka - 1));
    }
}

int pangkat(int basis, int eksponen)
{
    if (eksponen == 0)
    {

        return 1;
    }
    else
    {

        return basis * pangkat(basis, eksponen - 1);
    }
}

main()
{
    int a, b, c;
    cin >> a;
    cout << "Hasil Faktorial : " << faktorial(a) << endl;
    cin >> b >> c;
    cout << "Hasil Pangkat : " << pangkat(b, c);
    return 0;
}