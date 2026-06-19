from solution import count_occurrences

def test_count_occurrences():
    assert count_occurrences([1, 2, 3, 2, 4, 2, 5], 2) == 3
    assert count_occurrences([1, 2, 3], 4) == 0
    assert count_occurrences([], 1) == 0
    assert count_occurrences([5, 5, 5, 5], 5) == 4

if __name__ == "__main__":
    test_count_occurrences()
    print("All tests passed successfully!")
