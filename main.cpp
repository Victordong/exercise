#include <iostream>
#include <vector>
#include <stack>
#include <queue>
using namespace std;

class Prior {
public:
    int quickSortPart(vector<string> &strs, int begin, int end) {
        int left = begin;
        int right = end;
        string temp = strs[left];
        while (left < right) {
            while (temp + strs[right] < strs[right] + temp && right > left) {
                right--;
            }
            strs[left] = strs[right];
            while (temp + strs[left] > strs[left] + temp && right > left) {
                left++;
            }
            strs[right] = strs[left];
        }
        strs[left] = temp;
        return left;
    }

    void quickSort(vector<string> &strs, int left, int right) {
        if (left >= right) {
            return;
        }
        int mid = quickSortPart(strs, left, right);
        quickSort(strs, left, mid);
        quickSort(strs, mid + 1, right);
    }

    string findSmallest(vector<string> strs, int n) {
        // write code here
        string result;
        quickSort(strs, 0, n-1);
        for(int i =0;i<n;i++) {
            result += strs[i];
        }
        return result;
    }
};
int main() {
    string s[] = {"abc","de", "a"};
    vector<string> A;
    for(int i=0;i<3;i++) {
        A.push_back(s[i]);
    }
    Prior r;

    string result = r.findSmallest(A,3);
//    r.reverse(A, 4,9);
    cout<<result<<endl;
    return 0;
}