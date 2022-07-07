<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'
  import mime from 'mime'

  import ButtonSelect from '../../lib/components/button-select/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Table from '../../lib/components/table/index.svelte'
  import Spinner from '../../lib/components/spinner/index.svelte'

  import {
    downloadFileBlob,
    filterUnique,
    getColorFromDistinct,
  } from '../../lib/utils'
  import { spinCount } from '../../lib/stores'
  import * as api from '../../lib/api'
  import { colDefs } from './data'

  const { addNotification } = getNotificationsContext()

  let transactions = []
  let rangeValue = '720'
  let distinctKeys = []

  $: {
    distinctKeys = transactions.map((tx) => tx.api_key).filter(filterUnique)
  }

  function onRange(e) {
    getTransactions(e.detail.value)
  }

  function onAction(e) {
    download(e.detail.value)
  }

  const getRowIconActions = (tableName, item, idField = 'id') => {
    const icons = [
      {
        icon: 'download',
        type: 'download',
      },
    ]
    if (item?.secret) {
      icons.push({
        icon: 'key',
        render: 'icon',
        disabled: true,
      })
    }
    return icons
  }

  const getRenderProps = (name, colDef, idField, item) => {
    return {
      statusColor: getColorFromDistinct(item.api_key, distinctKeys),
    }
  }

  function getTransactions(hours) {
    api.getTransactions(
      hours,
      (data) => {
        transactions = data.transactions
      },
      (error) => {
        addNotification({
          text: `Error: ${error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })
      }
    )
  }

  function download(tx) {
    let contentType = ''

    fetch(`${BASE_URL}/api/v1/transactions/${tx.id}`, {
      method: 'GET',
      headers: {
        Authorization: 'Bearer ' + tx.api_key,
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not OK')
        }
        contentType = response.headers.get('content-type')
        return response.blob()
      })
      .then((blob) => {
        let filename = ''

        if (tx.filename === '') {
          const extension = mime.getExtension(contentType)
          const filenameSuffix = extension ? '.' + extension : ''

          filename = tx.id + filenameSuffix
        } else {
          filename = tx.filename
        }

        downloadFileBlob(filename, blob)

        addNotification({
          text: `Successfully downloaded: ${filename}`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 2000,
        })
      })
      .catch((err) =>
        addNotification({
          text: `Error: ${err}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })
      )
  }

  onMount(() => getTransactions(rangeValue))
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="History" />
    <Spacer h={24} />
    <div class="sub-row">
      <ButtonSelect
        name="range"
        value={rangeValue}
        items={[
          { label: '24H', value: '24' },
          { label: '1W', value: '168' },
          { label: '1M', value: '720' },
          { label: 'All', value: '' },
        ]}
        on:change={onRange}
      />
    </div>
    <Spacer h={24} />
    <Table
      {colDefs}
      data={transactions}
      pagination={{
        page: 1,
        pageSize: 15,
      }}
      {getRenderProps}
      {getRowIconActions}
      on:action={onAction}
    />
  </div>
</PageWithMenu>

{#if $spinCount > 0}
  <Spinner />
{/if}

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>
