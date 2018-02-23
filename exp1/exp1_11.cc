#include <iostream>
#include <climits>
#include <cstring>
#include <set>
#include <map>
#include <string>
#include <vector>
#include <queue>

using namespace std;

vector<size_t> parent;

vector<size_t> stud;
vector<size_t> prof;
vector<size_t> ts;

set<string> prof_time;
set<string> stud_time;

vector<vector<size_t>> adj;
vector<map<size_t, size_t>> rGraph;

bool bfs() {
    size_t t = adj.size() - 1; // t -> sink, 0 -> source
    vector<size_t> visited(t + 1, 0);
    queue<size_t> Q;
    Q.push(0);
    while (!Q.empty()) {
        size_t u = Q.front();
        Q.pop();
        for (auto v: adj[u]) {
            if (visited[v] == 0 && rGraph[u][v] > 0) {
                cout << u << " " << v << endl;
                Q.push(v);
                visited[v] = 1;
                parent[v] = u;
            }
        }
    }

    return (visited[t] == 1);
}

int fordFulkerson() {
    size_t t = adj.size() - 1;
    size_t s = 0;

    int max_flow = 0;

    while (bfs()) {
        size_t path_flow = INT_MAX;
        size_t u;
        for (auto v = t; v != s; v = parent[v]) {
            u = parent[v];
            path_flow = min(path_flow, rGraph[u][v]);
        }

        for (auto v = t; v != s; v = parent[v]) {
            u = parent[v];
            rGraph[u][v] -= path_flow;
            rGraph[v][u] += path_flow;
        }
        max_flow += path_flow;
    }

    return max_flow;
}

void mkGraph(size_t p) {
    map<string, size_t> stud_map;
    map<pair<size_t, size_t>, size_t> prof_map;

    size_t m = 1;
    for (auto iter = stud_time.begin(); iter != stud_time.end(); ++iter) {
        stud_map[*iter] = m++;
    }
    size_t n = m;
    for (auto iter = prof_time.begin(); iter != prof_time.end(); ++iter) {
        prof_map[*iter] = n++;
    }
    adj.resize(n + 1);
    parent.resize(n + 1);
    rGraph.resize(n + 1);
    for (size_t i = 0; i < p; ++i) {
        auto stud_pair = make_pair(stud[i], ts[i]);
        auto prof_pair = make_pair(prof[i], ts[i]);
        adj[stud_map[stud_pair]].push_back(prof_map[prof_pair]);
        rGraph[stud_map[stud_pair]][prof_map[prof_pair]] = 1;
    }
    for (size_t i = 1; i < m; ++i) {
        adj[0].push_back(i);
        rGraph[0][i] = 1;
    }
    for (size_t i = m; i < n; ++i) {
        adj[i].push_back(n);
        rGraph[i][n] = 1;
    }
}

int main() {
    size_t p, n, m, t;
    while (cin >> p >> n >> m >> t) {
        stud.resize(p);
        prof.resize(p);
        ts.resize(p);
        prof_time.clear();
        stud_time.clear();
        adj.clear();
        parent.clear();
        rGraph.clear();

        for (size_t i = 0; i < p; ++i) {
            cin >> stud[i] >> ts[i] >> prof[i];
            char stud_time_str[20];
            char prof_time_str[20];
            sprintf("%d:%d", )
            stud_time.insert(make_pair(stud[i], ts[i]));
            prof_time.insert(make_pair(prof[i], ts[i]));
        }
        mkGraph(p);
        size_t count = fordFulkerson();
        cout << count << endl;
    }
    return 0;
}