
import {incr, decr, get} from "lib.js";
var value = get();
for (var i = 0; i < 10; i++) {
    let val = incr();
}
global.print(get());
global.print(btoa);
global.print(btoa("Hello, World!"));
