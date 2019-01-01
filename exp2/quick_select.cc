#include <iostream>

using namespace std;

int partition1(int *input, int left, int right) {
    int pivot = (input[left] + input[right]) / 2;
    while (left < right) {
        while (input[left] < pivot) {
            ++left;
        }
        while (input[right] > pivot) {
            --right;
        }
        if (left <= right) {
            swap(input[left++], input[right--]);
        }
    }
    return left;
}

int partition(int *input, int left, int right) {
    int pivot = (input[left] + input[right]) >> 1;
    int i = left - 1;
    for (int j = left; j <= right; j++) {
        if (input[j] <= pivot) {
            swap(input[++i], input[j]);
        }
    }
    return min(i + 1, right);
}

int quick_select(int *input, int low, int high, int k) {
    if (low >= high) {
        return input[low];
    }
    int pi = partition(input, low, high);
    int length = pi - low;
    if (length < k) {
        return quick_select(input, pi, high, k - length);
    } else {
        return quick_select(input, low, pi - 1, k);
    }
}

void quick_sort(int *input, int low, int high) {
    if (low >= high) {
        return;
    }
    int pi = partition(input, low, high);
    quick_sort(input, low, pi - 1);
    quick_sort(input, pi, high);
}

int main() {
    int n, k;
    cin >> n >> k;
    int input[n];
    for (int i = 0; i < n; ++i) {
        cin >> input[i];
    }
    cout << k << "th minimum: " << quick_select(input, 0, n - 1, k) << endl;
    return 0;
}
