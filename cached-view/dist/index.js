var form = document.querySelector("form");
var url = document.querySelector("#url");
var q = document.querySelector("#query");
var googleCache = document.querySelector("#googleCache");
var archiveDotOrg = document.querySelector("#archiveDotOrg");
var live = document.querySelector("#live");
// Reset page
// form.addEventListener("submit", e => url.value = "")
url.addEventListener("keyup", updateFields);
function updateFields() {
    var value = new URL("https://example.com");
    if (url.value.length > 0 && url.willValidate) {
        value = new URL(url.value);
    }
    // Update dom elements
    q.value = "cache:".concat(value);
    googleCache.href = "https://webcache.googleusercontent.com/search?q=cache:".concat(value, "&strip=1&vwsrc=0");
    archiveDotOrg.href = "https://web.archive.org/web/*/".concat(value);
    live.href = " ".concat(value);
    console.log(q.value, googleCache.href);
}
updateFields();
