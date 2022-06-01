<script>
    import Health from './health.svelte'
    import TransactionStats from './transaction_stats.svelte'
</script>

<main>
    <h1>Status</h1>
    <Health/>
    <h1>Transaction stats</h1>
    <TransactionStats/>
</main>

<style>
  h1 {
    font-size: 1.4em;
    font-weight: bold;
    display: block;
  }
</style>