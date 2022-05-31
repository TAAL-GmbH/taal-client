<script context="module">
  export function GetFormatedTxDate(dateString) {
    var date = new Date(dateString)
    var year = date.getUTCFullYear()
    var month = date.getUTCMonth() + 1 // month is zero indexed
    var day = date.getUTCDate()
    var hours = date.getUTCHours()
    var minutes = date.getUTCMinutes()
    var seconds = date.getUTCSeconds()
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
  }

  export function TruncateTxID(str) {
    const n = 10
    return str.length > n ? str.slice(0, 5) + '...' + str.slice(-5) : str
  }

  export function TxDataSize(dataSizeBytes) {
    var nrOfDigits = dataSizeBytes.toString().length
    if (nrOfDigits >= 10) {
      // size >= 1 GB
      return numberWithCommas(dataSizeBytes / 1000000000) + ' GB'
    }
    if (nrOfDigits >= 7) {
      // size >= 1 MB
      return numberWithCommas(dataSizeBytes / 1000000) + ' MB'
    }
    if (nrOfDigits >= 4) {
      // size >= 1 kB
      return numberWithCommas(dataSizeBytes / 1000) + ' kB'
    }

    return numberWithCommas(dataSizeBytes) + ' B'
  }

  function numberWithCommas(x) {
    return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
  }

  export function TxFilename(filename) {
    if (filename == '') {
      return '-'
    }

    return filename

  }

</script>
