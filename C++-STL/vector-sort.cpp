#include <algorithm>
#include <vector>
#include <iostream>

int main() {
    int N;

    std::cin >> N;
    std::vector<int> vec(N);
    for (int i = 0; i < N; ++i) {
        std::cin >> vec[i];
    }

    std::sort(vec.begin(), vec.end());
    for (const auto e: vec) {
        std::cout << e << " ";
    }
}
