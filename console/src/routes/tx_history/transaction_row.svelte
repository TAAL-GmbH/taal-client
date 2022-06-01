<script>
  import {GetFormatedTxTimestamp, TruncateTxID, TxDataSize, TxFilename, GetColor} from './transaction_format_functions.svelte'
  import DownloadButton from './download_button.svelte'
  export let transaction
  export let distinctAPIKeys

  let color = GetColor(transaction.api_key, distinctAPIKeys)

</script>

<tr>
  <td>{GetFormatedTxTimestamp(transaction.created_at)}</td>
  <td
    ><a href="https://www.whatsonchain.com/tx/{transaction.id}"
      >{TruncateTxID(transaction.id)}</a
    ></td
  >
  <td style={color}>{transaction.api_key}</td>
  <td align="right">{TxDataSize(transaction.data_bytes)}</td>
  <td>{TxFilename(transaction.filename)}</td>
  <td
    ><DownloadButton
      txFilename={transaction.filename}
      txApiKey={transaction.api_key}
      txID={transaction.id}
    /></td
  >
</tr>
