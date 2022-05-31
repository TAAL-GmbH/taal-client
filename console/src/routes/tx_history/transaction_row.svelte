<script>
  import {GetFormatedTxDate, TruncateTxID, TxDataSize, TxFilename} from './transaction_format_functions.svelte'
  import DownloadButton from './download_button.svelte'
  export let transaction

</script>

<tr>
  <td>{GetFormatedTxDate(transaction.created_at)}</td>
  <td
    ><a href="https://www.whatsonchain.com/tx/{transaction.id}"
      >{TruncateTxID(transaction.id)}</a
    ></td
  >
  <td>{transaction.api_key}</td>
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
