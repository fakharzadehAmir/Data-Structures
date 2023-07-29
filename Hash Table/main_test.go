package hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func djb2Hash(key string) int {
	hash := 5381
	for _, ch := range key {
		hash = (hash * 33) ^ int(ch)
	}
	return hash & 0x7FFFFFFF
}

func TestInsert(t *testing.T) {
	t.Run("Insert and Retrieve", func(t *testing.T) {
		assert := assert.New(t)
		hash := NewHash(8)
		testData := []struct {
			data string
		}{
			{"apple"},
			{"orange"},
			{"banana"},
			{"grape"},
			{"melon"},
			{"kiwi"},
			{"pear"},
			{"peach"},
		}

		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
			expectedResult := hash.HashTable[hashKey].Data
			assert.Equal(data.data, expectedResult, "Retrieved data should be equal to inserted data")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete and Retrieve", func(t *testing.T) {
		assert := assert.New(t)
		hash := NewHash(8)
		testData := []struct {
			data string
		}{
			{"apple"},
			{"orange"},
			{"banana"},
			{"grape"},
			{"melon"},
			{"kiwi"},
			{"pear"},
			{"peach"},
		}

		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
		}

		deleteDataNonClosure := "orange"
		deleteHashKeyNonClosure := djb2Hash(deleteDataNonClosure) % 8
		hash.Delete(deleteDataNonClosure, deleteHashKeyNonClosure)
		deletedNodeNonClosure := hash.HashTable[deleteHashKeyNonClosure]
		for deletedNodeNonClosure != nil {
			assert.NotEqual(deleteDataNonClosure, deletedNodeNonClosure.Data, "Deleted data should not be present in the hashtable")
			deletedNodeNonClosure = deletedNodeNonClosure.Next
		}

		deleteDataClosure := "pear"
		deleteHashKeyClosure := djb2Hash(deleteDataClosure) % 8
		hash.Delete(deleteDataClosure, deleteHashKeyClosure)
		deletedNodeClosure := hash.HashTable[deleteHashKeyClosure]
		for deletedNodeClosure != nil {
			assert.NotEqual(deleteDataClosure, deletedNodeClosure.Data, "Deleted data should not be present in the hashtable")
			deletedNodeClosure = deletedNodeClosure.Next
		}

		nonExistingData := "strawberry"
		nonExistingHashKey := djb2Hash(nonExistingData) % 8
		result := hash.Delete(nonExistingData, nonExistingHashKey)
		expeced := false
		assert.Equal(result, expeced, "Retrieved data should be equal to inserted data")
	})
}

func TestSearch(t *testing.T) {
	t.Run("Insert and Search", func(t *testing.T) {
		assert := assert.New(t)
		hash := NewHash(8)
		testData := []struct {
			data string
		}{
			{"apple"},
			{"orange"},
			{"banana"},
			{"grape"},
			{"melon"},
			{"kiwi"},
			{"pear"},
			{"peach"},
		}

		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
		}

		searchData := "orange"
		searchHashKey := djb2Hash(searchData) % 8
		found := hash.Search(searchData, searchHashKey)
		assert.True(found, "The search should find the existing element")

		searchData = "apple"
		searchHashKey = djb2Hash(searchData) % 8
		found = hash.Search(searchData, searchHashKey)
		assert.True(found, "The search should find the existing element")

		nonExistingData := "strawberry"
		nonExistingHashKey := djb2Hash(nonExistingData) % 8
		found = hash.Search(nonExistingData, nonExistingHashKey)
		assert.False(found, "The search should not find the non-existing element")
	})
}

func TestClear(t *testing.T) {
	t.Run("Insert and Search", func(t *testing.T) {
		assert := assert.New(t)
		hash := NewHash(8)
		testData := []struct {
			data string
		}{
			{"apple"},
			{"orange"},
			{"banana"},
			{"grape"},
			{"melon"},
			{"kiwi"},
			{"pear"},
			{"peach"},
		}

		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
		}
		hash.Clear()
		nonExistingData := "apple"
		nonExistingHashKey := djb2Hash(nonExistingData) % 8
		found := hash.Search(nonExistingData, nonExistingHashKey)
		assert.False(found, "The search should not find the non-existing element")
	})
}

func TestCount(t *testing.T) {
	t.Run("Insert and Search", func(t *testing.T) {
		assert := assert.New(t)
		hash := NewHash(8)
		testData := []struct {
			data string
		}{
			{"apple"},
			{"orange"},
			{"banana"},
			{"grape"},
			{"melon"},
			{"kiwi"},
			{"pear"},
			{"peach"},
			{"strawberry"},
			{"blueberry"},
			{"watermelon"},
			{"mango"},
			{"pineapple"},
			{"cherry"},
			{"raspberry"},
			{"blackberry"},
			{"apricot"},
			{"plum"},
			{"pomegranate"},
			{"kiwifruit"},
			{"coconut"},
			{"lemon"},
			{"lime"},
			{"guava"},
			{"fig"},
			{"date"},
			{"papaya"},
			{"passionfruit"},
			{"dragonfruit"},
			{"avocado"},
			{"jackfruit"},
			{"persimmon"},
			{"cranberry"},
			{"apricot"},
			{"boysenberry"},
			{"kiwano"},
			{"lychee"},
			{"mangosteen"},
			{"mulberry"},
			{"nectarine"},
			{"quince"},
			{"rambutan"},
			{"tamarind"},
			{"ugli fruit"},
			{"yuzu"},
			{"zucchini"},
			{"bell pepper"},
			{"broccoli"},
			{"carrot"},
			{"eggplant"},
			{"lettuce"},
			{"tomato"},
		}

		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
		}
		count := hash.Count()
		assert.Equal(count, len(testData), "Retrieved counted data should be equal to inserted data")
	})
}
