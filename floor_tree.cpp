//
// Created by zhandong on 2020-01-03.
//
#include <vector>
#include <queue>
using namespace std;

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(NULL), right(NULL) {
    }
};

class TreePrinter {
public:
    vector<vector<int> > printTree(TreeNode *root) {
        vector<vector<int> > result;
        vector<int> result_line;
        TreeNode *last, *nlast = NULL;
        queue<TreeNode*> q;
        last = root;
        q.push(root);
        while(!q.empty()) {
            TreeNode* now = q.front();
            result_line.push_back(now->val);
            q.pop();
            if(now->left != NULL) {
                q.push(now->left);
                nlast = now->left;
            }
            if(now->right != NULL) {
                q.push(now->right);
                nlast = now->right;
            }
            if(now==last) {
                last = nlast;
                result.push_back(result_line);
                result_line.clear();
            }
        }
        return result;
    }
};