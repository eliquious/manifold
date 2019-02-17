
export function incr() {
	let resp = send("incr");
	return resp;
}

export function decr() {
	let resp = send("decr");
	return resp;
}

export function get() {
	let resp = send("get");
	return resp;
}
