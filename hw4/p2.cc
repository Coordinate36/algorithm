#include <iostream>
#include <algorithm>
#include <vector>
#include <string>

using namespace std;

int m;
int len1; 
int len2;

int minPenalty(vector<int> &gapPenalty, vector<vector<int>> &dismatchPenalty, vector<int> &s1, vector<int> &s2) {
    vector<vector<int>> dp(len1 + 1, vector<int>(len2 + 1, 0));
    for (int i = 1; i <= len1; ++i) {
        dp[i][0] = dp[i - 1][0] + gapPenalty[s1[i - 1]];
    }
    for (int j = 1; j <= len2; ++j) {
        dp[0][j] = dp[0][j - 1] + gapPenalty[s2[j - 1]];
    }
    for (int i = 1; i <= len1; ++i) {
        for (int j = 1; j <= len2; ++j) {
            dp[i][j] = min(dp[i - 1][j - 1] + dismatchPenalty[s1[i - 1]][s2[j - 1]], min(dp[i - 1][j] + gapPenalty[s1[i - 1]], dp[i][j - 1] + gapPenalty[s2[j - 1]]));
        }
    }
    return dp[len1][len2];
}

int main() {
    while (cin >> m) {
        vector<int> gapPenalty(m);
        vector<vector<int>> dismatchPenalty(m, vector<int>(m));
        for (int i = 0; i < m; ++i) {
            cin >> gapPenalty[i];
        }
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < m; ++j) {
                cin >> dismatchPenalty[i][j];
            }
        }
        cin >> len1;
        vector<int> s1(len1);
        for (int i = 0; i < len1; ++i) {
            cin >> s1[i];
        }
        cin >> len2;
        vector<int> s2(len2);
        for (int i = 0; i < len2; ++i) {
            cin >> s2[i];
        }
        int ans = minPenalty(gapPenalty, dismatchPenalty, s1, s2);
        cout << "minimum penalties: " << ans << endl;
    }
}