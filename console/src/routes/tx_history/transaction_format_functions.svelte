<script context="module">
  export function GetFormatedTxTimestamp(dateString) {
    var date = new Date(dateString)
    var dateString = date.toISOString().split('T')[0]
    var timeString = date.toISOString().split('T')[1].split('.')[0]
    return `${dateString} ${timeString}`
  }

  export function TruncateTxID(str) {
    const n = 10
    return str.length > n ? str.slice(0, 5) + '...' + str.slice(-5) : str
  }

  export function GetColor(apiKey, distinctAPIKeys) {
    var colorFactor = 0
    if (distinctAPIKeys.length > 1) {
      var index = distinctAPIKeys.indexOf(apiKey)
      var part = 360 / distinctAPIKeys.length

      colorFactor = index * part
    }

    var color = 'background-color: hsl(' + colorFactor + ', 100%, 85%);'
    return color
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
