"""https://the-algorithms.com/algorithm/fibonacci-numbers"""


class Fibonacci:
    def __init__(self) -> None:
        self.sequence = [0, 1]

    def get(self, index: int) -> list:
        difference = index - (len(self.sequence) - 2)
        if difference >= 1:
            for _ in range(difference):
                self.sequence.append(self.sequence[-1] + self.sequence[-2])
        return self.sequence[:index]


def main():
    fibonacci = Fibonacci()

    print(fibonacci.get(8))


if __name__ == "__main__":
    main()
