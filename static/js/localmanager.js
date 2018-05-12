function local_get(key) {
    return JSON.parse(window.localStorage.getItem(key))
}

function local_save(key, item) {
    window.localStorage.setItem(key, JSON.stringify(item))
}
