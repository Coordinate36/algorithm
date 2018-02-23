# Algorithm hw2-3
 
学号： U201514632 
姓名：王宏斌


## Parallel MST Algorithms

1. Boruvka algorithm:
```
Input: A graph G whose edges have distinct weights
 Initialize a forest F to be a set of one-vertex trees, one for each vertex of the graph.
 While F has more than one component:
   Find the connected components of F and label each vertex of G by its component
   Initialize the cheapest edge for each component to "None"
   For each edge uv of G:
     If u and v have different component labels:
       If uv is cheaper than the cheapest edge for the component of u:
         Set uv as the cheapest edge for the component of u
       If uv is cheaper than the cheapest edge for the component of v:
         Set uv as the cheapest edge for the component of v
   For each component whose cheapest edge is not "None":
     Add its cheapest edge to F
  Output: F is the minimum spanning forest of G.
```
2.  red-blue algorithm
```
Red rule: Let C be a cycle with no red edges. Select an uncolored edge of C
of max weight and color it red.
Blue rule: Let D be a cutset with no blue edges. Select an uncolored edge in
D of min weight and color it blue.
let n be the number of vertices of G
while the length of blue rules is less than n - 1:
	apply red and blue rule
output: the blue edges form a MST

```

## 

#### 分析
```
完成尽可能多的任务数，那么只要每次都选结束时间最小的就行，采用贪心算法。
```
#### 代码
```
#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

struct Task {
    double begin;
    double end;
};

void schedule(vector<struct Task> &tasks, int n) {
    sort(tasks.begin(), tasks.end(), [](auto &task1, auto &task2) {
        return task1.end < task2.end;
    });

    vector<struct Task> res;
    res.push_back(tasks[0]);
    for (int i = 1; i < n; ++i) {
        if (res.back().end <= tasks[i].begin) {
            res.push_back(tasks[i]);
        }
    }

    for (int i = 0; i < res.size(); ++i) {
        cout << res[i].begin << "    " << res[i].end << endl;
    }
}

int main() {
    int n;
    vector<struct Task> tasks;
    for (;;) {
        cin >> n;
        tasks.resize(n);

        for (int i = 0; i < n; ++i) {
            cin >> tasks[i].begin >> tasks[i].end;
        }
        cout << endl;
        schedule(tasks, n);
    }

    return 0;
}
```
#### 测试样例
input1:
```
4
0.1 0.2
0.2 0.3
0.3 0.4
0.4 0.5
```
output1:
```
0.1    0.2
0.2    0.3
0.3    0.4
0.4    0.5
```
input2:
```
4
0.1  0.2
0.15 0.16
0.17 0.5
0.4  0.8
```
output2:
```
0.15    0.16
0.17    0.5
```
input3:
```
3
0.1 0.4
0.3 0.5
0.4 0.7
```
output3:
```
0.1    0.4
0.4    0.7
```
input4:
```
5
0.01 0.1
0.9  0.99
0.05 0.12
0.88 0.95
0.11 0.89
```
output4:
```
0.01    0.1
0.11    0.89
0.9    0.99
```

## Interval schedule2

#### 分析
```
加权区间调度问题，显然dp算法，对任务按结束时间排序，以所有任务中最晚的结束时间为N，定义长度为N的状态数组，则dp[tasks[i].end] = max(dp[tasks[i].begin] + tasks[i].weight, dp[tasks[i].end - 1])
```
#### 代码
```cpp
#include <iostream>
#include <algorithm>
#include <vector>
#include <deque>

using namespace std;

struct Task {
    int begin;
    int end;
    int weight;
};

void schedule(vector<struct Task> &tasks, int n) {
    sort(tasks.begin(), tasks.end(), [](auto &task1, auto &task2) {
        return task1.end < task2.end;
    });

    int len = tasks[n - 1].end + 1;

    vector<int> path(len, -1);
    vector<int> last_task(len, -1);
    vector<int> dp(len, 0);
    
    deque<int> res;

    for (int i = 0; i < n; ++i) {
        if (dp[tasks[i].begin] + tasks[i].weight > dp[tasks[i].end]) {
            dp[tasks[i].end] = dp[tasks[i].begin] + tasks[i].weight;
            path[tasks[i].end] = tasks[i].begin;
            last_task[tasks[i].end] = i;
        }
        if (i + 1 < n) {
            for (int j = tasks[i].end + 1; j <= tasks[i + 1].end; ++j) {
                dp[j] = dp[j - 1];
                path[j] = path[j - 1];
                last_task[j] = last_task[j - 1];
            }
        }
    }

    for (int i = len - 1; last_task[i] != -1; i = path[i]) {
        res.push_front(last_task[i]);
    }

    for (int i = 0; i < res.size(); ++i) {
        cout << tasks[res[i]].begin << "  " << tasks[res[i]].end << endl;
    }
}

int main() {
    int n;
    vector<struct Task> tasks;
    for (;;) {
        cin >> n;
        tasks.resize(n);

        for (int i = 0; i < n; ++i) {
            cin >> tasks[i].begin >> tasks[i].end >> tasks[i].weight;
        }
        cout << endl;
        schedule(tasks, n);
        cout << endl << endl;
    }

    return 0;
}
```
#### 测试样例
input1:
```
3
1 3 1000
2 4 2001
3 5 1000
```
output1:
```
2  4
```
input2:
```
5
1 2 201
2 4 200
3 5 201
1 7 401
7 8 233
```
output2:
```
1  2
3  5
7  8
```
input3:
```
4
1 2 100
1 3 100
2 4 1
3 5 1
```
output3:
```
1  2
2  4
```
input4:
```
5
2 3 10
1 2 10
4 5 10
3 4 10
5 6 10
```
output4:
```
1  2
2  3
3  4
4  5
5  6
```

