package main

import (
	"bufio"
	"os"
	"strconv"
)
{% if mod or yes_str or no_str %}

{% endif %}
{% if mod %}
const mod = {{mod}}
{% endif %}
{% if yes_str %}
const yes = "{{ yes_str }}"
{% endif %}
{% if no_str %}
const no = "{{ no_str }}"
{% endif %}

func solve({{ formal_arguments }}) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()


}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	{% if prediction_success %}
	{{ input_part }}
	{% else %}
    // FIXME: Failed to predict input format
	{% endif %}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve({{ actual_arguments }})
}
