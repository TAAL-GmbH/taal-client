<script>
  import { onMount } from 'svelte'

  let filename = 'Data'
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

  function download() {
    fetch(`${BASE_URL}/api/v1/transactions/${transaction.id}`, {
      method: 'GET',
      headers: {
        Authorization: 'Bearer ' + transaction.api_key,
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not OK')
        }
        return response.blob()
      })
      .then((blob) => URL.createObjectURL(blob))
      .then((blobUrl) => {
        if (transaction.filename != '') {
          filename = transaction.filename
        }

        var a = document.createElement('a')
        a.href = blobUrl
        a.download = filename
        a.innerHTML = 'Click here to download the file'
        a.click()
      })
      .catch((err) => console.error(err))
  }
</script>

<tr>
  <td>{year}-{month}-{day} {hours}:{minutes}:{seconds}</td>
  <td
    ><a href="https://www.whatsonchain.com/tx/{transaction.id}"
      >{transaction.id}</a
    ></td
  >
  <td>{transaction.api_key}</td>
  <td>{transaction.data_bytes}</td>
  <td>{transaction.filename}</td>
  <td
    ><button class="button is-primary" on:click={download}>Download</button></td
  >
</tr>
