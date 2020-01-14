//
// Created by zhandong on 2020-01-06.
//

#include <string>
#include <vector>
#include <stack>

using namespace std;


struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;

    TreeNode(int x) :
            val(x), left(NULL), right(NULL) {
    }
};

class IdenticalTree {
public:
    string Serialize(TreeNode *A) {
        string result;
        stack<TreeNode *> tree;
        tree.push(A);
        while (!tree.empty()) {
            TreeNode *part = tree.top();
            tree.pop();
            if (part == NULL) {
                result += "#!";
            } else {
                tree.push(part->right);
                tree.push(part->left);
                result += to_string(part->val) + "!";
            }
        }
        return result;
    }

    bool chkIdentical(TreeNode *A, TreeNode *B) {
        // write code here
        string a = Serialize(A);
        string b = Serialize(B);
        return a.find(b) != string::npos;
    }
};

class Rotation {
public:
    bool chkRotation(string A, int lena, string B, int lenb) {
        // write code here
        if (lena != lenb) {
            return false;
        }
        string new_string = B + B;
        return new_string.find(A) != string::npos;
    }
};


class Transform {
public:
    bool chkTransform(string A, int lena, string B, int lenb) {
        // write code here
        if (lena != lenb) {
            return false;
        }
        int a[26];
        int b[26];
        for (int i = 0; i < 26; i++) {
            a[i] = 0;
            b[i] = 0;
        }
        for (int i = 0; i < lena; i++) {
            a[A[i] - 'a'] += 1;
        }
        for (int i = 0; i < lenb; i++) {
            a[B[i] - 'a'] += 1;
        }
        for (int i = 0; i < 26; i++) {
            if (a[i] != b[i]) {
                return false;
            }
        }
        return true;
    }
};

class Reverse {
public:
    void reverse(string &A, int begin, int end) {
        for (int i = begin; i < (begin + end) / 2; i++) {
            char temp = A[i];
            A[i] = A[end - i - 1 + begin];
            A[end - i - 1 + begin] = temp;
        }
    }

    string reverseSentence(string A, int n) {
        // write code here
        reverse(A, 0, n);
        int cur = 0;
        for (int i = 0; i < n; i++) {
            if (A[i] == ' ') {
                reverse(A, cur, i);
                cur = i + 1;
            }
        }
        reverse(A, cur, n);
        return A;
    }
};

class Translation {
public:
    void reverse(string &A, int begin, int end) {
        for (int i = begin; i < (begin + end) / 2; i++) {
            char temp = A[i];
            A[i] = A[end - i - 1 + begin];
            A[end - i - 1 + begin] = temp;
        }
    }

    string stringTranslation(string A, int n, int len) {
        // write code here
        reverse(A, 0, len);
        reverse(A, len, n);
        reverse(A, 0, n);
        return A;
    }
};

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

class Replacement {
public:
    string replaceSpace(string iniString, int length) {
        // write code here
        int n=0;
        for(int i=0;i<length;i++){
            if(iniString[i]==' '){
                n++;
            }
        }
        if (n==0) return iniString;
        else{
            int j=length+2*n-1;
            string s=iniString+string(2*n,'0');
            for(int i=length-1;i>=0;i--){
                if(iniString[i]==' '){
                    s[j--]='0';
                    s[j--]='2';
                    s[j--]='%';
                }
                else s[j--]=iniString[i];
            }
            return s;
        }
    }
};

class Parenthesis {
public:
    bool chkParenthesis(string A, int n) {
        // write code here
        int num = 0;
        for(int i=0;i<n;i++) {
            if(A[i]=='(') {
                num++;
            } else {
                num--;
            }
            if(num<0)
                return false;
        }
        if(num!=0)
            return false;
        else
            return true;
    }
};

class DistinctSubstring {
public:
    int longestSubstring(string A, int n) {
        // write code here
        vector<int> help[26];
        int length[n];
        for(int i=0;i<n;i++) {
            length[n] = 0;
        }
        int pre = 0;
        for(int i=0;i<n;i++) {
            int cur = A[i]-'a';
            if(help[cur].empty()){
                pre+=1;
                length[i] = pre;
            } else{
                int pos1 = help[cur].back();
                int pos2 = i - pre;
                if(pos1>=pos2) {
                    pre = i-pos1;
                } else {
                    pre += 1;
                }
                length[i] = pre;
            }
            help[cur].push_back(i);
        }
        int max=0;
        for(int i=0;i<n;i++) {
            if(max<length[i]) {
                max = length[i];
            }
        }
        return max;
    }
};