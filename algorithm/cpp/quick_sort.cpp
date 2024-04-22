#include <iostream>

using namespace std;

const int N = 100010;
int a[N];

void quickSort(int l, int r)
{
    if (l >= r)
        return;
    int i = l - 1, j = r + 1, mid = a[l + r >> 1];
    while (i < j)
    {
        do
            i++;
        while (a[i] < mid);
        do
            j--;
        while (a[j] > mid);
        if (i < j)
            swap(a[i], a[j]);
    }
    quickSort(l, j);
    quickSort(j + 1, r);
}

int main()
{
    int n;
    cin >> n;
    for (int i = 0; i < n; i++)
        cin >> a[i];
    quickSort(0, n - 1);
    for (int i = 0; i < n; i++)
        cout << a[i] << ' ';
    return 0;
}