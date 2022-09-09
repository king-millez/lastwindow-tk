import os
import zlib
from util.filehandler import FileHandler


class PackFileError(Exception):
    pass


class PackFile:
    def __init__(self, input_file: str) -> None:
        self.input_file = input_file
        if self._validate_file(self.input_file):
            self._read_header(self.input_file)
        else:
            raise PackFileError(f"{self.input_file} is not a valid .pack file")

    def _read_header(self, filepath: str) -> None:
        with open(filepath, "rb") as f:
            fh = FileHandler(f)
            fh.seek(4)  # Skip first null bytes
            self.file_count = fh.read_long()
            self.start_addr = (
                fh.read_long() + 8
            )  # We skip 8 bytes to move past unnecessary info later in the file
            fh.read_long()  # If anyone can figure out what this is, that'd be great
            self.files = []
            for _ in range(self.file_count):
                name_len = fh.read_byteint()
                name = fh.read_bytes(name_len).decode()
                zsize = fh.read_long()
                self.files.append({"size": zsize, "name": name})

    def _validate_file(self, filepath: str) -> bool:
        valid = False
        if os.path.isfile(filepath):
            with open(filepath, "rb") as f:
                fh = FileHandler(f)
                try:  # The .pack file header doesn't have any "magic numbers" like most files, so we have to do some mathematical improvisation to check for a valid file
                    assert fh.read_long() == 0
                    assert fh.read_long() > 0
                    data_start = fh.read_long()
                    assert data_start > 0
                    valid = True
                except AssertionError:
                    pass
        return valid

    def extract(self, output_dir: str) -> None:
        os.makedirs(output_dir, exist_ok=True)
        with open(self.input_file, "rb") as f:
            fh = FileHandler(f)
            fh.seek(self.start_addr)
            for dfile in self.files:
                with open(os.path.join(output_dir, f"{dfile['name']}"), "wb") as cf:
                    dat = fh.read_bytes(dfile["size"])
                    if (
                        os.path.splitext(dfile["name"])[1] == ".bra"
                    ):  # .bra files are uncompressed animation files
                        cf.write(dat)
                    else:
                        cf.write(zlib.decompress(dat))
                print(f"Extracted {dfile['name']}")
