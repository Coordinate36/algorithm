#include <iostream>
#include <algorithm>
#include <cstring>

using namespace std;

const int maxn = 1000;

struct Edge {
    int x;
    int y;
    int weight;
} edges[maxn];

int parent[maxn];
int depth[maxn];

void makeSet() {
    memset(parent, -1, sizeof(parent));
    memset(depth, 0, sizeof(depth));
}

int find(int x) {
    if (parent[x] == -1) {
        return x;
    }
    parent[x] = find(parent[x]);
    return parent[x];
}

void Union(int x, int y) {
    int fx = find(x);
    int fy = find(y);

    if (depth[fx] > depth[fy]) {
        parent[fy] = fx;
    } else {
        parent[fx] = fy;
        if (depth[fx] == depth[fy]) {
            ++depth[fy];
        }
    }
}

int main() {
    int n, m;
    cin >> n >> m;
    for (int i = 0; i < m; ++i) {
        cin >> edges[i].x >> edges[i].y >> edges[i].weight;
    }
    sort(edges, edges + m, [](struct Edge a, struct Edge b) {
        return a.weight < b.weight;
    });
    makeSet();
    for (int i = 0; i < m; ++i) {
        int fx = find(edges[i].x);
        int fy = find(edges[i].y);
        if (fx != fy) {
            cout << edges[i].x << " -> " << edges[i].y << " : " << edges[i].weight << endl;
            Union(edges[i].x, edges[i].y);
        }
    }
}