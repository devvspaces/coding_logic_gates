"""
Python implementation of logic gates.
From this implementation, basic gates can be chained together like
NAND and XNOR, where NOT Gate was chained with AND and XOR Gates
respectively.
"""


def AND(a, b):
    """
    AND Gate
    Returns True if only both values are True
    """
    return a and b


def OR(a, b):
    """
    OR Gate
    Returns True if any value is True
    """
    return a or b


def NOT(a):
    """
    NOT Gate
    Inverts the value
    """
    return not a


def XOR(a, b):
    """
    XOR Gate
    Returns True if only both values are different
    """
    return a != b


def NAND(a, b):
    """
    NAND Gate
    Returns True unless if both inputs are True
    """
    return NOT(AND(a, b))


def XNOR(a, b):
    """
    XNOR Gate
    Returns True when both inputs are the same
    """
    return NOT(XOR(a, b))
