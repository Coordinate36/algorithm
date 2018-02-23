#include <iostream>
#include <cstring>
#include <map>
#include <set>
#include <vector>

using namespace std;

vector<size_t> matching;
vector<size_t> check;

vector<size_t> stud;
vector<size_t> prof;
vector<size_t> ts;

set<pair<size_t, size_t>> prof_time;
set<pair<size_t, size_t>> stud_time;

vector<vector<size_t>> adj;

bool dfs(size_t u) {
    for (auto v: adj[u]) {
        if (check[v] == 0) {
            check[v] = true;
            if (matching[v] == -1 || dfs(matching[v])) {
                matching[v] = u;
                return true;
            }
        }
    }

    return false; 
}

size_t hungarian() {
    size_t left_num = stud_time.size();
    size_t node_num = left_num + prof_time.size();
    matching.resize(node_num);
    check.resize(node_num);
    memset(matching.data(), -1, node_num * sizeof(size_t));
    size_t count = 0;
    for (size_t i = 0; i < left_num; ++i) {
        memset(check.data(), 0, node_num * sizeof(size_t));
        if (dfs(i)) {
            ++count;
        }
    }
    return count;
}

void mkGraph(size_t p) {
    map<pair<size_t, size_t>, size_t> stud_map;
    map<pair<size_t, size_t>, size_t> prof_map;
    size_t i = 0;
    for (auto iter = stud_time.begin(); iter != stud_time.end(); ++iter) {
        stud_map[*iter] = i++;
    } 
    for (auto iter = prof_time.begin(); iter != prof_time.end(); ++iter) {
        prof_map[*iter] = i++;
    }
    adj.resize(i);
    for (i = 0; i < p; ++i) {
        auto stud_pair = make_pair(stud[i], ts[i]);
        auto prof_pair = make_pair(prof[i], ts[i]);
        adj[stud_map[stud_pair]].push_back(prof_map[prof_pair]);
        adj[prof_map[prof_pair]].push_back(stud_map[stud_pair]);
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
        for (size_t i = 0; i < p; ++i) {
            cin >> stud[i] >> ts[i] >> prof[i];
            stud_time.insert(make_pair(stud[i], ts[i]));
            prof_time.insert(make_pair(prof[i], ts[i]));
        }
        mkGraph(p);
        size_t count = hungarian();
        cout << count << endl;
    }
    return 0;
}