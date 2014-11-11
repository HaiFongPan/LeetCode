# -*- coding: utf-8 -*-

class LRUCache:

    # @param capacity, an integer
    def __init__(self, capacity):
        self.cache = {}
        self.record = []
        self.capacity = capacity

    # @return an integer
    def get(self, key):
        if self.cache.has_key(key):
            self.record.remove(key)
            self.record.append(key)
            return self.cache[key]
        return -1

    # @param key, an integer
    # @param value, an integer
    # @return nothing
    def set(self, key, value):
        if len(self.cache) == self.capacity:
            if self.cache.has_key(key):
                self.record.remove(key)
            else:
                r = self.record.pop(0)
                del self.cache[r]
            # self.record.append(key)
        elif key in self.record:
            self.record.remove(key)

        self.record.append(key)
        self.cache[key] = value




