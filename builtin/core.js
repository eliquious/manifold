
const str2ab = function(str) {
	var buf = new ArrayBuffer(str.length);
	var bufView = new Uint8Array(buf);
	for (var i = 0, strLen = str.length; i < strLen; i++) {
		bufView[i] = str.charCodeAt(i);
	}
	return buf;
}

const ab2str = function(buf) {
	return String.fromCharCode.apply(null, new Uint8Array(buf));
};

export function send(fn, args) {
	let resp = global.send(str2ab(JSON.stringify({
		fn,
		args
	})));

	let val = JSON.parse(ab2str(resp));
	if (val.err) throw (val.err);
	return val.data;
};

// (function (root, factory) {
//     if (typeof define === 'function' && define.amd) {
//         // AMD. Register as an anonymous module.
//         define([], function() {factory(root);});
//     } else factory(root);
// })(typeof exports !== "undefined" ? exports : this, function(root) {

// 	const str2ab = function(str) {
// 		var buf = new ArrayBuffer(str.length);
// 		var bufView = new Uint8Array(buf);
// 		for (var i = 0, strLen = str.length; i < strLen; i++) {
// 			bufView[i] = str.charCodeAt(i);
// 		}
// 		return buf;
// 	}

// 	const ab2str = function(buf) {
// 		return String.fromCharCode.apply(null, new Uint8Array(buf));
// 	};

// 	root.send = function(fn, args) {
// 		let resp = global.send(str2ab(JSON.stringify({
// 			fn,
// 			args
// 		})));

// 		let val = JSON.parse(ab2str(resp));
// 		if (val.err) throw (val.err);
// 		return val.data;
// 	};
// });
