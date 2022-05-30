<script>
  export let transaction

  let filename = 'Data'

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
  <td>{transaction.created_at}</td>
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
