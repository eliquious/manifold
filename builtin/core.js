
// str2ab converts a string to an arraybuffer
const str2ab = function(str) {
	var buf = new ArrayBuffer(str.length);
	var bufView = new Uint8Array(buf);
	for (var i = 0, strLen = str.length; i < strLen; i++) {
		bufView[i] = str.charCodeAt(i);
	}
	return buf;
}

// ab2str converts an arraybuffer to a string
const ab2str = function(buf) {
	return String.fromCharCode.apply(null, new Uint8Array(buf));
};

// Sends a msg to the API. Args needs to be an array of strings.
export function send(fn, args) {
	let resp = global.send(str2ab(JSON.stringify({
		fn,
		args
	})));

	let val = JSON.parse(ab2str(resp));
	if (val.err) throw (val.err);
	return val.data;
};
