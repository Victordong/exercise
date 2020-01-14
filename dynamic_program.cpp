//
// Created by zhandong on 2020-01-09.
//
#include <vector>
#include <string>

using namespace std;

class Exchange {
public:
    int countWays(vector<int> penny, int n, int aim) {
        // write code here
        int value[51][1001];
        for (int i = 0; i < n; i++) {
            value[i][0] = 1;
        }
        for (int j = 1; j <= aim; j++) {
            if (j % penny[0] == 0) {
                value[0][j] = 1;
            } else {
                value[0][j] = 0;
            }
        }
        for (int i = 1; i < n; i++) {
            for (int j = 1; j <= aim; j++) {
                if (j >= penny[i]) {
                    value[i][j] = value[i - 1][j] + value[i][j - penny[i]];
                } else {
                    value[i][j] = value[i - 1][j];
                }
            }
        }
        return value[n - 1][aim];
    }
};


class GoUpstairs {
public:
    int countWays(int n) {
        int a[100000];
        a[0] = 0;
        a[1] = 1;
        a[2] = 2;
        if (n <= 2) {
            return a[n];
        }
        for (int i = 3; i <= n; i++) {
            a[i] = (a[i - 2] + a[i - 1]) % 1000000007;
        }
        return a[n];
    }
};


class MinimumPath {
public:
    int getMin(vector<vector<int> > map, int n, int m) {
        int a[101][101];
        a[0][0] = map[0][0];
        for (int i = 1; i < m; i++) {
            a[0][i] = a[0][i - 1] + map[0][i];
        }
        for (int i = 1; i < n; i++) {
            a[i][0] = a[i - 1][0] + map[i][0];
        }
        for (int i = 1; i < n; i++) {
            for (int j = 1; j < m; j++) {
                a[i][j] = min(a[i - 1][j], a[i][j - 1]) + map[i][j];
            }
        }
        return a[n - 1][m - 1];
    }
};

class LongestIncreasingSubsequence {
public:
    int getLIS(vector<int> A, int n) {
        int a[501];
        a[0] = 1;
        for (int i = 1; i < n; i++) {
            int max = 0;
            for (int j = 0; j < i; j++) {
                if (A[i] > A[j] && max < a[j]) {
                    max = a[j];
                }
            }
            a[i] = max + 1;
        }
        int max = 0;
        for (int i = 0; i < n; i++) {
            if (max < a[i]) {
                max = a[i];
            }
        }
        return max;
    }
};

class LCS {
public:
    int findLCS(string A, int n, string B, int m) {
        // write code here
        int a[300][300];
        bool repeat = false;
        for (int i = 0; i < n; i++) {
            if (A[i] == B[0]) {
                repeat = true;
            }
            if (!repeat) {
                a[i][0] = 0;
            } else {
                a[i][0] = 1;
            }
        }
        repeat = false;
        for (int i = 0; i < m; i++) {
            if (B[i] == A[0]) {
                repeat = true;
            }
            if (!repeat) {
                a[0][i] = 0;
            } else {
                a[0][i] = 1;
            }
        }
        for (int i = 1; i < n; i++) {
            for (int j = 1; j < m; j++) {
                a[i][j] = max(max(a[i][j - 1], a[i - 1][j]), a[i - 1][j - 1] + 1);
            }
        }
        return a[n - 1][m - 1];
    }
};

class Backpack {
public:
    int maxValue(vector<int> w, vector<int> v, int n, int cap) {
        int a[1000][1000];
        for (int i = 0; i < n; i++) {
            a[i][0] = 0;
        }
        for (int i = 0; i < cap; i++) {
            if (i >= w[0]) {
                a[0][i] = v[0];
            } else {
                a[0][i] = 0;
            }
        }
        for (int i = 1; i < n; i++) {
            for (int j = 1; j <= cap; j++) {
                if (j >= w[i]) {
                    a[i][j] = max(a[i - 1][j], a[i - 1][j - w[i]] + v[i]);
                } else {
                    a[i][j] = a[i - 1][j];
                }
            }
        }
        return a[n - 1][cap];
    }
};

class MinCost {
public:
    int findMinCost(string A, int n, string B, int m, int c0, int c1, int c2) {
        // write code here
        int a[1000][1000];
        for (int i = 0; i <= n; i++) {
            a[i][0] = c1 * i;
        }
        for (int j = 0; j <= m; j++) {
            a[0][j] = c0 * j;
        }
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                int v1 = a[i][j - 1] + c0;
                int v2 = a[i - 1][j] + c1;
                int v3;
                if (A[i - 1] == B[j - 1]) {
                    v3 = a[i - 1][j - 1];
                } else {
                    v3 = a[i - 1][j - 1] + c2;
                }
                a[i][j] = min(min(v1, v2), v3);
            }
        }
        return a[n][m];
    }
};


class LongestPalindrome {
public:
    string longestPalindrome(string s) {

    }
};

class UniquePaths {
public:
    int uniquePaths(int m, int n) {
        int a[101][101];
        for (int i = 0; i < m; i++) {
            a[i][0] = 1;
        }
        for (int i = 0; i < n; i++) {
            a[0][i] = 1;
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                a[i][j] = a[i - 1][j] + a[i][j - 1];
            }
        }
        return a[m - 1][n - 1];
    }
};

class UniquePathsWithObstacles {
public:
    int uniquePathsWithObstacles(vector<vector<int>> &obstacleGrid) {
        int m = obstacleGrid.size();
        int n = obstacleGrid[0].size();
        int a[101][101];
        bool can = true;
        for (int i = 0; i < m; i++) {
            if (obstacleGrid[i][0] == 1) {
                can = false;
            }
            if (can) {
                a[i][0] = 1;
            } else {
                a[i][0] = 0;
            }
        }
        can = true;
        for (int i = 0; i < m; i++) {
            if (obstacleGrid[0][i] == 1) {
                can = false;
            }
            if (can) {
                a[0][i] = 1;
            } else {
                a[0][i] = 0;
            }
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (obstacleGrid[i][j] == 1) {
                    a[i][j] = 0;
                } else {
                    a[i][j] = a[i - 1][j] + a[i][j - 1];
                }
            }
        }
        return a[m - 1][n - 1];
    }
};

class Solution {
public:
    int minPathSum(vector<vector<int>> &grid) {
        int m = grid.size();
        int n = grid[0].size();
        for (int i = 1; i < m; i++) {
            grid[i][0]+=grid[i-1][0];
        }
        for (int i = 1; i < n; i++) {
            grid[0][i]+=grid[0][i-1];
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                int total = min(grid[i][j-1], grid[i-1][j]);
                grid[i][j] = total+grid[i][j];
            }
        }
        return grid[m - 1][n - 1];
    }
};
