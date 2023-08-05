import { For, createSignal, onMount } from 'solid-js'
import { Line } from 'solid-chartjs'
import { GetDaily, GetWeekly, GetAll, GetMonthly } from './getDataService'
import { Chart, ChartOptions, Tooltip } from 'chart.js'

import './App.css'

const emptyChart = { datasets: [{ label: "", data: [0] }, { label: "", data: [0] }, { label: "", data: [0] }], labels: ["", "", ""] }

const chartOptions: ChartOptions = {
    backgroundColor: "##ff93a4",
    color: "##ff6384",
    borderColor: "#ff6384",
    responsive: true,
    maintainAspectRatio: false,
    hover: {
        mode: 'index',
        intersect: true
    },
    indexAxis: 'x',
    scales: {
        y: {
            grid: {
                color: "#444"
            },
            ticks: {
                color: "#fff"
            }
        },
        x: {
            grid: {
                color: "#444"
            },
            ticks: {
                color: "#fff"
            }
        }
    }
}

function App() {
    onMount(() => {
        Chart.register([Tooltip])
    })

    const [chartData, setChartData] = createSignal(emptyChart)
    GetDaily().then((data) => setChartData(data))

    return (
        <>
            <nav>
                <h1>6obcy peple monitor</h1>
                <h4>See how many people are on <a href="https://6obcy.org/">6obcy.org</a> at selected time</h4>
            </nav>
            <main>
                <select onchange={e => GetData(e.currentTarget.selectedIndex)}>
                    <For each={['day', 'week', 'month', 'all']}>
                        {(e) => (
                            <option value={e}>{e}</option>
                        )}
                    </For>
                </select>
                <Line type='line' data={chartData()} options={chartOptions} />
            </main>
        </>
    )

    async function GetData(newState: number) {
        let data = emptyChart
        switch (newState) {
            case 0:
                data = await GetDaily()
                break
            case 1:
                data = await GetWeekly()
                break
            case 2:
                data = await GetMonthly()
                break
            case 3:
                data = await GetAll()
                break
        }
        setChartData(data)
    }
}


export default App
