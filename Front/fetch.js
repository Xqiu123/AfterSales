
export async function Fetch(path, method, body) {
    return await fetch('http://127.0.0.1:8080/api/v1' + path, {
        method: method,
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: body,
    }).then(res => {
        return res.json()
    })
}