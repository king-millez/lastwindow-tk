from io import BufferedReader
from ctypes import c_uint16


class FileHandler:
    def __init__(self, f: BufferedReader, bo: str = "big") -> None:
        self.f = f
        self.byte_order = bo

    def read_bytes(self, length: int):
        return self.f.read(length)

    def read_byteint(self):
        return int.from_bytes(self.read_bytes(1), self.byte_order)

    def read_short(self):
        return int.from_bytes(self.read_bytes(2), self.byte_order)

    def read_long(self):
        byte = self.read_bytes(4)
        print(byte)
        return int.from_bytes(byte, self.byte_order)

    def tell(self):
        return self.f.tell()

    def seek(self, bytec: int):
        self.f.seek(bytec)
