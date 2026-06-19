def count_occurrences(arr: list[int], x: int) -> int:
    """
    Counts the number of occurrences of an integer x in a list of integers arr.
    """
    return arr.count(x)

if __name__ == "__main__":
    arr = [1, 2, 3, 2, 4, 2, 5]
    x = 2
    result = count_occurrences(arr, x)
    print(f"Occurrences of {x}: {result}")
