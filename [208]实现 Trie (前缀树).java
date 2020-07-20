package pri.leetcode.leetcode.editor.cn;
//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
// 示例: 
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");   
//trie.search("app");     // 返回 true 
//
// 说明: 
//
// 
// 你可以假设所有的输入都是由小写字母 a-z 构成的。 
// 保证所有输入均为非空字符串。 
// 
// Related Topics 设计 字典树



//leetcode submit region begin(Prohibit modification and deletion)
class Trie {
    private TrieNode root;
    /** Initialize your data structure here. */
    public Trie() {
        root = new TrieNode('/');
    }
    
    /** Inserts a word into the trie. */
    public void insert(String word) {
        TrieNode point = root;
        for (int i = 0; i < word.length(); i++) {
            final int index = word.charAt(i) - 'a';
            if (point.children[index] == null) {
                TrieNode node = new TrieNode(word.charAt(i));
                point.children[index] = node;
            }

            point = point.children[index];
        }

        point.isEndChar = true;
    }
    
    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        TrieNode point = root;
        for (int i = 0; i < word.length(); i++) {
            final int index = word.charAt(i) - 'a';
            if (point.children[index] == null) {
                return false;
            }

            point = point.children[index];
        }

        return point.isEndChar;
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
        TrieNode point = root;
        for (int i = 0; i < prefix.length(); i++) {
            final int index = prefix.charAt(i) - 'a';
            if (point.children[index] == null) {
                return false;
            }

            point = point.children[index];
        }

        return point != null;
    }

    class TrieNode {
        char data;
        TrieNode[] children;
        boolean isEndChar;

        TrieNode(char data) {
            this.data = data;
            this.children = new TrieNode[26];
            this.isEndChar = false;
        }
    }

    public static void main(String[] args) {
        Trie trie = new Trie();
        trie.insert("apple");

        System.out.println(trie.search("apple"));
        System.out.println(trie.search("app"));
        System.out.println(trie.startsWith("app"));
        trie.insert("app");
        System.out.println(trie.search("app"));
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
