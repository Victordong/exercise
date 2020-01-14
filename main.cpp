#include <iostream>
#include <vector>
#include <stack>
#include <queue>
using namespace std;

class Solution {
public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
        int m = obstacleGrid.size();
        int n = obstacleGrid[0].size();
        int a[101][101];
        bool can = true;
        for(int i=0;i<m;i++) {
            if(obstacleGrid[i][0]==1) {
                can = false;
            }
            if(can){
                a[i][0] = 1;
            } else {
                a[i][0] = 0;
            }
        }
        can = true;
        for(int i=0;i<m;i++) {
            if(obstacleGrid[0][i]==1) {
                can = false;
            }
            if(can){
                a[0][i] = 1;
            } else {
                a[0][i] = 0;
            }
        }
        for(int i=1;i<m;i++) {
            for(int j=1;j<n;j++) {
                if(obstacleGrid[i][j]==1) {
                    a[i][j] = 0;
                } else {
                    a[i][j] = a[i-1][j]+a[i][j-1];
                }
            }
        }
        for(int i=0;i<m;i++) {
            for(int j=0;j<n;j++) {
                cout<<a[i][j]<<" ";
            }
            cout<<endl;
        }
        cout<<a[m-1][n-1]<<endl;
        return a[m-1][n-1];
    }
};
int main() {
    vector<vector<int> > A;
    vector<int> B;
    B.push_back(0);
    B.push_back(1);
    A.push_back(B);
    Solution i;
    i.uniquePathsWithObstacles(A);
//    r.reverse(A, 4,9);
    return 0;
}