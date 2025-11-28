export interface CharCount {
	interpreted: number;
	total: number;
}

export async function part1(filename: string) {
	const file = Bun.file(`input/${filename}`);
	const text = await file.text();
	const lines = text.split("\n").filter((line) => line);

	return lines
		.map((line) => countCharacters(line))
		.map((cc) => cc.total - cc.interpreted)
		.reduce((previous, current) => previous + current);
}

export async function part2(filename: string) {
	const file = Bun.file(`input/${filename}`);
	const text = await file.text();
	const lines = text.split("\n").filter((line) => line);

	return lines
		.map((line) => countEscapedCharacters(line))
		.map((cc) => cc.interpreted - cc.total)
		.reduce((previous, current) => previous + current);
}

export function countCharacters(raw: string): CharCount {
	const interpreted = interpretString(raw);
	return { interpreted: interpreted.length, total: raw.length };
}

function interpretString(raw: string): string {
	const content = raw.slice(1, -1);

	return content.replace(
		/\\(x([0-9A-Fa-f]{2})|u([0-9A-Fa-f]{4})|.)/g,
		(_, full, hex, unicode) => {
			if (hex) return String.fromCharCode(parseInt(hex, 16));
			if (unicode) return String.fromCharCode(parseInt(unicode, 16));
			const escapes: Record<string, string> = {
				n: "\n",
				r: "\r",
				t: "\t",
				"\\": "\\",
				"'": "'",
				'"': '"',
			};
			return escapes[full] ?? full;
		},
	);
}

export function countEscapedCharacters(raw: string): CharCount {
	const encoded = encodeString(raw);
	return { interpreted: encoded.length, total: raw.length };
}

export function encodeString(raw: string): string {
	const escaped = raw.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
	return `"${escaped}"`;
}
