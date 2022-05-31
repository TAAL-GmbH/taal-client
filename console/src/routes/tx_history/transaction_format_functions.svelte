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
</script>
