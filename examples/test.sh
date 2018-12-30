#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

fail=0
failed=()

export GO111MODULE=on

for example in $(find . -type f -name main.go | sort); do
	path=$(readlink -f $example)
	name=$(basename $(dirname $path))
	echo $path
	if ! go build -o ${DIR}/tmp/${name} ${path} ; then
		fail=1
		failed+=($name)
	fi
done

if (( $fail == 1 )); then
	echo "================"
	echo "Failed examples:"
	printf '  %s\n' "${failed[@]}"
fi

exit $fail
