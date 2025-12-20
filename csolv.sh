#! /bin/bash

TARGET_FOLDER=aoc
TARGET_PACKAGE=aoc

filename=$(basename "$0")
day="$1"
target_filename="${TARGET_FOLDER}/s${day}.go"

if [ -z "$day" ]; then
    echo "Usage: ${filename} DAY" 1>&2
    exit 1
fi

if ! [[ "$day" =~ ^[0-9]+$  ]]; then
    echo "'$day' has to be a number" 1>&2
    exit 1
fi

template=$(cat <<EOF
package $TARGET_PACKAGE

type Day${day}Solver struct{}

func (d *Day${day}Solver) GetDay() uint8 {
    return $day
}

func (d *Day${day}Solver) SolveOne(i []byte) (string, error) {
    return "", nil
}

func (d *Day${day}Solver) SolveTwo(i []byte) (string, error) {
    return "", nil
}
EOF
)

if [ -f "$target_filename" ]; then
    echo "file '$target_filename' already exists" 1>&2
    exit 1
fi

echo "$template" > $target_filename
