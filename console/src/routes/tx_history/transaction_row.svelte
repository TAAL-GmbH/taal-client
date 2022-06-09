<script>
  import {
    GetFormatedTxTimestamp,
    TruncateTxID,
    TxDataSize,
    TxFilename,
    GetColor,
  } from '../util/format_functions.svelte'
  import DownloadButton from './download_button.svelte'
  export let transaction
  export let distinctAPIKeys

  // FontAwesome icon...
  import Fa from 'svelte-fa'
  import { faCopy, faKey } from '@fortawesome/free-solid-svg-icons'

  let backgroundColor

  function OnDistinctKeyChange(distinctKeys) {
    backgroundColor = GetColor(transaction.api_key, distinctKeys)
  }

  $: OnDistinctKeyChange(distinctAPIKeys)

  function copyToClipboard() {
    navigator.clipboard.writeText(transaction.id)
  }
</script>

<tr>
  <td class="is-vcentered">{GetFormatedTxTimestamp(transaction.created_at)}</td>
  <td class="is-vcentered">
    <div style="margin-top: 6px;">
      <button
        id="copyButton"
        class="button is-small is-vcentered "
        on:click|preventDefault={copyToClipboard}
      >
        <Fa icon={faCopy} color="silver" />
      </button>
      <a
        href="https://www.whatsonchain.com/tx/{transaction.id}"
        target="_blank"
      >
        <span class="is-vcentered">
          {TruncateTxID(transaction.id)}
        </span>
      </a>
    </div>
  </td>
  <td class="is-vcentered">
    <span class="dot" style={backgroundColor} />
    {transaction.api_key}
  </td>
  <td class="is-vcentered" align="right"
    >{TxDataSize(transaction.data_bytes)}</td
  >
  <td class="is-vcentered">{TxFilename(transaction.filename)}</td>
  <td class="is-vcentered"
    ><DownloadButton
      txFilename={transaction.filename}
      txApiKey={transaction.api_key}
      txID={transaction.id}
    />
    {#if transaction.secret}
      <span class="key">
        <Fa icon={faKey} />
      </span>
    {/if}
  </td>
</tr>

<style>
  .dot {
    height: 10px;
    width: 10px;
    border-radius: 50%;
    display: inline-block;
    margin-right: 5px;
  }

  .key {
    display: inline-block;
    padding: 10px;
  }

  #copyButton {
    display: inline-block;
    border: none;
    padding-top: 1px;
    padding-left: 0;
    padding-right: 0;
  }
</style>
