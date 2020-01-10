#include <iostream>
#include <vector>

using namespace std;

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
        for (int i = 0; i <= n; i++) {
            for (int j = 0; j <= m; j++) {
                cout<<a[i][j]<<" ";
            }
            cout<<endl;
        }
        return a[n][m];
    }
};

int main() {
    string A = "bac";
    string B = "cbbbc";
    MinCost l;
    cout << l.findMinCost(A, 3, B, 5, 8, 3, 4);
    return 0;
}