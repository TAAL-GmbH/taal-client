<script>
  import { onMount } from 'svelte'
  import DownloadButton from './download_button.svelte'
  export let transaction
  let date = new Date()
  let year
  let month
  let day
  let hours
  let minutes
  let seconds

  onMount(() => {
    date = new Date(transaction.created_at)
    year = date.getUTCFullYear()
    month = date.getUTCMonth() + 1 // month is zero indexed
    day = date.getUTCDate()
    hours = date.getUTCHours()
    minutes = date.getUTCMinutes()
    seconds = date.getUTCSeconds()
  })

  function truncate(str) {
    const n = 10
    return str.length > n ? str.slice(0, 5) + '...' + str.slice(-5) : str
  }
</script>

<div class="column is-one-quarter">
  <div class="card">
    <div class="card-content">
      <div class="media">
        <div class="media-content">
          <p class="title is-4">
            <a href="https://www.whatsonchain.com/tx/{transaction.id}"
              >ID: {truncate(transaction.id)}</a
            >
          </p>
          <p class="subtitle is-6">
            Created at: {year}-{month}-{day}
            {hours}:{minutes}:{seconds}
          </p>
          {#if transaction.filename != ''}
            <p class="subtitle is-6">Filename: {transaction.filename}</p>
          {/if}
          <p class="subtitle is-6">API Key: {transaction.api_key}</p>
          <p class="subtitle is-6">
            Data size [bytes]: {transaction.data_bytes}
          </p>
          <DownloadButton
            txFilename={transaction.filename}
            txApiKey={transaction.api_key}
            txID={transaction.id}
          />
        </div>
      </div>
    </div>
  </div>
</div>
