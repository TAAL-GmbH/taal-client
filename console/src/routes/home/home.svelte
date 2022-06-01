<script>
  import Health from './health.svelte'
  import TransactionStats from './transaction_stats.svelte'
</script>

<main>
  <div class="columns">
    <div class="column is-three-quarters">
      <h1>Transaction stats</h1>
      <TransactionStats />
    </div>
    <div class="column ">
      <h1>Status</h1>
      <Health />
    </div>
  </div>
</main>

<style>
  h1 {
    font-size: 1.4em;
    font-weight: bold;
    display: block;
  }
</style>
