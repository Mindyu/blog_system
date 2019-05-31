package trie

import (
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils/trie"
	"github.com/labstack/gommon/log"
	"strings"
)

var keyTrie *trie.Trie

func init() {
	keyTrie = trie.NewTrie()
	res, err := stores.GetBlogSearchKey()
	if err != nil {
		log.Error(err)
	}
	//r, _ := json.Marshal(res)
	//log.Info(string(r))
	for _, v := range res {
		keyTrie.Insert(strings.Trim(v.BlogTitle, " "))
		keyTrie.Insert(strings.Trim(v.Author, ""))
		keyTrie.Insert(strings.Trim(v.TypeName, " "))
		for _, keyword := range strings.Split(v.Keywords, ",") {
			keyTrie.Insert(keyword)
		}
	}
	log.Info("成功生成前缀树")
}

// 获取keyTrie全局实例
func GetKeyTrie() *trie.Trie {
	return keyTrie
}
