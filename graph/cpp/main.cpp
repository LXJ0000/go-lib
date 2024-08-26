#include <iostream>
#include <queue>
#include <cstring>
#include <vector>
#include <algorithm>

using namespace std;

typedef pair<int, int> PII;

const int N = 150010;
int n, m;
int h[N], e[N], ne[N], w[N], idx;
int st[N], dist[N];

void add(int a, int b, int c) {
    e[idx] = b, w[idx] = c, ne[idx] = h[a], h[a] = idx ++;
}

int Dijkstra() {
    priority_queue<PII, vector<PII>, greater<PII>> heap;
    memset(dist, 0x3f, sizeof dist);
    heap.push({0, 1});
    dist[1] = 0;
    while(heap.size()) {
        auto t = heap.top();
        heap.pop();
        int ver = t.second, distance = t.first;
        if(st[ver]) continue;
        st[ver] = true;
        for(int i = h[ver]; i != -1; i = ne[i]) {
            int j = e[i];
            if(dist[j] > distance + w[i]) {
                dist[j] = distance + w[i];
                heap.push({dist[j], j});
            }
        }
    }
    return dist[n];
}

int main() {
    memset(h, -1, sizeof h);
    cin >> n >> m;
    while(m --) {
        int a, b, c;
        cin >> a >> b >> c;
        add(a, b, c); // 无向图需要再添加一次
    }
    int ret = Dijkstra();
    cout << (ret == 0x3f3f3f3f ? -1 : ret) << endl;
    return 0;
}