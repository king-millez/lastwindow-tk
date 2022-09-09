import argparse

from fmt.pack import PackFile


def main():
    parser = argparse.ArgumentParser(
        description="Extract data from Last Window .pack files"
    )
    parser.add_argument("-i", dest="input_file")
    parser.add_argument("-o", dest="output_dir")
    args = parser.parse_args()

    if not args.input_file:
        raise ValueError(f"An input .pack file must be specified")

    if not args.output_dir:
        raise ValueError("An output directory must be specified.")

    pf = PackFile(args.input_file)
    pf.extract(args.output_dir)


if __name__ == "__main__":
    main()
