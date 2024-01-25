def make_row_lookup(n: int) -> list[list[int]]:
    return [[x + y for y in range(n**2)] for x in range(0, n**4, n**2)]


def make_col_lookup(n: int) -> list[list[int]]:
    return [[x + y for y in range(0, n**4, n**2)] for x in range(n**2)]


def make_cel_lookup(n: int) -> list[list[int]]:
    return [
        [((z // n) * n**2) + (z % n) + (y * n) + (x * n**3) for z in range(n**2)]
        for x in range(n)
        for y in range(n)
    ]


def main() -> None:
    print("package autogen\n")
    print("// auto-generated lookup table\n")
    print(
        "var RelatedElements = [][][]int{RelatedElements2, RelatedElements3, RelatedElements4}\n"
    )
    for sz in [2, 3, 4]:
        row_lookup = make_row_lookup(sz)
        col_lookup = make_col_lookup(sz)
        cel_lookup = make_cel_lookup(sz)

        # We have to concatenate here because you can't have curly braces within a format string
        print(f"var RelatedElements{sz} = [][]int" + "{")

        for n in range(sz**4):
            row_element: int = n // sz**2
            col_element: int = n % sz**2
            cel_element: int = (((n // sz**3)) * sz) + ((n % sz**2) // sz)
            related: list = list(
                set(
                    row_lookup[row_element]
                    + col_lookup[col_element]
                    + cel_lookup[cel_element]
                )
            )
            related.remove(n)
            related.sort()
            related_str: str = "\t{" + ", ".join(str(x) for x in related) + "},"
            print(related_str)
        print("}\n")


if __name__ == "__main__":
    main()
