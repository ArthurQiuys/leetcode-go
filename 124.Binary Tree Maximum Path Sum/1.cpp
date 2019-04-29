/*
* @Author: qiuyu
* @Date:   2019-04-30 00:11:21
* @Last Modified by:   qiuyu
* @Last Modified time: 2019-04-30 00:14:02
*/
class Solution {
public:
    int maxPathSum(TreeNode* root) {
        int res = INT_MIN;
        maxpat(root, res);
        return res;
    }
    int maxpat(TreeNode* node, int& res) {
        if (!node) return 0;
        int left = max(maxpat(node->left, res), 0);
        int right = max(maxpat(node->right, res), 0);
        res = max(res, left + right + node->val);
        return max(left, right) + node->val;
    }
};