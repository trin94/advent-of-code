def main [file: string] {
	open $file
	| split row "\n"
	| each { |line|
		mut chars = $line | split chars
		let length = $chars | length;
		mut counter = 0;
		mut idx = 0;

		loop {
			$idx += 1

			if $idx > $length {
				break
			}

			if ($chars | length) == 0 {
				break
			}

			let nextChar = $chars | take 1 | first
			$chars = $chars | skip 1

			if $nextChar == "(" {
				$counter += 1
			} else if $nextChar == ")" {
				$counter -= 1
			}

			if $counter < 0 {
				break
			}

		};

		echo $idx


	} | str join ","
}

