"""Test gates with python unittest"""

from unittest import TestCase, main
from logic_gates import AND, NOT, NAND, OR, XNOR, XOR
from functools import wraps


def parametrize(args: str, values: list):
    """Like pytest parametrize

    Args:
        args (str): example `name, age`
        values (list of any): example `[('John Doe', 45)]`
    """

    # Convert args to dict
    keys = [key.strip() for key in args.split(',')]
    kwargs = []

    for value in values:
        kwarg = {key: input for key, input in zip(keys, value)}
        kwargs.append(kwarg)

    def decorator(fn):

        def wrapper(self):

            for kwarg in kwargs:
                fn(self, **kwarg)

        return wrapper

    return decorator


class TestGates(TestCase):

    @parametrize(
        args="input, output",
        values=[
            (
                [True, True], True
            ),
            (
                [True, False], False
            ),
            (
                [False, True], False
            ),
            (
                [False, False], False
            ),
        ]
    )
    def test_and_gate(self, input, output):
        self.assertEqual(output, AND(*input))

    @parametrize(
        args="input, output",
        values=[
            (
                [True, True], True
            ),
            (
                [True, False], True
            ),
            (
                [False, True], True
            ),
            (
                [False, False], False
            ),
        ]
    )
    def test_or_gate(self, input, output):
        self.assertEqual(output, OR(*input))

    @parametrize(
        args="input, output",
        values=[
            (
                [False], True
            ),
            (
                [True], False
            ),
        ]
    )
    def test_not_gate(self, input, output):
        self.assertEqual(output, NOT(*input))

    @parametrize(
        args="input, output",
        values=[
            (
                [True, True], False
            ),
            (
                [True, False], True
            ),
            (
                [False, True], True
            ),
            (
                [False, False], False
            ),
        ]
    )
    def test_xor_gate(self, input, output):
        self.assertEqual(output, XOR(*input))

    @parametrize(
        args="input, output",
        values=[
            (
                [True, True], False
            ),
            (
                [True, False], True
            ),
            (
                [False, True], True
            ),
            (
                [False, False], True
            ),
        ]
    )
    def test_nand_gate(self, input, output):
        self.assertEqual(output, NAND(*input))

    @parametrize(
        args="input, output",
        values=[
            (
                [True, True], True
            ),
            (
                [True, False], False
            ),
            (
                [False, True], False
            ),
            (
                [False, False], True
            ),
        ]
    )
    def test_xnor_gate(self, input, output):
        self.assertEqual(output, XNOR(*input))


if __name__ == "__main__":
    main()
