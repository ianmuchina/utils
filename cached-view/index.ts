let form: HTMLFormElement = document.querySelector("form")

let url: HTMLInputElement = document.querySelector("#url")
let q: HTMLInputElement = document.querySelector("#query")

let googleCache: HTMLAnchorElement = document.querySelector("#googleCache")
let archiveDotOrg: HTMLAnchorElement = document.querySelector("#archiveDotOrg")
let live: HTMLAnchorElement = document.querySelector("#live")

// Reset page
// form.addEventListener("submit", e => url.value = "")

url.addEventListener("keyup", updateFields)

function updateFields() {
    let value = new URL("https://example.com")

    if (url.value.length > 0 && url.willValidate) {
        value = new URL(url.value)
    }

    q.value = `cache:${value}`
    googleCache.href = `https://webcache.googleusercontent.com/search?q=cache:${value}&strip=1&vwsrc=0`
    archiveDotOrg.href = `https://web.archive.org/web/*/${value}`
    live.href = ` ${value}`
}

updateFields()