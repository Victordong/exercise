//
// Created by zhandong on 2020-01-06.
//

#include <string>
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
    string Serialize(TreeNode* A) {
        string result;
        stack<TreeNode*> tree;
        tree.push(A);
        while(!tree.empty()) {
            TreeNode *part = tree.top();
            tree.pop();
            if(part==NULL) {
                result += "#!";
            }else {
                tree.push(part->right);
                tree.push(part->left);
                result+=to_string(part->val)+"!";
            }
        }
        return result;
    }
    bool chkIdentical(TreeNode* A, TreeNode* B) {
        // write code here
        string a = Serialize(A);
        string b = Serialize(B);
        return a.find(b)!=string::npos;
    }
};

class Rotation {
public:
    bool chkRotation(string A, int lena, string B, int lenb) {
        // write code here
        if(lena!= lenb) {
            return false;
        }
        string new_string = B+B;
        return new_string.find(A)!= string::npos;
    }
};


class Transform {
public:
    bool chkTransform(string A, int lena, string B, int lenb) {
        // write code here
        if(lena!=lenb) {
            return false;
        }
    }
};