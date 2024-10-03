def selection_sort(arr):
    sorted_list = []
    while len(arr) > 0:
        lowest_int = min(arr)
        sorted_list.append(lowest_int)
        arr.remove(lowest_int)

    return sorted_list
        
print(selection_sort([3,7,5,9]))