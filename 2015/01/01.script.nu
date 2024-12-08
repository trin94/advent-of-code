def main [file: string] {
	open $file
	| split row "\n"
	| each { |line|
		let open = echo $line
			| split chars
			| each { |c| if $c == "(" { echo $c } }
	        | length

	    let close = echo $line
	        | split chars
	        | each { |c| if $c == ")" { echo $c } }
	        | length

		$open - $close

	} | str join ","
}

