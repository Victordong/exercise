//
// Created by zhandong on 2020-01-03.
//
#include <vector>
#include <queue>
#include <stack>
#include <string>

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
        queue<TreeNode *> q;
        last = root;
        q.push(root);
        while (!q.empty()) {
            TreeNode *now = q.front();
            result_line.push_back(now->val);
            q.pop();
            if (now->left != NULL) {
                q.push(now->left);
                nlast = now->left;
            }
            if (now->right != NULL) {
                q.push(now->right);
                nlast = now->right;
            }
            if (now == last) {
                last = nlast;
                result.push_back(result_line);
                result_line.clear();
            }
        }
        return result;
    }
};

class recursiveTreeToSequence {
public:
    void firstConvert(TreeNode *root, vector<int> &result) {
        if (root == NULL) {
            return;
        }
        result.push_back(root->val);
        firstConvert(root->left, result);
        firstConvert(root->right, result);
    }

    void secondConvert(TreeNode *root, vector<int> &result) {
        if (root == NULL) {
            return;
        }
        secondConvert(root->left, result);
        result.push_back(root->val);
        secondConvert(root->right, result);
    }

    void thirdConvert(TreeNode *root, vector<int> &result) {
        if (root == NULL) {
            return;
        }
        thirdConvert(root->left, result);
        thirdConvert(root->right, result);
        result.push_back(root->val);
    }

    vector<vector<int> > convert(TreeNode *root) {
        // write code here
        vector<int> result;
        vector<vector<int>> resultList;
        firstConvert(root, result);
        resultList.push_back(result);
        result.clear();
        secondConvert(root, result);
        resultList.push_back(result);
        result.clear();
        thirdConvert(root, result);
        resultList.push_back(result);
        return resultList;
    }
};


class TreeToSequence {
public:

    vector<int> firstConvert(TreeNode *root) {
        vector<int> result;
        stack<TreeNode *> tree;
        if (root == NULL) {
            return result;
        }
        tree.push(root);
        while (!tree.empty()) {
            TreeNode *cur = tree.top();
            tree.pop();
            result.push_back(cur->val);
            if (cur->right != NULL) {
                tree.push(cur->right);
            }
            if (cur->left != NULL) {
                tree.push(cur->left);
            }
        }
        return result;
    }

    vector<int> secondConvert(TreeNode *root) {
        vector<int> result;
        stack<TreeNode *> tree;
        if (root == NULL) {
            return result;
        }
        tree.push(root);
        TreeNode *cur = root;
        while (!tree.empty()) {
            while (cur->left != NULL) {
                tree.push(cur->left);
                cur = cur->left;
            }
            cur = tree.top();
            tree.pop();
            result.push_back(cur->val);
            while (cur->right == NULL && !tree.empty()) {
                cur = tree.top();
                tree.pop();
                result.push_back(cur->val);
            }
            cur = cur->right;
            if (cur != NULL) {
                tree.push(cur);
            }
        }
        return result;
    }

    vector<int> thirdConvert(TreeNode *root) {
        vector<int> result;
        stack<TreeNode *> treeBefore;
        stack<TreeNode *> treeEnd;
        if (root == NULL) {
            return result;
        }
        treeBefore.push(root);
        TreeNode *cur;
        while (!treeBefore.empty()) {
            cur = treeBefore.top();
            treeBefore.pop();
            if (cur->left != NULL) {
                treeBefore.push(cur->left);
            }
            if (cur->right != NULL) {
                treeBefore.push(cur->right);
            }
            treeEnd.push(cur);
        }
        while (!treeEnd.empty()) {
            cur = treeEnd.top();
            treeEnd.pop();
            result.push_back(cur->val);
        }
        return result;
    }

    vector<vector<int> > convert(TreeNode *root) {
        // write code here
        vector<int> result;
        vector<vector<int>> resultList;
        result = firstConvert(root);
        resultList.push_back(result);
        result.clear();
        result = secondConvert(root);
        resultList.push_back(result);
        result.clear();
        result = thirdConvert(root);
        resultList.push_back(result);
        return resultList;
    }
};

class TreeToString {
public:
    string toString(TreeNode *root) {
        // write code here
        string result;
        stack<TreeNode *> tree;
        TreeNode *cur;
        tree.push(root);
        while (!tree.empty()) {
            cur = tree.top();
            tree.pop();
            if (cur != NULL) {
                result += (to_string(cur->val) + "!");
                tree.push(cur->right);
                tree.push(cur->left);
            } else {
                result += "#!";
            }
        }
        return result;
    }
};

class CheckBalance {
public:
    int getHeight(TreeNode *root, int height, bool *isBalance) {
        if (root == NULL)
            return 0;
        int LH = getHeight(root->left, isBalance);
        int RH = getHeight(root->right, isBalance);
        if (((LH - RH) > 1) || ((RH - LH) > 1)) {
            *isBalance = false;
        }
        return max(LH, RH)+1;
    }

    bool check(TreeNode *root) {
        bool isBalance = true;
        getHeight(root, &isBalance);
        return isBalance;
    }
};


class CheckCompletion {
public:
    bool chk(TreeNode *root) {
        queue<TreeNode *> tree;
        TreeNode *cur;
        bool shouldLeaf = false;
        bool ifCompletion = true;
        tree.push(root);
        while (!tree.empty()) {
            cur = tree.front();
            tree.pop();
            if (cur->left != NULL && cur->right != NULL) {
                tree.push(cur->left);
                tree.push(cur->right);
            }
            if (cur->left == NULL && cur->right != NULL) {
                ifCompletion = false;
                break;
            }
            if (!shouldLeaf) {
                if (cur->left != NULL && cur->right == NULL) {
                    shouldLeaf = true;
                }
                if (cur->left == NULL && cur->right == NULL) {
                    shouldLeaf = true;
                }
            } else {
                if (cur->left != NULL || cur->right != NULL) {
                    ifCompletion = false;
                    break;
                }
            }

        }
        return ifCompletion;
    }
};

class FoldPaper {
public:
    void count(vector<string> &result, string now, int floor, int all) {
        if (floor == all) return;
        count(result, "down", floor + 1, all);
        result.push_back(now);
        count(result, "up", floor + 1, all);
    }

    vector<string> foldPaper(int n) {
        // write code here
        vector<string> result;
        count(result, "down", 0, n);
        return result;
    }
};


class FindErrorNode {
public:
    void sort(TreeNode *root, vector<int> &result, int &pre, int &cur) {
        if (root == NULL) return;
        sort(root->left, result, pre, cur);
        cur = root->val;
        if (cur < pre) {
            result.push_back(pre);
            result.push_back(cur);
        }
        pre = cur;
        sort(root->right, result, pre, cur);
    }

    vector<int> findError(TreeNode *root) {
        // write code here
        vector<int> result;
        int pre, cur = 0;
        sort(root, result, pre, cur);
        if (result.size() > 2)
            result.erase(result.begin() + 1, result.begin() + 3);
        reverse(result.begin(), result.end());
        return result;
    }
};

class LongestDistance {
public:
    int getHeight(TreeNode *root) {
        if (root == NULL) {
            return 0;
        }
        int LH = getHeight(root->left);
        int RH = getHeight(root->right);
        return max(LH, RH) + 1;
    }

    int getLong(TreeNode *root) {
        if (root == NULL) {
            return 0;
        }
        int LL = getLong(root->left);
        int RL = getLong(root->right);
        int LH = getHeight(root->left);
        int RH = getHeight(root->right);
        int maxL = max(LL,RL);
        return max(maxL, LH+RH+1);
    }

    int findLongest(TreeNode *root) {
        // write code here
        return getLong(root);
    }
};