<script>
  import { onMount } from "svelte";
  import Chart from "chart.js/auto";

  let portfolio = null;
  const data = {
    labels: ["Expenses", "Savings", "Investments"],
    datasets: [
      {
        label: "My First Dataset",
        data: [300, 50, 100],
        backgroundColor: ["#7000e1", "#fc8800", "#00b0e8"],
        hoverOffset: 10,
        borderWidth: 0,
      },
    ],
  };
  const config = {
    type: "doughnut",
    data: data,
    options: {
      responsive: true,
      cutout: "60%",
      spacing: 4,
      animation: { animateScale: true },
      plugins: {
        legend: {
          position: "top",
          display: true,
          labels: {
            usePointStyle: true,
            padding: 20,
            font: {
              size: 14,
            },
          },
        },
        title: {
          display: false,
          text: "My Personal Portfolio",
        },
      },
    },
  };
  onMount(() => {
    const ctx = portfolio.getContext("2d");
    let chart = new Chart(ctx, config);
  });
</script>

<div class="chart">
  <canvas bind:this={portfolio}/>
</div>
