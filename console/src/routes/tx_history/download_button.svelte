<script>
  let contentType
  let filename

  import mime from 'mime'

  export let txID
  export let txApiKey
  export let txFilename

  function download() {
    fetch(`${BASE_URL}/api/v1/transactions/${txID}`, {
      method: 'GET',
      headers: {
        Authorization: 'Bearer ' + txApiKey,
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not OK')
        }
        contentType = response.headers.get('content-type')
        return response.blob()
      })
      .then((blob) => URL.createObjectURL(blob))
      .then((blobUrl) => {
        if (txFilename == '') {
          var extension = mime.getExtension(contentType)
          filename = txID + '.' + extension
        } else {
          filename = txFilename
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

<button class="button is-primary" on:click={download}>Download</button>
