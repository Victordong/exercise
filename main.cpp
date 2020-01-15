#include <iostream>
#include <vector>
#include <stack>
#include <queue>
using namespace std;

class Solution {
public:
    string longestPalindrome(string s) {
        int a[1001][1001];
        int n = s.length();
        if (n == 1) {
            return s;
        }
        for (int i = 0; i < n; i++) {
            a[i][i] = true;
        }
        for (int i = 1; i < n; i++) {
            if (s[i] == s[i - 1]) {
                a[i - 1][i] = true;
            }
        }
        for(int m=2;m<n;m++) {
            for (int i = 0; i < n; i++) {
                int j = i+m;
                if(j>=n) {
                    break;
                }
                if(s[i]==s[j]&&a[i+1][j-1]) {
                    a[i][j] = true;
                } else {
                    a[i][j] = false;
                }
            }
        }

        int max = 0;
        int begin = 0;
        int end = 0;
        for(int i=0;i<n;i++) {
            for(int j=i;j<n;j++) {
                if(a[i][j] && j-i>max) {
                    max = j-i;
                    begin = i;
                    end = j;
                }
            }
        }
        for(int i=0;i<n;i++) {
            for(int j=0;j<i;j++) {
                cout<<0<<" ";
            }
            for(int j=i;j<n;j++) {
                cout<<a[i][j]<<" ";
            }
            cout<<endl;
        }
        return s.substr(begin, end-begin+1);
    }
};
int main() {
    string A = "abb";
    Solution i;
    cout<< i.longestPalindrome(A);
//    r.reverse(A, 4,9);
    return 0;
}