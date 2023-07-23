package hashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// djb2Hash calculates the hash value for the given key using the djb2 hash function.
func djb2Hash(key string) int {
	hash := 5381
	for _, ch := range key {
		hash = (hash * 33) ^ int(ch)
	}
	return hash & 0x7FFFFFFF // Ensure the hash is a positive value
}

func TestInsert(t *testing.T) {
	t.Run("Insert and Retrieve", func(t *testing.T) {
		assert := assert.New(t)
		hash := NewHash(8)
		testData := []struct {
			data string
		}{
			// Test data for insertion and retrieval test case
			{"apple"},
			{"orange"},
			{"banana"},
			{"grape"},
			{"melon"},
			{"kiwi"},
			{"pear"},
			{"peach"},
		}

		// Insert test data into the hashtable and ensure that the retrieved data is equal to the inserted data.
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
			// Test data for insertion and deletion test case
			{"apple"},
			{"orange"},
			{"banana"},
			{"grape"},
			{"melon"},
			{"kiwi"},
			{"pear"},
			{"peach"},
		}

		// Insert test data into the hashtable
		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
		}

		// Deleting a non-closure (first element in linked list)
		deleteDataNonClosure := "orange"
		deleteHashKeyNonClosure := djb2Hash(deleteDataNonClosure) % 8
		hash.Delete(deleteDataNonClosure, deleteHashKeyNonClosure)
		deletedNodeNonClosure := hash.HashTable[deleteHashKeyNonClosure]
		for deletedNodeNonClosure != nil {
			assert.NotEqual(deleteDataNonClosure, deletedNodeNonClosure.Data, "Deleted data should not be present in the hashtable")
			deletedNodeNonClosure = deletedNodeNonClosure.Next
		}

		// Deleting a closure (non-first element in linked list)
		deleteDataClosure := "pear"
		deleteHashKeyClosure := djb2Hash(deleteDataClosure) % 8
		hash.Delete(deleteDataClosure, deleteHashKeyClosure)
		deletedNodeClosure := hash.HashTable[deleteHashKeyClosure]
		for deletedNodeClosure != nil {
			assert.NotEqual(deleteDataClosure, deletedNodeClosure.Data, "Deleted data should not be present in the hashtable")
			deletedNodeClosure = deletedNodeClosure.Next
		}

		// Non-existing element
		nonExistingData := "strawberry"
		nonExistingHashKey := djb2Hash(nonExistingData) % 8
		result := hash.Delete(nonExistingData, nonExistingHashKey)
		expected := false
		assert.Equal(result, expected, "The result should be false for a non-existing element")
	})
}

// The TestSearch function has already been provided in the previous response and doesn't require additional comments.

func TestClear(t *testing.T) {
	t.Run("Insert and Search", func(t *testing.T) {
		assert := assert.New(t)
		hash := NewHash(8)
		testData := []struct {
			data string
		}{
			// Test data for clearing test case
			// ... (the provided test data)
		}

		// Insert test data into the hashtable
		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
		}
		// Test Case for testing Clear method
		hash.Clear()
		nonExistingData := "apple"
		nonExistingHashKey := djb2Hash(nonExistingData) % 8
		found := hash.Search(nonExistingData, nonExistingHashKey)
		assert.False(found, "The search should not find the non-existing element after clearing the hashtable")
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

		// Insert test data into the hashtable
		for _, data := range testData {
			hashKey := djb2Hash(data.data) % 8
			hash.Insert(data.data, hashKey)
		}
		// Test Case for testing Count method
		count := hash.Count()
		assert.Equal(count, len(testData), "Retrieved counted data should be equal to inserted data")
	})
}
