"""https://the-algorithms.com/algorithm/quick-sort"""

from random import randrange


def quick_sort(collection: list) -> list:
    if len(collection) < 2:
        return collection

    pivot_index = randrange(len(collection))
    pivot = collection[pivot_index]
    greater: list[int] = []
    lesser: list[int] = []

    for element in collection[:pivot_index]:
        (greater if element > pivot else lesser).append(element)

    for element in collection[pivot_index+1:]:
        (greater if element > pivot else lesser).append(element)

    return quick_sort(lesser) + [pivot] + quick_sort(greater)


if __name__ == "__main__":
    unsorted = [0,5,3,2,2]
    print(quick_sort(unsorted))
