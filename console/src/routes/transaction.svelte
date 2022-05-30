<script>
  export let transaction

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
        console.log(response.headers['Content-Type'])
        return response.blob()
      })
      .then((blob) => URL.createObjectURL(blob))
      .then((blobUrl) => {
        var a = document.createElement('a')
        a.href = blobUrl
        a.download = 'FILENAME.png'
        a.innerHTML = 'Click here to download the file'
        a.click()
      })
      .catch((err) => console.error(err))
  }
</script>

<tr>
  <td>{transaction.created_at}</td>
  <td
    ><a href="https://www.whatsonchain.com/tx/{transaction.id}"
      >{transaction.id}</a
    ></td
  >
  <td>{transaction.api_key}</td>
  <td>{transaction.data_bytes}</td>
  <td
    ><button class="button is-primary" on:click={download}>Download</button></td
  >
</tr>
