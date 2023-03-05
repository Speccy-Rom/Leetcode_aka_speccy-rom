# Design and implement a data structure for a Least Frequently Used (LFU) cache.
#
# Implement the LFUCache class:
#
# LFUCache(int capacity) Initializes the object with the capacity of the data structure.
# int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
# void put(int key, int value) Update the value of the key if present, or inserts the key if not already present. When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item. For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
# To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the smallest use counter is the least frequently used key.
#
# When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache is incremented either a get or put operation is called on it.
#
# The functions get and put must each run in O(1) average time complexity.
#
#
#
# Example 1:
#
# Input
# ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
# [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
# Output
# [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
#
# Explanation
# // cnt(x) = the use counter for key x
# // cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
# LFUCache lfu = new LFUCache(2);
# lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
# lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
# lfu.get(1);      // return 1
#                  // cache=[1,2], cnt(2)=1, cnt(1)=2
# lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
#                  // cache=[3,1], cnt(3)=1, cnt(1)=2
# lfu.get(2);      // return -1 (not found)
# lfu.get(3);      // return 3
#                  // cache=[3,1], cnt(3)=2, cnt(1)=2
# lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
#                  // cache=[4,3], cnt(4)=1, cnt(3)=2
# lfu.get(1);      // return -1 (not found)
# lfu.get(3);      // return 3
#                  // cache=[3,4], cnt(4)=1, cnt(3)=3
# lfu.get(4);      // return 4
#                  // cache=[4,3], cnt(4)=2, cnt(3)=3
import collections


class LFUCache(object):
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.k2f = {}
        self.f2kv = collections.defaultdict(collections.OrderedDict)
        self.min_freq = 0

    def get(self, key: int) -> int:
        if key not in self.k2f: return -1

        ofreq = self.k2f[key]
        nfreq = ofreq + 1

        # update f2kv to nf2kv
        val = self.f2kv[ofreq].pop(key)
        self.f2kv[nfreq][key] = val
        self.k2f[key] = nfreq
        if self.min_freq == ofreq and (not self.f2kv[ofreq]):
            self.min_freq = nfreq

        return val

    def put(self, key: int, value: int) -> None:
        if self.capacity <= 0: return

        if key in self.k2f:
            freq = self.k2f[key]
            self.f2kv[freq][key] = value
            self.get(key)
        else:  # new key
            if len(self.k2f) >= self.capacity:
                least_freq_k, least_freq_v = self.f2kv[self.min_freq].popitem(last=False)
                del self.k2f[least_freq_k]
            self.min_freq = 1  # new key
            self.k2f[key] = 1
            self.f2kv[1][key] = value

# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
