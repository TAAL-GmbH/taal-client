<script>
  import DownloadButton from './download_button.svelte'
  import {GetFormatedTxTimestamp, TruncateTxID, TxDataSize, GetColor} from './transaction_format_functions.svelte'
  export let transaction
  export let distinctAPIKeys

  let color = GetColor(transaction.api_key, distinctAPIKeys)
 </script>

<div class="column is-one-quarter">
  <div class="card">
    <div class="card-content">
      <div class="media">
        <div class="media-content">
          <p class="title is-4">
            <a href="https://www.whatsonchain.com/tx/{transaction.id}"
              >ID: {TruncateTxID(transaction.id)}</a
            >
          </p>
          <p class="subtitle is-6">
            Created at: {GetFormatedTxTimestamp(transaction.created_at)}
          </p>
          <p style={color} class="subtitle is-6">API Key: {transaction.api_key}</p>
          <p class="subtitle is-6">
            Data size: {TxDataSize(transaction.data_bytes)}
          </p>
          {#if transaction.filename != ''}
            <p class="subtitle is-6">Filename: {transaction.filename}</p>
          {/if}
          <DownloadButton
            txFilename={transaction.filename}
            txApiKey={transaction.api_key}
            txID={transaction.id}
          />
        </div>
      </div>
    </div>
  </div>
</div>
