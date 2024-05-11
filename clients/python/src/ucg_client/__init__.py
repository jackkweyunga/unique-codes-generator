import sys, os

try:
    from ._version import __version__
    __version__ = __version__
except:
    __version__ = "0.1"

# __version__ = "0.5"


sys.path.append(os.path.dirname(__file__))
