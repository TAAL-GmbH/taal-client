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

<tr>
  <td>{year}-{month}-{day} {hours}:{minutes}:{seconds}</td>
  <td
    ><a href="https://www.whatsonchain.com/tx/{transaction.id}"
      >{truncate(transaction.id)}</a
    ></td
  >
  <td>{transaction.api_key}</td>
  <td>{transaction.data_bytes}</td>
  <td>{transaction.filename}</td>
  <td
    ><DownloadButton
      txFilename={transaction.filename}
      txApiKey={transaction.api_key}
      txID={transaction.id}
    /></td
  >
</tr>
