const BASE_URL = import.meta.env.VITE_SERVER_URL || "http://localhost:8080"

interface Log {
    ID: number,
    CreatedAt: string
    count: number
}

export async function GetDaily() {
    try {
        console.log(BASE_URL)
        const res = await fetch(`${BASE_URL}/log/day`)
        const data =  await res.json()
        return remapData(data)
    }
    catch {
        throw alert("Error getting data")
    }
}

export async function GetWeekly() {
    try {
        const res = await fetch(`${BASE_URL}/log/week`)
        const data =  await res.json()
        return remapData(data)
    }
    catch {
        throw alert("Error getting data")
    }
}

export async function GetMonthly() {
    try {
        const res = await fetch(`${BASE_URL}/log/month`)
        const data =  await res.json()
        return remapData(data)
    }
    catch {
        throw alert("Error getting data")
    }
}

export async function GetAll() {
    try {
        const res = await fetch(`${BASE_URL}/log/all`)
        const data =  await res.json()
        return remapData(data)
    }
    catch {
        throw alert("Error getting data")
    }
}

function remapData(data: Log[]) {
    const mappedData = {
        labels: data.map((e: Log) => Intl.DateTimeFormat(undefined ,{ dateStyle: 'short', timeStyle: 'short',}).format(new Date(e.CreatedAt))),
        datasets: [
            {
                label: "People count",
                data: data.map((e: Log) => e.count),
                tension: 0.2,
                borderWidth: 7,
            },
        ]
    }
    return mappedData
}