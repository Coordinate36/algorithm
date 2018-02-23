#include <iostream>
#include <climits>
#include <vector>

using namespace std;

struct matrix {
    int row;
    int col;
};

int matrixChain(vector<struct matrix> &matrixes) {
    vector<vector<int>> dp(matrixes.size(), vector<int>(matrixes.size(), INT_MAX));
    for (int i = 0; i < matrixes.size(); ++i) {
        dp[i][i] = 0;
    }

    for (int r = 2; r <= matrixes.size(); ++r) {
        for (int i = 0; i + r - 1 < matrixes.size(); ++i) {
            int j = i + r - 1;
            for (int k = i; k < j; ++k) {
                dp[i][j] = min(dp[i][j], dp[i][k] + dp[k + 1][j] + matrixes[i].row * matrixes[k].col * matrixes[j].col);
            }
        }
    }
    return dp[0][matrixes.size() - 1];
}

int main() {
    int n;
    while (cin >> n) {
        vector<struct matrix> matrixes(n);
        for (int i = 0; i < n; ++i) {
            cin >> matrixes[i].row >> matrixes[i].col;
        }
        int ans = matrixChain(matrixes);
        cout << "minimum scalar multiplications: " << ans << endl << endl;
    }
    return 0;
}