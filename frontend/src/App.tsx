import { createSignal } from 'solid-js'
import { DefaultChart, Line } from 'solid-chartjs'
import './App.css'
import { GetDaily, GetWeekly, GetAll, GetMonthly } from './getDataService'

const emptyChart = {datasets: [{label: "", data: [0]}, {label: "", data: [0]}, {label: "", data: [0]}], labels: ["", "", ""]}

function App() {
  const [chartData, setChartData] = createSignal(emptyChart)
  GetDaily().then((data) => setChartData(data))
  console.log(chartData())

  const chartOptions = {
      backgroundColor: "#ff0000",
      responsive: true,
      maintainAspectRatio: false,
  }
  GetDaily()

  return (
    <>
    {/* <select onClick={e => GetData(e.target.selectedOptions[0].value)}>
        <option value="day" selected>Day</option>
        <option value="month">Month</option>
        <option value="week">Week</option>
        <option value="all">All</option>
      </select> */}
      <Line type='line' data={chartData()} options={chartOptions} />
    </>
  )

  function GetData(newState: string) {
      let data = emptyChart
      switch(newState) {
        case 'day':
          GetDaily().then(d => data = d)
          break
        case 'week':
          GetWeekly().then(d => data = d)
          break
        case 'month':
          GetMonthly().then(d => data = d)
          break
        case 'all':
          GetAll().then(d => data = d)
          break
      }
      setChartData(data)
  }
}


export default App
