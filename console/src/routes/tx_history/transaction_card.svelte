<script>
  import DownloadButton from './download_button.svelte'
  import {
    GetFormatedTxTimestamp,
    TruncateTxID,
    TxDataSize,
    GetColor,
  } from '../util/format_functions.svelte'
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

<div class="column is-one-quarter">
  <div class="card">
    <div class="card-content">
      <div class="media">
        <div class="media-content">
          <p class="title is-4">
            <button
              id="copyButton"
              class="button is-small is-vcentered "
              on:click|preventDefault={copyToClipboard}
            >
              <Fa icon={faCopy} color="silver" />
            </button>
            {#if transaction.api_key.startsWith('testnet')}
              <a
                href="https://test.whatsonchain.com/tx/{transaction.id}"
                target="_blank"
              >
                {TruncateTxID(transaction.id)}
              </a>
            {:else}
              <a
                href="https://www.whatsonchain.com/tx/{transaction.id}"
                target="_blank"
              >
                {TruncateTxID(transaction.id)}
              </a>
            {/if}
          </p>
          <p class="subtitle is-6">
            Created at: {GetFormatedTxTimestamp(transaction.created_at)}
          </p>

          <p class="is-6">
            <span class="dot" style={backgroundColor} />
            {transaction.api_key.slice(0, 8)}
            {transaction.api_key.slice(8)}
          </p>

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
          {#if transaction.secret}
            <span class="key">
              <Fa icon={faKey} />
            </span>
          {/if}
        </div>
      </div>
    </div>
  </div>
</div>

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
