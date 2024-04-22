#include <iostream>

using namespace std;

const int N = 100010;
int a[N];

int TopK(int l, int r, int k)
{
    if (l == r)
        return a[l];
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
    if (j - l + 1 >= k)
        return TopK(l, j, k);
    return TopK(j + 1, r, k - (j - l + 1));
}

int main()
{
    int n, k;
    cin >> n >> k;
    for (int i = 0; i < n; i++)
        cin >> a[i];
    cout << TopK(0, n - 1, k);
    return 0;
}